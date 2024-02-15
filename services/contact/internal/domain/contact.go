package domain

import (
	"encoding/json"
)

type Contact struct {
	Id         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Phone      string `json:"phone"`
	GroupId    int    `json:"group_id"`
}

func (c *Contact) JSON() string {
	json, err := json.MarshalIndent(c, "", "	")
	if err != nil {
		return ""
	}
	return string(json)
}
