package handler

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	vd "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

type (
	EditFigureResponse struct {
		ID             int     `json:"id"`
		User           string  `json:"user"`
		UserId         int     `json:"user_id"`
		FigureData     *Figure `json:"figure_data"`
		Script         string  `json:"script"`
		CommentRequest string  `json:"comment_request"`
	}
)

// PUT /api/v2/files/edit
func (h *Handler) EditFigure(c echo.Context) error {
	req := new(EditFigureResponse)
	if req.FigureData == nil {
		req.FigureData = new(Figure)
	}
	if err := c.Bind(req); err != nil {
		log.Printf("Error creating request: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
	}

	err := vd.ValidateStruct(
		req,
		// vd.Field(&req.ID, vd.Required),
		// vd.Field(&req.User, vd.Required),
		// vd.Field(&req.UserId, vd.Required),
		// vd.Field(&req.FigureData, vd.Required),
	)
	if err != nil {
		log.Printf("Error validating request: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err)).SetInternal(err)
	}

	fileName := setFileName()

	// if req.Script != nil {}

	script, outputPath, err := makeScript(fileName, req.FigureData)
	if err != nil {
		log.Printf("Error making script: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to make script").SetInternal(err)
	}

	err = makePlotFile(fileName, script)
	if err != nil {
		log.Printf("Error making plot file: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to make plot file").SetInternal(err)
	}

	err = sendPlotFile(outputPath, c, 0, 0)
	if err != nil {
		log.Printf("Error sending plot file: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to send plot file").SetInternal(err)
	}

	return c.JSON(http.StatusOK, &EditFigureResponse{
		ID:             req.ID,
		User:           req.User,
		UserId:         req.UserId,
		FigureData:     req.FigureData,
		Script:         script,
		CommentRequest: req.CommentRequest,
	})
}

func setFileName() string {
	currentTime := time.Now().String()
	removeChar := []string{":", " ", "+"}
	for _, char := range removeChar {
		currentTime = strings.Replace(currentTime, char, "_", -1)
	}
	return currentTime
}

func makePlotFile(fileName, script string) error {
	err := os.WriteFile("/app/scripts/"+fileName+".plt", []byte(script), fs.FileMode(0644))
	if err != nil {
		log.Printf("ファイルの書き込み中にエラーが発生しました: %v", err)
		return err
	}
	return nil
}

func sendPlotFile(outputPath string, c echo.Context, count int, beforeFileSize int64) error {
	if count > 30 {
		return errors.New("count over 100")
	}

	// ファイルを開く
	file, err := os.Open(outputPath)
	if os.IsNotExist(err) {
		// ファイルが存在しない場合は100ms待機して再度開く
		count++
		// 100ms待機
		time.Sleep(100 * time.Millisecond)
		return sendPlotFile(outputPath, c, count, 0)
	} else if err != nil {
		log.Printf("Error opening file: %v", err)
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Printf("failed to get file info: %v", err)
		return err
	}

	if fileInfo.Size() != beforeFileSize || fileInfo.Size() == 0 {
		// ファイルが書き込み中の場合は100ms待機して再度開く
		count++
		// 100ms待機
		time.Sleep(100 * time.Millisecond)
		return sendPlotFile(outputPath, c, count, fileInfo.Size())
	}

	// レスポンスにファイルの内容をコピー
	_, err = io.Copy(c.Response().Writer, file)
	if err != nil {
		log.Printf("Error copying file to response: %v", err)
		return err
	}
	return nil
}
