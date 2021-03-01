package main

import (
	"os"
	"bufio"
	"log"
	"fmt"

	"strings"
	"strconv"

)

func main() {
	file, err := os.Open("mGmgMKFFiJQ_chatlog.txt")
	if  err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := []string{}
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var count = 0
	for _, msg := range data {
		if strings.Contains(msg, "かわいい") {
			count++
		}
		if count >= 2 {
			idx := strings.Index(msg, "T")
			hms := msg[idx+1:idx+9]
			//fmt.Println(hms)
			hInt, err := strconv.Atoi(hms[:2]) //hours at UTC
			if err != nil {
				log.Fatal(err)
			}
			hStr :=  strconv.Itoa(hInt+9) //hours at JST
			fmt.Println(hStr+hms[2:])
			count = 0
		}
	}
}