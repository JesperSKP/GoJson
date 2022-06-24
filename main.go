package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Define struct used to hold API response data
type Response []struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

// Define struct used to hold data from people.json
type People []struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

func main() {
	go getData()
	fmt.Scanln()
}

func getData() {
	for true {
		// Send get request to API
		resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
		if err != nil {
			fmt.Println("No response from request")
		}

		// Get json response from API
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		// Unmarshal json to struct
		var APIData Response
		if err := json.Unmarshal(body, &APIData); err != nil {
			fmt.Println("Error during Unmarshal: ", err)
		}

		//Read json data from people.json
		content, err := ioutil.ReadFile("./people.json")
		if err != nil {
			fmt.Println("Error when opening file: ", err)
		}

		// Unmarshal json to struct
		var peopleData People
		if err := json.Unmarshal(content, &peopleData); err != nil {
			fmt.Println("Error during Unmarshal: ", err)
		}

		//Loop through API response and check if name and username exists in people.json
		for api := 0; api < len(APIData); api++ {
			for people := 0; people < len(peopleData); people++ {
				if APIData[api].Name == peopleData[people].Name {
					fmt.Println("Name:", peopleData[people].Name, "| Username:", peopleData[people].Username)
					break
				}
			}
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
