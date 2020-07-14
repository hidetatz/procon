package main

import "fmt"

func main() {
	fmt.Println(angleClock(1, 57))
	fmt.Println(angleClock(8, 7))
	fmt.Println(angleClock(3, 30))
	fmt.Println(angleClock(3, 15))
	fmt.Println(angleClock(4, 50))
	fmt.Println(angleClock(12, 0))
}

func angleClock(hour int, minutes int) float64 {
	if hour == 12 {
		hour = 0
	}

	minuteAngle := float64(6 * minutes)
	hourAngle := float64(30*hour) + 0.5*float64(minutes)

	ret := minuteAngle - hourAngle
	if ret < 0 {
		ret = -ret
	}

	if ret > 180 {
		return 360 - ret
	}

	return ret
}
