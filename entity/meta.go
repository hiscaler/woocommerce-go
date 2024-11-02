package entity

import "encoding/json"

type Meta struct {
	ID    int             `json:"id"`
	Key   string          `json:"key"`
	Value json.RawMessage `json:"value"`
}
