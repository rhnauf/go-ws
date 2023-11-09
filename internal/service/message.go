package service

import (
	"go-ws/internal/entity"
	"go-ws/internal/repository"
)

type MessageService interface {
	CreateMessage(message *entity.Message) error
	GetListMessage(request *entity.RequestFetchMessage) ([]*entity.Message, error)
}

type messageService struct {
	messageRepository repository.MessageRepository
}

func NewMessageService(messageRepository repository.MessageRepository) *messageService {
	return &messageService{messageRepository: messageRepository}
}

func (m *messageService) CreateMessage(message *entity.Message) error {
	return m.messageRepository.Create(message)
}

func (m *messageService) GetListMessage(request *entity.RequestFetchMessage) ([]*entity.Message, error) {
	messageList, err := m.messageRepository.Fetch(request)
	if err != nil {
		return nil, err
	}

	return messageList, nil
}
