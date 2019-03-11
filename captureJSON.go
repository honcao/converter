package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2017-03-30/compute"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
)

var (
	subscriptionID          = "3fab58ee-83c5-4147-858d-2c2a722e0e66"
	clientID                = "85115f84-ef7b-4ddb-b44d-b3a9d3b1990d"
	activeDirectoryEndpoint = "https://login.microsoftonline.com/"
	resourceManagerEndpoint = "https://management.azure.com/"
	resourceGroup           = "k111001"
	tenantID                = "72f988bf-86f1-41af-91ab-2d7cd011db47"
	clientSecret            = "o62HTZ5hufKkX4s7BxRb1OrzMHYt3R4v68cD5yW+DTY="
	tokenAudience           = "https://management.azure.com/"
)

func captureJSON() {
	captureVirtualMachine20170330()
}

func captureVirtualMachine20170330() {

	var tokenProvider adal.OAuthTokenProvider
	oauthConfig, err := adal.NewOAuthConfig(activeDirectoryEndpoint, tenantID)
	if err != nil {
		log.Fatalln(err)
	}
	tokenProvider, err = adal.NewServicePrincipalToken(
		*oauthConfig,
		clientID,
		clientSecret,
		tokenAudience)
	computeClient := compute.NewVirtualMachinesClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	computeClient.Authorizer = autorest.NewBearerAuthorizer(tokenProvider)
	page, err := computeClient.List(context.Background(), resourceGroup)
	values := page.Values()
	vmsJSON, _ := json.Marshal(values)
	fileName := "vmsAZS.json"
	fmt.Println(string(vmsJSON))

	if err = ioutil.WriteFile(fileName, vmsJSON, os.ModeAppend); err != nil {
		log.Fatalln(err)
	}
}
