package main

import "fmt"

type ID struct {
	id int
}
type ADDRESS struct {
	area    string
	country string
}
type Slice1 struct {
	ID
	name string
	ADDRESS
}
type TECHDETS struct {
	tech []string
	exp  []float64
}
type Slice2 struct {
	ID
	TECHDETS
}
type ContactDETS struct {
	email string
	phone string
}
type Slice3 struct {
	ID
	ContactDETS
}

func main() {
	listOfPeople := []Slice1{}
	listOfPeople = append(listOfPeople, Slice1{ID{1}, "Peter", ADDRESS{"London", "UK"}})
	listOfPeople = append(listOfPeople, Slice1{ID{2}, "David", ADDRESS{"Mumbai", "India"}})

	techDetails := []Slice2{}
	techDetails = append(techDetails, Slice2{ID{2}, TECHDETS{[]string{"JavaScript", "goLang"}, []float64{7.0, 8.0}}})
	techDetails = append(techDetails, Slice2{ID{1}, TECHDETS{[]string{"ReactJS", "VueJS"}, []float64{5.0, 7.0}}})

	contactDetails := []Slice3{}
	contactDetails = append(contactDetails, Slice3{ID{1}, ContactDETS{"peter@bacancy.com", "111111111"}})
	contactDetails = append(contactDetails, Slice3{ID{2}, ContactDETS{"david@bacancy.com", "222222222"}})

	var mapOfCountryCode = map[string]string{
		"IND": "+91",
		"UK":  "+41",
	}

	for i := 0; i < 2; i++ {
		var tech string
		for j := 0; j < 2; j++ {
			if listOfPeople[i].id == techDetails[j].id {
				tech = (techDetails[j].tech[0] + " " + techDetails[j].tech[1])
			}
		}
		var mail string
		for j := 0; j < 2; j++ {
			if listOfPeople[i].id == contactDetails[j].id {
				mail = contactDetails[j].email
			}
		}
		var contactNumber string
		if listOfPeople[i].country == "India" {
			contactNumber = mapOfCountryCode["IND"] + "-" + contactDetails[i].phone
		} else if listOfPeople[i].country == "UK" {
			contactNumber = mapOfCountryCode["UK"] + "-" + contactDetails[i].phone
		}
		finalMap := map[int]map[string]string{
			listOfPeople[i].id: map[string]string{
				"Name":        listOfPeople[i].name,
				"Address":     listOfPeople[i].area,
				"TechDetails": tech,
				"Email":       mail,
				"Phone":       contactNumber,
			},
		}
		fmt.Println(finalMap)
	}

}
