package engine

type Engine struct {
	sourceCollector *SourceCollector
}

func (engine *Engine) CreateActionPlan(source string, simple bool) (string, error) {
	actionPlan := NewActionPlan(nil)
	if err := engine.runSourceCollector(source); err != nil {
		return "", err
	}
	facts, err := ruleset.gatherFacts(source, engine.sourceCollector, nil)
	if err != nil {
		return "", err
	}
	for _, action := range registeredActions {
		pass := true
		for _, fact := range action.AdmissionCriteria {
			if !facts.Contains(fact) {
				pass = false
			}
		}
		for _, fact := range action.ExclusionCriteria {
			if facts.Contains(fact) {
				pass = false
			}
		}
		if pass {
			actionPlan.AddAction(action.Name, action.Inputs)
		}
	}
	actionPlanJson, err := actionPlan.ToJson()
	if err != nil {
		return "", err
	}
	return actionPlanJson, nil
}

func (engine *Engine) runSourceCollector(source string) error {
	return engine.sourceCollector.run(source)
}

func New() *Engine {
	return &Engine{sourceCollector: collector}
}
