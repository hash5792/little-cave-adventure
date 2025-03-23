package world

import "fmt"

type Location struct {
	Description string
	Tag         string
}

var locs = []Location{
	{
		Description: "an open field",
		Tag:         "field",
	},
	{
		Description: "a litte cave",
		Tag:         "cave",
	},
}

var locationOfPlayer = 0

func ExecuteLook(noun string) {
	if noun == "around" {
		fmt.Printf("You are in %v.\n", locs[locationOfPlayer].Description)
	} else {
		fmt.Printf("I don't understand what you want to see.\n")
	}
}

func ExecuteGo(noun string) {
	for i := range locs {
		if noun == locs[i].Tag {
			if i == locationOfPlayer {
				fmt.Printf("You can't get much closer than this.\n")
			} else {
				fmt.Printf("OK.\n")

				locationOfPlayer = i
				ExecuteLook("around")
			}
			return
		}
	}
	fmt.Printf("I don't understand where you want to go.\n")
}
