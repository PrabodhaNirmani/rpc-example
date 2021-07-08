package clientCore

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strconv"
	"strings"

	core "github.com/PrabodhaNirmani/vegetable-store/serverCore"
)

const STRING_TYPE = "string"
const FLOAT_TYPE = "float"

// UserInput | function to read input from the command line and validate it
func UserInput(message string, inputType string, acceptEmpty bool) (string, float64) {
	for {
		var inputString string
		fmt.Print(message)
		reader := bufio.NewReader(os.Stdin)
		inputString, err := reader.ReadString('\n')

		inputString = strings.TrimSuffix(inputString, "\n")

		if err != nil || (inputString == "" && !acceptEmpty) {
			fmt.Println("\nAn error occurred while reading input.")
			continue
		}

		// return `inputType` if the input accepts string values from the user
		if inputType == STRING_TYPE {
			return inputString, 0
		}

		// if the `inputString` is empty for inputs
		// which accepts empty values from the user
		// return `0.0`
		if inputString == "" && acceptEmpty {
			return "", 0.0
		} else {
			var value float64
			// convert string to `float64`
			value, err = strconv.ParseFloat(inputString, 64)
			if err != nil {
				fmt.Println("\nInput validation failed. ", err)
			} else {
				return "", value
			}
		}

	}
}

// GetAllVegetables | function to get all vegetables from the server
func GetAllVegetables(client *rpc.Client) {
	var vegetablesList []string

	fmt.Println("\n----------------Get all vegetables----------------")
	if err := client.Call("VegetableStore.GetAll", "", &vegetablesList); err != nil {
		fmt.Println("\nError :: VegetableStore.GetAll() |", err)
	} else {
		fmt.Println("\nSuccess :: Vegetables list | ", vegetablesList)
	}

}

// GetPrice | function to get per kg price of a vegetable from the server
func GetPrice(client *rpc.Client) {
	var price float64
	fmt.Println("\n----------------Get per kg price----------------")

	inputVeg, _ := UserInput("\nEnter vegetable name : ", STRING_TYPE, false)

	if err := client.Call("VegetableStore.Price", inputVeg, &price); err != nil {
		fmt.Println("\nError :: VegetableStore.Price() |", err)
	} else {
		fmt.Printf("\nSuccess :: Price of 1kg of '%s' | %v \n", inputVeg, price)
	}
}

// GetAmount | function to get available amount in kg of a vegetable from the server
func GetAmount(client *rpc.Client) {
	var amount float64
	fmt.Println("\n----------------Get available amount in kg----------------")

	inputVeg, _ := UserInput("\nEnter vegetable name : ", STRING_TYPE, false)

	if err := client.Call("VegetableStore.Amount", inputVeg, &amount); err != nil {
		fmt.Println("\nError :: VegetableStore.Amount() |", err)
	} else {
		fmt.Printf("\nSuccess :: Available amount of '%s' | %v kg\n", inputVeg, amount)
	}
}

// AddVegetable | function to add new vegetable record
func AddVegetable(client *rpc.Client) {
	fmt.Println("\n----------------Add new vegetable record----------------")

	inputVeg, _ := UserInput("\nEnter vegetable name : ", STRING_TYPE, false)
	_, inputPrice := UserInput("Enter per kg price : ", FLOAT_TYPE, false)
	_, inputAmount := UserInput("Enter amount in kg : ", FLOAT_TYPE, false)

	var vegetable core.Vegetable

	if err := client.Call("VegetableStore.Add", core.Vegetable{
		VegetableTag: inputVeg,
		Price:        inputPrice,
		Amount:       inputAmount,
	}, &vegetable); err != nil {
		fmt.Println("\nError :: VegetableStore.Add() |", err)
	} else {
		fmt.Printf("\nSuccess :: Vegetable '%s' added\n", inputVeg)
	}

}

// UpdateVegetable | function to update available vegetable record
func UpdateVegetable(client *rpc.Client) {
	fmt.Println("\n----------------Update existing vegetable record----------------")

	inputVeg, _ := UserInput("\nEnter vegetable name : ", STRING_TYPE, false)
	_, inputPrice := UserInput("Enter per kg price : ", FLOAT_TYPE, true)
	_, inputAmount := UserInput("Enter amount in kg : ", FLOAT_TYPE, true)

	var vegetable core.Vegetable

	if inputPrice == 0.0 && inputAmount == 0.0 {
		fmt.Println("\nPrice and amount both couldn't be empty")
	} else {
		if err := client.Call("VegetableStore.Update", core.Vegetable{
			VegetableTag: inputVeg,
			Price:        inputPrice,
			Amount:       inputAmount,
		}, &vegetable); err != nil {
			fmt.Println("\nError :: VegetableStore.Update() |", err)
		} else {
			fmt.Printf("\nSuccess :: Vegetable '%s' updated\n", inputVeg)
		}
	}

}
