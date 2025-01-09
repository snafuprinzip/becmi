package classes

import "becmi/dice"

func Age(race string) (age int, agespan string) {
	switch race {
	case "Human":
		return humanAge()
	case "Gnome":
		fallthrough
	case "Dwarf":
		return dwarfAge()
	case "Elf":
		return elfAge()
	case "Halfling":
		return halflingAge()
	}
	return 0, ""
}

func Body(race, sex string) (height string, weight int) {
	switch race {
	case "Human":
		return humanBody(sex)
	case "Gnome":
		fallthrough
	case "Dwarf":
		return dwarfBody(sex)
	case "Elf":
		return elfBody(sex)
	case "Halfling":
		return halflingBody()
	}
	return "", 0
}

func humanAge() (age int, agespan string) {
	age = dice.Roll("1d6+13")
	agespan = "14-100"
	return age, agespan
}

func dwarfAge() (age int, agespan string) {
	age = dice.Roll("1d20+19")
	agespan = "20-275"
	return age, agespan
}

func elfAge() (age int, agespan string) {

	age = dice.Roll("1d100")
	if age < 25 {
		age = 25
	}
	agespan = "25-800"
	return age, agespan
}

func halflingAge() (age int, agespan string) {
	age = dice.Roll("1d10+19")
	agespan = "20-200"
	return age, agespan
}

func humanBody(sex string) (height string, weight int) {
	roll := dice.RollDice(10)

	switch roll {
	case 1:
		height = "4'10\" / 1.47m"
		if sex == "male" {
			weight = 1100
		} else {
			weight = 1050
		}
	case 2:
		height = "5'0\" / 1.52m"
		if sex == "male" {
			weight = 1200
		} else {
			weight = 1100
		}
	case 3:
		height = "5'2\" / 1.57m"
		if sex == "male" {
			weight = 1300
		} else {
			weight = 1200
		}
	case 4:
		height = "5'4\" / 1.62m"
		if sex == "male" {
			weight = 1400
		} else {
			weight = 1250
		}
	case 5:
		height = "5'6\" / 1.67m"
		if sex == "male" {
			weight = 1500
		} else {
			weight = 1300
		}
	case 6:
		height = "5'8\" / 1.72m"
		if sex == "male" {
			weight = 1550
		} else {
			weight = 1400
		}
	case 7:
		height = "5'10\" / 1.77m"
		if sex == "male" {
			weight = 1650
		} else {
			weight = 1500
		}
	case 8:
		height = "6'0\" / 1.82m"
		if sex == "male" {
			weight = 1750
		} else {
			weight = 1550
		}
	case 9:
		height = "6'2\" / 1.87m"
		if sex == "male" {
			weight = 1850
		} else {
			weight = 1650
		}
	case 10:
		height = "6'4\" / 1.92m"
		if sex == "male" {
			weight = 2000
		} else {
			weight = 1750
		}
	}
	return height, weight
}

func dwarfBody(sex string) (height string, weight int) {
	roll := dice.RollDice(5)

	switch roll {
	case 1:
		height = "3'8\" / 1.12m"
		if sex == "male" {
			weight = 1300
		} else {
			weight = 1250
		}
	case 2:
		height = "3'10\" / 1.17m"
		if sex == "male" {
			weight = 1400
		} else {
			weight = 1350
		}
	case 3:
		height = "4'0\" / 1.22m"
		if sex == "male" {
			weight = 1500
		} else {
			weight = 1450
		}
	case 4:
		height = "4'2\" / 1.27m"
		if sex == "male" {
			weight = 1550
		} else {
			weight = 1500
		}
	case 5:
		height = "4'4\" / 1.32m"
		if sex == "male" {
			weight = 1650
		} else {
			weight = 1600
		}
	}
	return height, weight
}

func elfBody(sex string) (height string, weight int) {
	roll := dice.RollDice(6)

	switch roll {
	case 1:
		height = "4'8\" / 1.42m"
		if sex == "male" {
			weight = 900
		} else {
			weight = 750
		}
	case 2:
		height = "5'0\" / 1.52m"
		if sex == "male" {
			weight = 1000
		} else {
			weight = 800
		}
	case 3:
		height = "5'2\" / 1.57m"
		if sex == "male" {
			weight = 1100
		} else {
			weight = 900
		}
	case 4:
		height = "5'4\" / 1.62m"
		if sex == "male" {
			weight = 1200
		} else {
			weight = 1000
		}
	case 5:
		height = "5'6\" / 1.67m"
		if sex == "male" {
			weight = 1300
		} else {
			weight = 1100
		}
	case 6:
		height = "5'8\" / 1.72m"
		if sex == "male" {
			weight = 1400
		} else {
			weight = 1200
		}
	}
	return height, weight
}

func halflingBody() (height string, weight int) {
	roll := dice.RollDice(3)

	switch roll {
	case 1:
		height = "2'10\" / 0.87m"
		weight = 580
	case 2:
		height = "3'0\" / 0.92m"
		weight = 600
	case 3:
		height = "3'2\" / 0.97m"
		weight = 620
	}
	return height, weight
}
