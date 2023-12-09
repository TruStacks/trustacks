package engine

import (
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
)

func TestRulesetNodeAddChild(t *testing.T) {
	node := &RulesetNode{}
	node.addChild(&RulesetNode{})
	node.addChild(&RulesetNode{})
	node.addChild(&RulesetNode{})
	assert.Len(t, node.childNodes, 3)
}

func TestRulesetAppend(t *testing.T) {
	var objectIsPersonRule Rule = func(string, Collector, mapset.Set[Fact]) (Fact, error) {
		return NilFact, nil
	}
	var PersonHasBrownHairRule Rule = func(string, Collector, mapset.Set[Fact]) (Fact, error) {
		return NilFact, nil
	}

	rs := NewRuleset()
	rs.append(&objectIsPersonRule, &PersonHasBrownHairRule)

	parentNode := rs.getRulesetNode(&objectIsPersonRule)
	assert.Equal(t, parentNode.rule, &objectIsPersonRule)

	childNode := rs.getRulesetNode(&PersonHasBrownHairRule)
	assert.Equal(t, childNode.rule, &PersonHasBrownHairRule)
}

func TestGatherFacts(t *testing.T) {
	objectIsPersonFact := NewFact()
	PersonHasBrownHairFact := NewFact()
	var objectIsPersonRule Rule = func(string, Collector, mapset.Set[Fact]) (Fact, error) {
		return objectIsPersonFact, nil
	}
	var PersonHasBrownHairRule Rule = func(string, Collector, mapset.Set[Fact]) (Fact, error) {
		return PersonHasBrownHairFact, nil
	}
	rs := NewRuleset()
	rs.append(&objectIsPersonRule, &PersonHasBrownHairRule)
	facts, err := rs.gatherFacts("./", nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, facts.Contains(objectIsPersonFact))
	assert.True(t, facts.Contains(PersonHasBrownHairFact))
}
