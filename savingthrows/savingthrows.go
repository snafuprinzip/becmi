package savingthrows

import "fmt"

type SavingThrows [5]int
type ByLevel [36]SavingThrows

func (s SavingThrows) String() string {
	return fmt.Sprintf(""+
		"Death Ray, Poison         %2d\n"+
		"Magic Wands               %2d\n"+
		"Paralysis, Turn to Stone  %2d\n"+
		"Dragon Breath             %2d\n"+
		"Rod, Staff, Spells        %2d\n",
		s[0], s[1], s[2], s[3], s[4],
	)
}
