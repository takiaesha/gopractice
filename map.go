package main

import "fmt"

func mp() {
	var b map[string]int //declaring map
	fmt.Println(b)       // map[]

	//need to initialize before use, make_function use
	animeList := make(map[int]string)

	animeList[1] = "DeathNote"
	animeList[2] = "Arrietty"
	animeList[3] = "When Marnie was there"
	animeList[4] = "Grave of the Fireflies"

	fmt.Println(animeList[3])

	//shortest way
	toy := map[int]string{
		1: "Batman",
		2: "thor",
		3: "Antman",
		4: "Shinigami",
	}

	fmt.Println(toy[4])

	if el, ok := toy[2]; ok {
		fmt.Println(el)
	}

}
