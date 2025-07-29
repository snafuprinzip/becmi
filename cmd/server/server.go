package main

import (
	"becmi"
	"becmi/background"
	"becmi/classes"
	"becmi/localization"
	"becmi/magic"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

// CharacterForm represents the form data for character creation
type CharacterForm struct {
	Name      string
	Player    string
	Class     string
	Alignment string
	Sex       string
	BG        string
	XP        int
	Language  string
	Character *becmi.Character
}

// Form template HTML wird aktualisiert - nur der relevante Teil wird gezeigt
const formTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>BECMI Character Generator</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 20px;
            display: flex;
            flex-direction: column;
        }
        h1 {
            width: 100%;
            margin-bottom: 20px;
        }
        .container {
            display: flex;
            gap: 20px;
            width: 100%;
        }
        .form-container {
            flex: 0 0 200px;
            background-color: #f5f5f5;
            padding: 20px;
            border-radius: 5px;
        }
        .character-container {
            flex: 1;
            background-color: #f9f9f9;
            padding: 20px;
            border-radius: 5px;
        }
        label {
            display: block;
            margin-bottom: 5px;
            font-weight: bold;
        }
        input, select {
            width: 100%;
            padding: 8px;
            margin-bottom: 15px;
            border: 1px solid #ddd;
            border-radius: 4px;
            box-sizing: border-box;
        }
        button {
            background-color: #4CAF50;
            color: white;
            padding: 10px 15px;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            width: 100%;
        }
        button:hover {
            background-color: #45a049;
        }
        pre {
            background-color: #ffffff;
            padding: 15px;
            border: 1px solid #ddd;
            border-radius: 4px;
            white-space: pre-wrap;
            font-family: monospace;
            margin: 0;
        }
    </style>
</head>
<body>
    <h1>BECMI Character Generator</h1>
    <div class="container">
        <div class="form-container">
            <form method="post">
                <div>
                    <label for="name">Character Name:</label>
                    <input type="text" id="name" name="name" value="{{.Name}}">
                </div>
                <div>
                    <label for="player">Player Name:</label>
                    <input type="text" id="player" name="player" value="{{.Player}}">
                </div>
                <div>
                    <label for="class">Character Class:</label>
                    <select id="class" name="class">
                        {{range .Classes}}
                        <option value="{{.}}" {{if eq . $.SelectedClass}}selected{{end}}>{{.}}</option>
                        {{end}}
                    </select>
                </div>
                <div>
                    <label for="alignment">Alignment:</label>
                    <select id="alignment" name="alignment">
                        <option value="Lawful" {{if eq .Alignment "Lawful"}}selected{{end}}>Lawful</option>
                        <option value="Neutral" {{if eq .Alignment "Neutral"}}selected{{end}}>Neutral</option>
                        <option value="Chaotic" {{if eq .Alignment "Chaotic"}}selected{{end}}>Chaotic</option>
                    </select>
                </div>
                <div>
                    <label for="sex">Sex:</label>
                    <select id="sex" name="sex">
                        <option value="male" {{if eq .Sex "male"}}selected{{end}}>Male</option>
                        <option value="female" {{if eq .Sex "female"}}selected{{end}}>Female</option>
                        <option value="other" {{if eq .Sex "other"}}selected{{end}}>Other</option>
                    </select>
                </div>
                <div>
                    <label for="bg">Campaign Background:</label>
                    <select id="bg" name="bg">
                        {{range .Backgrounds}}
                        <option value="{{.}}" {{if eq . $.SelectedBG}}selected{{end}}>{{.}}</option>
                        {{end}}
                    </select>
                </div>
                <div>
                    <label for="xp">Experience Points:</label>
                    <input type="number" id="xp" name="xp" value="{{.XP}}">
                </div>
                <div>
                    <label for="language">Language:</label>
                    <select id="language" name="language">
                        <option value="en" {{if eq .Language "en"}}selected{{end}}>English</option>
                        <option value="de" {{if eq .Language "de"}}selected{{end}}>German</option>
                    </select>
                </div>
                <button type="submit">Generate Character</button>
            </form>
        </div>
        <div class="character-container">
            {{if .Character}}
            <pre>{{.CharacterSheet}}</pre>
            {{end}}
        </div>
    </div>
</body>
</html>
`

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize magic spells
	magic.LoadSpells()

	// Get available classes and backgrounds
	availableClasses := classes.ClassIndices
	availableBackgrounds := background.BackgroundIndices

	// Set default language
	language := "en"
	if lang, ok := os.LookupEnv("LANG"); ok {
		language = lang[:2]
	}

	// Create template data
	data := struct {
		Name           string
		Player         string
		Classes        []string
		SelectedClass  string
		Alignment      string
		Sex            string
		Backgrounds    []string
		SelectedBG     string
		XP             int
		Language       string
		Character      *becmi.Character
		CharacterSheet string
	}{
		Name:          "Bargle",
		Player:        "NPC",
		Classes:       availableClasses,
		SelectedClass: "Cleric",
		Alignment:     "Lawful",
		Sex:           "male",
		Backgrounds:   availableBackgrounds,
		SelectedBG:    "Karameikos",
		XP:            0,
		Language:      language,
	}

	// Handle form submission
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		// Get form values
		name := r.FormValue("name")
		player := r.FormValue("player")
		class := r.FormValue("class")
		alignment := r.FormValue("alignment")
		sex := r.FormValue("sex")
		bg := r.FormValue("bg")
		xpStr := r.FormValue("xp")
		language := r.FormValue("language")

		// Parse XP
		xp := 0
		if xpStr != "" {
			var err error
			xp, err = strconv.Atoi(xpStr)
			if err != nil {
				http.Error(w, "Invalid XP value", http.StatusBadRequest)
				return
			}
		}

		// Set language
		localization.LanguageSetting = language

		// Create character
		char := becmi.NewCharacter(name, player, alignment, sex, class, bg, xp)

		// Update template data
		data.Name = name
		data.Player = player
		data.SelectedClass = class
		data.Alignment = alignment
		data.Sex = sex
		data.SelectedBG = bg
		data.XP = xp
		data.Language = language
		data.Character = char
		data.CharacterSheet = char.String()
	}

	fmt.Printf("%+v\n", data)

	// Parse and execute template
	tmpl, err := template.New("form").Parse(formTemplate)
	if err != nil {
		http.Error(w, "Error creating template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/", IndexHandler)

	// Start server
	port := "8080"
	fmt.Printf("Server starting on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
