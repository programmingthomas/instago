//Mirrors the functions described on http://instagram.com/developer/endpoints/locations
package instago

import "fmt"

//Gets basic information such as name and coordinates for a location
//locationId: The id of a location to lookup
func (api InstagramAPI) Location(locationId string) Location {
	params := getEmptyMap()
	response := api.DoRequest("locations/" + locationId, params)
	return LocationFromAPI(response.Object("data"))
}

//Gets media posted from that location
//locationId: The id of the location
//beforePost: (optional = "") posts before this ID
//afterPost: (optional = "") posts after this ID
func (api InstagramAPI) LocationPosts(locationId, beforePost, afterPost string) []Image {
	return api.GenericImageListRequest("locations/" + locationId + "/media/recent", beforePost, afterPost, 0)
}

//Gets a list of locations near a give latitude/longitude within a certain distance
//lat: The latitude to search near
//long: The longitude to search near
//distance: (optional = 0) The number of meters to search within
func (api InstagramAPI) LocationsNear(lat, long, distance float64) []Location {
	params := getEmptyMap()
	if distance > 0 {
		params["distance"] = fmt.Sprintf("%f", distance)
	}
	params["lat"] = fmt.Sprintf("%f", lat)
	params["lng"] = fmt.Sprintf("%f", long)
	results := api.DoRequest("locations/search", params)
	data := results.ObjectArray("data")
	locations := make([]Location, 0)
	for _, loc := range data {
		locations = append(locations, LocationFromAPI(loc))
	}
	return locations
}