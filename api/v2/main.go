package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/anaskhan96/soup"
)

type SponsorsCount struct {
	Current int `json:"current"`
	Past    int `json:"past"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Result string `json:"result"`
}

type SuccessResponseCount struct {
	Status   string        `json:"status"`
	Sponsors SponsorsCount `json:"sponsors"`
}

func generateErrorResponse(result string) string {
	resp := ErrorResponse{
		Status: "error",
		Result: result,
	}

	jsonData, _ := json.Marshal(resp)
	return string(jsonData)
}

func getSponsorCount(username string) string {
	if username == "" {
		return generateErrorResponse("No user specified.")
	}

	url := fmt.Sprintf("https://github.com/sponsors/%s", username)
	htmlResponse, err := soup.Get(url)

	if err != nil {
		return generateErrorResponse("Unable to fetch the page.")
	}

	doc := soup.HTMLParse(htmlResponse)

	currentCount := 0
	pastCount := 0

	section := doc.Find("div", "id", "sponsors-section-list")
	if section.Error != nil {
		inactiveUrl := url + "/sponsors_partial?filter=inactive"
		htmlResponse, err = soup.Get(inactiveUrl)
		if err != nil {
			if err.Error() == io.EOF.Error() {
				pastCount = 0
			} else {
				return generateErrorResponse("Unable to fetch the page." + err.Error())
			}
		} else {
			doc = soup.HTMLParse(htmlResponse)
			svgs := doc.FindAll("svg")
			imgs := doc.FindAll("img")

			pastCount = len(svgs) + len(imgs)
		}

		activeUrl := url + "/sponsors_partial?filter=active"
		htmlResponse, err = soup.Get(activeUrl)
		if err != nil {
			if err.Error() == io.EOF.Error() {
				currentCount = 0
			} else {
				return generateErrorResponse("Unable to fetch the page." + err.Error())
			}
		} else {
			doc = soup.HTMLParse(htmlResponse)
			svgs := doc.FindAll("svg")
			imgs := doc.FindAll("img")

			currentCount = len(svgs) + len(imgs)
		}

		if pastCount == 0 && currentCount == 0 {
			return generateErrorResponse("GitHub Sponsors aren't setup with this user. Error: " + section.Error.Error())
		}
	} else {
		soup.SetDebug(true)

		sponsorSection := section.Find("div")
		currentCountElement := sponsorSection.Find("span")
		pastCountElement := sponsorSection.FindNextElementSibling().Find("span")
	
		if pastCountElement.Error != nil {
			return generateErrorResponse("An error occured: " + currentCountElement.Error.Error())
		}
		pastCount, err = strconv.Atoi(pastCountElement.Text())
		if err != nil {
			return generateErrorResponse(err.Error())
		}
	
		if currentCountElement.Error != nil {
			return generateErrorResponse("An error occured: " + currentCountElement.Error.Error())
		}
		currentCount, err = strconv.Atoi(currentCountElement.Text())
		if err != nil {
			return generateErrorResponse(err.Error())
		}
	}

	resp := SuccessResponseCount{
		Status: "success",
		Sponsors: SponsorsCount{
			Current: currentCount,
			Past:    pastCount,
		},
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		return generateErrorResponse("Failed to marshal response")
	}

	return string(jsonData)
}

func Main(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := r.URL.Query()
	username := params.Get("u")

	if username == "" {
		fmt.Fprint(w, generateErrorResponse("No user specified."))
		return
	}

	sponsors := getSponsorCount(username)
	fmt.Fprint(w, sponsors)
}
