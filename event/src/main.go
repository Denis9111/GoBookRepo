// инкапсуляция в go на примере структуры Event
package main

import (
	"fmt"
	"log"
	"test/pkg/calindar"
)

func main() {
	var event calindar.Event

	if err := event.SetYear(2024); err != nil {
		log.Fatal(err)
	}
	if err := event.SetMonth(11); err != nil {
		log.Fatal(err)
	}
	if err := event.SetDay(6); err != nil {
		log.Fatal(err)
	}
	if err := event.SetTitle("Day of crash"); err != nil {
		log.Fatal(err)
	}

	fmt.Println(event.GetYear())
	fmt.Println(event.GetMonth())
	fmt.Println(event.GetDay())
	fmt.Println(event)

}
