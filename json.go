package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type User struct {
	Userid  int     `json:"id"`
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

type Address struct {
	Area    string `json:"area"`
	Country string `json:"country"`
}

type TechDets struct {
	Technolgy  string  `json:"techdata"`
	Experience float64 `json:"exp"`
}

type Tech struct {
	Userid   int        `json:"id"`
	TechDets []TechDets `json:"techDets"`
}

type ContactDets struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Contact struct {
	Userid      int         `json:"id"`
	ContactDets ContactDets `json:contactDets`
}

type mergeStruct struct {
	Userid      int
	Name        string
	Address     Address
	TechDetails []TechDets
	Email       string
	Phone       string
}

var mapofCountryCode = map[string]string{
	"IND": "+91",
	"UK":  "+41",
}

func createJSON(users []User, techDetails []Tech, contactDetails []Contact) {

	merged := make([]mergeStruct, len(users))
	for i, userdata := range users {

		newdata := mergeStruct{}
		newdata.Userid = userdata.Userid
		newdata.Name = userdata.Name
		newdata.Address.Area = userdata.Address.Area
		newdata.Address.Country = userdata.Address.Country

		for _, techdata := range techDetails {

			if userdata.Userid == techdata.Userid {
				newdata.TechDetails = techdata.TechDets
			}

		}

		for _, contactdata := range contactDetails {

			if userdata.Userid == contactdata.Userid {
				newdata.Email = contactdata.ContactDets.Email
				countryCode := mapofCountryCode[newdata.Address.Country]
				newdata.Phone = countryCode + "-" + contactdata.ContactDets.Phone
			}
		}
		merged[i] = newdata
	}
	file, _ := json.MarshalIndent(merged, "", " ")
	_ = ioutil.WriteFile("newdata.json", file, 0644)
}

func main() {

	usercontent, error1 := ioutil.ReadFile("user.json")
	if error1 != nil {
		fmt.Println(error1.Error())
	}

	var users []User
	error2 := json.Unmarshal([]byte(usercontent), &users)
	if error2 != nil {
		fmt.Println(error2.Error())

	}

	techcontent, error3 := ioutil.ReadFile("tech.json")
	if error3 != nil {
		fmt.Println(error3.Error())
	}

	var techDetails []Tech
	error4 := json.Unmarshal([]byte(techcontent), &techDetails)
	if error4 != nil {
		fmt.Println(error4.Error())

	}

	contactcontent, error5 := ioutil.ReadFile("contact.json")
	if error5 != nil {
		fmt.Println(error5.Error())
	}

	var contactDetails []Contact
	error6 := json.Unmarshal([]byte(contactcontent), &contactDetails)
	if error6 != nil {
		fmt.Println(error6.Error())
	}
	createJSON(users, techDetails, contactDetails)
}
