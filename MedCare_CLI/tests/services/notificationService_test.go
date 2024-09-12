package services_test

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/services"
	"doctor-patient-cli/tests/mocks"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNotificationService_GetNotificationsByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockNotificationRepository(ctrl)
	service := services.NewNotificationService(mockRepo)

	testCases := []struct {
		name           string
		userID         string
		mockReturnData []models.Notification
		mockReturnErr  error
		expectedResult []models.Notification
		expectedErr    error
	}{
		{
			name:           "Success case - notifications retrieved",
			userID:         "123",
			mockReturnData: []models.Notification{{UserID: "123", Content: "Test Notification", Timestamp: []uint8("2024-09-01")}},
			mockReturnErr:  nil,
			expectedResult: []models.Notification{{UserID: "123", Content: "Test Notification", Timestamp: []uint8("2024-09-01")}},
			expectedErr:    nil,
		},
		{
			name:           "Failure case - repository error",
			userID:         "123",
			mockReturnData: nil,
			mockReturnErr:  errors.New("database error"),
			expectedResult: nil,
			expectedErr:    errors.New("database error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.EXPECT().GetNotificationsByUserID(tc.userID).Return(tc.mockReturnData, tc.mockReturnErr)

			notifications, err := service.GetNotificationsByUserID(tc.userID)

			assert.Equal(t, tc.expectedResult, notifications)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestNotificationService_GetAllNotifications(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockNotificationRepository(ctrl)
	service := services.NewNotificationService(mockRepo)

	testCases := []struct {
		name           string
		mockReturnData []models.Notification
		mockReturnErr  error
		expectedResult []models.Notification
		expectedErr    error
	}{
		{
			name:           "Success case - all notifications retrieved",
			mockReturnData: []models.Notification{{UserID: "123", Content: "Test Notification 1", Timestamp: []uint8("2024-09-01")}},
			mockReturnErr:  nil,
			expectedResult: []models.Notification{{UserID: "123", Content: "Test Notification 1", Timestamp: []uint8("2024-09-01")}},
			expectedErr:    nil,
		},
		{
			name:           "Failure case - repository error",
			mockReturnData: nil,
			mockReturnErr:  errors.New("database error"),
			expectedResult: nil,
			expectedErr:    errors.New("database error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.EXPECT().GetAllNotifications().Return(tc.mockReturnData, tc.mockReturnErr)

			notifications, err := service.GetAllNotifications()

			assert.Equal(t, tc.expectedResult, notifications)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

//package services
//
//import (
//	"doctor-patient-cli/services"
//	"doctor-patient-cli/tests/mockDB"
//	"doctor-patient-cli/utils"
//	"fmt"
//	"testing"
//
//	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/stretchr/testify/assert"
//)
//
//func TestGetNotificationsByUserID(t *testing.T) {
//	// Initialize the mock database
//	mockDB.MockInitDB(t)
//	defer utils.CloseDB()
//
//	t.Run("GetNotificationsByUserID Success", func(t *testing.T) {
//		// Set up mock rows to return
//		rows := sqlmock.NewRows([]string{"user_id", "content", "timestamp"}).
//			AddRow("user1", "Notification 1", "2023-01-01 10:00:00").
//			AddRow("user1", "Notification 2", "2023-01-02 11:00:00")
//
//		// Expect the query and return the mock rows
//		mockDB.Mock.ExpectQuery("SELECT user_id, content, timestamp FROM notifications WHERE user_id = ?").
//			WithArgs("user1").
//			WillReturnRows(rows)
//
//		// Call the GetNotificationsByUserID function
//		notifications, err := services.GetNotificationsByUserID("user1")
//
//		// Check for no errors
//		assert.NoError(t, err)
//		assert.Len(t, notifications, 2)
//		assert.Equal(t, "Notification 1", notifications[0].Content)
//
//		// Ensure all expectations are met
//		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
//			t.Errorf("there were unfulfilled expectations: %v", err)
//		}
//	})
//
//	t.Run("GetNotificationsByUserID Query Error", func(t *testing.T) {
//		// Expect the query and simulate an error
//		mockDB.Mock.ExpectQuery("SELECT user_id, content, timestamp FROM notifications WHERE user_id = ?").
//			WithArgs("user1").
//			WillReturnError(fmt.Errorf("query error"))
//
//		// Call the GetNotificationsByUserID function
//		_, err := services.GetNotificationsByUserID("user1")
//
//		// Check that an error was returned
//		assert.Error(t, err)
//
//		// Ensure all expectations are met
//		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
//			t.Errorf("there were unfulfilled expectations: %v", err)
//		}
//	})
//}
//
//func TestGetAllNotifications(t *testing.T) {
//	// Initialize the mock database
//	mockDB.MockInitDB(t)
//	defer utils.CloseDB()
//
//	t.Run("GetAllNotifications Success", func(t *testing.T) {
//		// Set up mock rows to return
//		rows := sqlmock.NewRows([]string{"user_id", "content", "timestamp"}).
//			AddRow("user1", "Notification 1", "2023-01-01 10:00:00").
//			AddRow("user2", "Notification 2", "2023-01-02 11:00:00")
//
//		// Expect the query and return the mock rows
//		mockDB.Mock.ExpectQuery("SELECT user_id, content, timestamp FROM notifications").
//			WillReturnRows(rows)
//
//		// Call the GetAllNotifications function
//		notifications, err := services.GetAllNotifications()
//
//		// Check for no errors
//		assert.NoError(t, err)
//		assert.Len(t, notifications, 2)
//		assert.Equal(t, "Notification 1", notifications[0].Content)
//
//		// Ensure all expectations are met
//		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
//			t.Errorf("there were unfulfilled expectations: %v", err)
//		}
//	})
//
//	t.Run("GetAllNotifications Query Error", func(t *testing.T) {
//		// Expect the query and simulate an error
//		mockDB.Mock.ExpectQuery("SELECT user_id, content, timestamp FROM notifications").
//			WillReturnError(fmt.Errorf("query error"))
//
//		// Call the GetAllNotifications function
//		_, err := services.GetAllNotifications()
//
//		// Check that an error was returned
//		assert.Error(t, err)
//
//		// Ensure all expectations are met
//		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
//			t.Errorf("there were unfulfilled expectations: %v", err)
//		}
//	})
//}
