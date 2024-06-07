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
		ResponseBody: test.ArenaHTMLTestFile,
	}
	client := test.CreateMockClient(mockResponse)

	// test
	result, err := Spider(client, ARENA)

	// assert
	expectedItemCount := 86
	if len(result) != expectedItemCount {
		t.Errorf("result length is %d while expected was %d", len(result), expectedItemCount)
	}
	if err != nil {
		t.Errorf("found error %+v", err)
	}
}

func Test_Spider_Arena_item(t *testing.T) {
	mockResponse := test.ResponseSummery{
		ResponseCode: 200,
		ResponseBody: test.ArenaHTMLTestFile,
	}
	client := test.CreateMockClient(mockResponse)

	result, err := Spider(client, ARENA)

	if err != nil {
		t.Errorf("found error %+v", err)
	}

	testItem := result[0]
	expectedTime := time.Date(2024, 06, 10, 18, 00, 0, 0, time.UTC)
	if testItem.Start.Compare(expectedTime) != 0 {
		t.Errorf("found '%s', but expected '%s'", testItem.Start, expectedTime)
	}
	assert(t, testItem.Name, "IVE")
	assert(t, testItem.PictureUrl, "https://www.uber-arena.de/assets/img/IVE_thumb-8805576581.jpg")
	assert(t, testItem.Link, "https://www.uber-arena.de/en/events/detail/ive/2024-06-10-2000")
}

func Test_Spider_Hall_item(t *testing.T) {
	mockResponse := test.ResponseSummery{
		ResponseCode: 200,
		ResponseBody: test.HallHTMLTestFile,
	}
	client := test.CreateMockClient(mockResponse)

	result, err := Spider(client, HALL)

	if err != nil {
		t.Errorf("found error %+v", err)
	}

	testItem := result[0]
	expectedTime := time.Date(2024, 06, 9, 18, 00, 0, 0, time.UTC)
	if testItem.Start.Compare(expectedTime) != 0 {
		t.Errorf("found '%s', but expected '%s'", testItem.Start, expectedTime)
	}
	assert(t, testItem.Name, "Big Time Rush")
	assert(t, testItem.PictureUrl, "https://uber-arena.production.carbonhouse.com/assets/img/BTR_Thumb-35d52c0b31.jpg")
	assert(t, testItem.Link, "https://www.uber-eats-music-hall.de/en/events/detail/big-time-rush/2024-06-09-2000")
}

func assert(t *testing.T, actual string, expected string) {
	if expected != actual {
		t.Errorf("found '%s', but expected '%s'", actual, expected)
	}
}
