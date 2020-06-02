package main

import "fmt"

func main() {
	grade := ""
	score := 78.5

	switch {
	case score >= 70:
		grade = "优秀"
	case score >= 60:
		grade = "良好"
	case score >= 50:
		grade = "不及格"

	}
	fmt.Println(grade)
}
