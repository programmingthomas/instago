//Instago provides a simple library that makes it easier to interact with Instagram through
//their API directly from Go
package instago

import (
	"net/http"
	"net/url"
	"io"
	"os"
	"io/ioutil"
	"encoding/json"
	"time"
	"strconv"
	"fmt"
)

//The InstagramAPI object stores your credentials. You can obtain a ClientID from
//http://instagram.com/developer. If you want to interact directly with a user's account
//you can also obtain an AccessToken through OAuth, however this library currently doesn't
//support obtaining the ClientID. If the AccessToken is present the ClientID will be
//ignored (even if the request fails). You should create an InstagramAPI struct with
//at least one of these values
type InstagramAPI struct {
	ClientID string
	AccessToken string
}

//Represents an image response from Instagram's servers including key details about the
//image. Comments are currently not included.
type Image struct {
	Filter string
	Tags []string
	Link string
	LowResolution string
	Thumbnail string
	StandardResolution string
	User string
	UserID string
	Name string
	Caption string
	CreationTime time.Time
	ID string
	Likes int
	Comments int
	Location Location
}

//Represents a user response from Instagram's servers. This may come from an image,
//comment or directly from a user request (N.B. these kind of requests require OAuth)
type User struct {
	ID string
	Username string
	FullName string
	ProfilePicture string
	Bio string
	Website string
	TotalImages int
	TotalFollows int
	TotalFollowers int
}

//Represents a tag and the total number of images with that tag
type Tag struct {
	Tag string
	MediaCount int
}

//As well as being able to look near a specific longitude/latitude, you can also look at
//a specific location, such as a bar, museum, company, etc. This type represents the
//responses from Instagram's servers.
type Location struct {
	ID string
	Name string
	Latitude float64
	Longitude float64
}

//This will does all GET requests (all Instagram API requests that do not require 
//authentication are GET requests anyway). It returns the JSON object in case of success
//or an empty object in case of failure

//endpoint: The api request that you want to do on Instagram

//params: The parameters you may want to add
func (api InstagramAPI) DoRequest(endpoint string, params map[string]string) JSON {
	var contents []byte
	
	fullURL := api.GetURLForRequest(endpoint, params)
	resp, err := http.Get(fullURL)
	if err != nil {
		contents = []byte("{}")
	}
	defer resp.Body.Close()
	contents, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		contents = []byte("{}")
	}
	//fmt.Println(string(contents))
	var jsonResponse JSON
	json.Unmarshal(contents, &jsonResponse)
	
	return jsonResponse
}

//This function will build the request URL so that you can add extra parameters to
//requests.

//endpoint: The API request that you are planning on doing; such as tags/{x}/media/recent

//params: A map of the extra parameters (aside from client_id) that you want to add to
//the query
func (api InstagramAPI) GetURLForRequest(endpoint string, params map[string]string) string {
	u, err := url.Parse("https://api.instagram.com/v1/" + endpoint)
	if err != nil {
		return ""
	}
	q := u.Query()
	//If you have an AccessToken (from OAuth), use it
	if api.AccessToken != "" {
		q.Set("access_token", api.AccessToken)
	} else {
		q.Set("client_id", api.ClientID)
	}
	for key, value := range params {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

//This will take API a JSON object that includes the details for an image and puts it into
//the Go data structure for Images.

//data: a JSON object that represents an image
func ImageFromAPI(data JSON) Image {
	var image Image
	
	//Basic information on the image
	image.Filter = data.String("filter")
	image.Tags = data.StringArray("tags")
	image.Link = data.String("link")
	image.ID = data.String("id")
	image.Likes = data.Object("likes").Int("count")
	image.Comments = data.Object("comments").Int("count")
	
	//Image caption
	caption := data.Object("caption")
	image.Caption = caption.String("text")
	
	//Creation time
	t, _ := strconv.ParseInt(data.String("created_time"), 0, 0)
	image.CreationTime = time.Unix(t, 0)
	
	//User detail
	user := data.Object("user")
	image.User = user.String("username")
	image.Name = user.String("full_name")
	image.UserID = user.String("id")
	
	images := data.Object("images")
	
	lowRes := images.Object("low_resolution")
	image.LowResolution = lowRes.String("url")
	
	thumbnail := images.Object("thumbnail")
	image.Thumbnail = thumbnail.String("url")
	
	standardRes := images.Object("standard_resolution")
	image.StandardResolution = standardRes.String("url")
	
	location := data.Object("location")
	image.Location = LocationFromAPI(location)
	
	return image
}

//Takes a generic location API JSON response and returns a Location
func LocationFromAPI(location JSON) Location {
	loc := Location{}
	loc.Longitude = location.Float("longitude")
	loc.Latitude = location.Float("latitude")
	loc.Name = location.String("name")
	loc.ID = location.String("id")
	return loc
}

//This will take an API JSON response that includes some user detail and return a more
//usable Go User object
func UserFromAPI(data JSON) User {
	user := User{}
	user.ID = data.String("id")
	user.Username = data.String("username")
	user.FullName = data.String("full_name")
	//Oddly some responses include full_name, but others split it up...
	if user.FullName == "" {
		user.FullName = data.String("first_name") + " " + data.String("last_name")
	}
	user.ProfilePicture = data.String("profile_picture")
	user.Bio = data.String("bio")
	user.Website = data.String("website")
	user.TotalImages = data.Object("counts").Int("media")
	user.TotalFollows = data.Object("counts").Int("follows")
	user.TotalFollowers = data.Object("counts").Int("followed_by")
	return user
}

//Many queries to Instagram's API simply return a list of images (tag, user, location, etc)
//so this function handles the request to simplify things a little. Note that Intago
//functions provide wrappers around this function so you need not call it, however it is
//exported in case Instagram adds to their API in the future and you want to add to this
//library

//endPoint: The API endpoint, such as /tags/tag/media/recent

//before: (optional) Search for images before this image ID

//after: (optional) Search for images after this image ID

//max: (optional) The great number of images to return (there is an imposed limit on this)
func (api InstagramAPI) GenericImageListRequest(endPoint, before, after string, max int) []Image {
	params := getEmptyMap()
	if max > 0 {
		params["count"] = fmt.Sprintf("%d", max)
	}
	if before != "" {
		params["max_id"] = before
	}
	if after != "" {
		params["min_id"] = after
	}
	results := api.DoRequest(endPoint, params)
	data := results.ObjectArray("data")
	images := make([]Image, 0)
	for _, image := range data {
		images = append(images, ImageFromAPI(image))
	}
	return images
}

//Download a file from the given URL and save it to the given file
//Note that the Instagram API encourages you to take into account the IP of Instagram
//users, so you shouldn't download user's images with this
func Download(url, saveFile string) {
	out, err := os.Create(saveFile)
	if err != nil {
		return
	}
	defer out.Close()
	
	resp, err := http.Get(url)
	if err == nil {
		defer resp.Body.Close()
		io.Copy(out, resp.Body)
	}
}

//Most of the API functions have to get make a map[string] string for parameters so this
//utlility function saves them all having to do it
func getEmptyMap() map[string]string{
	return make(map[string]string, 0)
}