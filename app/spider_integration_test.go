package app

import (
	"net/http"
	"testing"
)

func Test_integration_spider_arena(t *testing.T) {
	result, err := Spider(http.DefaultClient, ARENA)

	if len(result) <= 0 {
		t.Errorf("result length is %d", len(result))
	}

	if err != nil {
		t.Errorf("found error %+v", err)
	}
}

func Test_integration_spider_hall(t *testing.T) {
	result, err := Spider(http.DefaultClient, HALL)

	if len(result) <= 0 {
		t.Errorf("result length is %d", len(result))
	}

	if err != nil {
		t.Errorf("found error %+v", err)
	}
}