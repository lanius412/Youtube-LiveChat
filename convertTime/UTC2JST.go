package convertTime

import (
	"time"
	"log"
)

func UTC2JST(publishedAt string) (liveStartJST string) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	liveStartTimeUTC, err := time.ParseInLocation("2006-01-02T15:04:05", publishedAt[:19], loc)
	if err != nil {
		log.Fatal(err)
	}
	liveStartJST = liveStartTimeUTC.Add(9 * time.Hour).Format("Mon Jan _2 15:04:05 2006")
	return liveStartJST
}