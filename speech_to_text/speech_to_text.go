package speech_to_text

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/tidwall/gjson"
	"github.com/valyala/fasthttp"
)

const (
	azureAPIKey = "12e94df29f0b439782741f88d5bef9a8"
	azureRegion = "westus"
)

func SpeechToText(audioFile string) string {
	// Read audio file
	audioData, err := ioutil.ReadFile(audioFile)
	if err != nil {
		fmt.Println("Error reading audio file:", err)
		return ""
	}

	// Encode audio data to base64
	audioDataB64 := base64.StdEncoding.EncodeToString(audioData)

	// Prepare the API request
	requestBody := fmt.Sprintf(`{
		"audio": {
			"bytes": "%s"
		},
		"config": {
			"language": "en-US",
			"format": "detailed"
		}
	}`, audioDataB64)

	// Prepare the API URL
	apiURL := fmt.Sprintf("https://westus.api.cognitive.microsoft.com/sts/v1.0/issuetoken", azureRegion)

	// Create a new HTTP request
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(apiURL)
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Ocp-Apim-Subscription-Key", azureAPIKey)
	req.SetBodyString(requestBody)

	// Send the HTTP request
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	err = fasthttp.Do(req, resp)
	if err != nil {
		fmt.Println("Error making API request:", err)
		return ""
	}

	// Parse the API response
	responseBody := resp.Body()
	transcription := gjson.Get(string(responseBody), "NBest.0.Display").String()
	return transcription
}
