package slrealtid

import (
	"encoding/json"
	_ "github.com/mrevilme/slrealtid"
	"github.com/mrevilme/slrealtid/versions/base"
	"io"
	"time"
	"strconv"
	"fmt"
)

func DecodeResponse(reader io.Reader) (base.Departures,error) {
	dep := RealtidResponse{}
	err := json.NewDecoder(reader).Decode(&dep)
	if err != nil {
		return &Departures{}, err
	}

	return dep.ResponseData,nil;
}

type Deviation struct {
	base.Deviation
	Consequence string `json:"Consequence"`
	ImportanceLevel int64 `json:"ImportanceLevel"`
	Text    string `json:"Text"`
}

func (dev Deviation) GetConsequence() string {
	return dev.Consequence
}

func (dev Deviation) GetImportanceLevel() int64 {
	return dev.ImportanceLevel
}

func (dev Deviation) GetText() string {
	return dev.Text
}

type Transports []Transport

func (t Transports) AsBase() []base.Transport {
	bt := make([]base.Transport, len(t))
	for idx, itm := range t {
		bt[idx] = itm
	}
	return bt
}

type Transport struct {
	GroupOfLine          string      					`json:"GroupOfLine"`
	TransportMode        base.TransportModeString     	`json:"TransportMode"`
	LineNumber           string      					`json:"LineNumber"`
	Destination          string      					`json:"Destination"`
	JourneyDirection     int64         					`json:"JourneyDirection"`
	StopAreaName         string      					`json:"StopAreaName"`
	StopAreaNumber       int64         					`json:"StopAreaNumber"`
	StopPointNumber      int64         					`json:"StopPointNumber"`
	StopPointDesignation string      					`json:"StopPointDesignation"`
	TimeTabledDateTime   string      					`json:"TimeTabledDateTime"`
	ExpectedDateTime     string   					`json:"ExpectedDateTime"`
	DisplayTime          string      					`json:"DisplayTime"`
	JourneyNumber        int64        					`json:"JourneyNumber"`
	Deviations           []Deviation 					`json:"Deviations"`
}

func (trans Transport) GetGroupOfLine() string {
	return trans.GroupOfLine
}
func (trans Transport) GetTransportMode() base.TransportMode {
	return trans.TransportMode.AsType()
}
func (trans Transport) GetLineNumber() int64 {
	lineNumber,err := strconv.ParseInt(trans.LineNumber, 10, 64);
	if err != nil {
		fmt.Println("Cannot parse linenumber: %s | %s", trans.LineNumber, err)
		return 0
	}
	return lineNumber
}
func (trans Transport) GetDestination() string {
	return trans.Destination
}
func (trans Transport) GetJourneyDirection() int64{
	return trans.JourneyDirection
}
func (trans Transport) GetStopAreaName() string {
	return trans.StopAreaName
}
func (trans Transport) GetStopAreaNumber() int64 {
	return trans.StopAreaNumber
}
func (trans Transport) GetStopPointNumber() int64 {
	return trans.StopPointNumber
}
func (trans Transport) GetStopPointDesignation() string {
	return trans.StopPointDesignation
}
func (trans Transport) GetTimeTabledDateTime() time.Time {
	date,err := time.Parse("2006-01-02T15:04:05", trans.TimeTabledDateTime)
	if err != nil {
		fmt.Println("Cannot parse TimeTabledDateTime: %s | %s", trans.TimeTabledDateTime, err)
		return time.Now()
	}
	return date
}
func (trans Transport) GetExpectedDateTime() time.Time {
	date,err := time.Parse("2006-01-02T15:04:05", trans.ExpectedDateTime)
	if err != nil {
		fmt.Println("Cannot parse ExpectedDateTime: %s | %s", trans.ExpectedDateTime, err)
		return time.Now()
	}
	return date
}
func (trans Transport) GetDisplayTime() string {
	return trans.DisplayTime
}
func (trans Transport) GetDeviations() []base.Deviation {
	return nil
}

func (trans Transport) GetJourneyNumber() int64 {
	return trans.JourneyNumber
}

type Departures struct {
	LatestUpdate 		string      `json:"LatestUpdate"`
	DataAge      		int64       `json:"DataAge"`
	Metros       		Transports `json:"Metros"`
	Buses        		Transports `json:"Buses"`
	Trains              Transports `json:"Trains"`
	Trams               Transports `json:"Trams"`
	Ships               Transports `json:"Ships"`
	StopPointDeviation []StopPointDeviation `json:"StopPointDeviations"`

}

type StopInfo struct{
	GroupOfLine string `json:"GroupOfLine"`
	StopAreaName string `json:"StopAreaName"`
	StopAreaNumber int64 `json:"StopAreaNumber"`
	TransportMode base.TransportModeString `json: "TransportMode"`
}
func (si StopInfo) GetGroupOfLine() string {
	return si.GroupOfLine
}
func (si StopInfo) GetStopAreaName() string {
	return si.StopAreaName
}
func (si StopInfo) GetStopAreaNumber() int64 {
	return si.StopAreaNumber
}

func (si StopInfo) GetTransportMode() base.TransportMode {
	return si.TransportMode.AsType()
}

type StopPointDeviation struct{
	StopInfo StopInfo `json:"StopInfo"`
	Deviation Deviation `json:"Deviation"`
}

func (spd StopPointDeviation) GetDeviation() base.Deviation {
	return spd.Deviation
}

func (spd StopPointDeviation) GetStopInfo() base.StopInfo {
	return spd.StopInfo
}

func (dep Departures) GetMetros() []base.Transport {
	return dep.Metros.AsBase()
}

func (dep Departures) GetBuses() []base.Transport {
	return dep.Buses.AsBase();
}

func (dep Departures) GetTrains() []base.Transport  {
	return dep.Trains.AsBase();
}

func (dep Departures) GetTrams() []base.Transport  {
	return dep.Trams.AsBase();
}

func (dep Departures) GetShips() []base.Transport  {
	return dep.Ships.AsBase();
}

func (dep Departures) GetAllDepartures() []base.Transport {
	allDepartures := make([]base.Transport,0)
	allDepartures = append(allDepartures,dep.GetMetros()...)
	allDepartures = append(allDepartures,dep.GetShips()...)
	allDepartures = append(allDepartures,dep.GetTrains()...)
	allDepartures = append(allDepartures,dep.GetTrams()...)
	allDepartures = append(allDepartures,dep.GetBuses()...)
	return allDepartures
}

func (dep Departures) GetStopPointDeviations() []base.StopPointDeviation {
	bt := make([]base.StopPointDeviation, len(dep.StopPointDeviation))
	for idx, itm := range dep.StopPointDeviation {
		bt[idx] = itm
	}
	return bt
}

type RealtidResponse struct {
	StatusCode    int         `json:"StatusCode"`
	Message       string	  `json:"Message"`
	ExecutionTime int         `json:"ExecutionTime"`
	ResponseData  Departures `json:"ResponseData"`
}