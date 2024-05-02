package main

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/synadia-io/control-plane-sdk-go/syncp"
)

const (
	systemId = "2frdT6yEyFDDu5nuYobGvMD2j6Q"
	baseUrl  = "https://cloud.synadia.com"
	pat      = "uat_tYE3gCpcteSXmehE9Ll5enl47KTQiqHph49Ba65hVk9ayW2xBhKqx4WbXa5UWXkT"
)

func main() {
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

	log.Printf("Accounts in System: %s\n", strings.Join(accountNames, ", "))

}

func handleApiError(err error) {
	// error with body
	apiErr := &syncp.GenericOpenAPIError{}

	if errors.As(err, &apiErr) {
		log.Fatal(apiErr.Error(), string(apiErr.Body()))
	}

	log.Fatal(err)
}
