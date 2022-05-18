package woocommerce

import (
	"github.com/hiscaler/woocommerce-go/entity/systemstatus"
	jsoniter "github.com/json-iterator/go"
)

type systemService service

func (s systemService) Status() (item systemstatus.SystemStatus, err error) {
	resp, err := s.httpClient.R().Get("/system_status")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}
