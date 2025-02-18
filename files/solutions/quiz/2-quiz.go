package main

import "fmt"

type employee struct {
	name     string
	taskDone int
	score    float64
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
	return func(tasks int) float64 {
		switch {
		case tasks <= 5:
			return 2.0
		case tasks <= 10:
			return 3.5
		default:
			return 5.0
		}
	}
}

func trackPerformance(employees map[string]*employee) {
	evaluator := getPerformanceEvaluator()

	for _, emp := range employees {
		emp.score = evaluator(emp.taskDone)
		fmt.Printf("Employee: %s, Tasks Completed: %d, Performance Score: %.1f\n", emp.name, emp.taskDone, emp.score)
	}
}
