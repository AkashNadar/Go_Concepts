package main

import (
	"encoding/json"
	"fmt"
)

type book struct {
	Name   string   `json:"book_name"`
	Price  int      `json:"-"` //It will not show this is will be helpful in case of passwords
	Author string   `json:"author"`
	Tags   []string `json:"tags,omitempty"` //omitempty -> if the list is empty then it will not show
}

func EncodeJson(bookData []book) {
	// data, err := json.Marshal(bookData)
	data, err := json.MarshalIndent(bookData, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(bookData)
	fmt.Printf("%s\n", data)

}

func DecodeJson() {
	data := []byte(`
	{
		"book_name": "maths",
		"Price": 200,
		"Author": "kumboj",
		"Tags": ["11","12","13"]
	}
	`)

	var result book

	isValid := json.Valid(data)

	if isValid {
		json.Unmarshal(data, &result)
		fmt.Printf("%#v\n", result)
		// fmt.Println(result.Name)
	} else {
		fmt.Println("Error decoding data")
	}

	// We can also use the map
	var jsonData map[string]interface{}
	json.Unmarshal(data, &jsonData)
	fmt.Println(jsonData)
	fmt.Println(jsonData["book_name"])

}

func main() {

	// bookData := []book{
	// 	{"maths", 200, "kumboj", []string{"11", "12", "13"}},
	// 	{"physics", 400, "ram", []string{"fun", "hard", "tough"}},
	// 	{"maths", 200, "karan", nil},
	// }

	// EncodeJson(bookData)
	DecodeJson()

}
