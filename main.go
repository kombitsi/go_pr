package main

import (
"go_pr/actions"
)

func main() {
	habrTags := []string{"go", "python", "kubernetes"}
	actions.Gophers()
	for _, i := range habrTags {
		actions.HabrGo(i)
	}
}

