package main

// TODO: define the emplyee struct
type employee struct {
}

func main() {
	employees := map[string]*employee{
		"Alice":   {name: "Alice", taskDone: 12},
		"Bob":     {name: "Bob", taskDone: 4},
		"Charlie": {name: "Charlie", taskDone: 7},
	}

	trackPerformance(employees)
}

func getPerformanceEvaluator() func(int) float64 {
	// Implement the function closure for performance evaluation
	// - If 5 or less tasks were done, then it's a score of 2
	// - If 6 to 10 tasks were done, then it's a score of 3.5
	// - otherwise, it's a score of 5
}

func trackPerformance(employees map[string]*employee) {
	// Implement trackPerformance method (using pointer receiver)
	// It should use the 'getPerformanceEvaluator' to get the score and then print it:
	// using the format: "Employee: NAME, Tasks Completed: TASK, Performance Score: SCORE\n", only print 1 decimal of the score

}
