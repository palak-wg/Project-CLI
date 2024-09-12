package services_test

import (
	"database/sql"
	"doctor-patient-cli/models"
	"doctor-patient-cli/services"
	"doctor-patient-cli/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestReviewService_AddReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockReviewRepository(ctrl)
	service := services.NewReviewService(mockRepo)

	review := &models.Review{
		PatientID: "p123",
		DoctorID:  strconv.Itoa(1),
		Content:   "Great doctor!",
		Rating:    5,
		Timestamp: []uint8("2024-09-09 10:00:00"),
	}

	mockRepo.EXPECT().AddReview(review).Return(nil)

	err := service.AddReview(review)
	assert.NoError(t, err)
}

func TestReviewService_AddReview_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockReviewRepository(ctrl)
	service := services.NewReviewService(mockRepo)

	review := &models.Review{
		PatientID: "p123",
		DoctorID:  strconv.Itoa(1),
		Content:   "Great doctor!",
		Rating:    5,
		Timestamp: []uint8("2024-09-09 10:00:00"),
	}

	mockRepo.EXPECT().AddReview(review).Return(sql.ErrConnDone)

	err := service.AddReview(review)
	assert.Error(t, err)
}

func TestReviewService_GetAllReviews(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockReviewRepository(ctrl)
	service := services.NewReviewService(mockRepo)

	reviews := []models.Review{
		{PatientID: "p123", DoctorID: strconv.Itoa(1), Content: "Great doctor!", Rating: 5, Timestamp: []uint8("2024-09-09 10:00:00")},
	}

	mockRepo.EXPECT().GetAllReviews().Return(reviews, nil)

	result, err := service.GetAllReviews()
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "p123", result[0].PatientID)
}

func TestReviewService_GetAllReviews_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockReviewRepository(ctrl)
	service := services.NewReviewService(mockRepo)

	mockRepo.EXPECT().GetAllReviews().Return(nil, sql.ErrConnDone)

	result, err := service.GetAllReviews()
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestReviewService_GetReviewsByDoctorID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockReviewRepository(ctrl)
	service := services.NewReviewService(mockRepo)

	reviews := []models.Review{
		{PatientID: "p123", DoctorID: strconv.Itoa(1), Content: "Great doctor!", Rating: 5, Timestamp: []uint8("2024-09-09 10:00:00")},
	}

	mockRepo.EXPECT().GetReviewsByDoctorID(1).Return(reviews, nil)

	result, err := service.GetReviewsByDoctorID(1)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "p123", result[0].PatientID)
}

func TestReviewService_GetReviewsByDoctorID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockReviewRepository(ctrl)
	service := services.NewReviewService(mockRepo)

	mockRepo.EXPECT().GetReviewsByDoctorID(1).Return(nil, sql.ErrConnDone)

	result, err := service.GetReviewsByDoctorID(1)
	assert.Error(t, err)
	assert.Nil(t, result)
}
