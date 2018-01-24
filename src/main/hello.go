package main 

import ("fmt"
		"github.com/rackspace/gophercloud/openstack"
		"github.com/rackspace/gophercloud")




func main(){

	opts := gophercloud.AuthOptions{
	IdentityEndpoint: "http://10.0.2.15/identity/v3",
	Username: "admin",
	Password: "Devsecret1",
	DomainID: "default",
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil{
			fmt.Println("Error: %v", err)
	} else{
			fmt.Println("Provide: ", provider)
	}


}