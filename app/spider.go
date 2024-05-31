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
	HOST       = "https://www.uber-arena.de"
	ALL_EVENTS = HOST + "/en/events-tickets/"
	ASSET_PATH = HOST + "/assets/img/"
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

var nonNumericRegex = regexp.MustCompile(`[^0-9]+`)
var timeRegex = regexp.MustCompile(`([0-9]{2}):([0-9]{2})`)

func Spider(httpClient *http.Client) (result []MBEvent, err error) {
	result = make([]MBEvent, 0)

	responseString, err := getResponseAsString(httpClient, ALL_EVENTS)
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
		imageUrl = strings.Replace(imageUrl, "./Events _ Mercedes-Benz Arena Berlin_files/", ASSET_PATH, 1)

		yearElements := htmlquery.Find(eventHtml, XPATH_YEAR)
		year := clearString(getInnerText(yearElements[0]))

		monthElements := htmlquery.Find(eventHtml, XPATH_MONTH)
		month := clearString(getInnerText(monthElements[0]))

		dayElements := htmlquery.Find(eventHtml, XPATH_DAY)
		day := clearString(getInnerText(dayElements[0]))

		hourElements := htmlquery.Find(eventHtml, XPATH_HOURS)
		hours := getInnerText(hourElements[0])
		timeElements := timeRegex.FindStringSubmatch(hours)

		var timestamp time.Time
		location, err := time.LoadLocation("Europe/Berlin")
		if err != nil {
			log.Printf("could not load time location, error: %+v", err)
			continue
		}
		if len(timeElements) == 3 {
			// time was available
			timestamp, err = time.ParseInLocation(time.DateTime, fmt.Sprintf("%s-%s-%s %s:%s:00", year, month, day, timeElements[1], timeElements[2]), location)
		} else {
			// not time available
			timestamp, err = time.ParseInLocation(time.DateOnly, fmt.Sprintf("%s-%s-%s", year, month, day), location)
		}
		if err != nil {
			log.Printf("could not read time for event '%s', error: %+v", title, err)
			continue
		}

		result = append(result, MBEvent{
			Name:       title,
			Link:       link,
			PictureUrl: imageUrl,
			Start:      timestamp,
		})
	}

	return result, nil
}

func clearString(str string) string {
	return nonNumericRegex.ReplaceAllString(str, "")
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
