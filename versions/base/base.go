package base

import (
	"time"
)

type TransportMode int64

const (
	TransportModeUnknown TransportMode = iota
	TransportModeMetro
	TransportModeBus
	TransportModeTrain
	TransportModeTram
)

type TransportModeString string
func (tm TransportModeString) AsType() TransportMode{
	if tm == "METRO" {
		return TransportModeMetro
	} else if tm == "BUS" {
		return TransportModeBus
	} else if tm == "TRAIN" {
		return TransportModeTrain
	} else if tm == "TRAM" {
		return TransportModeTram
	} else {
		return TransportModeUnknown
	}
}

type StopPointDeviation interface{
	GetStopInfo() StopInfo;
	GetDeviation() Deviation;
}

type Deviation interface{
	GetConsequence() string
 	GetImportanceLevel() int64
	GetText() string
}

type StopInfo interface{
	GetGroupOfLine() string
	GetStopAreaName() string
	GetStopAreaNumber() int64
	GetTransportMode() TransportMode
}

type Departures interface{
	GetMetros() []Transport;
	GetBuses() []Transport;
	GetTrams() []Transport;
	GetShips() []Transport;
	GetStopPointDeviations() []StopPointDeviation
	GetAllDepartures() []Transport;
}

type Transport interface {
	GetLineNumber()		 		int64
	GetTransportMode() 			TransportMode
	GetDestination()			string
	GetGroupOfLine() 			string
	GetJourneyDirection()       int64
	GetStopAreaName() 	        string
	GetStopAreaNumber()	 	    int64
	GetStopPointNumber()	    int64
	GetStopPointDesignation() 	string
	GetTimeTabledDateTime() 	time.Time
	GetExpectedDateTime()   	time.Time
	GetDisplayTime()        	string
	GetJourneyNumber()			int64
	GetDeviations()				[]Deviation
}