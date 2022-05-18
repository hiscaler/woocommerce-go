package woocommerce

import jsoniter "github.com/json-iterator/go"

type settingService service

type Group struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
	ParentId    string `json:"parent_id"`
	SubGroups   string `json:"sub_groups"`
}

func (s settingService) All() (items []Group, err error) {
	resp, err := s.httpClient.R().Get("/settings")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}
