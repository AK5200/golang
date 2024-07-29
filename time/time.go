package main

import (
	"fmt"
	"time"
)

func main() {
	// 02/01/2006 15:04:05  || 3:04 PM
	// 2 jan 2006
	// same numbers need to be typed for month, day, year etc

	currentTime := time.Now()
	formatted := currentTime.Format("02/01/2006") // day/month/year  // prints current data and time not 02/01/2006
	formatted2 := currentTime.Format("2 Jan 2006, 3:04 PM")
	formatted3 := currentTime.Format(" Jan / 2 / 2006")
	fmt.Println(formatted, formatted2, formatted3)

	// if not need to current time
	go_str := "2006-01-02"
	data_str := "2002-10-22"
	formatted1, _ := time.Parse(go_str, data_str)
	fmt.Println(formatted1) // 22/10/2002

}
