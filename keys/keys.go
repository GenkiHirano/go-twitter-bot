package keys

import (
	"github.com/ChimeraCoder/anaconda"
)

func GetTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey("xxx")
	anaconda.SetConsumerSecret("xxx")
	api := anaconda.NewTwitterApi("xxx", "xxx")
	return api
}
