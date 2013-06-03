package instago

//Gets all (16) recent photos with the given tag
//
//tag: The tag (don't include the # hash) that you want to fetch
//
//before: (optional - use "") find photos posted before this ID (use Image.ID)
//
//after: (optional - use "") find photos posted after this ID (use Image.ID)
func (api InstagramAPI) TagRecent(tag, before, after string) []Image {
	return api.GenericImageListRequest("tags/" + tag + "/media/recent", before, after, 0)
}

//Gets the total number of images on Instagram with a given tag
//
//tag: a string that represents the tag that you want to search for
func (api InstagramAPI) TagInfo(tag string) Tag {
	params := getEmptyMap()
	result := api.DoRequest("tags/" + tag, params)
	return tagObject(result.Object("data"))
}

//Will fetch the tag along with similar tags from Instagram so you can see the number of
//images with that tag
//
//tag: a string that represents the tag you want to search for 
func (api InstagramAPI) TagSearch(tag string) []Tag {
	params := getEmptyMap()
	params["q"] = tag
	result := api.DoRequest("tags/search", params)
	tags := make([]Tag, 0)
	for _, tag := range result.ObjectArray("data") {
		tags = append(tags, tagObject(tag))
	}
	return tags
}

//Both TagInfo and TagSearch need to create Tag objects
func tagObject(json JSON) Tag {
	return Tag{Tag: json.String("name"), MediaCount: json.Int("media_count")}
}