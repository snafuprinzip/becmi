package background

type Background interface {
	Nation() string
	Nationality() string
	//Ethnicity() string
	//SocialStatus() (string, bool)
	//Hometown() string
	String() string
}

var BackgroundIndices []string
var Backgrounds map[string]Background

func AvailableBackgrounds() string {
	backgroundsstr := "["

	for idx, b := range BackgroundIndices {
		backgroundsstr += b
		if idx != len(BackgroundIndices)-1 {
			backgroundsstr += ", "
		} else {
			backgroundsstr += "]"
		}
	}
	return backgroundsstr
}
