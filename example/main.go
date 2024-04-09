package main

import (
	"encoding/json"
	"fmt"
	jsonql "github.com/NubeIO/jql"
	"log"
)

type Client struct {
	Name     string            `json:"name"`
	Gender   string            `json:"gender"`
	Age      int               `json:"age"`
	Hobby    *string           `json:"hobby"`
	Skills   []string          `json:"skills"`
	MetaTags map[string]string `json:"metaTags"`
	Tags     []string          `json:"tags"`
}

func main() {
	clients := []Client{
		{
			Name:     "elgs",
			Gender:   "m",
			Age:      111,
			Skills:   []string{"Golang", "Java", "C"},
			MetaTags: map[string]string{"city": "Sydney"},
			Tags:     []string{"developer", "engineer"},
		},
		{
			Name:     "enny",
			Gender:   "f",
			Age:      99,
			Hobby:    nil,
			Skills:   []string{"IC", "Electric design", "Verification"},
			MetaTags: map[string]string{"city": "Melbourne"},
			Tags:     []string{"designer", "engineer"},
		},
		{
			Name:     "sam",
			Gender:   "m",
			Age:      1,
			Hobby:    strPtr("dancing"),
			Skills:   []string{"Eating", "Sleeping", "Crawling"},
			MetaTags: map[string]string{"city": "Brisbane"},
			Tags:     []string{"baby", "dancer"},
		},
	}

	jsonBytes, err := json.Marshal(clients)
	if err != nil {
		log.Fatalf("Error marshaling clients to JSON: %v", err)
	}

	jsonString := string(jsonBytes)

	parser, err := jsonql.NewStringQuery(jsonString)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(parser.Query("name!=''")) // all

	fmt.Println(parser.Query("metaTags.city='Sydney' && tags contains 'developer'"))
	fmt.Println(parser.Query("age > 40 && ((metaTags.city='Sydney' && tags contains 'developer') || (metaTags.city='Melbourne' && tags contains 'engineer'))"))

	// This should return the first client, elgs, who is a developer in Sydney.
}

func strPtr(s string) *string {
	return &s
}
