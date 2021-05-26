package text

import (
    "time"
    "fmt"
)

func ChooseTweet()string{

    TweetList :=  [31]string{
        "golangを使ってtwitter botを作ってみた" ,
        "成功してることを祈ります。"
    }

    d := time.Now().Day()
    m := time.Now().Month()

    TweetContent := TweetList[d-1]
    TweetOfToday := fmt.Sprintf("【%d月%d日】\n %s",m,d, TweetContent)
    return TweetOfToday
}