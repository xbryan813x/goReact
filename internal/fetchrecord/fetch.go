package fetchrecord

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/xbryan813x/goReact/notes"
)

func compareName(recordString string, record notes.Record) string {
	userName := notes.Store[len(notes.Store)-1].Name
	isMatch := strings.Contains(recordString, userName)
	if isMatch {
		return fmt.Sprint(record)
	}

	return fmt.Sprint("Sorry No Match Found... Try Again")
}

// FetchRecord is used to fetch the record for the user
func FetchRecord(note notes.Note) notes.Record {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// fmt.Printf("%s\n", body)
	var records []notes.Record

	error := json.Unmarshal([]byte(body), &records)

	if error != nil {
		log.Fatal(error)
	}
	notes.RecordList = records
	for _, record := range records {
		for i := 0; i < reflect.TypeOf(record).NumField(); i++ {
			if fmt.Sprint(reflect.ValueOf(record).Field(i).Kind()) != "struct" {
				if compareName(fmt.Sprint(reflect.ValueOf(record).Field(i)), record) != "Sorry No Match Found... Try Again" {
					fmt.Println(compareName(fmt.Sprint(reflect.ValueOf(record).Field(i)), record))
					break
				}
			}
		}
	}
	// fmt.Println("Record list in state => ", notes.RecordList)
	// SearchRecord(note, records)
	return records[0]
}

// We are going to store all the user input into a centralized pkg
// Convert data into proper Go data structures for us to consume
// We will then compute if we get any match on the records

// We need to filter the recordList with the query inputted by user
// and then return the most closely matching record by name

// SearchRecord returns the record that closest matches the profile given
func SearchRecord(query notes.Note, recordList []notes.Record) {
	for _, e := range recordList {

		fmt.Println("element ====> \n", reflect.TypeOf(e))
		fmt.Println("element ====> \n", reflect.TypeOf(e))
	}
	// fmt.Println("THis is the records!!! = >> ", query)
}
