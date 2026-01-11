package service

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/faizallmaullana/lenteng-agung/backend/internal/domains/repo"
	"github.com/faizallmaullana/lenteng-agung/backend/internal/pkg/utils"
)

type UploadService struct {
	repo repo.FileUploadRepo
}

func NewUploadService(repo repo.FileUploadRepo) *UploadService {
	return &UploadService{repo: repo}
}

func (s *UploadService) UploadFile(fileData []byte, fileName string, fileType string) (string, error) {
	uploadsDir := filepath.Join("static", "uploads", fileType)
	if err := os.MkdirAll(uploadsDir, 0755); err != nil {
		return "", err
	}

	generated := utils.GenerateFilename(fileName)
	outPath := filepath.Join(uploadsDir, generated)

	if err := os.WriteFile(outPath, fileData, 0644); err != nil {
		return "", err
	}

	// persist metadata (userID currently not tracked here)
	if err := s.repo.SaveFileMetadata(generated, outPath, ""); err != nil {
		return "", err
	}

	return generated, nil
}

// OcrExtract uploads the file to the external AI OCR service (configured via AI_OCR env var)
// It performs a simple healthcheck and then POSTs the file to /extract. Returns the service response body.
func (s *UploadService) OcrExtract(fileData []byte, fileName string, fileType string) (string, error) {
	uploadsDir := filepath.Join("static", "uploads", fileType)
	if err := os.MkdirAll(uploadsDir, 0755); err != nil {
		return "", err
	}

	generated := utils.GenerateFilename(fileName)
	outPath := filepath.Join(uploadsDir, generated)

	if err := os.WriteFile(outPath, fileData, 0644); err != nil {
		return "", err
	}

	// persist metadata (save uploaded file record)
	if err := s.repo.SaveFileMetadata(generated, outPath, ""); err != nil {
		return "", err
	}

	aiBase := os.Getenv("AI_OCR")
	if strings.TrimSpace(aiBase) == "" {
		return "", fmt.Errorf("AI_OCR environment variable not set")
	}

	// healthcheck: try base URL then /health
	client := &http.Client{Timeout: 5 * time.Second}
	healthURL := strings.TrimRight(aiBase, "/")
	resp, err := client.Get(healthURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		// try /health
		resp, err = client.Get(healthURL + "/health")
		if err != nil || resp.StatusCode != http.StatusOK {
			return "", fmt.Errorf("AI_OCR service healthcheck failed")
		}
	}
	if resp.Body != nil {
		resp.Body.Close()
	}

	// prepare multipart form
	var b bytes.Buffer
	mw := multipart.NewWriter(&b)
	fw, err := mw.CreateFormFile("file", generated)
	if err != nil {
		return "", err
	}
	if _, err := fw.Write(fileData); err != nil {
		return "", err
	}
	mw.Close()

	extractURL := healthURL + "/extract"
	req, err := http.NewRequest("POST", extractURL, &b)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", mw.FormDataContentType())

	// send request
	resp, err = client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("AI_OCR error (%d): %s", resp.StatusCode, string(body))
	}

	return string(body), nil
}
