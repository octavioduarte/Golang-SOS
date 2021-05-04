package main

import "fmt"

func main() {
	handleMessageWithCustomParam := handleCustomMessage("Janis Joplin")

	fmt.Println(handleMessageWithCustomParam)
}

func handleCustomMessage(singer string) string {
	return "The best singer of world is " + singer
}
