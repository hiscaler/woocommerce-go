package woocommerce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportService_All(t *testing.T) {
	_, err := wooClient.Services.Report.All()
	assert.Equal(t, nil, err)
}

func TestReportService_SalesReports(t *testing.T) {
	req := SalesReportsQueryParams{
		DateMin: "2022-01-01",
		DateMax: "2022-01-01",
	}
	_, err := wooClient.Services.Report.SalesReports(req)
	assert.Equal(t, nil, err)
}
