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
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var records []notes.Record

	error := json.Unmarshal([]byte(body), &records)

	if error != nil {
		log.Fatal(error)
	}
	notes.RecordList = records
	var recordStruct []notes.Record

	for _, record := range records {
		for i := 0; i < reflect.TypeOf(record).NumField(); i++ {
			if fmt.Sprint(reflect.ValueOf(record).Field(i).Kind()) != "struct" {
				if compareName(fmt.Sprint(reflect.ValueOf(record).Field(i)), record) != "Sorry No Match Found... Try Again" {
					fmt.Println(compareName(fmt.Sprint(reflect.ValueOf(record).Field(i)), record))
					// recordFound = compareName(fmt.Sprint(reflect.ValueOf(record).Field(i)), record)
					recordStruct = append(recordStruct, record)
					break
				}
			}
		}
	}
	return recordStruct[len(recordStruct)-1]
}

// SearchRecord returns the record that closest matches the profile given
func SearchRecord(query notes.Note, recordList []notes.Record) {
	for _, e := range recordList {

		fmt.Println("element ====> \n", reflect.TypeOf(e))
		fmt.Println("element ====> \n", reflect.TypeOf(e))
	}
}
