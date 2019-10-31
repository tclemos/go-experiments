package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		y := strconv.Itoa(i)
		f := ""
		if i < 10 {
			f = "0"
		}
		toDate2(f + y + "0501")
	}
}

func toDate2(dt string) {
	test, err := time.Parse("060102", dt)
	if err != nil {
		panic(err)
	}

	fmt.Println(test)
}

func toDate(dt string) {
	currentYear := time.Now().Year()

	y, _ := strconv.Atoi(dt[0:2])
	m, _ := strconv.Atoi(dt[2:4])
	d, _ := strconv.Atoi(dt[4:6])

	y += 2000

	// if y >= (currentYear-20) && y <= (currentYear+20) {

	// } else {
	// 	y -= 100
	// }

	if y < (currentYear-20) || y > (currentYear+20) {
		y -= 100
	}

	l, _ := time.LoadLocation("")

	td := time.Date(y, time.Month(m), d, 0, 0, 0, 0, l)
	fmt.Println(td.Format("2006-01-02"))
}
