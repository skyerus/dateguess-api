package logger

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

// DiscordWebhook ...
type DiscordWebhook struct {
	Content string `json:"content"`
}

// Log ...
func Log(err error) {
	if os.Getenv("ENV") == "dev" {
		fmt.Println(err)
		return
	}
	var b bytes.Buffer
	log.SetOutput(&b)
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println(err)
	sendToDiscord(&DiscordWebhook{Content: b.String()})
}

func sendToDiscord(discWebhook *DiscordWebhook) {
	client := &http.Client{}
	byteData, _ := json.Marshal(discWebhook)

	request, _ := http.NewRequest("POST", os.Getenv("DISCORD_WEBHOOK"), bytes.NewReader(byteData))
	request.Header.Add("Content-Type", "application/json")
	_, erro := client.Do(request)
	if erro != nil {
		fmt.Println(erro)
	}
}

// LogStruct ...
func LogStruct(s interface{}, prefix string) {
	payload, err := json.Marshal(s)
	if err != nil {
		go Log(err)
	}
	fmt.Println(errors.New(prefix + string(payload)))
	Log(errors.New(prefix + string(payload)))
}
