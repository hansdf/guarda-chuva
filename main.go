package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

const (
    apiKey  = "571b9e483fa9fb3022d7c5e06accf948" // I do not know how to hide my apikey on github, so it's just this for now
    baseURL = "http://api.openweathermap.org/data/2.5/forecast"
    lat     = "-30.033056"
    lon     = "-51.230000"
)

type WeatherResponse struct {
    List []struct {
        Dt     int64 `json:"dt"`
        Main   struct {
            Temp      float64 `json:"temp"`
            FeelsLike float64 `json:"feels_like"`
            TempMin   float64 `json:"temp_min"`
            TempMax   float64 `json:"temp_max"`
            Pressure  int     `json:"pressure"`
            Humidity  int     `json:"humidity"`
        } `json:"main"`
        Weather []struct {
            ID          int    `json:"id"`
            Main        string `json:"main"`
            Description string `json:"description"`
            Icon        string `json:"icon"`
        } `json:"weather"`
        Clouds struct {
            All int `json:"all"`
        } `json:"clouds"`
        Wind struct {
            Speed float64 `json:"speed"`
            Deg   int     `json:"deg"`
        } `json:"wind"`
        Visibility int `json:"visibility"`
        Pop        float64 `json:"pop"`
        Rain       struct {
            ThreeH float64 `json:"3h"`
        } `json:"rain,omitempty"`
        DtTxt string `json:"dt_txt"`
    } `json:"list"`
}

func main() {
    url := fmt.Sprintf("%s?lat=%s&lon=%s&appid=%s&units=metric", baseURL, lat, lon, apiKey)

    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("Failed to make API request: %v\n", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Failed to read response body: %v\n", err)
        return
    }

    var weatherResponse WeatherResponse
    err = json.Unmarshal(body, &weatherResponse)
    if err != nil {
        fmt.Printf("Failed to unmarshal JSON response: %v\n", err)
        return
    }

    willRain := false
    rainDays := make(map[string]bool)

    for _, forecast := range weatherResponse.List {
        forecastTime := time.Unix(forecast.Dt, 0)
        date := forecastTime.Format("2006-01-02")

        for _, weather := range forecast.Weather {
            if weather.Main == "Rain" {
                willRain = true
                rainDays[date] = true
            }
        }
    }

    if willRain {
        fmt.Println("Over the next 5 days, you should bring an umbrella to go out on the following:")
        for day := range rainDays {
            fmt.Println(day)
        }
    } else {
        fmt.Println("No rain is expected over the next 5 days.")
    }
}
