package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    fmt.Println("Starting the application...")
    response, err := http.Get("https://data.wa.gov/resource/2tkm-ssw6.geojson?%24where=within_circle(location,%2047.59,%20-122.33,%20250)")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
    fmt.Println("Terminating the application...")	
}