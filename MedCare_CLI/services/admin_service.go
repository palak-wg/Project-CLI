package services

import (
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/models"
	"doctor-patient-cli/utils"
	"log"
)

type AdminService struct {
	adminRepo interfaces.AdminRepository
	userRepo  interfaces.UserRepository
}

//func (service *AdminService) PendingDoctorSignupRequest() ([]models.Doctor, error) {
//	//TODO implement me
//	panic("implement me")
//}

func NewAdminService(adminRepo interfaces.AdminRepository, userRepo interfaces.UserRepository) *AdminService {
	return &AdminService{
		adminRepo: adminRepo,
		userRepo:  userRepo,
	}
}

// ApproveDoctorSignup approves a doctor signup request and sends a notification
func (service *AdminService) ApproveDoctorSignup(userID string) error {
	// Step 1: Fetch user details
	user, err := service.userRepo.GetUserByID(userID)
	if err != nil {
		log.Printf("Service: Error fetching user details for userID %s: %v", userID, err)
		return err
	}

	// Step 2: Approve the doctor signup
	err = service.adminRepo.ApproveDoctorSignup(userID)
	if err != nil {
		log.Printf("Service: Error approving doctor signup: %v", err)
		return err
	}

	// Step 3: Create a notification for the user
	err = service.adminRepo.CreateNotificationForUser(userID, "Your signup request has been approved by the admin.")
	if err != nil {
		log.Printf("Service: Error creating notification for userID %s: %v", userID, err)
		return err
	}

	// Step 4: Email the user
	go utils.SendEmail(user.Email, "MedCare Signup Approved", "Your signup request has been approved by the admin.")

	return nil
}

// GetPendingDoctorRequests retrieves pending doctor signup requests
func (service *AdminService) GetPendingDoctorRequests() ([]models.Doctor, error) {
	return service.adminRepo.GetPendingDoctorRequests()
}

// GetAllUsers retrieves all users from the repository
func (service *AdminService) GetAllUsers() ([]models.User, error) {
	return service.adminRepo.GetAllUsers()
}

func (service *AdminService) CreateNotificationForUser(userID string, content string) error {
	return service.adminRepo.CreateNotificationForUser(userID, content)
}
