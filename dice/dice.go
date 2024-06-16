package dice

import (
	"becmi/attributes"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

const debug = false

func RollDice(dice int) int {
	x := rand.Intn(dice) + 1
	if debug {
		fmt.Printf("1w%d = %d\n", dice, x)
	}
	return x
}

func Roll(roll string) int {
	var numDice, diceType, modifier int
	var err error

	if strings.Contains(roll, "w") {
		roll = strings.ReplaceAll(roll, "w", "d")
	}

	if strings.Contains(roll, "d") {
		splitstr := strings.Split(roll, "d")
		if len(splitstr[0]) > 0 {
			numDice, err = strconv.Atoi(splitstr[0])
			if err != nil {
				log.Printf("Error converting dice number to integer %s", splitstr[0])
				return 0
			}
		}

		if strings.Contains(splitstr[1], "+") {
			modstr := strings.Split(splitstr[1], "+")
			if len(modstr[0]) > 0 {
				diceType, err = strconv.Atoi(modstr[0])
				if err != nil {
					log.Printf("Error converting dice type to integer %s", modstr[0])
					return 0
				}
			}
			if len(modstr[1]) > 0 {
				modifier, err = strconv.Atoi(modstr[1])
				if err != nil {
					log.Printf("Error converting modifier to integer %s", modstr[0])
					return 0
				}
			}
		} else if strings.Contains(splitstr[1], "-") {
			modstr := strings.Split(splitstr[1], "-")
			if len(modstr[0]) > 0 {
				diceType, err = strconv.Atoi(modstr[0])
				if err != nil {
					log.Printf("Error converting dice type to integer %s", modstr[0])
					return 0
				}
			}
			if len(modstr[1]) > 0 {
				modifier, err := strconv.Atoi(modstr[1])
				if err != nil {
					log.Printf("Error converting negative modifier to integer %s", modstr[0])
					return 0
				}
				modifier = -modifier
			}

		} else {
			if len(splitstr[1]) > 0 {
				diceType, err = strconv.Atoi(splitstr[1])
				if err != nil {
					log.Printf("Error converting dice type to integer %s", splitstr[1])
					return 0
				}
			}
			modifier = 0
		}
	}

	result := 0
	for i := 0; i < numDice; i++ {
		result += RollDice(diceType)
	}

	if debug {
		fmt.Printf("%s (%d d %d + %d) == %d\n", roll, numDice, diceType, modifier, result+modifier)
	}

	return result + modifier
}

func RollAttributes() attributes.Attributes {
	var attr attributes.Attributes = make(attributes.Attributes)

	for _, str := range attributes.AttributeIndices {
		switch str {
		case "STR":
			attr[str] = attributes.Attribute{"Strength", Roll("3d6")}
		case "INT":
			attr[str] = attributes.Attribute{"Intelligence", Roll("3d6")}
		case "WIS":
			attr[str] = attributes.Attribute{"Wisdom", Roll("3d6")}
		case "DEX":
			attr[str] = attributes.Attribute{"Dexterity", Roll("3d6")}
		case "CON":
			attr[str] = attributes.Attribute{"Constitution", Roll("3d6")}
		case "CHA":
			attr[str] = attributes.Attribute{"Charisma", Roll("3d6")}
		}
	}

	if debug {
		fmt.Printf("%s\n", attr)
	}

	return attr
}
