package models

type User struct {
	UserID      string `json:"username"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	UserType    string `json:"role"`
	IsApproved  bool   `json:"is_approved"`
}

type Doctor struct {
	User
	Specialization string  `json:"specialization"`
	Experience     int     `json:"experience"`
	Rating         float64 `json:"rating"`
}

type Patient struct {
	User
	MedicalHistory string `json:"medical_history"`
}
type Review struct {
	PatientID string  `json:"patient_id"`
	DoctorID  string  `json:"doctor_id"`
	Content   string  `json:"content"`
	Rating    int     `json:"rating"`
	Timestamp []uint8 `json:"time"`
}

type Notification struct {
	UserID    string  `json:"user_id"`
	Content   string  `json:"content"`
	Timestamp []uint8 `json:"time"`
}

type Appointment struct {
	AppointmentID int     `json:"appointment_id"`
	DoctorID      string  `json:"doctor_id"`
	PatientID     string  `json:"patient_id"`
	DateTime      []uint8 `json:"time"`
	IsApproved    bool    `json:"is_approved"`
}

type Message struct {
	Sender    string
	Content   string
	Receiver  string
	Timestamp []uint8
	Status    string
}
