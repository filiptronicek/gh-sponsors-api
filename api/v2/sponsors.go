package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anaskhan96/soup"
)

type Sponsor struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type Sponsors struct {
	Current []Sponsor `json:"current"`
	Past    []Sponsor `json:"past"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Result string `json:"result"`
}

type SuccessResponseSponsors struct {
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

func getSponsors(username string) string {
	if username == "" {
		return generateErrorResponse("No user specified.")
	}

	url := fmt.Sprintf("https://github.com/sponsors/%s", username)
	htmlResponse, err := soup.Get(url)

	if err != nil {
		return generateErrorResponse("Unable to fetch the page.")
	}

	doc := soup.HTMLParse(htmlResponse)

	soup.SetDebug(true)

	section := doc.Find("div", "id", "sponsors-section-list")
	if section.Error != nil {
		return generateErrorResponse("GitHub Sponsors aren't setup with this user. Error: " + section.Error.Error())
	}

	sponsorContainers := section.FindAll("remote-pagination")

	pastSponsorsContainer := sponsorContainers[1]
	currentSponsorsContainer := sponsorContainers[0]

	pastSponsors := append(
		pastSponsorsContainer.FindAll("a", "data-hovercard-type", "user"),
		pastSponsorsContainer.FindAll("a", "data-hovercard-type", "organization")...,
	)
	currentSponsors := append(
		currentSponsorsContainer.FindAll("a", "data-hovercard-type", "user"),
		currentSponsorsContainer.FindAll("a", "data-hovercard-type", "organization")...,
	)

	if pastSponsorsContainer.Error != nil {
		return generateErrorResponse("GitHub Sponsors aren't setup with this user. Error: " + pastSponsorsContainer.Error.Error())
	}

	if currentSponsorsContainer.Error != nil {
		return generateErrorResponse("GitHub Sponsors aren't setup with this user. Error: " + currentSponsorsContainer.Error.Error())
	}

	pastSponsorsList := []Sponsor{}

	for _, sponsor := range pastSponsors {
		sponsorUsername := sponsor.Find("img").Attrs()["alt"][1:]
		sponsorAvatar := sponsor.Find("img").Attrs()["src"]

		pastSponsorsList = append(pastSponsorsList, Sponsor{
			Username: sponsorUsername,
			Avatar:   sponsorAvatar,
		})
	}

	currentSponsorsList := []Sponsor{}

	for _, sponsor := range currentSponsors {
		sponsorUsername := sponsor.Find("img").Attrs()["alt"][1:]
		sponsorAvatar := sponsor.Find("img").Attrs()["src"]

		currentSponsorsList = append(currentSponsorsList, Sponsor{
			Username: sponsorUsername,
			Avatar:   sponsorAvatar,
		})
	}

	resp := SuccessResponseSponsors{
		Status: "success",
		Sponsors: Sponsors{
			Current: currentSponsorsList,
			Past:    pastSponsorsList,
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

	sponsors := getSponsors(username)
	fmt.Fprint(w, sponsors)
}
