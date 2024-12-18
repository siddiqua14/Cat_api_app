package models

import (
    "io/ioutil"
    "net/http"
    "time"
    "encoding/json"
	
    "github.com/astaxie/beego" // Keep this import for accessing AppConfig
)

type CatData struct {
    ID        string   `json:"id"`
    URL       string   `json:"url"`
    Width     int      `json:"width"`
    Height    int      `json:"height"`
    Breeds    []Breed  `json:"breeds"`
    Favourite Favourite `json:"favourite"`
}

type Breed struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

type Favourite struct {
    ID string `json:"id"`
}

func GetCatData() (*CatData, error) {
    client := &http.Client{Timeout: 10 * time.Second}
    req, err := http.NewRequest("GET", "https://api.thecatapi.com/v1/images/search?limit=1&size=med&mime_types=jpg&format=json&has_breeds=true", nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("x-api-key", beego.AppConfig.String("apikey"))

    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var catData []CatData
    err = json.Unmarshal(body, &catData)
    if err != nil {
        return nil, err
    }

    return &catData[0], nil
}
