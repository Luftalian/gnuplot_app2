package handler

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func (h *Handler) UploadFile(c echo.Context) error {
	// フォームからファイルを取得
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "No file is received"})
	}

	src, err := file.Open()
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return err
	}
	defer src.Close()

	// 保存するディレクトリとファイル名を決定
	dstDir := "/app/data"
	if _, err := os.Stat(dstDir); os.IsNotExist(err) {
		err := os.Mkdir(dstDir, os.ModePerm)
		if err != nil {
			log.Printf("Error creating directory: %v", err)
			return err
		}
	}
	dstPath := filepath.Join(dstDir, file.Filename)

	// ファイルを保存
	dst, err := os.Create(dstPath)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		log.Printf("Error copying file: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "File uploaded successfully", "filePath": dstPath})
}
