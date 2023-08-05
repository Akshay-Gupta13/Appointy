package main

import(
	"time"
	"net/http"
	"fmt"
	// "net/httptest"

)

// entities

type Resource struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BusinessHour struct {
	Id         string    `json:"id"`
	ResourceId string    `json:"resource_id"`
	Quantity   int64     `json:"quantity"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
}

type BlockHour struct {
	Id         string    `json:"id"`
	ResourceId string    `json:"resource_id"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
}

type Appointment struct {
	Id         string    `json:"id"`
	ResourceId string    `json:"resource_id"`
	Quantity   int64     `json:"quantity"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
}

type Duration struct {
	Seconds int64 `json:"seconds"`
}

// endpoint request structs

type ListBusinessHoursRequest struct {
	ResourceId string `json:"resourceId"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}

type ListBlockHoursRequest struct {
	ResourceId string `json:"resourceId"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}

type ListAppointmentRequest struct {
	ResourceId string `json:"resourceId"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}


// helper functions
type Slot struct {
	ResourceId string `json:"resourceId"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
	Quantity   int64  `json:"quantity"`

}

func TimeToString(tm time.Time) string {
	return tm.Format(time.RFC3339)
}

func StringToTime(timeStr string) (time.Time, error) {
	s, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return time.Time{}, err
	}

	return s, nil
}
func (s *Server) routes(){
	s:HandleFunc("/available-slots", s.listnewUserSlot()),Method("GET")
	s:HandleFunc("/available-slots", s.createnewUserSlot()),Method("POST")
	s:HandleFunc("/available-slots/{id}", s.removenewUserSlot()),Method("DELETE")
   
   }
   
func (s *Server) newUserSlot() http.HandleFunc {
	   return func(w http.ResponseWriter, r *http.Request){
		   var i Slot 
		   if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			   http.Error(w, err.Error(), http.StatusBadRequest)
			   return
		   }
		   i.ID = uuid.New()
		   s.newUserSlot = append(s.newUserSlot, i)
   
		   w.Header().Set("Content-Type", "application/json")
		   if err := json.NewEncoder(w).Encode(i); err != nil {
		   http.Error(w, err.Error(), http.StatusInternalServerError)
		   return 
		   }
	   }
   }
//    the core logic i am implementing was first to check and fetch details hierarchy as first import data from first five given API's then at last using the Availability API to assign/notassign slots which depends on factor: the time duration of previous slots, the quantity of previous slots beacuse if the remaining slots (total - slots occupied by previous slot) and the new slot quantity is less than or equal to left over slot space then we can assign slot parallel to previous slot other wise we have to wait till it end and another condition of block hours also to be considered as if we assign slot but it overlaps with beginning time of block hours then we cant assign that window and base case for each case before assigning is it lies in time range of after 9:00 to before 5:00 if ending time of any slot is beyond 5:00 then we cant assign slots. 
//    first run individually each API then check for the new created one and after that using an iterator to check for each slot in interval of given 30 min , 60 min and 120 min then as number of working hours is fixed time constraints might not be issue with this approach and after that we check it as repeating sub problem for each window where a user ask for slot on 3 pitchs along with the quantity(count).