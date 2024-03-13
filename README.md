# Saludo go

Este paquete proporciona una forma simple de obtener saludos personalizados en Go.

## Instalación

Ejecute el siguiente comando para instalar el paquete:

````bash
go get -u github.com/camilobayonag/go-greetings
````

## Uso

Aquí tienes un ejemplo de cómo utilizar el paquete en tu código:

````go
package main

import (
	"fmt"
	"github.com/camilobayonag/go-greetings"
	"log"
)

func main() {
	var message, err = greetings.Hello("Camilo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}

````

Este ejemplo importa el paquete github.com/camilobayonag/go-greetings y llama a la función Hello