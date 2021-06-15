package main

import (
	"demoweb/model"
	"fmt"
	"html/template"
	"net/http"
)

type User model.User

func (u *User) getAllInfo() string {
	return fmt.Sprintf(
		"Position: %s, \n Name is: %s, \n Age is %d,\n balance: %d",
		"CEO", u.Name, u.Age, u.Money)
}

func (u *User) setName(name string) {
	u.Name = name
}

func index(page http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../templates/index.html")
	if err != nil {
		fmt.Fprintf(page, err.Error())
	} else {
		tmpl.Execute(page, "")
	}

}

func contactsPage(page http.ResponseWriter, r *http.Request) {
	// Todo change to get from db
	bob := User{"Bob", 20, 100, 4.3, 1.3, []string{"Football", "Programming"}}
	// bob.setName("Alex")
	// fmt.Fprintf(page, "Here is our contacts: \n"+bob.getAllInfo())
	tmpl, _ := template.ParseFiles("../templates/contacts.html")
	tmpl.Execute(page, bob)
}

func makeRequest() {
	http.HandleFunc("/", index)
	http.HandleFunc("/contacts/", contactsPage)
	http.ListenAndServe(":8080", nil)
}

func main() {
	makeRequest()
}
