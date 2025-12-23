package services

import (
	"errors"
	"readingtracker/internal/dto"
	"readingtracker/internal/models"
	"testing"
)

// fake repo
type FakeProgressRepo struct {
	Progress *models.ReadingProgress
	Err      error
}

func (f *FakeProgressRepo) GetByBookID(bookID uint) (*models.ReadingProgress, error) {

	if f.Progress == nil {
		return nil, errors.New("record not found")
	}
	return f.Progress, nil
}

func (f *FakeProgressRepo) Save(p *models.ReadingProgress) error {
	f.Progress = p
	return nil
}

func TestUpdateProgress_NewEntry(t *testing.T) {

	repo := &FakeProgressRepo{Progress: nil}
	service := NewProgressService(repo)

	req := dto.UpdateProgressRequest{
		Status:      "Reading",
		CurrentPage: 50,
	}

	err := service.UpdateProgress(1, req)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if repo.Progress.CurrentPage != 50 {
		t.Errorf("Expected CurrentPage 50, but got %d", repo.Progress.CurrentPage)
	}
	if repo.Progress.Status != "Reading" {
		t.Errorf("Expected Status 'Reading', but got %s", repo.Progress.Status)
	}
}

func TestUpdateProgress_ExistingEntry(t *testing.T) {

	existing := &models.ReadingProgress{
		ID:          1,
		BookID:      1,
		Status:      "Reading",
		CurrentPage: 20,
	}
	repo := &FakeProgressRepo{Progress: existing}
	service := NewProgressService(repo)

	req := dto.UpdateProgressRequest{
		Status:      "Completed",
		CurrentPage: 200,
	}

	err := service.UpdateProgress(1, req)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if repo.Progress.Status != "Completed" {
		t.Errorf("Expected Status 'Completed', but got %s", repo.Progress.Status)
	}
	if repo.Progress.CurrentPage != 200 {
		t.Errorf("Expected CurrentPage 200, but got %d", repo.Progress.CurrentPage)
	}
}

func TestGetProgress_Success(t *testing.T) {

	existing := &models.ReadingProgress{BookID: 1, Status: "Reading"}
	repo := &FakeProgressRepo{Progress: existing}
	service := NewProgressService(repo)

	result, err := service.GetProgress(1)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if result.Status != "Reading" {
		t.Errorf("Expected Status 'Reading', but got %s", result.Status)
	}
}
