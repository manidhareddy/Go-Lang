package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "https://dummyjson.com/products/categories"

type ProductCategory struct {
	Slug string `json:"slug"` // lowercase means private field |  uppercase means public field
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {
	response, error := http.Get(url)
	if error != nil {
		fmt.Println("Error:", error)
		panic(error)
	}
	defer response.Body.Close()
	fmt.Println("Response Status Code:", response.StatusCode)
	data, error := io.ReadAll(response.Body)
	if error != nil {
		fmt.Println("Error:", error)
		panic(error)
	}
	if response.StatusCode == http.StatusOK {
		var categories []ProductCategory
		fmt.Println("Data:", string(data))
		error = json.Unmarshal(data, &categories)
		if error != nil {
			panic(error)
		}
		for _, category := range categories {
			fmt.Printf("%v\n", category.Name)
		}
	} else {
		fmt.Println("Failed to fetch categories")
	}
}
