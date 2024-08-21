package main

import (
	"fmt"
	"runtime"
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

	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("Mañana")
	case t.Hour() < 17:
		fmt.Println("Tarde")
	default:
		fmt.Println("Noche")
	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Go run -> Windows")
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> macOS")
	default:
		fmt.Println("Go run -> otro OS ")
	}

	for i := 1; i <= 10; i++ {

		if i == 3 {
			continue // no ejcuta lo que esta por debajo de la linea y sigue en el bucle
		}
		fmt.Println(i)
		if i == 5 {
			break // se sale de bucle
		}

	}
}
