package handler

import (
	"io"

	"github.com/faizallmaullana/lenteng-agung/backend/internal/domains/service"
	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	svc service.UploadService
}

func NewUploadHandler(svc service.UploadService) *UploadHandler {
	return &UploadHandler{svc: svc}
}

func (h *UploadHandler) UploadFile(c *gin.Context) {
	fileType := c.Param("file_type")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "file is required"})
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to open uploaded file"})
		return
	}
	defer openedFile.Close()

	fileData := make([]byte, file.Size)
	_, err = openedFile.Read(fileData)
	if err != nil && err != io.EOF {
		c.JSON(500, gin.H{"error": "failed to read uploaded file"})
		return
	}

	savedFileName, err := h.svc.UploadFile(fileData, file.Filename, fileType)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to save uploaded file"})
		return
	}

	c.JSON(200, gin.H{"file_name": savedFileName})
}

func (h *UploadHandler) OcrExtract(c *gin.Context) {
	fileType := c.Param("file_type")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "file is required"})
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to open uploaded file"})
		return
	}
	defer openedFile.Close()

	fileData := make([]byte, file.Size)
	_, err = openedFile.Read(fileData)
	if err != nil && err != io.EOF {
		c.JSON(500, gin.H{"error": "failed to read uploaded file"})
		return
	}

	response, err := h.svc.OcrExtract(fileData, file.Filename, fileType)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to extract OCR"})
		return
	}

	c.JSON(200, gin.H{"response": response})
}
