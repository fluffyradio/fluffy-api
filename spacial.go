package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"fmt"

	"strings"

	"net/url"

	"github.com/labstack/echo"
)

// FluffyAPIURL is a string representing the Spacial Audio URL for interacting with the control system
const FluffyAPIURL = "http://widgets-proxy.cloudapp.net/webapi/station/"

func getAPIURL() string {
	return FluffyAPIURL + *spacialID + "/"
}

type (
	// Song represents an instance of a song from spacial
	Song struct {
		ID       string `json:"id"`
		Artist   string `json:"artist"`
		Title    string `json:"title"`
		Album    string `json:"album"`
		AlbumArt string `json:"album_art_url"`
		Website  string `json:"website_url"`
	}

	// RequestStatus represents an instance of the status of a request
	RequestStatus struct {
		Status string `json:"status"`
		SongID string `json:"song_id"`
	}
)

// songs returns an array of songs in the library
func songs(c echo.Context) error {
	// Parse query params
	take := c.QueryParam("take")
	top := c.QueryParam("top")
	q := c.QueryParam("q")

	if take == "" {
		take = "0"
	}

	if top == "" {
		top = "10"
	}

	// Get current song
	res, err := http.Get(getAPIURL() + "/library?format=json&start=" + take + "&top=" + top + "&search=" + url.QueryEscape(q) + "&mediaTypeCodes=MUS&token=" + *spacialToken)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Close the response, but defer to everything else
	defer res.Body.Close()

	// Unmarshal the body to the JSON object
	var data []map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	// Log the response from spacial
	if !*productionMode {
		fmt.Println(data)
	}

	// Convert the array into a []Song
	songs := make([]Song, len(data))
	for i, v := range data {
		r := Song{}
		r.ID = v["MediaItemId"].(string)
		r.Artist = v["Artist"].(string)
		r.Title = v["Title"].(string)
		r.Album = v["Album"].(string)

		if v["Picture"] != nil {
			r.AlbumArt = v["Picture"].(string)
		}

		if v["Website"] != nil {
			r.Website = v["Website"].(string)
		}

		songs[i] = r
	}

	//Return object
	return c.JSON(http.StatusOK, songs)
}

// currentSong returns the currently playing song
func currentSong(c echo.Context) error {

	// Get current song
	res, err := http.Get(getAPIURL() + "/history/npe?format=json&token=" + *spacialToken)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Close the response, but defer to everything else
	defer res.Body.Close()

	// Unmarshal the body to the JSON object
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	// Log the response from spacial
	if !*productionMode {
		fmt.Println(data)
	}

	// Grab the specific JSON object from the data payload
	song := data["m_Item2"].(map[string]interface{})

	// Build the result
	r := Song{}
	r.ID = song["MediaItemId"].(string)
	r.Artist = song["Artist"].(string)
	r.Title = song["Title"].(string)
	r.Album = song["Album"].(string)

	if song["Picture"] != nil {
		r.AlbumArt = song["Picture"].(string)
	}

	if song["Website"] != nil {
		r.Website = song["Website"].(string)
	}

	// Return the result
	return c.JSON(http.StatusOK, r)
}

// requestSong accepts a song request and returns status details
func requestSong(c echo.Context) error {
	id := c.Param("id")

	// Check for empty ID
	if id == "" {
		log.Println("Request media id is empty")
		return echo.NewHTTPError(http.StatusNotFound)
	}

	// Call Spacial
	res, err := http.Post(getAPIURL()+"/request/"+id+"?format=json&token="+*spacialToken, "text/plain", strings.NewReader(""))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Parse the response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Close the response, but defer to everything else
	defer res.Body.Close()

	// Normally this only fails when the rate limit is hit.
	if res.StatusCode != 200 {
		log.Println(res.Status + ": " + string(body))
		return echo.NewHTTPError(http.StatusTooManyRequests, string(body))
	}

	// Unmarshal the body to the JSON object
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	// Log the response from spacial
	if !*productionMode {
		fmt.Println(data)
	}

	// Success, let the user know
	s := new(RequestStatus)
	s.Status = "Pending"
	s.SongID = id
	return c.JSON(http.StatusAccepted, s)
}
