package duration

import (
	"fmt"
	"time"
)

func ExampleDuration() {

	d := Duration{Duration: time.Duration(12*year + 365*day + 5*day + 5*time.Hour)}
	fmt.Println(d.Format())
	fmt.Println(fmt.Sprintf("%s", time.Duration(1503*time.Minute)))
	d = Duration{Duration: time.Duration(1503 * time.Minute)}
	fmt.Println(d.Format())

	// Output:
	// 13y5d5h
	// 25h3m0s
	// 1d1h3m
}
