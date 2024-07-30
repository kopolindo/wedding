package catapush

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"wedding/log"
)

type MessageJSON struct {
	MobileAppID      int          `json:"mobileAppId"`
	NotificationText string       `json:"notificationText"`
	Text             string       `json:"text"`
	Recipients       []Identifier `json:"recipients"`
}

type Identifier struct {
	ID string `json:"identifier"`
}

var (
	CATAPUSH_MOBILE_APP_ID_STRING string
	CATAPUSH_MOBILE_APP_ID        int
	CATAPUSH_IDENTIFIER           string
	CATAPUSH_API_KEY              string
	Payload                       *MessageJSON
)

func init() {
	log.Debugf("catapush initialization")
	var err error
	CATAPUSH_MOBILE_APP_ID_STRING = os.Getenv("CATAPUSH_MOBILE_APP_ID")
	log.Debugf("CATAPUSH_MOBILE_APP_ID_STRING: %s", CATAPUSH_MOBILE_APP_ID_STRING)
	CATAPUSH_MOBILE_APP_ID, err = strconv.Atoi(CATAPUSH_MOBILE_APP_ID_STRING)
	if err != nil {
		log.Errorf("error during catapush initialization: %s", err.Error())
	}
	CATAPUSH_IDENTIFIER = os.Getenv("CATAPUSH_IDENTIFIER")
	log.Debugf("CATAPUSH_IDENTIFIER: %s", CATAPUSH_IDENTIFIER)
	identifier := &Identifier{
		ID: CATAPUSH_IDENTIFIER,
	}
	var recipients []Identifier
	recipients = append(recipients, *identifier)
	CATAPUSH_API_KEY = os.Getenv("CATAPUSH_API_KEY")
	log.Debugf("CATAPUSH_API_KEY: %s", CATAPUSH_API_KEY)
	Payload = &MessageJSON{
		MobileAppID:      CATAPUSH_MOBILE_APP_ID,
		NotificationText: "wedding app monitoring",
		Text:             "",
		Recipients:       recipients,
	}
	log.Infof("catapush correct initialization")
}

func SendNotification(message string) error {

	url := "https://api.catapush.com/1/messages"

	Payload.Text = message

	marshalledPayload, err := json.Marshal(Payload)
	if err != nil {
		return fmt.Errorf("error marshaling catapush payload. errors: %v", err)
	}

	stringPayload := string(marshalledPayload)
	log.Debugf("catapush string payload: %s", stringPayload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(stringPayload)))
	if err != nil {
		return fmt.Errorf("error creating request. errors: %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", CATAPUSH_API_KEY))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request. errors: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	log.Debugf("catapush response: %v", resp)
	log.Debugf("catapush response body: %s", string(body))
	return nil
}
