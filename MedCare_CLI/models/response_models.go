package models

type APIResponse struct {
	Status int `json:"status_code"`
	//ErrorCode int 	`json:"error_code"`
	Data any `json:"data"`
}

/*
status --> error/ info
status code
error code	--> should be of us
msg		--> should map with error code
data
*/

type APIResponseUser struct {
	UserID      string `json:"username"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type APIResponseDoctor struct {
	DoctorID       string  `json:"doctor_id"`
	Specialization string  `json:"specialization"`
	Experience     int     `json:"experience"`
	Rating         float64 `json:"rating"`
}

type APIResponsePendingSignup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
