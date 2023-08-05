package main

import (
	// "fmt"
	// "log"
	// "os"
	// "testing"
	// "log"
	// "net/http"
	// "net/http/httptest"
	// "strconv"
	// "bytes"
)

type ResourceId struct {	
  Id   string `json:"id"`
  Name string `json:"name"`
  Type string `json:"type"`
//   key value pair, upper case -> datatype then json format
}

type Date struct {
	Date int64 `json:"date"`
// date cant be string but integer type.
}
// during business hours the occuiped slot with some q1 quantity and remain with less than total-q1 space for people.
type TotalSlots struct {
  Id        string `json:"id"`
  StartTime string `json:"start_time"`
  EndTime   string `json:"end_time"`


}
type UnavailableSlot struct{
// related to minutes or seconds , slots which are not available for given quantitites.
	Id           string `json:"id"`
	ResourceId   string `json:"resource"`
	StartTime    string `json:"starttime"`
	EndTime      string `json:"endtime"`
}


type Quantity struct {
     Id               string `json:"id"`
	 ResourceId       string `json:"resource"` 	  
	 UnavailableSlot  string `json:"unavailableslot"`
}
type Appointment struct {
	Id         string    `json:"id"`
	ResourceId string    `json:"resource_id"`
	Quantity   int64     `json:"quantity"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
}

type BlockHours struct {
	Id         string `json:"id"`
	ResourceId string `json:"resource_id"`
	StartTime  string `json:"starttime"`
	EndTime    string `json:"endtime"`
// maintainance time or server reset time
}
// string to time conversion needed ?..
// slot avalibilty depends on : occupied users 
// : space required for new user (total-occupied)>= newusers
//    start and end duration (overlap with other might be possible)
// output part..
  
 func main() {
 

 }