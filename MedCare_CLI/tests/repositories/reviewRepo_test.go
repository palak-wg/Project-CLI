package repositories_test

import (
	"database/sql"
	"doctor-patient-cli/models"
	"doctor-patient-cli/repositories"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddReview_Success(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Initialize the repository with the mock DB
	repo := repositories.NewReviewRepository(db)

	// Create a sample review
	review := &models.Review{
		PatientID: "patient123",
		DoctorID:  "doctor456",
		Content:   "Excellent service",
		Rating:    5,
		Timestamp: "09-09-2024",
	}

	// Expect a successful INSERT into the reviews table
	mock.ExpectExec("INSERT INTO reviews").
		WithArgs(review.PatientID, review.DoctorID, review.Content, review.Rating).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the AddReview function
	err = repo.AddReview(review)

	// Assert no errors
	assert.NoError(t, err)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAddReview_Failure(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Initialize the repository with the mock DB
	repo := repositories.NewReviewRepository(db)

	// Create a sample review
	review := &models.Review{
		PatientID: "patient123",
		DoctorID:  "doctor456",
		Content:   "Excellent service",
		Rating:    5,
		Timestamp: "09-09-2024",
	}

	// Simulate an error during the INSERT operation
	mock.ExpectExec("INSERT INTO reviews").
		WithArgs(review.PatientID, review.DoctorID, review.Content, review.Rating).
		WillReturnError(errors.New("insert failed"))

	// Call the AddReview function
	err = repo.AddReview(review)

	// Assert that an error occurred
	assert.Error(t, err)

	// Ensure the error message is what we expect
	assert.EqualError(t, err, "insert failed")

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestReviewRepository_GetAllReviews(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock db: %v", err)
	}
	defer db.Close()

	repo := repositories.NewReviewRepository(db)

	rows := sqlmock.NewRows([]string{"patient_id", "doctor_id", "content", "rating", "timestamp"}).
		AddRow("p123", 1, "Great doctor!", 5, "2024-09-09 10:00:00").
		AddRow("p124", 1, "Good experience.", 4, "2024-09-09 11:00:00")

	mock.ExpectQuery("SELECT patient_id, doctor_id, content, rating, timestamp FROM reviews").
		WillReturnRows(rows)

	reviews, err := repo.GetAllReviews()
	assert.NoError(t, err)
	assert.Len(t, reviews, 2)
	assert.Equal(t, "p123", reviews[0].PatientID)
	assert.Equal(t, "1", reviews[0].DoctorID)
}

func TestReviewRepository_GetAllReviews_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock db: %v", err)
	}
	defer db.Close()

	repo := repositories.NewReviewRepository(db)

	mock.ExpectQuery("SELECT patient_id, doctor_id, content, rating, timestamp FROM reviews").
		WillReturnError(sql.ErrConnDone)

	reviews, err := repo.GetAllReviews()
	assert.Error(t, err)
	assert.Nil(t, reviews)
}

func TestReviewRepository_GetReviewsByDoctorID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock db: %v", err)
	}
	defer db.Close()

	repo := repositories.NewReviewRepository(db)

	rows := sqlmock.NewRows([]string{"patient_id", "doctor_id", "content", "rating", "timestamp"}).
		AddRow("p123", 1, "Great doctor!", 5, "2024-09-09 10:00:00")

	mock.ExpectQuery("SELECT patient_id, doctor_id, content, rating, timestamp FROM reviews WHERE doctor_id = ?").
		WithArgs(1).
		WillReturnRows(rows)

	reviews, err := repo.GetReviewsByDoctorID("1")
	assert.NoError(t, err)
	assert.Len(t, reviews, 1)
	assert.Equal(t, "p123", reviews[0].PatientID)
	assert.Equal(t, "1", reviews[0].DoctorID)
}

func TestReviewRepository_GetReviewsByDoctorID_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock db: %v", err)
	}
	defer db.Close()

	repo := repositories.NewReviewRepository(db)

	mock.ExpectQuery("SELECT patient_id, doctor_id, content, rating, timestamp FROM reviews WHERE doctor_id = ?").
		WithArgs(1).
		WillReturnError(sql.ErrConnDone)

	reviews, err := repo.GetReviewsByDoctorID("1")
	assert.Error(t, err)
	assert.Nil(t, reviews)
}
