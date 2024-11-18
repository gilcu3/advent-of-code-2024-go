package aoc

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func WebInput(year, day int) []byte {
	// Fetch from Advent of Code website
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	c := http.Client{Timeout: time.Duration(3) * time.Second}

	cookie, err := os.ReadFile("/home/rey/.adventofcode.session")
	if err != nil {
		logrus.Fatal(err)
	}

	session := &http.Cookie{
		Name:   "session",
		Value:  string(cookie),
		MaxAge: 0,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.Fatalf("Error with HTTP request: %s", err.Error())
	}
	req.AddCookie(session)

	resp, err := c.Do(req)
	if err != nil {
		logrus.Errorf("Error: %s", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		logrus.Error("Warning: Could not authenticate with AOC server")
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("Error: %s", err)
		return nil
	}
	return body
}
