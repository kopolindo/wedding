package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"wedding/log"
)

var (
	BOT_TOKEN       string
	CHATID          int
	telegramMessage *TelegramMessage
)

// Struct to hold the message data
type TelegramMessage struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func init() {
	log.Debugf("telegram bot initialization")
	var err error
	BOT_TOKEN = os.Getenv("BOT_TOKEN")
	CHATID, err = strconv.Atoi(os.Getenv("CHATID"))
	if err != nil {
		log.Errorf("error during telegram bot initialization: %s", err.Error())
	}
	telegramMessage = &TelegramMessage{
		ChatID: int64(CHATID),
	}
	log.Infof("telegram bot correct initialization")
}

// Function to send message to a Telegram bot
func SendNotification(message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", BOT_TOKEN)

	log.Debugf("[telegram] sending message '%s' to bot", message)
	telegramMessage.Text = message

	messageBytes, err := json.Marshal(telegramMessage)
	if err != nil {
		return fmt.Errorf("[telegram] could not marshal message: %s", err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(messageBytes))
	if err != nil {
		return fmt.Errorf("[telegram] failed to create request: %s", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	// Dump the outgoing request
	requestDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return fmt.Errorf("[telegram] could not dump request: %s", err.Error())
	}
	fmt.Printf("[telegram] request dump: %s", string(requestDump))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("[telegram] failed to send message: %s", err.Error())
	}
	defer resp.Body.Close()

	// Dump the response
	responseDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return fmt.Errorf("[telegram] could not dump response: %s", err.Error())
	}
	fmt.Printf("[telegram] response dump: %s", string(responseDump))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("[telegram] unexpected response status: %s", resp.Status)
	}
	log.Infof("[telegram] message sent to bot")

	return nil
}
