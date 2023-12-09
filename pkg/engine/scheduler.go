package engine

import (
	"fmt"
	"strings"

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
		if assignments[OnDemand] == nil || assignments[OnDemand].Cardinality() == 0 {
			break
		}
		for _, action := range assignments[OnDemand].ToSlice() {
			matched := false
			for _, artifact := range action.OutputArtifacts {
				stages := []Stage{}
				requiredInputStages, ok := s.requiredInputs[artifact]
				if ok {
					stages = append(stages, requiredInputStages.ToSlice()...)
				}
				optionalInputStages, ok := s.optionalInputs[artifact]
				if ok {
					stages = append(stages, optionalInputStages.ToSlice()...)
				}
				if len(stages) == 0 {
					break
				}
				matched = true
				firstOccurance := stages[0]
				for i := 1; i < len(stages); i++ {
					if stages[i] < firstOccurance {
						firstOccurance = stages[i]
					}
				}
				if firstOccurance == OnDemand {
					if unmatchableWatchList.Contains(action) {
						return fmt.Errorf("outputs for on-demand action '%s' cannot be resolved", action.Name)
					}
					unmatchableWatchList.Add(action)
					continue
				}
				assignments[firstOccurance].Add(action)
				assignments[OnDemand].Remove(action)
				for _, artifact := range action.InputArtifacts {
					s.requiredInputs[artifact].Remove(OnDemand)
					s.requiredInputs[artifact].Add(firstOccurance)
				}
			}
			if !matched {
				assignments[OnDemand].Remove(action)
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
			optionallyDeferredActions := mapset.NewSet[string]()
			for {
				startingCardinality := actions.Cardinality()
			ActionLoop:
				for _, action := range actions.ToSlice() {
					if len(action.OptionalInputArtifacts) > 0 {
						optionallyDeferredActions.Add(action.Name)
					}
					if len(action.InputArtifacts) == 0 && len(action.OptionalInputArtifacts) == 0 {
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
						for _, input := range action.OptionalInputArtifacts {
							var inputIsFulfilled bool
							for _, output := range assignedActionOutputs {
								if input == output {
									inputIsFulfilled = true
								}
							}
							if !inputIsFulfilled {
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
					// assign optionally deferred actions
					hasUnresolvedActions := false
					for _, action := range actions.ToSlice() {
						if !optionallyDeferredActions.Contains(action.Name) {
							hasUnresolvedActions = true
						} else {
							sortedAssignments[stage] = append(sortedAssignments[stage], action)
							assignedActionOutputs = append(assignedActionOutputs, action.OutputArtifacts...)
							actions.Remove(action)
						}
					}
					if hasUnresolvedActions {
						return nil, fmt.Errorf("the following action has inputs that cannot be resolved: '%s'", strings.Join(actionsWithUnresolvedInputs.ToSlice(), ","))
					}
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
