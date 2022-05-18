package woocommerce

import (
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type systemService service

func (s systemService) Status() (item entity.SystemStatus, err error) {
	resp, err := s.httpClient.R().Get("/system_status")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}
