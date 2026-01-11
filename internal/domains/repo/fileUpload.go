package repo

import (
	"time"

	"github.com/faizallmaullana/lenteng-agung/backend/internal/models"
	"gorm.io/gorm"
)

type FileUploadRepo struct {
	db *gorm.DB
}

func NewFileUploadRepo(db *gorm.DB) *FileUploadRepo {
	return &FileUploadRepo{db: db}
}

func (r *FileUploadRepo) SaveFileMetadata(fileName string, filePath string, userID string) error {
	model := &models.FileUpload{
		FileName:  fileName,
		FilePath:  filePath,
		Timestamp: time.Now(),
	}

	if err := r.db.Create(&model).Error; err != nil {
		return err
	}

	return nil
}

func (r *FileUploadRepo) GetFileMetadata(fileID string) (string, string, error) {
	return "", "", nil
}
