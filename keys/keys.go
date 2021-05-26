package keys

import (
	"github.com/ChimeraCoder/anaconda"
)

func GetTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey("FHQ9hHYK3nP959weYdp8WvBFF")
	anaconda.SetConsumerSecret("lwO5AGwyYqNmGVN2b0BQyIy2wfC820jHk3JWi0lJkpp6yteny9")
	api := anaconda.NewTwitterApi("AAAAAAAAAAAAAAAAAAAAABEvQAEAAAAA5Dar6zhKgUAwe4UNsqknjTIbZFs%3DMGBbwv0a8Eh27YywgmyWhiQqSsY6Cl173qNfy3D0HNKEv0HKjC")
	return api
}
