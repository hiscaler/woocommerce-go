package woocommerce

import (
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/hiscaler/woocommerce-go/constant"
	"time"
)

// ToISOTimeString Convert to iso time string
// If date format is invalid, then return original value
func ToISOTimeString(dateStr string, addMinTimeString, addMaxTimeString bool) (s string) {
	s = dateStr
	if dateStr == "" {
		return
	}
	format, err := dateparse.ParseFormat(dateStr)
	if err == nil && (format == constant.DateFormat || format == constant.DatetimeFormat || format == constant.WooDatetimeFormat) {
		if addMinTimeString {
			dateStr += " 00:00:00"
		}
		if addMaxTimeString {
			dateStr += " 23:59:59"
		}
		if t, err := dateparse.ParseAny(dateStr); err == nil {
			return t.Format(time.RFC3339)
		}
	}
	return s
}

// IsValidateTime Is validate time
func IsValidateTime(dateStr string) error {
	format, err := dateparse.ParseFormat(dateStr)
	if err != nil {
		return err
	}
	switch format {
	case constant.DateFormat, constant.DatetimeFormat, constant.WooDatetimeFormat:
		return nil
	default:
		return fmt.Errorf("%s 日期格式无效", dateStr)
	}
}
