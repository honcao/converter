package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2017-03-30/compute"
	azcompute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
)

var (
	subscriptionID          = "3fab58ee-83c5-4147-858d-2c2a722e0e66"
	clientID                = "85115f84-ef7b-4ddb-b44d-b3a9d3b1990d"
	activeDirectoryEndpoint = "https://login.microsoftonline.com/"
	resourceManagerEndpoint = "https://management.azure.com/"
	resourceGroup           = "k111001"
	resourceGroupss         = "k111004"
	tenantID                = "72f988bf-86f1-41af-91ab-2d7cd011db47"
	clientSecret            = ""
	tokenAudience           = "https://management.azure.com/"
)

func captureJSON() {
	//captureVirtualMachine20170330()
	//captureVirtualMachine20181001()
	captureVirtualMachineScaleSet20170330()
	captureVirtualMachineScaleSet20181001()

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

	if err = ioutil.WriteFile(fileName, vmsJSON, 0644); err != nil {
		log.Fatalln(err)
	}

	diskClient := compute.NewDisksClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	diskClient.Authorizer = autorest.NewBearerAuthorizer(tokenProvider)
	diskpage, err := diskClient.ListByResourceGroup(context.Background(), resourceGroup)
	vmsJSON, _ = json.Marshal(diskpage.Values())
	diskFileName := "disksAZS.json"
	fmt.Println(string(vmsJSON))

	if err = ioutil.WriteFile(diskFileName, vmsJSON, 0644); err != nil {
		log.Fatalln(err)
	}
}

func captureVirtualMachine20181001() {

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
	computeClient := azcompute.NewVirtualMachinesClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	computeClient.Authorizer = autorest.NewBearerAuthorizer(tokenProvider)
	page, err := computeClient.List(context.Background(), resourceGroup)
	values := page.Values()
	vmsJSON, _ := json.Marshal(values)
	fileName := "vmsAZ.json"
	fmt.Println(string(vmsJSON))

	if err = ioutil.WriteFile(fileName, vmsJSON, 0644); err != nil {
		log.Fatalln(err)
	}

	diskClient := azcompute.NewDisksClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	diskClient.Authorizer = autorest.NewBearerAuthorizer(tokenProvider)
	diskpage, err := diskClient.ListByResourceGroup(context.Background(), resourceGroup)
	vmsJSON, _ = json.Marshal(diskpage.Values())
	diskFileName := "disksAZ.json"
	fmt.Println(string(vmsJSON))

	if err = ioutil.WriteFile(diskFileName, vmsJSON, 0644); err != nil {
		log.Fatalln(err)
	}
}

func captureVirtualMachineScaleSet20170330() {

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
	vmssClient := compute.NewVirtualMachineScaleSetsClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	vmssClient.Authorizer = autorest.NewBearerAuthorizer(tokenProvider)
	page, err := vmssClient.List(context.Background(), resourceGroupss)
	values := page.Values()
	vmsJSON, _ := json.Marshal(values)
	fileName := "vmssAZS.json"
	fmt.Println(string(vmsJSON))

	if err = ioutil.WriteFile(fileName, vmsJSON, 0644); err != nil {
		log.Fatalln(err)
	}

	vmssvmClient := compute.NewVirtualMachineScaleSetVMsClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	vmssvmClient.Authorizer = autorest.NewBearerAuthorizer(tokenProvider)
	for _, ss := range page.Values() {
		sspage, err := vmssvmClient.List(context.Background(), resourceGroupss, *ss.Name, "", "", "")
		vmsJSON, _ := json.Marshal(sspage.Values())
		fileName := "vmssvmAZS.json"
		fmt.Println(string(vmsJSON))

		if err = ioutil.WriteFile(fileName, vmsJSON, 0644); err != nil {
			log.Fatalln(err)
		}
	}
}

func captureVirtualMachineScaleSet20181001() {

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
	vmssClient := azcompute.NewVirtualMachineScaleSetsClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	vmssClient.Authorizer = autorest.NewBearerAuthorizer(tokenProvider)
	page, err := vmssClient.List(context.Background(), resourceGroupss)
	values := page.Values()
	vmsJSON, _ := json.Marshal(values)
	fileName := "vmssAZ.json"
	fmt.Println(string(vmsJSON))

	if err = ioutil.WriteFile(fileName, vmsJSON, 0644); err != nil {
		log.Fatalln(err)
	}

	vmssvmClient := azcompute.NewVirtualMachineScaleSetVMsClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	vmssvmClient.Authorizer = autorest.NewBearerAuthorizer(tokenProvider)
	for _, ss := range page.Values() {
		sspage, err := vmssvmClient.List(context.Background(), resourceGroupss, *ss.Name, "", "", "")
		vmsJSON, _ := json.Marshal(sspage.Values())
		fileName := "vmssvmAZ.json"
		fmt.Println(string(vmsJSON))

		if err = ioutil.WriteFile(fileName, vmsJSON, 0644); err != nil {
			log.Fatalln(err)
		}
	}
}
