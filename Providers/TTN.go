package Providers

import (
	"encoding/json"
	"fmt"
	ttnsdk "github.com/TheThingsNetwork/go-app-sdk"
	ttnlog "github.com/TheThingsNetwork/go-utils/log"
	"github.com/TheThingsNetwork/go-utils/log/apex"
	"github.com/TheThingsNetwork/ttn/core/types"
	"github.com/apex/log"
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




func (T *TTN_Info) NewDevice() {
	
	log := apex.Stdout()
	ttnlog.Set(log)
	
	
	setAppID:
	fmt.Println("What is the TTN application ID (appID)? ")
	fmt.Scan(&T.AppID)
	
	if T.AppID == "" {
		goto setAppID
	}
	
	setAccKey:
	fmt.Println("What is the TTN access key? ")
	fmt.Scan(&T.AppAccessKey)
	
	if T.AppAccessKey == "" {
	goto setAccKey
	}
	
	
	T.config =ttnsdk.NewCommunityConfig(T.AppID)
	T.config.ClientVersion = "2.0.5"
	T.client = T.config.NewClient(T.AppID, T.AppAccessKey)
	//defer T.client.Close()
	
	devices, err := T.client.ManageDevices()
	if err != nil {
		log.WithError(err).Fatalf("%s: could not read CA certificate file", T.AppID)
	}
	T.dev = new(ttnsdk.Device)
	T.dev.AppID = T.AppID
	
	
	fmt.Println("What is your device name?")
	fmt.Scan(&T.dev.DevID)
	T.DevID = T.dev.DevID
	
	fmt.Println("What is the description of this device?")
	fmt.Scan(&T.dev.Description)
	
	
	fmt.Println("Please enter the Application APPEui in comma separted hex values (0xFF)")
	test := []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	fmt.Scan(  &test[0], &test[1], &test[2], &test[3], &test[4], &test[5], &test[6], &test[7])
	T.dev.AppEUI = types.AppEUI{ test[0], test[1], test[2], test[3], test[4], test[5], test[6], test[7]}
	T.AppEUID = test
	
	
	fmt.Println("Please enter the device APPEui in space separated hex values (0xFF)")
	test2 := []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	fmt.Scan(  &test2[0], &test2[1], &test2[2], &test2[3], &test2[4], &test2[5], &test2[6], &test2[7])
	T.dev.DevEUI = types.DevEUI{ test2[0], test2[1], test2[2], test2[3], test2[4], test2[5], test2[6], test2[7]}
	T.DevEUID = test2
	
	T.dev, err = devices.Get(T.DevID)
	if err != nil {
		log.WithError(err).Fatalf("%s: could not get device", T.AppID)
	}
	

	T.pubsub, err = T.client.PubSub()
	if err != nil {
		log.WithError(err).Fatalf("%s: could not get application pub/sub", T.AppID)
	}
	
	
	allDebvicePubSub := T.pubsub.AllDevices()
	
	activations, err := allDebvicePubSub.SubscribeActivations()
	if err != nil {
		log.WithError(err).Fatalf("%s: could not subscribe to activations", T.AppID)
	}
	
	go func() {
		for activation := range activations {
			log.WithFields(ttnlog.Fields{
				"appEUI":  activation.AppEUI.String(),
				"devEUI":  activation.DevEUI.String(),
				"devAddr": activation.DevAddr.String(),
			}).Info("my-amazing-app: received activation")
		}
	}()
}



func (T *TTN_Info)GetCoordinates() (float32, float32) {
	if T.dev.Latitude == 0 || T.dev.Longitude == 0 {
		fmt.Println("Longitude and Latitude seems to be unset... might be useless for this device unless its set.")
		return 0.0, 0.0
	}
	return T.dev.Latitude, T.dev.Longitude

	
}


func (T *TTN_Info) LoadDevices() {
	
	
	log := apex.Stdout()
	ttnlog.Set(log)
	
	T.config =ttnsdk.NewCommunityConfig(T.AppID)
	T.config.ClientVersion = "2.0.5"
	T.client = T.config.NewClient(T.AppID, T.AppAccessKey)
	
	devices, err := T.client.ManageDevices()
	if err != nil {
		log.WithError(err).Fatalf("%s: could not read CA certificate file", T.AppID)
	}
	T.dev = new(ttnsdk.Device)
	T.dev.AppID = T.AppID
	
	T.dev.AppEUI = types.AppEUI{ T.AppEUID[0], T.AppEUID[1], T.AppEUID[2], T.AppEUID[3], T.AppEUID[4], T.AppEUID[5], T.AppEUID[6], T.AppEUID[7]}
	T.dev.DevEUI = types.DevEUI{ T.DevEUID[0], T.DevEUID[1], T.DevEUID[2], T.DevEUID[3], T.DevEUID[4], T.DevEUID[5], T.DevEUID[6], T.DevEUID[7]}
	
	T.dev, err = devices.Get(T.DevID)
	if err != nil {
		log.WithError(err).Fatalf("%s: could not get device", T.AppID)
	}
	
	T.pubsub, err = T.client.PubSub()
	if err != nil {
		log.WithError(err).Fatalf("%s: could not get application pub/sub", T.AppID)
	}
	
	allDebvicePubSub := T.pubsub.AllDevices()
	
	activations, err := allDebvicePubSub.SubscribeActivations()
	if err != nil {
		log.WithError(err).Fatalf("%s: could not subscribe to activations", T.AppID)
	}
	
	go func() {
		for activation := range activations {
			log.WithFields(ttnlog.Fields{
				"appEUI":  activation.AppEUI.String(),
				"devEUI":  activation.DevEUI.String(),
				"devAddr": activation.DevAddr.String(),
			}).Info("my-amazing-app: received activation")
		}
	}()
}


func (T *TTN_Info) GetData(tmp *map[string]interface{}, fmts map[string]interface{} ) {
	var err error
	var container map[string]interface{}
	
	T.myNewDevicePubSub = T.pubsub.Device(T.dev.DevID)
	
	defer T.myNewDevicePubSub.Close()
	T.uplink, err  = T.myNewDevicePubSub.SubscribeUplink()
	if err != nil {
		log.WithError(err).Fatalf("%s: could not subscribe to uplink messages", T.AppID)
	}
	
	for message := range T.uplink {
		
		json.Unmarshal(message.PayloadRaw, &container)
		
		for fm  := range fmts {
			if (*tmp)== nil {
				//fmt.Println("nil")
				(*tmp) = map[string]interface{}{}
			}
			
			for _, key := range fmts[fm].([]interface{}) {
				(*tmp)[key.(string)] = container[key.(string)]
			}

		}
	}
	
}


func (T *TTN_Info) WatchFeed() {
	var err error
	
	//Subscribes to the device
	T.myNewDevicePubSub = T.pubsub.Device(T.dev.DevID)
	
	defer T.myNewDevicePubSub.Close()
	T.uplink, err  = T.myNewDevicePubSub.SubscribeUplink()
	if err != nil {
		log.WithError(err).Fatalf("%s: could not subscribe to uplink messages", T.AppID)
	}
	
	//Prints the message stream
	for message := range T.uplink {
		
		stringPayload := string(message.PayloadRaw)
		log.WithField("data", stringPayload).Infof("%s: received uplink", T.AppID)
		
	}
}

