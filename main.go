package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	list, _ := searchByName("https://api.thecatapi.com/v1/breeds/search/?q=sib")
	fmt.Println("Gato Siberiano", list)
	list1, _ := listCategories("https://api.thecatapi.com/v1/categories")
	fmt.Println("Categorias: ", list1)
}

type Lists []List
type List struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	URLWikipedia string `json:"wikipedia_url"`
}

func searchByName(url string) (Lists, error) {
	response, err := http.Get("https://api.thecatapi.com/v1/breeds/search/?q=sib")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	var list Lists
	json.Unmarshal(bytes, &list)
	return list, nil
}

type Lists1 []List1
type List1 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func listCategories(url string) (Lists1, error) {

	response, err := http.Get("https://api.thecatapi.com/v1/categories")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	var list1 Lists1
	json.Unmarshal(bytes, &list1)
	return list1, nil

}
