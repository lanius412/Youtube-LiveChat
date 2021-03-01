package setup

import (
	"os"
	"log"
)

func Create_log_file(channelName string, liveVideoId string) (logFile *os.File){
	cDIr, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	logDir := cDIr+"/chat_log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.Mkdir(logDir, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
	channelDir := logDir+"/"+channelName
	if _, err := os.Stat(channelDir); os.IsNotExist(err) {
		err = os.Mkdir(channelDir, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
	logFile, err = os.Create(channelDir+"/"+liveVideoId+".txt")
	if err != nil {
		log.Fatalf("Error Can't Create log file: %w", err)
	}

	return logFile
}