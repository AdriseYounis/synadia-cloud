package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/synadia-io/control-plane-sdk-go/syncp"
)

const (
	systemId = "2frdT6yEyFDDu5nuYobGvMD2j6Q"
	baseUrl  = "https://cloud.synadia.com"
	pat      = "uat_tYE3gCpcteSXmehE9Ll5enl47KTQiqHph49Ba65hVk9ayW2xBhKqx4WbXa5UWXkT"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)

	myRouter.HandleFunc("/accountNames", getAccountNames)

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func getAccountNames(w http.ResponseWriter, r *http.Request) {
	client := syncp.NewAPIClient(syncp.NewConfiguration())

	ctx := context.WithValue(context.Background(), syncp.ContextServerVariables, map[string]string{
		"baseUrl": baseUrl,
	})

	ctx = context.WithValue(ctx, syncp.ContextAccessToken, pat)

	accountList, _, err := client.SystemAPI.ListAccounts(ctx, systemId).Execute()

	if err != nil {
		handleApiError(err)
	}

	var accountNames []string
	for _, account := range accountList.Items {
		accountNames = append(accountNames, account.Name)
	}

	json.NewEncoder(w).Encode(accountNames)

	log.Printf("Accounts in System: %s\n", strings.Join(accountNames, ", "))
}

func main() {
	handleRequests()
}

func handleApiError(err error) {
	// error with body
	apiErr := &syncp.GenericOpenAPIError{}

	if errors.As(err, &apiErr) {
		log.Fatal(apiErr.Error(), string(apiErr.Body()))
	}

	log.Fatal(err)
}
