package main

import (
	"log"

	sdk "github.com/scottwinkler/go-petstore"
)

func main() {
	config := &sdk.Config{
		//Address: "insert-your-petstore-address-here",
		Address: "https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/petstore",
	}

	client, err := sdk.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new organization
	options := sdk.PetCreateOptions{
		Name:    "mittens",
		Species: "cat",
		Age:     2,
	}

	pet, err := client.Pets.Create(options)
	if err != nil {
		log.Fatal(err)
	}

	// Delete an organization
	err = client.Pets.Delete(pet.ID)
	if err != nil {
		log.Fatal(err)
	}
}
