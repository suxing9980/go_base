package main

import "fmt"

func GetScore(scores ...float64) (sum, avg float64, count int) {
	for _, value := range scores {
		sum += value
		count++
	}
	avg = sum / float64(count)
	return
}

func main() {
	sum, avg, count := GetScore(90, 82.5, 73, 64.8)
	fmt.Println("sum, avg, count分别是：", sum, avg, count)
	fmt.Println("**************************************")
	scores := []float64{92, 72.5, 93, 74.5, 89, 87, 74}
	sum, avg, count = GetScore(scores...)
	fmt.Println("sum, avg, count分别是：", sum, avg, count)
}
