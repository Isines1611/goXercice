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
		},
	}

	hintsContent := map[string][]string{
		"001-variables.go": {
			"Hint 1: Variables in Go use 'var' or ':='.",
			"Hint 2: The type can be inferred automatically.",
		},
		"002-functions.go": {
			"Hint 1: Functions in Go start with 'func'.",
			"Hint 2: Functions need a return type if they return a value.",
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
