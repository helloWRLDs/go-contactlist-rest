package domain

type Contact struct {
	Id         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Phone      string `json:"phone"`
}

func NewContact(id int, firstName, lastName, middleName, phone string) (*Contact, error) {
	return &Contact{
		Id:         id,
		FirstName:  firstName,
		LastName:   lastName,
		MiddleName: middleName,
		Phone:      phone,
	}, nil
}

func (c *Contact) Validate() error {
	if len(c.FirstName) == 0 || len(c.LastName) == 0 || len(c.MiddleName) == 0 {
		return ErrEmptyNameFields
	}
	return nil
}
