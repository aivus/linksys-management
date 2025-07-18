package main

import (
	"fmt"
	"log"

	"github.com/reujab/linksys"
)

func main() {
	// Replace with your router's IP, username, and password
	client := linksys.NewClient()
	client.Endpoint = "http://192.168.0.1/JNAP/"
	err := client.Authorize("password") // Replace with your router's password
	if err != nil {
		log.Fatalf("Failed to authorize router: %v", err)
		return
	}

	err = client.MakeRequest("core/Reboot", nil, nil)
	if err != nil {
		log.Fatalf("Failed to reboot router: %v", err)
		return
	}

	fmt.Println("Router reboot initiated successfully.")
}
