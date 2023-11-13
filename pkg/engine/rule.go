package engine

import (
	mapset "github.com/deckarep/golang-set/v2"
)

type Rule func(string, Collector, mapset.Set[Fact]) (Fact, error)

var ruleset = NewRuleset()

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

func (rs *Ruleset) getRulesetNode(rule *Rule) *RulesetNode {
	if node, ok := rs.index[rule]; ok {
		return node
	}
	return nil
}

func (rs *Ruleset) append(parentRule, childRule *Rule) {
	parentNode := rs.getRulesetNode(parentRule)
	if parentNode == nil {
		parentNode = &RulesetNode{rule: parentRule}
		rs.index[parentRule] = parentNode
		rs.root = append(rs.root, parentNode)
	}
	if childRule == nil {
		return
	}
	childNode := rs.getRulesetNode(childRule)
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

// gatherFacts sets source facts by recursing each branch of the
// rule tree until an end node or nil fact is reached.
func (rs *Ruleset) gatherFacts(source string, collector *SourceCollector, nodes []*RulesetNode) (mapset.Set[Fact], error) {
	facts := mapset.NewSet[Fact]()
	if nodes == nil {
		nodes = rs.root
	}
	for _, node := range nodes {
		fact, err := (*node.rule)(source, collector, facts)
		if err != nil {
			return facts, err
		}
		if fact != NilFact {
			facts.Add(fact)
			if node.childNodes != nil {
				childFacts, err := rs.gatherFacts(source, collector, node.childNodes)
				if err != nil {
					return facts, err
				}
				facts.Append(childFacts.ToSlice()...)
			}
		}
	}
	return facts, nil
}

func NewRuleset() *Ruleset {
	return &Ruleset{
		root:  make([]*RulesetNode, 0),
		index: make(map[*Rule]*RulesetNode),
	}
}

func AddToRuleset(parent, child *Rule) {
	ruleset.append(parent, child)
}
