package main 

import ("fmt"
		"github.com/rackspace/gophercloud/openstack"
		"github.com/rackspace/gophercloud"
		//"github.com/rackspace/gophercloud/pagination"
		"github.com/rackspace/gophercloud/openstack/identity/v3/services")

/*

type ProviderClient struct {
	// IdentityBase is the base URL used for a particular provider's identity
	// service - it will be used when issuing authenticatation requests. It
	// should point to the root resource of the identity service, not a specific
	// identity version.
	IdentityBase string

	// IdentityEndpoint is the identity endpoint. This may be a specific version
	// of the identity service. If this is the case, this endpoint is used rather
	// than querying versions first.
	IdentityEndpoint string

	// TokenID is the ID of the most recently issued valid token.
	TokenID string

	// EndpointLocator describes how this provider discovers the endpoints for
	// its constituent services.
	EndpointLocator EndpointLocator

	// HTTPClient allows users to interject arbitrary http, https, or other transit behaviors.
	HTTPClient http.Client

	// UserAgent represents the User-Agent header in the HTTP request.
	UserAgent UserAgent

	// ReauthFunc is the function used to re-authenticate the user if the request
	// fails with a 401 HTTP response code. This a needed because there may be multiple
	// authentication functions for different Identity service versions.

*/


func main(){

	//uname, passwd := getLoginDetails()

	provider, err := 	getProvider("admin", "Devsecret1")

	if err != nil{
			fmt.Println("\n\t Error: %v", err)
	} else{
			fmt.Println("\n\t Provide: ", provider.TokenID)
			client:= openstack.NewIdentityV3(provider)
			fmt.Println("\n\t Client: ", client)

			// We have the option of filtering the service list. If we want the full
			// collection, leave it as an empty struct
			opts := services.ListOpts{}

			// Retrieve a pager (i.e. a paginated collection)
			pager := services.List(client, opts)

			fmt.Println(pager)
		}

}



func getProvider(username, passwd string) (provider *gophercloud.ProviderClient, err error){
	
	opts := gophercloud.AuthOptions{
	IdentityEndpoint: "http://10.0.2.15/identity/v3",
	Username: username,
	Password: passwd,
	DomainID: "default",
	}

	provider, err = openstack.AuthenticatedClient(opts)
	return provider, err

}

func getLoginDetails() (uname, passwd string){

	fmt.Print("\n\t Enter Username: ")
	fmt.Scanf("%s", &uname)
	fmt.Println("")
	fmt.Print("\t Enter Password: ")
	fmt.Scanf("%s", &passwd)
	return
}
