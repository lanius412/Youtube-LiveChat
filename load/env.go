package load

import (
	"embed"
	"bufio"
	"log"
)

//go:embed keys.txt
var f embed.FS

func Read_key(keyFlag *int) (developerKey string){
	file, err := f.Open("keys.txt")
	if  err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	keys := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		keys = append(keys, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	developerKey = keys[*keyFlag]

	return developerKey
}