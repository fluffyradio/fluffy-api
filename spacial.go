package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"fmt"

	"github.com/labstack/echo"
)

// FluffyAPI is a string representing the Spacial Audio URL for interacting with the control system
const FluffyAPI = "http://widgets-proxy.cloudapp.net/webapi/station/65655/"

type (
	// Song represents an instance of a song from spacial
	Song struct {
		ID       string `json:"id"`
		Artist   string `json:"artist"`
		Title    string `json:"title"`
		Album    string `json:"album"`
		AlbumArt string `json:"album_art_url"`
	}
)

// songs returns an array of songs in the library
func songs(c echo.Context) error {
	return c.JSON(http.StatusOK, "Songs")
}

// currentSong returns the currently playing song
func currentSong(c echo.Context) error {

	// Get current song
	res, err := http.Get(FluffyAPI + "/history/npe?format=json&token=" + spacialToken)
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
	fmt.Println(data)

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

	// Return the result
	return c.JSON(http.StatusOK, r)
}
