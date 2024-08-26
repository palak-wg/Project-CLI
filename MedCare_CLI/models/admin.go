package models

import (
	"doctor-patient-cli/utils"
	"fmt"
	"time"
)

// ApproveDoctorSignup update the unapproved doctors based on the provided user ID to approved ones
func ApproveDoctorSignup(userID string) error {
	// Update the doctor record to set IsApproved to true
	db := utils.GetDB()
	_, err := db.Exec("UPDATE users SET is_approved = ? WHERE user_id = ?", true, userID)

	// making entry to doctor table
	_, _ = db.Exec("INSERT INTO doctors (user_id, specialization, experience, rating) VALUES (?, ?, ?,?)",
		userID, "xxx", 0, 2)

	if err != nil {
		return fmt.Errorf("error approving doctor signup: %v", err)
	}

	// Create a notification for the doctor
	_, err = db.Exec("INSERT INTO notifications (user_id, content, timestamp) VALUES (?, ?, ?)",
		userID, "Your signup request has been approved by the admin.", time.Now())
	if err != nil {
		return fmt.Errorf("error creating notification: %v", err)
	}

	// Assuming we have a function to fetch doctor email to send notification
	doctor, err := GetUserByID(userID)
	if err != nil {
		return fmt.Errorf("error fetching doctor: %v", err)
	}

	// Send email notification to the doctor
	go utils.SendEmail(doctor.Email, "Signup Approved", "Your signup request has been approved by the admin.")

	return nil
}

// PendingDoctorSignupRequest display unapproved doctor signup request
func PendingDoctorSignupRequest() {
	db := utils.GetDB()

	// Fetching all pending requests
	rows, err := db.Query("SELECT user_id FROM users WHERE user_type ='doctor' AND is_approved=0 ")
	if err != nil {
		fmt.Println("Error getting pending request.")
	}
	defer rows.Close()

	// Displaying all pending requests
	for rows.Next() {
		var ID string
		err = rows.Scan(&ID)
		if err != nil {
			fmt.Println("Error getting pending request.")
		}
		fmt.Println("Request pending for Doctor ID: ", ID)
	}
}
