package woocommerce

import (
	"fmt"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/dashboard-bg/woocommerce-go/constant"
)

// ToISOTimeString Convert to iso time string
// If date format is invalid, then return original value
// If dateStr include time part, and you set addMinTimeString/addMaxTimeString to true,
// but still return original dateStr value.
func ToISOTimeString(dateStr string, addMinTimeString, addMaxTimeString bool) (s string) {
	dateStr = strings.TrimSpace(dateStr)
	if dateStr == "" {
		return
	}

	s = dateStr
	format, err := dateparse.ParseFormat(dateStr)
	if err == nil && (format == constant.DateFormat || format == constant.DatetimeFormat || format == constant.WooDatetimeFormat) {
		if strings.Index(dateStr, " ") == -1 {
			if addMinTimeString {
				dateStr += " 00:00:00"
			}
			if addMaxTimeString {
				dateStr += " 23:59:59"
			}
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
		return fmt.Errorf("%s invalid date format", dateStr)
	}
}
