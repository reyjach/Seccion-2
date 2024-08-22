package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {

	jugar()

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

	saludo := hello("Teto")
	fmt.Println(saludo)
	//ciclo()
	sum, mul := calc(4, 5)

	fmt.Println("la suma es: ", sum)
	fmt.Println("la multiplicacion es: ", mul)

	aleatorio()
}

func hello(name string) string {
	return "hola desde la funcion " + name
}

func ciclo() {

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

func calc(a, b int) (int, int) {
	sum := a + b
	mul := a * b
	return sum, mul
}

func calc2(a, b int) (sum, mul int) { // otra forma de retornar funciones
	sum = a + b
	mul = a * b
	return
}

func aleatorio() {

	fmt.Println(rand.Intn(100))
}

func jugar() {
	numAleatorio := rand.Intn(100)

	var numIngresado int

	var intentos int

	const maxIntentos = 10

	for intentos < maxIntentos {

		fmt.Printf("Ingresa un numero (intentos restantes: %d): ", maxIntentos-intentos)
		fmt.Scanln(&numIngresado)

		if numIngresado == numAleatorio {
			fmt.Println("!Felicitaciones, adivinaste el numero")
			jugarNuevamente()
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("El numero a adivinar es mayor.")
		} else {
			fmt.Println("El numero a adivinar es menor.")
		}

		intentos++
	}

	fmt.Println("Se acabaron los intentos. El numero era: ", numAleatorio)
	jugarNuevamente()
}
func jugarNuevamente() {
	var eleccion string
	fmt.Println("¿Quieres jugar nuevamente? (s/n)")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("Gracias por jugar!")
	default:
		fmt.Println("Ellecion invalida. Intenta nuevamente. ")
		jugarNuevamente()
	}
}
