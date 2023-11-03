package main

import "log"

func main() {

	var line string = "Once"

	for i, l := range line {
		log.Println(i, l)
	}
}
