package app

import (
	"testing"
	"time"

	"github.com/jo-hoe/mb-arena-service/app/test"
)

func Test_Spider_items(t *testing.T) {
	// prepare
	mockResponse := test.ResponseSummery{
		ResponseCode: 200,
		ResponseBody: test.HTMLTestFile,
	}
	client := test.CreateMockClient(mockResponse)

	// test
	result, err := Spider(client)

	// assert
	expectedItemCount := 131
	if len(result) != expectedItemCount {
		t.Errorf("result length is %d while expected was %d", len(result), expectedItemCount)
	}
	if err != nil {
		t.Errorf("found error %+v", err)
	}
}

func Test_Spider_item(t *testing.T) {
	mockResponse := test.ResponseSummery{
		ResponseCode: 200,
		ResponseBody: test.HTMLTestFile,
	}
	client := test.CreateMockClient(mockResponse)

	result, err := Spider(client)

	if err != nil {
		t.Errorf("found error %+v", err)
	}

	testItem := result[0]
	expectedTime := time.Date(2023, 10, 14, 17, 00, 0, 0, time.UTC)
	if testItem.Start.Compare(expectedTime) != 0 {
		t.Errorf("found '%s', but expected '%s'", testItem.Start, expectedTime)
	}
	assert(t, testItem.Name, "50 Cent")
	assert(t, testItem.PictureUrl, "https://www.mercedes-benz-arena-berlin.de/assets/img/AnOp_50_Cent_MBA_WS_460x205px_01_58-d2a8f966b7.jpg")
	assert(t, testItem.Link, "https://www.mercedes-benz-arena-berlin.de/en/events/detail/50-cent/2023-10-14-1900")
}

func assert(t *testing.T, actual string, expected string) {
	if expected != actual {
		t.Errorf("found '%s', but expected '%s'", actual, expected)
	}
}
