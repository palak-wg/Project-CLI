package main

import (
	"doctor-patient-cli/controllers"
	"doctor-patient-cli/handlers"
	"doctor-patient-cli/middlewares"
	"doctor-patient-cli/repositories"
	"doctor-patient-cli/services"
	"doctor-patient-cli/utils"
	"fmt"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func init() {
	utils.InitDB()
}

func main() {
	defer utils.CloseDB()
	loggerZap, _ := zap.NewProduction()

	go func() {
		r := mux.NewRouter()

		userHandler := handlers.NewUserHandler(services.NewUserService(repositories.NewUserRepository(utils.GetDB())))
		doctorHandler := handlers.NewDoctorHandler(services.NewDoctorService(repositories.NewDoctorRepository(utils.GetDB())))
		adminHandler := handlers.NewAdminHandler(services.NewAdminService(repositories.
			NewAdminRepository(utils.GetDB()), repositories.NewUserRepository(utils.GetDB())))
		notificationHandler := handlers.NewNotificationHandler(services.
			NewNotificationService(repositories.NewNotificationRepository(utils.GetDB())))
		reviewHandler := handlers.NewReviewHandler(services.NewReviewService(repositories.NewReviewRepository(utils.GetDB())))
		appointmentHandler := handlers.NewAppointmentHandler(services.NewAppointmentService(repositories.
			NewAppointmentRepository(utils.GetDB())))
		patientHandler := handlers.NewPatientHandler(services.NewPatientService(repositories.NewPatientRepository(utils.GetDB())))
		messageHandler := handlers.NewMessageHandler(services.NewMessageService(repositories.NewMessageRepository(utils.GetDB())))

		// basic authentication and token generation
		r.HandleFunc("/signup", userHandler.Signup).Methods("POST")
		r.HandleFunc("/login", userHandler.Login).Methods("POST")

		// auth middleware
		apiRouter := r.PathPrefix("/api").Subrouter()
		apiRouter.Use(middlewares.AuthenticationMiddleware)

		apiRouter.HandleFunc("/users/{user_id}", userHandler.GetUser).Methods("GET")
		apiRouter.HandleFunc("/doctors", doctorHandler.GetDoctors).Methods("GET")
		apiRouter.HandleFunc("/doctors/{doctor_id}", doctorHandler.GetDoctor).Methods("GET")
		apiRouter.HandleFunc("/notifications/{user_id}", notificationHandler.GetNotifications).Methods("GET")
		apiRouter.HandleFunc("/reviews/{doctor_id}", reviewHandler.GetDoctorSpecificReviews).Methods("GET")
		apiRouter.HandleFunc("/reviews", reviewHandler.CreateReview).Methods("POST")
		apiRouter.HandleFunc("/appointments", appointmentHandler.GetAppointments).Methods("GET")
		apiRouter.HandleFunc("/appointments", appointmentHandler.CreateAppointment).Methods("POST")
		apiRouter.HandleFunc("/appointments", appointmentHandler.UpdateAppointment).Methods("PATCH")
		apiRouter.HandleFunc("/users", patientHandler.UpdateProfile).Methods("PUT")
		apiRouter.HandleFunc("/messages", messageHandler.GetMessages).Methods("GET")
		apiRouter.HandleFunc("/messages", messageHandler.AddMessage).Methods("POST")

		// admin middleware
		adminRouter := apiRouter.PathPrefix("/admin").Subrouter()
		adminRouter.Use(middlewares.AuthenticationMiddleware)

		adminRouter.HandleFunc("/users", adminHandler.GetUsers).Methods("GET")
		adminRouter.HandleFunc("/notifications", notificationHandler.AllNotifications).Methods("GET")
		adminRouter.HandleFunc("/doctors/approval", adminHandler.GetPendingDoctors).Methods("GET")
		adminRouter.HandleFunc("/doctors/approval", adminHandler.ApprovePendingDoctors).Methods("PATCH")
		adminRouter.HandleFunc("/reviews", reviewHandler.GetAllReview).Methods("GET")

		// starting of the server
		loggerZap.Info("Server listening on port 8075")
		if err := http.ListenAndServe(":8075", r); err != nil {
			loggerZap.Fatal("Could not start server:", zap.Error(err))
		}
	}()
	defer loggerZap.Sync()

	// Application via command line
	StartApp()
}

// StartApp runs the application logic
func StartApp() {
	for {
		// Title and Welcome Message
		color.Cyan("===========================================")
		color.Cyan("\t🚀 Welcome To The MedCare 🚀")
		color.Cyan("===========================================")

		// Menu Options
		color.Magenta("\nPlease choose an option:")
		fmt.Println("1. Login")
		fmt.Println("2. Signup")
		fmt.Println("3. Exit")
		fmt.Print("\nEnter your choice: ")

		// User input
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			color.Blue("🔑 Logging in...")
			user := controllers.Login()
			switch user.UserType {
			case "admin":
				color.Yellow("👨‍💼 Welcome, Admin!")
				controllers.AdminMenu()
			case "doctor":
				color.Yellow("👨‍⚕️ Welcome, Doctor!")
				controllers.DoctorMenu(user)
			case "patient":
				color.Yellow("🧑‍⚕️ Welcome, Patient!")
				controllers.PatientMenu(user)
			default:
				color.Red("🚨 Invalid user type")
			}
		case 2:
			color.Blue("📝 Signing up...")
			controllers.Signup()
		case 3:
			color.Green("👋 Exiting... Goodbye!")
			return
		default:
			color.Red("🚫 Invalid choice. Please try again.")
		}
	}
}
