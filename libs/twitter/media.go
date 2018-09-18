package twitter

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type MediaUploadInitResponse struct {
	MediaID          int64  `json:"media_id"`
	MediaIDString    string `json:"media_id_string"`
	ExpiresAfterSecs int    `json:"expires_after_secs"`
	MediaKey         string `json:"media_key"`
}

func NewMediaUploadInitResponse() *MediaUploadInitResponse {
	return &MediaUploadInitResponse{}
}

type MediaUploadFinalizeResponse struct {
	MediaID          int64  `json:"media_id"`
	MediaIDString    string `json:"media_id_string"`
	Size             int    `json:"size"`
	ExpiresAfterSecs int    `json:"expires_after_secs"`
	MediaKey         string `json:"media_key"`
}

func NewMediaUploadFinalizeResponse() *MediaUploadFinalizeResponse {
	return &MediaUploadFinalizeResponse{}
}

func MediaUpload(client *http.Client, imagePath string) (string, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	byte, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	base64 := base64.StdEncoding.EncodeToString(byte)

	initParams := url.Values{}
	initParams.Set("command", "INIT")
	initParams.Set("media_type", "image/png")
	initParams.Set("media_category", "dm_image")
	initParams.Set("total_bytes", fmt.Sprint(len(byte)))
	initParams.Set("shared", "true")

	// INIT
	initResponse, err := client.Post(postMediaUploadUrl, ContentTypeFormUrlEncode, strings.NewReader(initParams.Encode()))
	if err != nil {
		return "", err
	}
	defer initResponse.Body.Close()

	body, err := ioutil.ReadAll(initResponse.Body)
	if err != nil {
		return "", err
	}

	resp := NewMediaUploadInitResponse()
	if err := json.Unmarshal(body, resp); err != nil {
		return "", err
	}

	// APPEND
	for i := 0; i*chunkSize < len(base64); i++ {
		begin := i * chunkSize
		end := (i + 1) * chunkSize
		if end > len(base64) {
			end = len(base64)
		}

		appendParams := url.Values{}
		appendParams.Set("command", "APPEND")
		appendParams.Set("media_id", resp.MediaIDString)
		appendParams.Set("media_data", base64[begin:end])
		appendParams.Set("segment_index", fmt.Sprint(i))
		_, err := client.Post(postMediaUploadUrl, ContentTypeFormUrlEncode, strings.NewReader(appendParams.Encode()))
		if err != nil {
			return "", err
		}
	}

	// FINALIZE
	finalizeParams := url.Values{}
	finalizeParams.Set("command", "FINALIZE")
	finalizeParams.Set("media_id", resp.MediaIDString)
	finalizeResponse, err := client.Post(postMediaUploadUrl, ContentTypeFormUrlEncode, strings.NewReader(finalizeParams.Encode()))
	if err != nil {
		return "", err
	}
	defer finalizeResponse.Body.Close()

	finalizeBody, err := ioutil.ReadAll(finalizeResponse.Body)
	if err != nil {
		return "", err
	}

	finalizeResp := NewMediaUploadFinalizeResponse()
	if err := json.Unmarshal(finalizeBody, finalizeResp); err != nil {
		return "", err
	}

	return finalizeResp.MediaIDString, nil
}
