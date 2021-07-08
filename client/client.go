package main

import (
	"fmt"
	"net/rpc"

	clientCore "github.com/PrabodhaNirmani/vegetable-store/clientCore"
)

func main() {

	// get RPC client by dialing at `rpc.DefaultRPCPath` endpoint
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000")

	fmt.Println("\n---------- Online Vegetable Store ----------")

ExitMainLoop:
	for {

		inputOp, _ := clientCore.UserInput("\n1. Get all vegetables \n2. Get per kg price \n3. Get available amount in kg \n4. Add new vegetable record \n5. Update existing vegetable record \n6. Exit \nPlease enter your choice in single digit : ", clientCore.STRING_TYPE, false)

		switch inputOp {
		case "1":
			// Get all vegetable names
			clientCore.GetAllVegetables(client)
		case "2":
			// Get per kg price of a vegetable
			clientCore.GetPrice(client)
		case "3":
			// Get available amount of a vegetable
			clientCore.GetAmount(client)
		case "4":
			// Add vegetable to the store
			clientCore.AddVegetable(client)
		case "5":
			// Update vegetable in the store
			clientCore.UpdateVegetable(client)
		case "6":
			// Exit
			break ExitMainLoop
		default:
			fmt.Print("\nPlease enter valid option number\n\n")
		}
	}
}
