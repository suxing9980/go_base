package main

import "fmt"

func main() {
	// 总结：
	// 1、 slice[low:high:max]中，len = high - low，cap = max - low
	// 2、 切片后的容量不可以超过切片前容量，如果从0位置开始，那么最多等同前面cap
	// 3、 基于切片的切片，开始位置是基准切片开始往后读取，不可以超过最初的数组界限
	var arr [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("arr:", arr)

	s := arr[1:5:8]
	fmt.Println("s:", s)
	fmt.Println("len(s):", len(s))
	fmt.Println("cap(s):", cap(s))

	fmt.Println("***************************************")

	// 容量不能超过基准的总量减去low
	s1 := s[1:7:7]
	fmt.Println("s1:", s1)
	fmt.Println("len(s1):", len(s1))
	fmt.Println("cap(s1):", cap(s1))

	fmt.Println("***************************************")

	s2 := arr[1:5]
	fmt.Println("s2:", s2)
	fmt.Println("len(s2):", len(s2))
	fmt.Println("cap(s2):", cap(s2))

	fmt.Println("***************************************")

	s3 := s2[1:5]
	fmt.Println("s3:", s3)
	fmt.Println("len(s3):", len(s3))
	fmt.Println("cap(s3):", cap(s3))
}
