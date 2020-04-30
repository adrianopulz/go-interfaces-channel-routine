package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type person interface {
	writeData() string
}

type user struct {
	name      string
	birthDate string
	phone     int
}

type candidate struct {
	user
	skills    []string
	education string
	links     []string
}

func main() {
	u := user{
		name:      "Adriano Pulz",
		birthDate: "10/07/1987",
		phone:     51989890087,
	}

	c := candidate{
		user: u,
		skills: []string{
			"PHP",
			"GO",
			"ReactJs",
			"CSS",
			"SASS",
			"HTML",
			"GIT",
		},
		education: "Bachelor in Information Systems",
		links: []string{
			"https://www.linkedin.com/in/adriano-pulz",
			"https://www.drupal.org/u/adrianopulz",
			"https://github.com/adrianopulz",
			"http://www.adrianopulz.com.br",
		},
	}

	c.writeData()
}

// User Struct writeDate implementation
func (u user) writeData() {
	fmt.Println("Name:", u.name)
	fmt.Println("Birth date:", u.birthDate)
	fmt.Println("Phone:", u.phone)
}

// Candidate Struct writeDate implementation
func (c candidate) writeData() {
	c.user.writeData()
	fmt.Println("Skills:", strings.Join(c.skills, ", "))
	fmt.Println("Education:", c.education)
	fmt.Println("Links:")

	// Using a channel and routine to check all the URLs status at the same time
	channel := make(chan string)
	for _, url := range c.links {
		go checkURLStatus(url, channel)
	}

	for i := 0; i < len(c.links); i++ {
		fmt.Println(<-channel)
	}
}

// Function that will load and check each URLs usign GO Routine
func checkURLStatus(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		c <- url + " [down]"
		return
	}

	if resp.StatusCode != 200 {
		c <- url + " [" + strconv.Itoa(resp.StatusCode) + "]"
		return
	}

	c <- url + " [up]"
}
