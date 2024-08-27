package models

type User struct {
	UserID      string
	Password    string
	Username    string
	Age         int
	Gender      string
	Email       string
	PhoneNumber string
	UserType    string
	IsApproved  bool
}

type Doctor struct {
	User
	Specialization string
	Experience     int
	Rating         float64
}

type Patient struct {
	User
	MedicalHistory string
}

type Review struct {
	PatientID string
	DoctorID  string
	Content   string
	Rating    int
	Timestamp []uint8
}

type Notification struct {
	UserID    string
	Content   string
	Timestamp []uint8
}

type Appointment struct {
	AppointmentID int
	DoctorID      string
	PatientID     string
	DateTime      []uint8
	IsApproved    bool
}

type Message struct {
	Sender    string
	Content   string
	Receiver  string
	Timestamp []uint8
	Status    string
}
