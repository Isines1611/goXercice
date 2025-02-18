package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Next int    `json:"next"`
	Hint int    `json:"hint"`
	Path string `json:"path"`
}

func Initialize(workingPath string) bool {
	// Create directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("❌ Error: Unable to get home directory:", err)
		return false
	}
	goxerciceDir := filepath.Join(homeDir, ".goxercice")

	if err := os.MkdirAll(goxerciceDir, 0755); err != nil {
		fmt.Println("❌ Error creating goxercice directory:", err)
		return false
	}

	// declare files
	exercicePath := filepath.Join(goxerciceDir, "exercices.json")
	hintsPath := filepath.Join(goxerciceDir, "hints.json")
	exercicesContent := map[string][]string{
		"exercices": {
			"00-intro/1-intro.go",
			"00-intro/2-intro.go",
			"01-variable/1-variable.go",
			"01-variable/2-variable.go",
			"01-variable/3-variable.go",
			"01-variable/4-variable.go",
			"02-functions/1-functions.go",
			"02-functions/2-functions.go",
			"02-functions/3-functions.go",
			"03-if/1-if.go",
			"03-if/2-if.go",
			"03-if/3-switch.go",
			"quiz/1-quiz.go",
			"04-for/1-for.go",
			"04-for/2-for.go",
			"04-for/3-for.go",
			"05-defer/1-defer.go",
			"06-pointers/1-pointers.go",
			"06-pointers/2-pointers.go",
			"07-struct/1-struct.go",
			"07-struct/2-struct.go",
			"08-arrays/1-arrays.go",
			"08-arrays/2-arrays.go",
			"08-arrays/3-arrays.go",
			"09-map/1-map.go",
			"09-map/2-map.go",
			"09-map/3-map.go",
			"10-func-var/1-func-var.go",
			"10-func-var/2-func-var.go",
			"quiz/2-quiz.go",
			"11-method/1-method.go",
			"11-method/2-method.go",
			"12-interface/1-interface.go",
			"12-interface/2-interface.go",
			"12-interface/3-interface.go",
			"13-generics/1-generics.go",
			"13-generics/2-generics.go",
			"14-goroutine/1-goroutine.go",
			"15-channels/1-channels.go",
			"15-channels/2-channels.go",
			"15-channels/3-channels.go",
			"16-mutex/1-mutex.go",
			"16-mutex/1-mutex.go",
			"quiz/3-quiz.go",
		},
	}

	hintsContent := map[string][]string{
		"00-intro/1-intro.go": {
			"Hint 1: Nothing to say",
		},
		"00-intro/2-intro.go": {
			"Hint 1: Check out the capital letter",
		},
		"01-variable/1-variable.go": {
			"Hint 1: 'var' is used for classic declaration",
			"Hint2: ':=' is used for short declaration",
		},
		"01-variable/2-variable.go": {
			"Hint 1: Declare all in the same line",
		},
		"01-variable/3-variable.go": {
			"Hint 1: Use 'T(x)' to convert 'x' to type 'T'",
		},
		"01-variable/4-variable.go": {
			"Hint 1: 'const' cannot use short declaration",
		},
		"02-functions/1-functions.go": {
			"Hint 1: 'call_me()' is used to call a function",
		},
		"02-functions/2-functions.go": {
			"Hint 1: Check how to write a function, where to set each information",
			"Hint 2: It should look like: 'func say_hello() {}'",
		},
		"02-functions/3-functions.go": {
			"Hint 1: You can return a computation using 'return'",
		},
		"03-if/1-if.go": {
			"Hint 1: Check the README.md for links",
		},
		"03-if/2-if.go": {
			"Hint 1: Check the README.md for links",
			"Hint 2: You can write '-x' to instead of 'x * -1'",
		},
		"03-if/3-switch.go": {
			"Hint 1: Check the README.md for links",
			"Hint 2: Do not set anything after the switch: 'switch {}'",
		},
		"quiz/1-quiz.go": {
			"Hint 1: Here you can switch over 'fidelity': 'switch fidelity {}'",
			"Hint 2: Each 'case' can check for a number",
		},
		"04-for/1-for.go": {
			"Hint 1: 'for' has 3 parts: initialization, condition, end statement",
			"Hint 2: Check the README.md for links",
		},
		"04-for/2-for.go": {
			"Hint 1: Check the README.md for links",
		},
		"04-for/3-for.go": {
			"Hint 1: Check the README.md for links",
		},
		"05-defer/1-defer.go": {
			"Hint 1: 'defer' are push on the stack",
			"Hint 2: This means it should be done in a back order",
		},
		"06-pointers/1-pointers.go": {
			"Hint 1: '&' gives the memory address of a variable",
		},
		"06-pointers/2-pointers.go": {
			"Hint 1: Use several '*' to get to the value",
		},
		"07-struct/1-struct.go": {
			"Hint 1: Define all variables as int",
			"Hint 2: use brackets '{}' to declare them",
		},
		"07-struct/2-struct.go": {
			"Hint 1: use dot '.' to access fields",
		},
		"08-arrays/1-arrays.go": {
			"Hint 1: Use brackets '{}' to declare an array",
			"Hint 2: This way: 'a := []int{1, 1, 1}'",
			"Hint 3: Use range to loop over an array",
		},
		"08-arrays/2-arrays.go": {
			"Hint 1: Check the README.md for links about slices",
		},
		"08-arrays/3-arrays.go": {
			"Hint 1: Use 'make' to easly initialize the array",
			"Hint 2: Elements in matrix are arrays too",
		},
		"09-map/1-map.go": {
			"Hint 1: Check the README.md for links",
		},
		"09-map/2-map.go": {
			"Hint 1: Use 'delete' to delete from a map",
		},
		"09-map/3-map.go": {
			"Hint 1: Use 'strings.Fields(s)'",
			"Hint 2: You can loop over the returned value using 'range'",
		},
		"10-func-var/1-func-var.go": {
			"Hint 1: Check the README.md for links",
		},
		"10-func-var/2-func-var.go": {
			"Hint 1: Check the README.md for links",
		},
		"quiz/2-quiz.go": {
			"Hint 1: This exercices uses struct, function variables, switch and loops. You can check the README.md of those chapters",
		},
		"11-method/1-method.go": {
			"Hint 1: Check the README.md for links",
		},
		"11-method/2-method.go": {
			"Hint 1: Remember the type conversion",
		},
		"12-interface/1-interface.go": {
			"Hint 1: Check the README.md for links",
		},
		"12-interface/2-interface.go": {
			"Hint 1: Check the README.md for links",
		},
		"12-interface/3-interface.go": {
			"Hint 1: Check the README.md for links",
			"Hint 2: When returning an int and error, use the zero value: 'return 0, DivideZero{}'",
		},
		"13-generics/1-generics.go": {
			"Hint 1: Check the README.md for links",
		},
		"13-generics/2-generics.go": {
			"Hint 1: This exercice is classified as harder, you can check the solution if you're really stuck.",
		},
		"14-goroutine/1-goroutine.go": {
			"Hint 1: Use the keyword 'go' to start a goroutine",
		},
		"15-channels/1-channels.go": {
			"Hint 1: Check the README.md for links",
		},
		"15-channels/2-channels.go": {
			"Hint 1: Check the README.md for links",
		},
		"15-channels/3-channels.go": {
			"Hint 1: Check the README.md for links",
		},
		"16-mutex/1-mutex.go": {
			"Hint 1: Check the README.md for links",
			"Hint 2: Remember to unlock the mutex at the end or use 'defer'",
		},
		"16-mutex/2-mutex.go": {
			"Hint 1: Check the README.md for links",
			"Hint 2: Remember to unlock the mutex at the end or use 'defer'",
		},
		"quiz/3-quiz.go": {
			"Use all you've learn and Gook Luck!",
		},
	}

	// writing exercices.json content
	exercicesFile, err := os.Create(exercicePath)
	if err != nil {
		fmt.Println("❌ Error creating exercices.json:", err)
		return false
	}

	defer exercicesFile.Close()

	encoder := json.NewEncoder(exercicesFile)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(exercicesContent)
	if err != nil {
		fmt.Println("❌ Error writing to exercices.json:", err)
		return false
	}

	// writing hints.json content
	hintsFile, err := os.Create(hintsPath)
	if err != nil {
		fmt.Println("❌ Error creating hints.json:", err)
		return false
	}

	defer hintsFile.Close()

	encoder = json.NewEncoder(hintsFile)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(hintsContent)
	if err != nil {
		fmt.Println("❌ Error writing to hints.json:", err)
		return false
	}

	// writing config
	SaveConfig(Config{Next: 0, Hint: 0, Path: filepath.Join(workingPath, "files")})

	return true
}

// Called when reset to delete all files
func ResetFiles() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("❌ Error: Unable to get home directory:", err)
		return
	}

	goxerciceDir := filepath.Join(homeDir, ".goxercice")

	err = os.RemoveAll(goxerciceDir)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
}

func SaveConfig(config Config) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("❌ Error: Unable to get home directory:", err)
		return err
	}

	configPath := filepath.Join(homeDir, ".goxercice", "config.json")

	file, err := os.Create(configPath)
	if err != nil {
		return err
	}

	defer file.Close()

	return json.NewEncoder(file).Encode(config)
}

func LoadConfig() Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("❌ Error: Unable to get home directory:", err)
		return Config{}
	}

	configPath := filepath.Join(homeDir, ".goxercice", "config.json")

	file, err := os.Open(configPath)
	if err != nil {
		fmt.Println("❌ Error loading config file:", err)
		return Config{}
	}
	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		fmt.Println("❌ Error reading config file:", err)
		return Config{}
	}

	return config
}

// Get the next hint of the pending exercice, if no more hint, a message is displayed
func GetHint() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("❌ Error: Unable to get home directory:", err)
		return "", err
	}

	hintsPath := filepath.Join(homeDir, ".goxercice", "hints.json")
	config := LoadConfig()

	hintsFile, err := os.Open(hintsPath)
	if err != nil {
		return "", err
	}
	defer hintsFile.Close()

	var hintsData map[string][]string
	decoder := json.NewDecoder(hintsFile)
	err = decoder.Decode(&hintsData)
	if err != nil {
		fmt.Println("❌ Error: Unable read hints data:", err)
		return "", err
	}

	var fileName string
	var counter int
	for key := range hintsData {
		if counter == config.Next {
			fileName = key
			break
		}
		counter++
	}

	// Check if the file exists and retrieve the hint at counter
	if fileHints, exists := hintsData[fileName]; exists {
		if config.Hint >= 0 && config.Hint < len(fileHints) {
			hints := fileHints[:config.Hint+1]

			var text string
			for i, v := range hints {
				text += fmt.Sprintf("%d/%d - %s\n", i+1, len(fileHints), v)
			}

			config.Hint++
			SaveConfig(config)

			return text, nil
		} else {
			var text string
			for i, v := range fileHints {
				text += fmt.Sprintf("%d/%d - %s\n", i+1, len(fileHints), v)
			}

			return text, nil
		}
	}
	return "", fmt.Errorf("file not found")

}

// GetNextExercise gets the next exercise from exercices.json
func GetNextExercise() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("❌ Error: Unable to get home directory:", err)
		return "", err
	}

	exercicePath := filepath.Join(homeDir, ".goxercice", "exercices.json")
	config := LoadConfig()

	exerciceFile, err := os.Open(exercicePath)
	if err != nil {
		return "", err
	}
	defer exerciceFile.Close()

	var data struct {
		Exercices []string `json:"exercices"`
	}

	decoder := json.NewDecoder(exerciceFile)
	if err := decoder.Decode(&data); err != nil {
		return "err", err
	}

	if config.Next < 0 || config.Next >= len(data.Exercices) {
		return "index out of range", fmt.Errorf("index out of range")
	}

	return data.Exercices[config.Next], nil
}
