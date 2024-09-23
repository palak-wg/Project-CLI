package controllers

import (
	"bufio"
	"doctor-patient-cli/repositories"
	"doctor-patient-cli/services"
	"doctor-patient-cli/utils"
	"os"
)

var (
	db = utils.GetDB()

	adminRepo        = repositories.NewAdminRepository(db)
	patientRepo      = repositories.NewPatientRepository(db)
	userRepo         = repositories.NewUserRepository(db)
	doctorRepo       = repositories.NewDoctorRepository(db)
	messageRepo      = repositories.NewMessageRepository(db)
	notificationRepo = repositories.NewNotificationRepository(db)
	appointmentRepo  = repositories.NewAppointmentRepository(db)
	reviewRepo       = repositories.NewReviewRepository(db)

	adminService        = services.NewAdminService(adminRepo, userRepo)
	patientService      = services.NewPatientService(patientRepo)
	userService         = services.NewUserService(userRepo)
	doctorService       = services.NewDoctorService(doctorRepo)
	messageService      = services.NewMessageService(messageRepo)
	notificationService = services.NewNotificationService(notificationRepo)
	appointmentService  = services.NewAppointmentService(appointmentRepo)
	reviewService       = services.NewReviewService(reviewRepo)

	//userHandler = handlers.NewUserHandler(adminService)

	reader = bufio.NewReader(os.Stdin)
)
