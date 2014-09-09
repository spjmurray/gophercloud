// +build acceptance

package openstack

import (
	"os"
	"testing"

	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud/openstack/utils"
)

func TestAuthenticatedClient(t *testing.T) {
	// Obtain credentials from the environment.
	ao, err := utils.AuthOptions()
	if err != nil {
		t.Fatalf("Unable to acquire credentials: %v", err)
	}

	// Trim out unused fields.
	ao.TenantID, ao.TenantName = "", ""

	client, err := openstack.AuthenticatedClient(ao)
	if err != nil {
		t.Fatalf("Unable to authenticate: %v", err)
	}

	if client.TokenID == "" {
		t.Errorf("No token ID assigned to the client")
	}

	t.Logf("Client successfully acquired a token: %v", client.TokenID)

	// Find the storage service in the service catalog.
	storage, err := openstack.NewStorageV1(client, os.Getenv("OS_REGION_NAME"))
	if err != nil {
		t.Errorf("Unable to locate a storage service: %v", err)
	} else {
		t.Logf("Located a storage service at endpoint: [%s]", storage.Endpoint)
	}
}
