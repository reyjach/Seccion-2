package main

import (
	"fmt"
	"time"
)

func main() {

	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("Mañana")

	} else if t.Hour() < 17 {
		fmt.Println("Tarde")
	} else {
		fmt.Println("Noche")
	}
}
