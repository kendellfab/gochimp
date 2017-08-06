package stats

import "encoding/json"

type Count struct {
	count int `json:"count"`
}

func (c Count) Json() ([]byte, error) {
	return json.Marshal(c)
}