package main

import (
	"encoding/json"
	"fmt"
)

type TestField struct {
	Id   int    `json:"id,string"`
	Name string `json:"name"`
}

type Item struct {
	Id    int         `json:"id,string"`
	Title string      `json:"title"`
	Test  []TestField `json:"test"`
}

func main() {
	data := []byte(`
		[
			{
				"id": "1",
				"title": "Title 1",
				"test": [
					{
						"id": "1",
						"name": "Alfonz"
					},
					{
						"id": "2",
						"name": "Mucha"
					}
				]
			},
			{
				"id": "2",
				"title": "Title 2",
				"test": [
					{
						"id": "3",
						"name": "Karel"
					},
					{
						"id": "4",
						"name": "Drda"
					}
				]
			}
		]
	`)

	// Known structure
	var test []Item
	err := json.Unmarshal(data, &test)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%+v\n", test)

	fmt.Println()
	fmt.Println("---------------------------------------------------")
	fmt.Println()

	// Unknown strucuture
	var test2 interface{}

	err = json.Unmarshal(data, &test2)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", test2)
}

/*
Output:

[
	{
		Id:1
		Title:Title 1
		Test:[
			{
				Id:1
				Name:Alfonz
			}
			{
				Id:2 Name:Mucha
			}
		]
	}
	{
		Id:2
		Title:Title 2
		Test:[
			{
				Id:3
				Name:Karel
			}
			{
				Id:4
				Name:Drda
			}
		]
	}
]

---------------------------------------------------

[]interface {}{
	map[string]interface {}{
		"id":"1",
		"title":"Title1",
		"test":[]interface {}{
			map[string]interface {}{
				"id":"1",
				"name":"Alfonz"
			},
			map[string]interface {}{
				"id":"2",
				"name":"Mucha"
			}
		}
	},
	map[string]interface {}{
		"id":"2",
		"title":"Title 2",
		"test":[]interface {}{
			map[string]interface {}{
				"id":"3",
				"name":"Karel"
			},
			map[string]interface {}{
				"id":"4",
				"name":"Drda"
			}
		}
	}
}
*/
