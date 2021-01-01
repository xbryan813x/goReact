package fetchrecord

import (
	"io/ioutil"
	"net/http"
)

// FetchRecord is used to fetch the record for the user
func FetchRecord() string {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// fmt.Printf("%s\n", body)
	return string(body)
}

// We are going to store all the user input into a centralized pkg
// Convert data into proper Go data structures for us to consume
// We will then compute if we get any match on the records
