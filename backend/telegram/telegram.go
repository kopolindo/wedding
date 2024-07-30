package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"strings"
	"wedding/log"
)

var (
	BOT_TOKEN string
	CHATID    int
	TGMessage *TelegramMessage
)

// Struct to hold the message data
type TelegramMessage struct {
	ChatID    int64  `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func init() {
	log.Debugf("telegram bot initialization")
	var err error
	BOT_TOKEN = os.Getenv("BOT_TOKEN")
	CHATID, err = strconv.Atoi(os.Getenv("CHATID"))
	if err != nil {
		log.Errorf("error during telegram bot initialization: %s", err.Error())
	}
	TGMessage = &TelegramMessage{
		ChatID: int64(CHATID),
	}
	log.Infof("telegram bot correct initialization")
}

func escapeDoubleQuotes(s string) string {
	return strings.ReplaceAll(s, `"`, `\"`)
}

// Helper function to escape characters
func escapeCharacter(c rune) string {
	return "\\" + string(c)
}

// Function to prepare the markup text according to the given rules
func PrepareMarkup(text string) string {
	var result strings.Builder

	// Track if we are inside pre/code entities or inline link/custom emoji definitions
	inPreOrCode := false
	inLinkOrEmoji := false

	for i := 0; i < len(text); i++ {
		c := rune(text[i])

		switch {
		// Check if we're entering or leaving pre/code entities
		case strings.HasPrefix(text[i:], "```"):
			inPreOrCode = !inPreOrCode
			result.WriteString("```")
			i += 2

		case c == '`':
			if inPreOrCode {
				result.WriteString(escapeCharacter(c))
			} else {
				result.WriteRune(c)
			}

		// Check if we're entering or leaving inline link/custom emoji definitions
		case c == '(' && !inLinkOrEmoji:
			inLinkOrEmoji = true
			result.WriteRune(c)

		case c == ')' && inLinkOrEmoji:
			inLinkOrEmoji = false
			result.WriteString(escapeCharacter(c))

		case inLinkOrEmoji && (c == ')' || c == '\\'):
			result.WriteString(escapeCharacter(c))

		case inPreOrCode && (c == '`' || c == '\\'):
			result.WriteString(escapeCharacter(c))

		// Escape special characters outside of pre/code entities and inline link/custom emoji definitions
		case !inPreOrCode && !inLinkOrEmoji && (c == '_' || c == '*' || c == '[' || c == ']' || c == '(' || c == ')' || c == '~' ||
			c == '`' || c == '>' || c == '#' || c == '+' || c == '-' || c == '=' || c == '|' || c == '{' || c == '}' ||
			c == '.' || c == '!'):
			result.WriteString(escapeCharacter(c))

		// General escaping for any character with code between 1 and 126 inclusively
		case c >= 1 && c <= 126:
			result.WriteRune(c)

		default:
			result.WriteRune(c)
		}
	}

	return result.String()
}

// Function to send message to a Telegram bot
func SendNotification(message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", BOT_TOKEN)

	log.Debugf("[telegram] sending message '%s' to bot", escapeDoubleQuotes(message))
	TGMessage.Text = message

	messageBytes, err := json.Marshal(TGMessage)
	if err != nil {
		return fmt.Errorf("[telegram] could not marshal message: %s", err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(messageBytes))
	if err != nil {
		return fmt.Errorf("[telegram] failed to create request: %s", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	//////////////////// Dump the outgoing request ////////////////////
	requestDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return fmt.Errorf("[telegram] could not dump request: %s", err.Error())
	}
	fmt.Printf("[telegram] request dump: %s", string(requestDump))
	///////////////////////////////////////////////////////////

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("[telegram] failed to send message: %s", err.Error())
	}
	defer resp.Body.Close()

	//////////////////// Dump the response ////////////////////
	responseDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return fmt.Errorf("[telegram] could not dump response: %s", err.Error())
	}
	fmt.Printf("[telegram] response dump: %s", string(responseDump))
	///////////////////////////////////////////////////////////

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("[telegram] unexpected response status: %s", resp.Status)
	}
	log.Infof("[telegram] message sent to bot")

	return nil
}
