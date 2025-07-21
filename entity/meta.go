package entity

import "encoding/json"

type Meta struct {
	ID    int             `json:"id"`
	Key   string          `json:"key"`
	Value json.RawMessage `json:"value"`
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
		var b []byte
		b, err = json.Marshal(str)
		if err != nil {
			return err
		}
		m.Value = b
		return nil
	}

	m.Value = aux.Value
	return nil
}
