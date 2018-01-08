package slrealtid

import (
	"net/url"
	"net/http"
	"fmt"
	"github.com/mrevilme/slrealtid"
	"github.com/mrevilme/slrealtid/versions/base"
	"encoding/json"
)
const BASEURL="http://api.sl.se"
const ENDPOINT="/api2/realtimedeparturesV4.json"

type Client slrealtid.Client

func NewClient(key string) (Client,error) {
	client := Client{}
	client.Key = key
	return client, nil
}

func (client Client) Departures(siteid int64, timewindow int64) (base.Departures,error) {
	//http://api.sl.se/api2/realtimedeparturesV4.<FORMAT>?key=<DIN API NYCKEL>&siteid=<SITEID>&timewindow=<TIMEWINDOW>
	params := url.Values{}
	params.Add("key", client.Key)
	params.Add("siteid", fmt.Sprintf("%d",siteid))
	params.Add("timewindow", fmt.Sprintf("%d",timewindow))
	resp,err := http.Get(fmt.Sprintf("%s%s?%s", BASEURL, ENDPOINT, params.Encode()))
	if err != nil {
		return nil, err
	}

	response := RealtidResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return response.ResponseData,nil
}

func (client Client) DeparturesNow(siteid int64) (base.Departures,error) {
	return client.Departures(siteid, 0);
}