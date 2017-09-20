package go_flarum

import (
	"errors"
	"strconv"
)

// LogIn creates a RememberMe token that can be added as a cookie to authenticate in Flarum forum
func (fc FlarumClient) LogIn(username string, password string) (token string, userId string, err error) {
	payload := map[string]interface{}{
		"identification": username,
		"password":       password,
		"lifetime":       strconv.Itoa(fc.maxAge * DAYS_MULTIPLIER),
	}

	response, err := fc.sendApiRequest("POST", "/token", payload)
	if err != nil {
		return "", "", errors.New(flarumErrors.ERROR_GENERAL)
	}

	responseToken, ok := response["token"]
	if !ok {
		return "", "", errors.New(flarumErrors.ERROR_INVALID_CREDENTIALS)
	}

	returnedUserId, ok := response["userId"].(float64)
	if !ok {
		return "", "", errors.New(flarumErrors.ERROR_INVALID_CREDENTIALS)
	}

	return responseToken.(string), strconv.FormatFloat(returnedUserId, 'f', 0, 64), nil
}

// SignUp creates a new account in Flarum with the provided username, email and password
// Password needs to be at least 8 characters long
func (fc FlarumClient) SignUp(username string, email string, password string) error {
	var payload = map[string]interface{}{
		"data": map[string]interface{}{
			"type": "users",
			"attributes": map[string]interface{}{
				"username": username,
				"email":    email,
				"password": password,
			},
		},
	}

	_, err := fc.sendApiRequest(request.POST, "/users", payload)
	return err
}

func (fc FlarumClient) ActivateUser(userId string) error {
	return fc.toggleUserActivation(userId, true)
}

func (fc FlarumClient) DeactivateUser(userId string) error {
	return fc.toggleUserActivation(userId, false)
}

func (fc FlarumClient) toggleUserActivation(userId string, isActivated bool) error {
	var payload = map[string]interface{}{
		"data": map[string]interface{}{
			"type": "users",
			"id":   userId,
			"attributes": map[string]interface{}{
				"isActivated": isActivated,
			},
		},
	}
	_, err := fc.sendApiRequest(request.PATCH, "/users/"+userId, payload)
	return err
}

// UpdateUserBio update user field for a given attribute key and value
func (fc FlarumClient) UpdateUserAttribute(userId string, attribute string, value string) error {
	var payload = map[string]interface{}{
		"data": map[string]interface{}{
			"type": "users",
			"id":   userId,
			"attributes": map[string]interface{}{
				attribute: value,
			},
		},
	}

	_, err := fc.sendApiRequest(request.PATCH, "/users/"+userId, payload)
	return err
}

// DeleteUser removes a user form the database. This is action cannot be undone!
func (fc FlarumClient) DeleteUser(userId string) error {
	var payload = map[string]interface{}{
		"data": map[string]interface{}{
			"type": "users",
			"id":   userId,
		},
	}

	_, err := fc.sendApiRequest(request.DELETE, "/users/"+userId, payload)
	return err
}

// GetUserByUsername fetches the user data by a given userId
func (fc FlarumClient) GetUserByUsername(username string) (user map[string]interface{}, err error) {
	response, err := fc.sendApiRequest(request.GET, "/users?filter[q]=\""+username + "\"", map[string]interface{}{})
	return response, err
}



