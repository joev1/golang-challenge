package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)


func main() {

	// Download image using API key
	// res, err := client.Get("https://api.pexels.com/v1/")
	// client := &http.Client{}
	// req, err := http.NewRequest("GET", "https://wallup.net/wp-content/uploads/2018/10/07/39886-animals-cats-kittens-cute-eyes-babies.jpg", nil)
	// req.Header.Set("Authorization", "563492ad6f917000010000015ce91e4af6c647cb8ed9cfce6a000100")
	// res, err := client.Do(req)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// body, err := io.ReadAll(res.Body)
	// res.Body.Close()
	// if res.StatusCode > 299 {
	// 	log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	// }
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Download image
    url := "http://i.imgur.com/m1UIjW1.jpg"
    response, e := http.Get(url)
    if e != nil {
        log.Fatal(e)
    }
    defer response.Body.Close()

	// Save image
	// To-do create folder if not exists
	file, err := os.Create("./tmp/asdf.jpg")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
    _, err = io.Copy(file, response.Body)
    if err != nil {
        log.Fatal(err)
    }

	image, _ := filepath.Abs("./tmp/asdf.jpg")
	cmd := exec.Command("JPEGView", image)
	log.Printf("Running command and waiting for it to finish...")
	err = cmd.Run()
	log.Printf("Command finished with error: %v", err)
}