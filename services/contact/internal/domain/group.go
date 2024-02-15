package domain

import "encoding/json"

type Group struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (g *Group) JSON() string {
	json, err := json.MarshalIndent(g, "", "	")
	if err != nil {
		return ""
	}
	return string(json)
}
