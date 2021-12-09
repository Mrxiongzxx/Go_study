package main

import "fmt"

var days = [13]int{0,31,29,31,30,31,30,31,31,30,31,30,31}

func check(x int) bool {
	year  := x / 10000
	month := x / 1000000 % 10
	day   := x / 100000000 % 10
	fmt.Println(year,month,day)
	return true
}

func main()  {
	startDate := 10000101
	endDate   := 92211229

	for i:=startDate;i<=endDate;i++{
		if check(i) {
			fmt.Println(i)
		}
	}
}