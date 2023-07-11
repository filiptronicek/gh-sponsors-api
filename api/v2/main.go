package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/anaskhan96/soup"
)

type SponsorsCount struct {
	Current int `json:"current"`
	Past    int `json:"past"`
}

type SuccessResponseCount struct {
	Status   string        `json:"status"`
	Sponsors SponsorsCount `json:"sponsors"`
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

	section := doc.Find("div", "id", "sponsors-section-list")
	if section.Error != nil {
		return generateErrorResponse("GitHub Sponsors aren't setup with this user. Error: " + section.Error.Error())
	}

	soup.SetDebug(true)

	sponsorSection := section.Find("div")
	currentCountElement := sponsorSection.Find("span")
	pastCountElement := sponsorSection.FindNextElementSibling().Find("span")

	if pastCountElement.Error != nil {
		return generateErrorResponse("GitHub Sponsors aren't setup with this user. Error: " + pastCountElement.Error.Error())
	}

	if currentCountElement.Error != nil {
		return generateErrorResponse("GitHub Sponsors aren't setup with this user. Error: " + currentCountElement.Error.Error())
	}

	currentCount, err := strconv.Atoi(currentCountElement.Text())
	if err != nil {
		return generateErrorResponse(err.Error())
	}

	pastCount, err := strconv.Atoi(pastCountElement.Text())
	if err != nil {
		return generateErrorResponse(err.Error())
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
