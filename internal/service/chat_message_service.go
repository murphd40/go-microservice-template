package service

import (
	"github.com/murphd40/go-microservice-template/internal/dao/model"
	"github.com/murphd40/go-microservice-template/internal/dao/repository"
)

type ChatMessageService interface {
	CreateChatMessage(model.ChatMessage) model.ChatMessage
	GetChatMessageById(string) (model.ChatMessage, bool)
}

type chatMessageServiceImpl struct {
	chatMessageRepository repository.ChatMessageRepository
}

func NewChatMessageService(chatMessageRepository repository.ChatMessageRepository) ChatMessageService {
	return &chatMessageServiceImpl{
		chatMessageRepository: chatMessageRepository,
	}
}

func (c *chatMessageServiceImpl) CreateChatMessage(chatMessage model.ChatMessage) model.ChatMessage {
	result, _ := c.chatMessageRepository.Insert(chatMessage)
	return result
}

func (c *chatMessageServiceImpl) GetChatMessageById(chatMessageId string) (model.ChatMessage, bool) {
	return c.chatMessageRepository.FindById(chatMessageId)
}
