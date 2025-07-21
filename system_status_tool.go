package woocommerce

import (
	"fmt"

	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type systemStatusToolService service

func (s systemStatusToolService) All() (item entity.SystemStatusTool, err error) {
	resp, err := s.httpClient.R().Get("/system_status/tools")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

func (s systemStatusToolService) One(id string) (item entity.SystemStatusTool, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/system_status/tools/%s", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

func (s systemStatusToolService) Run(id string) (item entity.SystemStatusTool, err error) {
	resp, err := s.httpClient.R().Put(fmt.Sprintf("/system_status/tools/%s", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}
