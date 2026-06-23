package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

	pastSponsorsList, err := paginateSponsors(username, "inactive")
	if err != nil {
		return generateErrorResponse("Failed to fetch past sponsors")
	}

	currentSponsorsList, err := paginateSponsors(username, "active")
	if err != nil {
		return generateErrorResponse("Failed to fetch current sponsors")
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

func paginateSponsors(username, filter string) ([]Sponsor, error) {
	var sponsorsList []Sponsor
	page := 1

	for {
		url := fmt.Sprintf("https://github.com/sponsors/%s/sponsors_partial?filter=%s&page=%d", username, filter, page)
		htmlResponse, err := soup.Get(url)
		if err != nil || strings.Contains(htmlResponse, "sponsors not found") {
			break
		}

		doc := soup.HTMLParse(htmlResponse)
		sponsors := parseSponsors(doc)
		if len(sponsors) == 0 {
			break
		}
		sponsorsList = append(sponsorsList, sponsors...)
		page++
	}

	return sponsorsList, nil
}

func parseSponsors(doc soup.Root) []Sponsor {
	sponsorsList := []Sponsor{}
	userSponsors := doc.FindAll("a", "data-hovercard-type", "user")
	orgSponsors := doc.FindAll("a", "data-hovercard-type", "organization")
	sponsorContainers := append(userSponsors, orgSponsors...)

	for _, sponsor := range sponsorContainers {
		img := sponsor.Find("img")
		if img.Error != nil {
			continue // Skip if there's no image
		}

		sponsorUsername := img.Attrs()["alt"][1:] // Remove '@' at the beginning
		sponsorAvatar := img.Attrs()["src"]

		sponsorsList = append(sponsorsList, Sponsor{
			Username: sponsorUsername,
			Avatar:   sponsorAvatar,
		})
	}

	return sponsorsList
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
