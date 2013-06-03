package instago

import "fmt"

//Gets basic information about a given user
//
//userID: a string representing the ID (not the username) of a given user
func (api InstagramAPI) UserDetail(userID string) User {
	params := getEmptyMap()
	result := api.DoRequest("users/" + userID, params)
	data := result.Object("data")
	return UserFromAPI(data)
}

//Query the users on Instagram and get a list of them back
//
//query: The description such as 'jack' or 'thomas' to search for
//
//max: (optional, default = 0) the number of users to return
func (api InstagramAPI) SearchUsers(query string, max int) []User {
	params := getEmptyMap()
	params["q"] = query
	if max > 0 {
		params["count"] = fmt.Sprintf("%d", max)
	}
	result := api.DoRequest("users/search", params)
	data := result.ObjectArray("data")
	users := make([]User, 0)
	for _, user := range data {
		users = append(users, UserFromAPI(user))
	}
	return users
}

//Will return an array of recently posted images by a user. Requires OAuth
//
//userId: string representing the user
//
//max: the greatest number of images to return
//
//before: (optional = "") posts before a certain ID
//
//after: (optional = "") posts after a certain ID
func (api InstagramAPI) RecentPostsByUser(userId string, max int, before, after string) []Image {
	return api.GenericImageListRequest("users/" + userId + "/media/recent", before, after, max)
}

//Gets the current user's feed (requires OAuth)
//
//before: (optional = "") posts before a certain ID
//
//after: (optional = "") posts after a certain ID
//
//max: (optional = 0) the greatest number of images to return
func (api InstagramAPI) Feed(before, after string, max int) []Image {
	return api.GenericImageListRequest("users/self/feed", before, after, max)
}

//Gets the posts like by the current user (requires OAuth)
//
//max: (optional = 0) the greatest number of posts to return
//
//before: (optional = 0) posts liked before a certain image ID
func (api InstagramAPI) Liked(max int, before string) []Image {
	return api.GenericImageListRequest("users/self/media/liked", before, "", max)
}