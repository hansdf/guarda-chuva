package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    // "os"
)

const apiKey = "placeholder" // since I don't know how to hide the api key on github D:

type WeatherResponse struct {
    Current struct {
        Temp      float64 `json:"temp"`
        FeelsLike float64 `json:"feels_like"`
        Pressure  int     `json:"pressure"`
        Humidity  int     `json:"humidity"`
        WindSpeed float64 `json:"wind_speed"`
        Weather   []struct {
            Main        string `json:"main"`
            Description string `json:"description"`
        } `json:"weather"`
    } `json:"current"`
}

func fetchWeather(lat, lon string) (*WeatherResponse, error) {
    url := fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=%s&lon=%s&exclude=minutely,hourly,daily,alerts&appid=%s&units=metric", lat, lon, apiKey)

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var weatherResponse WeatherResponse
    if err := json.Unmarshal(body, &weatherResponse); err != nil {
        return nil, err
    }

    return &weatherResponse, nil
}

func main() {
    lat := "33.44" 
    lon := "-94.04" 

    weather, err := fetchWeather(lat, lon)
    if err != nil {
        log.Fatalf("Error fetching weather: %v", err)
    }

    fmt.Printf("Temperature: %.2f°C\n", weather.Current.Temp)
    fmt.Printf("Feels Like: %.2f°C\n", weather.Current.FeelsLike)
    fmt.Printf("Pressure: %d hPa\n", weather.Current.Pressure)
    fmt.Printf("Humidity: %d%%\n", weather.Current.Humidity)
    fmt.Printf("Wind Speed: %.2f m/s\n", weather.Current.WindSpeed)
    if len(weather.Current.Weather) > 0 {
        fmt.Printf("Weather: %s (%s)\n", weather.Current.Weather[0].Main, weather.Current.Weather[0].Description)
    }
}
