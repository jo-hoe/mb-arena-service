package app

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

var (
	ARENA      = "https://www.uber-arena.de"
	HALL       = "https://www.uber-eats-music-hall.de"
	ALL_EVENTS = "/en/events-tickets/"

	ARENA_NAME = "Uber Arena"
	HALL_NAME  = "Uber Eats Music Hall"
)

var (
	XPATH_EVENT = "//div[@class='active-date entry clearfix' or @class='active-date entry alt clearfix']"
	XPATH_TITLE = "//h3[@class='event-title']//a"
	XPATH_IMAGE = "//img"
	XPATH_YEAR  = "//span[@class='m-date__year']"
	XPATH_MONTH = "//span[@class='m-date__month']"
	XPATH_DAY   = "//span[@class='m-date__day']"
	XPATH_HOURS = "//span[@class='m-date__hour']"
)

var replacer = strings.NewReplacer(
	"jan", "01",
	"feb", "02",
	"mar", "03",
	"apr", "04",
	"may", "05",
	"jun", "06",
	"jul", "07",
	"aug", "08",
	"sep", "09",
	"oct", "10",
	"nov", "11",
	"dec", "12",
)

var nonNumericRegex = regexp.MustCompile(`[^0-9]+`)
var nonAlphabeticRegex = regexp.MustCompile(`[^A-Za-z]+`)
var timeRegex = regexp.MustCompile(`([0-9]{2}):([0-9]{2})`)

func Spider(httpClient *http.Client, host string) (result []Event, err error) {
	result = make([]Event, 0)

	responseString, err := getResponseAsString(httpClient, fmt.Sprintf("%s%s", host, ALL_EVENTS))
	if err != nil {
		return result, err
	}

	doc, err := htmlquery.Parse(strings.NewReader(responseString))
	if err != nil {
		return result, err
	}
	eventHtmls := htmlquery.Find(doc, XPATH_EVENT)

	for _, eventHtml := range eventHtmls {
		titleElements := htmlquery.Find(eventHtml, XPATH_TITLE)
		link := getAttribute(titleElements[0], "href")
		title := getInnerText(titleElements[0])

		imageElements := htmlquery.Find(eventHtml, XPATH_IMAGE)
		imageUrl := getAttribute(imageElements[0], "src")

		yearElements := htmlquery.Find(eventHtml, XPATH_YEAR)
		year := clearNumericString(getInnerText(yearElements[0]))

		monthElements := htmlquery.Find(eventHtml, XPATH_MONTH)
		month := ""
		if host == ARENA {
			month = clearNumericString(getInnerText(monthElements[0]))
		}
		if host == HALL {
			month = clearString(getInnerText(monthElements[0]))
			month = replacer.Replace(strings.ToLower(month))
		}

		dayElements := htmlquery.Find(eventHtml, XPATH_DAY)
		day := clearNumericString(getInnerText(dayElements[0]))

		hourElements := htmlquery.Find(eventHtml, XPATH_HOURS)
		hours := getInnerText(hourElements[0])
		timeElements := timeRegex.FindStringSubmatch(hours)

		var timestamp time.Time
		region, err := time.LoadLocation("Europe/Berlin")
		if err != nil {
			log.Printf("could not load time location, error: %+v", err)
			continue
		}
		if len(timeElements) == 3 {
			// time was available
			timestamp, err = time.ParseInLocation(time.DateTime, fmt.Sprintf("%s-%s-%s %s:%s:00", year, month, day, timeElements[1], timeElements[2]), region)
		} else {
			// not time available
			timestamp, err = time.ParseInLocation(time.DateOnly, fmt.Sprintf("%s-%s-%s", year, month, day), region)
		}
		if err != nil {
			log.Printf("could not read time for event '%s', error: %+v", title, err)
			continue
		}

		location := ""
		if host == HALL {
			location = HALL_NAME
		} else if host == ARENA {
			location = ARENA_NAME
		}

		result = append(result, Event{
			Name:       title,
			Link:       link,
			PictureUrl: imageUrl,
			Start:      timestamp,
			Location:   location,
		})
	}

	return result, nil
}

func clearNumericString(str string) string {
	return nonNumericRegex.ReplaceAllString(str, "")
}

func clearString(str string) string {
	return nonAlphabeticRegex.ReplaceAllString(str, "")
}

func getAttribute(node *html.Node, attributeName string) (result string) {
	result = ""
	for _, attribute := range node.Attr {
		if attribute.Key == attributeName {
			result = attribute.Val
			break
		}
	}
	return result
}

func getInnerText(node *html.Node) (result string) {
	result = ""
	if node.FirstChild != nil {
		result = node.FirstChild.Data
	}
	return result
}

func getResponseAsString(httpClient *http.Client, url string) (result string, err error) {
	result = ""

	response, err := httpClient.Get(url)
	if err != nil {
		return result, err
	}
	if response.StatusCode != http.StatusOK {
		return result, fmt.Errorf("return code of '%s' was %d", url, response.StatusCode)
	}
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return result, err
	}
	return string(bodyBytes), err
}
