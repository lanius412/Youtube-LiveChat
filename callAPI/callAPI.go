package callAPI

import (
	"fmt"
	"os"
	"log"
	"time"
	"net/http"

	"google.golang.org/api/youtube/v3"
	"google.golang.org/api/googleapi/transport"

	"github.com/ry0-suke/Youtube-LiveChat/load"
	"github.com/ry0-suke/Youtube-LiveChat/convertTime"
)

var developerKey string

func Get_key(keyFlag *int) {
	developerKey = load.Read_key(keyFlag)
}

func Get_chat(liveChatId string, pageToken string, logFile *os.File) string {
	client := &http.Client {
		Transport : &transport.APIKey{Key: developerKey},
	}
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error Creating New Youtube Client(get_chat):  %w", err)
	}

	var call *youtube.LiveChatMessagesListCall
	if  pageToken == "" {
		call = service.LiveChatMessages.List(liveChatId, []string{"snippet"})
	} else {
		call = service.LiveChatMessages.List(liveChatId, []string{"snippet"}).PageToken(pageToken)
	}

	res, err := call.Do()
	if err != nil {
		time.Sleep(time.Second * 5)
		return pageToken
	}

	if res.PageInfo.TotalResults == 0 {
		time.Sleep(time.Second * 5)
		return pageToken
	}

	//chatText := res.Items[0].Snippet.TextMessageDetails.MessageText
	chatText := res.Items[0].Snippet.DisplayMessage
	fmt.Println(chatText)
	chatDateUTC := res.Items[0].Snippet.PublishedAt
	chatDateJST := convertTime.UTC2JST(chatDateUTC)
	_, err = logFile.Write([]byte(chatDateJST+" :: "+chatText+"\n"))
	if err != nil {
		log.Fatal(err)
	}
	
	return res.NextPageToken
}

func Get_chat_id(liveVideoId string) (liveChatId string) {
	client := &http.Client {
		Transport : &transport.APIKey{Key: developerKey},
	}
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error Creating New Youtube Client(get_chat_id):  %w", err)
	}
	call := service.Videos.List([]string{"liveStreamingDetails"}).Id(liveVideoId)
	res, err := call.Do()
	if err != nil {
		log.Fatalf("Error Response : %w", err)
	}
	liveChatId = res.Items[0].LiveStreamingDetails.ActiveLiveChatId

	return liveChatId
}

func Get_video_info(channelId string) (channelName string, liveVideoId string, title string, liveStartTime string) {
	client := &http.Client {
		Transport : &transport.APIKey{Key: developerKey},
	}
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error Creating New Youtube Client(get_video_id):  %w", err)
	}
	call := service.Search.List([]string{"id", "snippet"}).ChannelId(channelId).Type("video").EventType("Live").MaxResults(1)
	res, err :=  call.Do()
	if err != nil {
		log.Fatalf("Error Response : %w", err)
	}
	channelName = res.Items[0].Snippet.ChannelTitle
	liveVideoId = res.Items[0].Id.VideoId
	title = res.Items[0].Snippet.Title
	publishedAt := res.Items[0].Snippet.PublishedAt
	liveStartTime = convertTime.UTC2JST(publishedAt)
	
	return channelName, liveVideoId, title, liveStartTime
}

func IsLive(channelId string) bool {
	client := &http.Client {
		Transport : &transport.APIKey{Key: developerKey},
	}
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error Creating New Youtube Client(isLive):  %w", err)
	}
	call := service.Search.List([]string{"snippet"}).ChannelId(channelId).Type("channel").MaxResults(1)
	res, err := call.Do()
	if err != nil {
		log.Fatalf("Error Response : %w", err)
	}
	liveState := res.Items[0].Snippet.LiveBroadcastContent
	if liveState == "live" {
		return true
	} else {
		return false
	}
}
