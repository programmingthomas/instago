package main

import (
	"github.com/programmingthomas/instago"
	"fmt"
	"io/ioutil"
)

func main() {
	//Load the Client ID from a file called config.txt
	api := instago.InstagramAPI{}
	clientId, _ := ioutil.ReadFile("config.txt")
	api.ClientID = string(clientId)
	
	fmt.Println("INSTAGO  DEMO")
	fmt.Println("=============")
	fmt.Println("Enter a user:")
	
	var query string
	fmt.Scan(&query)
	
	//Search the users
	users := api.SearchUsers(query, 0)
	for _, user := range users {
		fmt.Println("Username:", user.Username, "Full Name:", user.FullName)
	}
	
	//Present basic inforamtion about the user
	fmt.Println("More detail on @" + users[0].Username)
	user := api.UserDetail(users[0].ID)
	fmt.Println("ID:", user.ID)
	fmt.Println("Username:", user.Username)
	fmt.Println("Full name:", user.FullName)
	fmt.Println("Bio:", user.Bio)
	fmt.Println("Website:", user.Website)
	fmt.Println("Follows:", user.TotalFollows)
	fmt.Println("Followers:", user.TotalFollowers)
	fmt.Println("Images:", user.TotalImages)
}