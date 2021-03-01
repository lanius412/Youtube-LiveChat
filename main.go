package main

import (
	"flag"
	"os"
	"fmt"
	"log"
	"time"

	"LiveChat_yt/callAPI"
	"LiveChat_yt/setup"
)
//"UCrWkFUu3DK5XpYBUPDC5uLA"//ぽっぽ
//"UChLfthKoUV502J7gU9STArg" //しぶはる
//"UCIu-aUArYq_H84dBpCAokMA" //れいど
//"UCXXEOK4erP8E3cbszN22bEQ" //福岡放送
//"UC2GuoutVyegg6PUK88lLpjw" //2bro

func main() {
	var (
		keyFlag = flag.Int("k", 0, "which developerKey use")
		intervalFlag = flag.Int("i", 1, "access time interval[seconds]")
		channelFlag = flag.String("c", "UCgdHxnHSXvcAi4PaMIY1Ltg", "Channel ID")
	)
	flag.Parse()

	callAPI.Get_key(keyFlag)
	var channelId = *channelFlag

	if callAPI.IsLive(channelId) == false {
		fmt.Println("No livestreaming in this channel")
		os.Exit(0)
	}

	channelName, liveVideoId, liveTitle, liveStartTime := callAPI.Get_video_info(channelId)
	fmt.Println(liveTitle)
	liveChatId := callAPI.Get_chat_id(liveVideoId)

	logFile := setup.Create_log_file(channelName, liveVideoId)
	defer logFile.Close()

	_, err := logFile.Write([]byte(liveTitle+" ["+liveStartTime+" ~] -chat log-\n"))
	if err != nil {
		log.Fatal(err)
	}

	//var countGET int
	var nextPageToken string
	for {
		//countGET++
		//fmt.Println(countGET)
		nextPageToken = callAPI.Get_chat(liveChatId, nextPageToken, logFile)
		time.Sleep(time.Second * time.Duration(*intervalFlag))
	}
}