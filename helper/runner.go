package helper

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func CorrectExercice(exercicePath string, saveIt bool, output bool) bool {
	fmt.Println("🛠 Checking:", exercicePath)

	solutionFile := strings.Replace(exercicePath, "exercices", "solutions", 1)

	// 1 - run user code
	cmd := exec.Command("go", "run", exercicePath)
	var out, errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Runtime Error!")
		fmt.Println(errOut.String())
		return false
	}

	userOutput := strings.TrimSpace(out.String())

	// 2 - run solution code
	cmd = exec.Command("go", "run", solutionFile)
	var sout, serrOut bytes.Buffer
	cmd.Stdout = &sout
	cmd.Stderr = &serrOut

	if serr := cmd.Run(); serr != nil {
		fmt.Println("❌ Solution Error!. Please report it.")
		fmt.Println(serrOut.String())
		return false
	}

	solutionOutput := strings.TrimSpace(sout.String())

	if saveIt && userOutput == solutionOutput {
		fmt.Println("✅ Correct!")
		config := LoadConfig()
		SaveConfig(Config{Next: config.Next + 1, Hint: 0, Path: config.Path})
	}

	if output && userOutput != solutionOutput {
		fmt.Println("❌ Incorrect Output!")
		fmt.Println("🔹 Expected:", solutionOutput)
		fmt.Println("🔻 Got:", userOutput)
	}

	return userOutput == solutionOutput
}
