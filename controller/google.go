package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mcmuralishclint/personal_tutor/lecturer-service/config"
	"github.com/mcmuralishclint/personal_tutor/lecturer-service/middleware"
	"github.com/mcmuralishclint/personal_tutor/lecturer-service/models"
)

func GoogleLogin(res http.ResponseWriter, req *http.Request) {
	googleConfig := config.SetupConfig()
	url := googleConfig.AuthCodeURL("randomstate")
	http.Redirect(res, req, url, http.StatusSeeOther)
}

func GoogleCallback(res http.ResponseWriter, req *http.Request) {
	response := make(map[string]string)

	state := req.URL.Query()["state"][0]
	if state != "randomstate" {
		response["email"] = ""
		json.NewEncoder(res).Encode(response)
		fmt.Fprintln(res, "states dont match")
		return
	}
	code := req.URL.Query()["code"][0]

	googleConfig := config.SetupConfig()

	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		response["email"] = ""
		json.NewEncoder(res).Encode(response)
		fmt.Fprintln(res, "Code Token Exchange Failed")
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		response["email"] = ""
		json.NewEncoder(res).Encode(response)
		fmt.Fprintln(res, "Failed to fetch user data")
		return
	}

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response["email"] = ""
		json.NewEncoder(res).Encode(response)
		fmt.Fprintln(res, "JSON parsing failed")
		return
	}

	existing, email, err := createOrQueryUser(string(userData))
	if err != nil {
		response["email"] = ""
		json.NewEncoder(res).Encode(response)
		fmt.Fprintln(res, "Unable to create or retreive user")
		return
	}

	jwtToken, _ := middleware.GenerateJWT(email)

	if existing {
		fmt.Println("Welcome back " + email)
		res.WriteHeader(http.StatusOK)
	} else {
		fmt.Println("Greetings " + email)
		res.WriteHeader(http.StatusCreated)
	}

	response["email"] = email
	response["token"] = jwtToken
	json.NewEncoder(res).Encode(response)

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
