package app

import (
	"net/http"
	"testing"
)

func Test_integration_spider(t *testing.T) {
	result, err := Spider(http.DefaultClient)

	if len(result) <= 0 {
		t.Errorf("result length is %d", len(result))
	}

	if err != nil {
		t.Errorf("found error %+v", err)
	}
}
