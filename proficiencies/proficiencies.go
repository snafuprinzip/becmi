package proficiencies

type Proficiency struct {
	Name      string
	Attribute string
	Value     int
}

type Proficiencies struct {
	Proficiencies []Proficiency
	FreeSlots     int
}
