package main

import (
	"fmt"

	"github.com/stianeikeland/go-rpio/v4"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Printf("hello\n")
	err := rpio.Open()
	check(err)

	pins := []rpio.Pin{rpio.Pin(14), rpio.Pin(15), rpio.Pin(18), rpio.Pin(17), rpio.Pin(27), rpio.Pin(22)}

	for _, pin := range pins {
		pin.Input()
		pin.PullDown()
	}
}
