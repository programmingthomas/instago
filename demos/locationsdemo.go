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
	
	fmt.Println("   INSTAGO  DEMO   ")
	fmt.Println("===================")
	fmt.Println("Posts at Instagram:")
	
	//Instagram HQ
	imagesInstagram := api.LocationPosts("514276", "", "")
	
	for _, image := range imagesInstagram {
		fmt.Println("User:", image.User, "Filter:", image.Filter, "Likes:", image.Likes)
	}
	
	fmt.Println("===============================")
	fmt.Println("Locations near the Eiffel Tower")
	
	//Locations near the Eiffel Tower
	locationsInParis := api.LocationsNear(48.858844, 2.294351, 0)
	
	for _, loc := range locationsInParis {
		fmt.Println("Name:", loc.Name, "Coords:", loc.Latitude, loc.Longitude)
	}
}
