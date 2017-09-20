package go_flarum

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const COOKIE_REMEMBER_ME = "flarum_remember"
const COOKIE_SESSION = "flarum_session"
const DAYS_MULTIPLIER = 60 * 60 * 24
const FORUM_API_SUFFIX = "/api"

type FlarumClient struct {
	url    string
	token  string
	maxAge int
}

/*
 * Creates a new client for the Flarum forum.
	Requires:
	- Forum URL: without /api suffix and without trailing slash. eg. http://myforum.example.com
	- Token: The API token that you set in your `_api_keys` table
	- maxAge: Number of days for the cookie to be valid
*/
func NewClient(url string, token string, maxAge int) *FlarumClient {

	client := FlarumClient{
		url:    url,
		token:  token,
		maxAge: maxAge,
	}
	return &client
}

/*
 * Sends HTTP request to forum API instance
 */
func (fc FlarumClient) sendApiRequest(method string, path string, payload map[string]interface{}) (response map[string]interface{}, err error) {
	url := request.PROTOCOL_HTTP + fc.url + FORUM_API_SUFFIX + path

	// Convert map to a JSON string
	payloadString, err := json.Marshal(payload)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadString))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Token "+fc.token+"; userId=1")

	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var bodyMap map[string]interface{}
	json.Unmarshal(body, &bodyMap)
	return bodyMap, err
}
