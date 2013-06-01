package main

import (
	"instago"
	"fmt"
	"io/ioutil"
)

func main() {
	//Load the Client ID from a file called config.txt
	api := instago.InstagramAPI{}
	clientId, _ := ioutil.ReadFile("config.txt")
	api.ClientID = string(clientId)
	
	fmt.Println("INSTAGO DEMO")
	fmt.Println("============")
	fmt.Println("Enter a tag:")
	
	var tag string
	fmt.Scan(&tag)
	
	tagInfo := api.TagInfo(tag)
	fmt.Println("Tag: ", tagInfo.Tag, "Total: ", tagInfo.MediaCount)
	
	images := api.TagRecent(tag, "", "")
	for _, img := range images {
		fmt.Println(img.User, img.Filter)
	}
	
	fmt.Println("Similar tags")
	tags := api.TagSearch(tag)
	for _, tag := range tags {
		fmt.Println("Tag: ", tag.Tag, "Total: ", tag.MediaCount)
	}
}