package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var day = 1
var year = 2023

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	sessionCookie := os.Getenv("ADVENT_SESSION_COOKIE")
	if sessionCookie == "" {
		fmt.Println("ADVENT_SESSION_COOKIE is not set in the environment file")
		return
	}

	input, err := downloadInput(sessionCookie)
	if err != nil {
		fmt.Println("Error downloading input:", err)
		return
	}

	filename := fmt.Sprintf("input.txt")
	err = saveToFile(input, filename)
	if err != nil {
		fmt.Println("Error saving to file:", err)
		return
	}
}

func downloadInput(sessionCookie string) ([]byte, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Cookie", fmt.Sprintf("session=%s", sessionCookie))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func saveToFile(content []byte, filename string) error {
	err := ioutil.WriteFile("day"+strconv.Itoa(day)+"/"+filename, content, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("Input saved to: %s\n", filename)
	return nil
}
