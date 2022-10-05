package main

import (
	"fmt"
	"strings"
)

type JRD struct {
	Subject    string
	Aliases    []string
	Properties map[string]string
	Links      []Link
}

type Link struct {
	Rel        string
	Href       string
	Titles     map[string]string
	Properties map[string]interface{}
}

func getAccountName(resource string) string {
	resource = strings.TrimPrefix(resource, "acct:")
	fmt.Println(" => ", resource)
	resource = strings.Trim(resource, " ")
	return resource
}

func getJRD(userName string) JRD {
	return data[userName]
}

func setup() {
	data = make(map[string]JRD)
	data["@test"] = JRD{
		Subject: "acct: @test",
		Aliases: []string{"pelle@pelle.com", "www.facebook.com/pelle"},
		Properties: map[string]string{
			"status":  "On vacation in Philippines",
			"job":     "Male - emeperor",
			"GPS":     "GPS 76.3434334,32.345435",
			"@shiela": "Will be going from work at 16:30",
			"@anyone": "GPS: 76.3434334,32.345435",
		},
		Links: []Link{
			{
				Rel:        "",
				Href:       "http://www.substack.com/blog/pelle",
				Titles:     map[string]string{},
				Properties: map[string]interface{}{},
			},
			{
				Rel:        "",
				Href:       "spotify://link.to.pod",
				Titles:     map[string]string{},
				Properties: map[string]interface{}{},
			},
			{
				Rel:        "",
				Href:       "insta://link.to.pod",
				Titles:     map[string]string{},
				Properties: map[string]interface{}{},
			},
		},
	}
	data["@test2"] = JRD{
		Subject: "acct: @test2",
	}
}

func createSimpleResponse() JRD {
	return JRD{
		Subject: "acct: @test2",
	}
}

func createFullResponse() JRD {
	return JRD{
		Subject: "acct: @test",
		Aliases: []string{"pelle@pelle.com", "www.facebook.com/pelle"},
		Properties: map[string]string{
			"status":  "Male - emeperor",
			"GPS":     "GPS 76.3434334,32.345435",
			"@shiela": "Will be going from work at 16:30",
			"@anyone": "GPS: 76.3434334,32.345435",
		},
		Links: []Link{
			{
				Rel:        "",
				Href:       "http://www.substack.com/blog/pelle",
				Titles:     map[string]string{},
				Properties: map[string]interface{}{},
			},
			{
				Rel:        "",
				Href:       "spotify://link.to.pod",
				Titles:     map[string]string{},
				Properties: map[string]interface{}{},
			},
			{
				Rel:        "",
				Href:       "insta://link.to.pod",
				Titles:     map[string]string{},
				Properties: map[string]interface{}{},
			},
		},
	}
}
