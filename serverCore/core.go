package serverCore

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Vegetable | struct represents a vegetable
type Vegetable struct {
	VegetableTag  string
	Price, Amount float64
}

// VegetableStore | struct represents a vegetable store
type VegetableStore struct {
	data map[string]Vegetable
}

// WriteToFile | function to write data to the data.txt file
func WriteToFile(fileName string, fileOpenType int, data string) (int, error) {
	f, err := os.OpenFile("data.txt", fileOpenType, 0644)
	if err != nil {
		return -1, err
	}
	bytes, err := f.WriteString(data)
	if err != nil {
		return bytes, err
	}
	return bytes, nil
}

// Add | function to add new vegetable to the store
func (store *VegetableStore) Add(payload Vegetable, reply *Vegetable) error {

	tag := strings.ToLower(payload.VegetableTag)
	// check whether the given vegetable already exists in the store
	if _, ok := store.data[tag]; ok {
		return fmt.Errorf("Vegetable '%s' already exists", payload.VegetableTag)
	}
	// add vegetable to the `store`
	store.data[tag] = payload

	NewRecord := "\n" + payload.VegetableTag + ":" + fmt.Sprintf("%v", payload.Price) + ":" + fmt.Sprintf("%v", payload.Amount)
	// append new record to the file
	_, err := WriteToFile("data.txt", os.O_APPEND|os.O_WRONLY, NewRecord)
	if err != nil {
		return err
	}
	*reply = payload
	return nil

}

// Update | function to update price or amount of available vegetable
func (store *VegetableStore) Update(payload Vegetable, reply *Vegetable) error {

	tag := strings.ToLower(payload.VegetableTag)

	// check whether the given vegetable not exist in the store
	if _, ok := store.data[tag]; !ok {
		return fmt.Errorf("Vegetable '%s' does not exist", payload.VegetableTag)
	}

	payload.VegetableTag = store.data[tag].VegetableTag

	// set the `Amount` to existing amount if the user input is empty
	if payload.Amount == 0 {
		payload.Amount = store.data[tag].Amount
	}

	// set the `Price` to existing price if the user input is empty
	if payload.Price == 0 {
		payload.Price = store.data[tag].Price
	}

	// update the `store`
	store.data[tag] = payload

	// prepare the `Data` string to be write in the file the store
	Data := ""
	for _, vegetable := range store.data {

		record := "\n" + vegetable.VegetableTag + ":" + fmt.Sprintf("%v", vegetable.Price) + ":" + fmt.Sprintf("%v", vegetable.Amount)
		Data = Data + record
	}
	// write data to the file
	_, err := WriteToFile("data.txt", os.O_WRONLY, Data[1:])
	if err != nil {
		return err
	}
	*reply = payload
	return nil

}

// Price | function to get price per kilogram of given vegetable
func (store *VegetableStore) Price(payload string, reply *float64) error {

	// Get vegetable in the store by `VegetableTag`
	result, ok := store.data[strings.ToLower(payload)]
	// return error if the vegetable not exit in the store
	if !ok {
		return fmt.Errorf("Vegetable '%s' does not exist", payload)
	}

	// set `Price` to the `reply`
	*reply = result.Price
	return nil

}

// Amount | function to get available amount of given vegetable in kilograms
func (store *VegetableStore) Amount(payload string, reply *float64) error {

	// Get vegetable in the store by `VegetableTag`
	result, ok := store.data[strings.ToLower(payload)]

	// return error if the vegetable not exit in the store
	if !ok {
		return fmt.Errorf("Vegetable '%s' does not exist", payload)
	}

	// set `Amount` to the `reply`
	*reply = result.Amount
	return nil

}

// GetAll | function to get all vegetable names available in the store
func (store *VegetableStore) GetAll(payload string, reply *[]string) error {

	vegetables := make([]string, 0, len(store.data))

	// append all `VegetableTag` to the `vegetables` list
	for _, v := range store.data {
		vegetables = append(vegetables, v.VegetableTag)
	}

	// set `vegetables` list to the `reply`
	*reply = vegetables
	return nil

}

// NewVegetableStore | function to build VegetableStore
func NewVegetableStore() *VegetableStore {
	const BIT_SIZE = 64
	content, err := ioutil.ReadFile("data.txt")

	if err != nil {
		log.Fatal(err)
		return &VegetableStore{
			data: make(map[string]Vegetable),
		}
	}
	vegesArray := strings.Split(string(content), "\n")

	vegetableMap := make(map[string]Vegetable)

	// prepare a `vegetableMap` from the vegetables in the file
	for _, veg := range vegesArray {
		item := strings.Split(veg, ":")
		price, errPrice := strconv.ParseFloat(item[1], BIT_SIZE)
		amount, errAmount := strconv.ParseFloat(item[2], BIT_SIZE)
		if errPrice != nil || errAmount != nil {
			return &VegetableStore{
				data: make(map[string]Vegetable),
			}
		}

		vegetableMap[strings.ToLower(item[0])] = Vegetable{
			VegetableTag: item[0],
			Price:        price,
			Amount:       amount,
		}

	}
	return &VegetableStore{
		data: vegetableMap,
	}
}
