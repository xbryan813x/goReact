package notes

// A Note is
type Note struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	Note string `json:"note"`
}

// Record contains all the properties for each record
type Record struct {
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

// A SearchRecord contains all data from request from client
type SearchRecord Note

// Empty determines if struct is empty
func (r Record) IsEmpty() bool {
	return (r.Name == "")
}
