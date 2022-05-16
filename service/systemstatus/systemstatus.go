package systemstatus

import (
	"github.com/hiscaler/woocommerce-go/entity/systemstatus"
	jsoniter "github.com/json-iterator/go"
)

func (s service) SystemStatus() (item systemstatus.SystemStatus, err error) {
	resp, err := s.woo.Client.R().Get("/system_status")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}
