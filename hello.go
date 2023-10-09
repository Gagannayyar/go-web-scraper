package main

import "log"



func main() {

	names := []string{"One", "Two", "Three"}

	for i,j := range names {
		log.Println(i, j)
	}
	
}