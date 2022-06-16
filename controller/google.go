package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mcmuralishclint/personal_tutor/lecturer-service/config"
	"github.com/mcmuralishclint/personal_tutor/lecturer-service/models"
)

func GoogleLogin(res http.ResponseWriter, req *http.Request) {
	googleConfig := config.SetupConfig()
	url := googleConfig.AuthCodeURL("randomstate")
	http.Redirect(res, req, url, http.StatusSeeOther)
}

func GoogleCallback(res http.ResponseWriter, req *http.Request) {
	state := req.URL.Query()["state"][0]
	if state != "randomstate" {
		fmt.Fprintln(res, "states dont match")
		return
	}
	code := req.URL.Query()["code"][0]

	googleConfig := config.SetupConfig()

	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Fprintln(res, "Code Token Exchange Failed")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Fprintln(res, "Failed to fetch user data")
	}

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(res, "JSON parsing failed")
	}

	existing, email, err := createOrQueryUser(string(userData))
	if err != nil {
		fmt.Fprintln(res, "Unable to create or retreive user")
	}

	if existing {
		fmt.Fprintln(res, "Welcome back "+email)
	} else {
		fmt.Fprintln(res, "Greetings "+email)
	}

}

func createOrQueryUser(userData string) (bool, string, error) {
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(userData), &jsonMap)
	email := jsonMap["email"].(string)
	lecturer, err := models.FindLecturer(email)
	if err != nil {
		return false, "", err
	}
	if (lecturer == models.Lecturer{}) {
		//Create User
		lecturer := models.Lecturer{
			Email:      email,
			FullName:   jsonMap["name"].(string),
			GivenName:  jsonMap["given_name"].(string),
			FamilyName: jsonMap["family_name"].(string),
			Verified:   jsonMap["verified_email"].(bool),
			Picture:    jsonMap["picture"].(string),
		}
		_, err := models.CreateLecturer(lecturer)
		if err != nil {
			return false, email, err
		}
		return false, email, nil
	} else {
		return true, email, nil
	}
}
