package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	name      string
	age       uint16
	money     int16
	avgGrades float64
	happiness float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf(
		"Position: %s, \n Name is: %s, \n Age is %d,\n balance: %d",
		"CEO", u.name, u.age, u.money)
}

func (u *User) setName(name string) {
	u.name = name
}

func homePage(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Go main page here")
}

func contactsPage(page http.ResponseWriter, r *http.Request) {
	bob := User{name: "Bob", age: 20, money: 100, avgGrades: 4.3, happiness: 1.3}
	// bob.setName("Alex")
	// fmt.Fprintf(page, "Here is our contacts: \n"+bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/home.html")
	tmpl.Execute(page, bob)
}

func makeRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts", contactsPage)
	http.ListenAndServe(":8080", nil)
}

func main() {
	makeRequest()
}
