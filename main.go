package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name      string
	Age       uint16
	Money     int16
	AvgGrades float64
	Happiness float64
	Hobbies   []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf(
		"Position: %s, \n Name is: %s, \n Age is %d,\n balance: %d",
		"CEO", u.Name, u.Age, u.Money)
}

func (u *User) setName(name string) {
	u.Name = name
}

func homePage(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Go main page here")
}

func contactsPage(page http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 20, 100, 4.3, 1.3, []string{"Football", "Programming"}}
	// bob.setName("Alex")
	// fmt.Fprintf(page, "Here is our contacts: \n"+bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/home.html")
	tmpl.Execute(page, bob)
}

func makeRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	http.ListenAndServe(":8080", nil)
}

func main() {
	makeRequest()
}
