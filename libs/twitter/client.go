package twitter

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"

	"github.com/mrjones/oauth"
)

const (
	baseUrl              = "https://api.twitter.com/1.1/"
	PostDirectMessageUrl = baseUrl + "direct_messages/events/new.json"
	baseUrlUploader      = "https://upload.twitter.com/1.1/"
	postMediaUploadUrl   = baseUrlUploader + "media/upload.json"
	chunkSize            = 1024 * 500
	RetryMetaData        = "100"

	ContentTypeFormUrlEncode = "application/x-www-form-urlencoded"
	ContentTypeJson          = "application/json"
)

var (
	ck = os.Getenv("CONSUMER_KEY")
	cs = os.Getenv("CONSUMER_SECRET")
	at = os.Getenv("ACCESS_TOKEN")
	as = os.Getenv("ACCESS_TOKEN_SECRET")
)

func CreateTwitterClient() (*http.Client, error) {
	c := oauth.NewConsumer(
		ck,
		cs,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})
	c.Debug(true)

	t := oauth.AccessToken{
		Token:  at,
		Secret: as,
	}

	client, err := c.MakeHttpClient(&t)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func CreateCRCToken(crcToken string) string {
	mac := hmac.New(sha256.New, []byte(cs))
	mac.Write([]byte(crcToken))
	return "sha256=" + base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func PostDM(client *http.Client, request interface{}) error {
	byte, err := json.Marshal(request)
	if err != nil {
		return err
	}

	response, err := client.Post(PostDirectMessageUrl, ContentTypeJson, bytes.NewBuffer(byte))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}
