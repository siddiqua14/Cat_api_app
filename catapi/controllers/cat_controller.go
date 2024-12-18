package controllers

import (
	"encoding/json"
	"fmt"  // Import fmt
	"io/ioutil"
	"net/http"
	"github.com/beego/beego/v2/server/web"
)

type CatController struct {
	web.Controller
}

type CatImage struct {
	URL string `json:"url"`
}

func (c *CatController) GetCatImage() {
	// Fetch the API key and URL from configuration
	apiKey, _ := web.AppConfig.String("catapi.key")
	apiURL, _ := web.AppConfig.String("catapi.url")

	// Fetch cat images
	imageURL, err := fetchCatImage(apiURL, apiKey)
	if err != nil {
		c.Data["CatImage"] = ""
	} else {
		c.Data["CatImage"] = imageURL
	}
	c.TplName = "index.tpl"
}

func (c *CatController) GetCatImageAPI() {
	// API response for AJAX requests
	apiKey, _ := web.AppConfig.String("catapi.key")
	apiURL, _ := web.AppConfig.String("catapi.url")

	imageURL, err := fetchCatImage(apiURL, apiKey)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to fetch image"}
	} else {
		c.Data["json"] = map[string]string{"url": imageURL}
	}
	c.ServeJSON()
}

func fetchCatImage(apiURL, apiKey string) (string, error) {
	// Correct API URL with '/images/search' endpoint
	reqURL := apiURL cat+ "/images/search?size=med&mime_types=jpg&format=json&has_breeds=true&order=RANDOM&page=0&limit=1"

	// Create the request with the correct API URL
	req, _ := http.NewRequest("GET", reqURL, nil)
	req.Header.Add("x-api-key", apiKey)

	// Initialize HTTP client and perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, _ := ioutil.ReadAll(resp.Body)

	// Check if response is empty or an error occurs
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("API returned status code %d", resp.StatusCode)  // Use fmt to format error
	}

	// Parse the response body into a list of CatImage
	var result []CatImage
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	// If a valid image is returned, return its URL
	if len(result) > 0 {
		return result[0].URL, nil
	}

	// Return an empty string if no image is found
	return "", nil
}
