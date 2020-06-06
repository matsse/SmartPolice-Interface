package Providers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type TS_info struct {
	DevID               string     `json:"devID,omitempty"`
	ChannelID           string     `json:"channelID,omitempty"`
	URL                 string     `json:"url,omitempty"`
	Author              string     `json:"author,omitempty"`
}



func (T *TS_info) NewDevice() {
	
	baseAddress := "https://thingspeak.com/channels/"
	
	
	setAppID:
	fmt.Println("What is the channel ID of your data source?")
	fmt.Scan(&T.ChannelID)
	if _, err := http.Get(baseAddress+ T.ChannelID); err != nil {
		log.Println(err)
		goto setAppID
	}

	selectFields:
	
	
	request, _ := http.Get(baseAddress+ T.ChannelID)
	defer request.Body.Close()
	x, err := ioutil.ReadAll(request.Body)
	if  err != nil {
		log.Println(err)
		goto selectFields
	}
	
	
	fmt.Println(string(x))
	
	
	
	
}