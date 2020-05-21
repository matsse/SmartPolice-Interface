package Tests

import (
	ttnsdk "github.com/TheThingsNetwork/go-app-sdk"
	ttnlog "github.com/TheThingsNetwork/go-utils/log"
	"github.com/TheThingsNetwork/go-utils/log/apex"
	"github.com/TheThingsNetwork/ttn/core/types"
	"testing"
)



type TTN_Info struct {
	
	AppID               string                          `json:"appID"`
	DevID               string                          `json:"devID"`
	AppAccessKey        string                          `json:"appAccesskey"`
	AppEUID             []byte                          `json:"appEUID"`
	DevEUID             []byte                          `json:"devEUID"`
	config              ttnsdk.ClientConfig             // Unexported fields are not encoded
	client              ttnsdk.Client                   // Unexported fields are not encoded
	dev                 *ttnsdk.Device                  // Unexported fields are not encoded
	pubsub              ttnsdk.ApplicationPubSub        // Unexported fields are not encoded
	uplink              <- chan *types.UplinkMessage    // Unexported fields are not encoded
	myNewDevicePubSub   ttnsdk.DevicePubSub             // Unexported fields are not encoded
}




func TestConnection(t *testing.T) {
	T := TTN_Info{}
	
	log := apex.Stdout()
	ttnlog.Set(log)
	
	T.AppID = "go_home"
	T.AppAccessKey = "ttn-account-v2.nIxlEQvAfmYB5W8OFFtj_mghjh8HnLcX7ebbKV9QgCs"
	
	T.config =ttnsdk.NewCommunityConfig(T.AppID)
	T.config.ClientVersion = "2.0.5"
	T.client = T.config.NewClient(T.AppID, T.AppAccessKey)
	defer T.client.Close()
	
	if _, err := T.client.ManageApplication() ;  err != nil {
		log.WithError(err).Fatalf("%s: could not connect to TTN server", T.AppID)
	}


}





