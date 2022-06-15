package controller

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mcmuralishclint/personal_tutor/lecturer-service/config"
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

	fmt.Fprintln(res, string(userData))
}
