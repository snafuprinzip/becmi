package becmi

import (
	"fmt"
	"testing"
)

func TestCharacter_String(t *testing.T) {
	char := NewCharacter("Test", "DM", "Neutral", "male", "Cleric", 15000)
	fmt.Printf("%s\n", char)
}
