package classes

import (
	"becmi/attributes"
	"fmt"
	"testing"
)

func TestCleric_Name(t *testing.T) {
	var c Cleric
	if c.Name() != "Cleric" {
		t.Errorf("Name is incorrect. Expected Cleric, got %s", c.Name())
	}
}

func TestCleric_Level(t *testing.T) {
	var c Cleric
	level := c.Level(150000)
	expectedLevel := 8
	if level != expectedLevel {
		t.Errorf("Level is incorrect. Expected %d, got %d", expectedLevel, level)
	}
}

func TestCleric_NextLevelAt(t *testing.T) {
	var c Cleric
	level := 5
	expectedXP := 25000
	if nextLevel := c.NextLevelAt(level); nextLevel != expectedXP {
		t.Errorf("NextLevelAt is incorrect. Expected %d, got %d", expectedXP, nextLevel)
	}
}

func TestCleric_CheckXPModifier(t *testing.T) {
	var c Cleric
	a := attributes.Attributes{
		"STR": attributes.Attribute{Name: "STR", Value: 12},
		"INT": attributes.Attribute{Name: "INT", Value: 10},
		"WIS": attributes.Attribute{Name: "WIS", Value: 16},
		"DEX": attributes.Attribute{Name: "DEX", Value: 8},
		"CON": attributes.Attribute{Name: "CON", Value: 14},
		"CHA": attributes.Attribute{Name: "CHA", Value: 10},
	}
	expectedModifier := 10
	if modifier := c.CheckXPModifier(a); modifier != expectedModifier {
		t.Errorf("CheckXPModifier is incorrect. Expected %d, got %d", expectedModifier, modifier)
	}
	setAttributeValue(a, "WIS", 5)
	expectedModifier = -20
	if modifier := c.CheckXPModifier(a); modifier != expectedModifier {
		t.Errorf("CheckXPModifier is incorrect. Expected %d, got %d", expectedModifier, modifier)
	}
}

func setAttributeValue(attr attributes.Attributes, attrKey string, attrValue int) {
	if a, ok := attr[attrKey]; ok {
		a.Value = attrValue
		attr[attrKey] = a
	}
}

func TestTurnUndeadAbilities(t *testing.T) {
	var a TurnUndeadAbilities
	level := 5
	a = TurnUndeadAbilitiesPerLevel[level-1]
	expectedTurnMummy := "9"
	fmt.Printf("%s\n", a)
	if a[5] != expectedTurnMummy {
		t.Errorf("NextLevelAt is incorrect. Expected %s, got %s", expectedTurnMummy, a[5])
	}
}
