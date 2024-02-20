package domain

type Contact struct {
	Id         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Phone      string `json:"phone"`
}

func NewContact(id int, firstName, lastName, middleName, phone string) (*Contact, error) {
	//Validation in future
	return &Contact{
		Id:         id,
		FirstName:  firstName,
		LastName:   lastName,
		MiddleName: middleName,
		Phone:      phone,
	}, nil
}
