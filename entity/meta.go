package entity

import "encoding/json"

type Meta struct {
	ID    int    `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (m *Meta) UnmarshalJSON(data []byte) error {
	type Alias Meta
	aux := &struct {
		Value json.RawMessage `json:"value"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var str string
	if err := json.Unmarshal(aux.Value, &str); err == nil {
		m.Value = str
		return nil
	}

	m.Value = string(aux.Value)
	return nil
}
