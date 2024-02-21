package domain

type Group struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func NewGroup(id int, name string) (*Group, error) {
	return &Group{
		Id:   id,
		Name: name,
	}, nil
}