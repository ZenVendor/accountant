package main

import "log"

func main() {
	_, _, err := prepare()
	if err != nil {
		log.Fatal(err)
	}

}
