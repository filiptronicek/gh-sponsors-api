package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/anaskhan96/soup"
)

type Sponsors struct {
	Current int `json:"current"`
	Past    int `json:"past"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Result string `json:"result"`
}

type SuccessResponse struct {
	Status   string   `json:"status"`
	Sponsors Sponsors `json:"sponsors"`
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

	resp := SuccessResponse{
		Status: "success",
		Sponsors: Sponsors{
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

func Handler(w http.ResponseWriter, r *http.Request) {
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
