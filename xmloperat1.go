package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	file, err := os.Open("sites.xml")
	if err != nil {
		fmt.Printf("error %v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error %v", err)
		return
	}

	v := objectsites{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error %v", err)
		return
	}

	fmt.Println(v)

}

type objectsites struct {
	xmlname     xml.Name `xml: "sites"`
	version     string   `xml:"version,attr"`
	Sites       []site   `xml:"site"`
	description string   `xml:",innerxml"`
}

type site struct {
	xmlname     xml.Name `xml: "site"`
	name        string   `xml:"name"`
	description string   `xml:"description"`
	category    string   `xml:"category"`
}
