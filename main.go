package main

import "time"

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

func TimeToString(tm time.Time) string {
	return tm.Format(time.RFC3339)
}

func StringToTime(timeStr string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
func (s *Server) routes(){
	s.HandleFunc("/available-slots", s.listnewUserSlot()),Method("GET")
	s.HandleFunc("/available-slots", s.createnewUserSlot()),Method("POST")
	s.HandleFunc("/available-slots/{id}", s.removenewUserSlot()),Method("DELETE")
   
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