package services

import (
	"doctor-patient-cli/utils"
	"github.com/fatih/color"
)

// ApproveDoctorSignup update the unapproved doctors based on the provided user ID to approved ones
func ApproveDoctorSignup(userID string) error {
	db := utils.GetDB()

	// Update the doctor record to set IsApproved to true
	_, err := db.Exec("UPDATE users SET is_approved = ? WHERE user_id = ?", true, userID)

	// making entry to doctor table
	_, _ = db.Exec("INSERT INTO doctors (user_id, specialization, experience, rating) VALUES (?, ?, ?,?)",
		userID, "xxx", 0, 2)

	if err != nil {
		color.Red("error approving doctor signup: %v", err)
		return err
	}

	// Create a notification for the doctor
	_, err = db.Exec("INSERT INTO notifications (user_id, content) VALUES (?, ?)",
		userID, "Your signup request has been approved by the admin.")
	if err != nil {
		color.Red("Error creating notification: %v", err)
		return err
	}

	// Assuming we have a function to fetch doctor email to send notification
	doctor, err := GetUserByID(userID)
	if err != nil {
		color.Red("Error fetching doctor: %v", err)
		return err
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
		color.Red("Error getting pending requests: %v", err)
		return
	}
	defer rows.Close()

	// Displaying all pending requests
	color.Cyan("\n============== PENDING DOCTOR SIGNUPS ================")
	for rows.Next() {
		var ID string
		_ = rows.Scan(&ID)

		color.Magenta("Request pending for Doctor ID: %s", ID)
	}
}
