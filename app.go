package main

import (
	"fmt"

	"github.com/tomlawson/travelspy/gorc"
)

func main() {
	fmt.Println("Hello World!🌎")
	rc := gorc.New()
	countries := rc.Named("aruba")

	fmt.Println(countries.Capital)
}
