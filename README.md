InstaGo
=======

InstaGo is a simple and easy to use Go library that simplifies communicating with the Instagram JSON API. I currently only intend for it to use the public functions of the Instagram API that don't require user authentication, however I support some of them if you have an access token.

To use this library, clone the repo into your Go workspace and have a look at some of the demo apps. You'll need to build and install the instago package and get an Instagram API developer ID. Paste your client ID in a file called config.txt in the demos folder to run the samples.

Implemented methods:
* /tags/tagname/media/recent
* /tags/tagname
* /tags/search
* /users/user-id
* /users/search
* /media/media-id
* /media/popular
* /media/search
* /locations/location-id
* /locations/location-id/media/recent
* /locations/search/

Implemented methods that require OAuth (not demoed):
* /users/user-id/media/recent
* /users/self/feed
* /users/self/media/liked
