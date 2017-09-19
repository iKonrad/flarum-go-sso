package lib

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"encoding/json"
	"strconv"
	"log"
)

const REMEMBER_ME_KEY = "flarum_remember"
const DAYS_MULTIPLIER = 60 * 60 * 24
const FORUM_API_SUFFIX = "/api"

type FlarumClient struct {
	url string
	token string
	maxAge int
}

/*
 * Creates a new client for the Flarum forum.
	Requires:
	- Forum URL: without /api suffix and without trailing slash. eg. http://myforum.example.com
	- Token: The API token that you set in your `_api_keys` table
	- maxAge: Number of days for the cookie to be valid
 */
func NewClient (url string, token string, maxAge int) *FlarumClient {

	 client := FlarumClient{
		url: url,
		token: token,
		maxAge: maxAge,
	}
	return &client
}

func (fc FlarumClient) LogIn (username string, password string) (token string, userId string, err error) {

	payload := map[string]interface{}{
		"identification": username,
		"password": password,
		"lifetime": strconv.Itoa(fc.maxAge * DAYS_MULTIPLIER),
	}

	response, err := fc.sendApiRequest("POST", "/token", "", payload)

	if err != nil{
		return "", "", errors.New("Error while logging in")
	}

	responseToken, ok := response["token"]

	if !ok {
		return "", "", errors.New("Invalid username or password")
	}

	returnedUserId, ok := response["userId"].(float64)


	if !ok {
		return "", "", errors.New("Invalid username or password")
	}

	return responseToken.(string), strconv.FormatFloat(returnedUserId, 'f', 0, 64), nil
}



func (fc FlarumClient) SignUp (profile *FlarumProfile) error {


	return errors.New("lol")

}

func (fc FlarumClient) UpdateBio(token string, userId string, bio string) error {

	var payload = map[string]interface{}{
		"data": map[string]interface{}{
			"type": "users",
			"id": userId,
			"attributes": map[string]interface{}{
				"bio": bio,
			},
		},
	}

	response, err := fc.sendApiRequest("PATCH", "/users/" + userId, token, payload)

	log.Println(response);
	return err
}

//
//func (fc FlarumClient) UpdateProfile (profile *FlarumProfile) error {
//
//}

/*
 * Sends HTTP request to forum API instance
 */
func (fc FlarumClient) sendApiRequest(method string, path string, token string, payload map[string]interface{}) (response map[string]interface{}, err error) {

	url := fc.url + FORUM_API_SUFFIX + path
	fmt.Println("URL:>", url)

	// Convert map to a JSON string
	payloadString, err := json.Marshal(payload)

	fmt.Println(bytes.NewBuffer(payloadString))

	fmt.Println(url)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadString))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Token " + token)

	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)

	var bodyMap map[string]interface{}
	json.Unmarshal(body, &bodyMap)
	fmt.Println("response Body:", bodyMap)

	return bodyMap, err
}



