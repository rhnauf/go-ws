package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-ws/internal/entity"
	"go-ws/internal/service"
	"log"
	"net/http"
	"strings"
	"time"
)

type MessageHandler interface {
	GetListMessage(ctx *fiber.Ctx) error
}

type messageHandler struct {
	messageService service.MessageService
}

func NewMessageHandler(messageService service.MessageService) *messageHandler {
	return &messageHandler{messageService: messageService}
}

type responseMessage struct {
	StatusCode int32       `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func (m *messageHandler) GetListMessage(ctx *fiber.Ctx) error {
	// validate
	now := time.Now()
	startDateParam := ctx.Query("start_date")
	endDateParam := ctx.Query("end_date")

	startDate, err := time.Parse("2006-01-02 15:04:05", startDateParam)
	if err != nil {
		log.Println(err)
		startDate = now
	}

	endDate, err := time.Parse("2006-01-02 15:04:05", endDateParam)
	if err != nil {
		log.Println(err)
		endDate = now
	}

	req := &entity.RequestFetchMessage{
		Username:  strings.Trim(ctx.Query("username"), " "),
		Recipient: strings.Trim(ctx.Query("recipient"), " "),
		StartDate: startDate.Format("2006-01-02 15:04:05"),
		EndDate:   endDate.Format("2006-01-02 15:04:05"),
	}

	res, err := m.messageService.GetListMessage(req)
	if err != nil {
		return ctx.JSON(responseMessage{
			StatusCode: http.StatusInternalServerError,
			Message:    "internal server error",
			Data:       nil,
		})
	}

	return ctx.JSON(responseMessage{
		StatusCode: http.StatusOK,
		Message:    "success get list messages",
		Data:       res,
	})
}
