package engine

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/pkg/plan"
)

type Engine struct {
	sourceCollector *SourceCollector
}

// CreateActionPlan .
func (engine *Engine) CreateActionPlan(source string) (string, error) {
	facts := mapset.NewSet[Fact]()
	actionPlan := plan.NewActionPlan(nil)
	if err := engine.runSourceCollector(source); err != nil {
		return "", err
	}
	ruleset.gatherFacts(source, engine.sourceCollector, facts, nil)
	for name, resolver := range admissionResolvers {
		pass := true
		for _, fact := range resolver.criteria {
			if !facts.Contains(fact) {
				pass = false
			}
		}
		if pass {
			actionPlan.AddAction(name, resolver.userInputs)
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
	return &Engine{sourceCollector: newSourceCollector()}
}

var admissionResolvers = map[string]AdmissionResolver{}

type AdmissionResolver struct {
	criteria   []Fact
	userInputs []string
}

// RegisterAdmissionTest .
func RegisterAdmissionResolver(name string, criteria []Fact, userInputs []string) {
	admissionResolvers[name] = AdmissionResolver{criteria, userInputs}
}

type Rule func(string, *SourceCollector, mapset.Set[Fact]) (Fact, error)

type RulesetNode struct {
	rule       *Rule
	childNodes []*RulesetNode
}

func (n *RulesetNode) addChild(r *RulesetNode) {
	n.childNodes = append(n.childNodes, r)
}

type Ruleset struct {
	root  []*RulesetNode
	index map[*Rule]*RulesetNode
}

func (rs *Ruleset) getRuleNode(rule *Rule) *RulesetNode {
	if node, ok := rs.index[rule]; ok {
		return node
	}
	return nil
}

func (rs *Ruleset) append(parentRule, childRule *Rule) {
	parentNode := rs.getRuleNode(parentRule)
	if parentNode == nil {
		parentNode = &RulesetNode{rule: parentRule}
		rs.index[parentRule] = parentNode
		rs.root = append(rs.root, parentNode)
	}
	if childRule == nil {
		return
	}
	childNode := rs.getRuleNode(childRule)
	if childNode == nil {
		childNode = &RulesetNode{rule: childRule}
		rs.index[childRule] = childNode
	} else {
		for i, rootNode := range rs.root {
			if rootNode.rule == childRule {
				rs.root = append(rs.root[:i], rs.root[i+1:]...)
				break
			}
		}
	}
	parentNode.addChild(childNode)
}

func (rs *Ruleset) gatherFacts(source string, collector *SourceCollector, facts mapset.Set[Fact], nodes []*RulesetNode) error {
	if nodes == nil {
		nodes = rs.root
	}
	for _, node := range nodes {
		fact, err := (*node.rule)(source, collector, facts)
		if err != nil {
			return err
		}
		if fact != NilFact {
			facts.Add(fact)
			if node.childNodes != nil {
				if err := rs.gatherFacts(source, collector, facts, node.childNodes); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewRuleset() *Ruleset {
	return &Ruleset{
		root:  make([]*RulesetNode, 0),
		index: make(map[*Rule]*RulesetNode),
	}
}

var ruleset = NewRuleset()

func addToRuleset(parent, child *Rule) {
	ruleset.append(parent, child)
}
