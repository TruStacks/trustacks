package engine

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
	mapset "github.com/deckarep/golang-set/v2"
)

type scheduler struct {
	requiredInputs map[Artifact]mapset.Set[Stage]
	optionalInputs map[Artifact]mapset.Set[Stage]
}

func newScheduler() *scheduler {
	return &scheduler{
		requiredInputs: map[Artifact]mapset.Set[Stage]{},
		optionalInputs: map[Artifact]mapset.Set[Stage]{},
	}
}

func (s *scheduler) assignActivityStage(actions []string) map[Stage]mapset.Set[*Action] {
	assignments := map[Stage]mapset.Set[*Action]{}
	for _, actionName := range actions {
		for registeredActionName, registeredAction := range registeredActions {
			if actionName == registeredActionName {
				if _, ok := assignments[registeredAction.Stage]; !ok {
					assignments[registeredAction.Stage] = mapset.NewSet[*Action]()
				}
				assignments[registeredAction.Stage].Add(registeredAction)
			}
		}
	}
	return assignments
}

func (s *scheduler) bindActionInputs(assignments map[Stage]mapset.Set[*Action]) {
	for stage, actions := range assignments {
		for _, action := range actions.ToSlice() {
			for _, artifact := range action.InputArtifacts {
				if _, ok := s.requiredInputs[artifact]; !ok {
					s.requiredInputs[artifact] = mapset.NewSet[Stage]()
				}
				s.requiredInputs[artifact].Add(stage)
			}
			for _, artifact := range action.OptionalInputArtifacts {
				if _, ok := s.optionalInputs[artifact]; !ok {
					s.optionalInputs[artifact] = mapset.NewSet[Stage]()
				}
				s.optionalInputs[artifact].Add(stage)
			}
		}
	}
}

func (s *scheduler) assignOnDemandActions(assignments map[Stage]mapset.Set[*Action]) error {
	unmatchableWatchList := mapset.NewSet[*Action]()
	for {
		// if on demand actions are nil or empty then all actions are assigned
		// or no on demand actions exist in the action plan schedule.
		if assignments[OnDemandStage] == nil || assignments[OnDemandStage].Cardinality() == 0 {
			break
		}
		for _, action := range assignments[OnDemandStage].ToSlice() {
			matched := false
			// check if the on demand action output artifacts have actions
			// in an activity stage that require their outputs.
			for _, artifact := range action.OutputArtifacts {
				if stages, ok := s.requiredInputs[artifact]; ok {
					matched = true
					stages := stages.ToSlice()
					firstOccurance := stages[0]
					for i := 1; i < len(stages); i++ {
						if stages[i] < firstOccurance {
							firstOccurance = stages[i]
						}
					}
					// if the first dependency of the input is in the on demand
					// stage, then we must wait until the next loop iteration
					// for the dependant action to be assigned to a finite
					// activity stage.
					if firstOccurance == OnDemandStage {
						// if the action is already in the unmatched list then
						// the matching failed for two loop cycles.
						// This means that the associated input could not be
						// tranferred to to a finite activity stage and cannot
						// be resolved.
						if unmatchableWatchList.Contains(action) {
							return fmt.Errorf("outputs for on-demand action '%s' cannot be resolved", action.Name)
						}
						// add the action to the watch list until the next loop
						// cycle.
						unmatchableWatchList.Add(action)
						continue
					}
					// transfer the action from the on demand stage to the
					// required activity stage.
					assignments[firstOccurance].Add(action)
					assignments[OnDemandStage].Remove(action)
					// it is possible for an on demand action to have inputs.
					// transfer inputs along with the action to the appropriate
					// stage and remove those inputs from the on demand stage
					// since the action no longer exists there.
					for _, artifact := range action.InputArtifacts {
						s.requiredInputs[artifact].Remove(OnDemandStage)
						s.requiredInputs[artifact].Add(firstOccurance)
					}
				}
			}
			if !matched {
				err := fmt.Errorf("one or more outputs from on-demand action '%s' could not be matched to an action in a finite stage. this action will be excluded from the action plan", action.Name)
				log.Info("", "err", err)
				assignments[OnDemandStage].Remove(action)
			}
		}
	}
	return nil
}

func (s *scheduler) sortActions(assignments map[Stage]mapset.Set[*Action]) (map[Stage][]*Action, error) {
	sortedAssignments := map[Stage][]*Action{}
	assignedActionOutputs := []Artifact{}
	for stageIndex := range actionStages {
		stage := Stage(stageIndex)
		if actions, ok := assignments[stage]; ok {
			actionsWithUnresolvedInputs := mapset.NewSet[string]()
			for {
				startingCardinality := actions.Cardinality()
			ActionLoop:
				for _, action := range actions.ToSlice() {
					if len(action.InputArtifacts) == 0 {
						sortedAssignments[stage] = append(sortedAssignments[stage], action)
					} else {
						for _, input := range action.InputArtifacts {
							var inputIsFulfilled bool
							for _, output := range assignedActionOutputs {
								if input == output {
									inputIsFulfilled = true
								}
							}
							if !inputIsFulfilled {
								actionsWithUnresolvedInputs.Add(action.Name)
								continue ActionLoop
							}
						}
						sortedAssignments[stage] = append(sortedAssignments[stage], action)
					}
					assignedActionOutputs = append(assignedActionOutputs, action.OutputArtifacts...)
					actions.Remove(action)
				}
				if actions.Cardinality() == 0 {
					break
				} else if startingCardinality == actions.Cardinality() {
					return nil, fmt.Errorf("the scheduler has detected unresolved inputs for the following actions: '%s'", strings.Join(actionsWithUnresolvedInputs.ToSlice(), ","))
				}
			}
		}
	}
	return sortedAssignments, nil
}

func (s *scheduler) schedule(actions []string) (map[Stage][]*Action, error) {
	assignments := s.assignActivityStage(actions)
	s.bindActionInputs(assignments)
	if err := s.assignOnDemandActions(assignments); err != nil {
		return nil, err
	}
	sortedAssignments, err := s.sortActions(assignments)
	if err != nil {
		return nil, err
	}
	return sortedAssignments, nil
}
