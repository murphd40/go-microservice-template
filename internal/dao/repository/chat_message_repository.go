package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/murphd40/go-microservice-template/internal/dao/model"
)

type ChatMessageRepository interface {
	Insert(model.ChatMessage) (model.ChatMessage, error)
	FindById(string) (model.ChatMessage, bool)
	DeleteById(string) bool
}

func NewChatMessageRepository() ChatMessageRepository {
	return &chatMessageRepositoryImpl{
		chatMessages: make(map[string]model.ChatMessage),
	}
}

type chatMessageRepositoryImpl struct {
	chatMessages map[string]model.ChatMessage
}

func (c *chatMessageRepositoryImpl) Insert(chatMessage model.ChatMessage) (model.ChatMessage, error) {
	chatMessage.Id = uuid.NewString()
	chatMessage.CreatedAt = time.Now()

	c.chatMessages[chatMessage.Id] = chatMessage

	return chatMessage, nil
}

func (c *chatMessageRepositoryImpl) FindById(chatMessageId string) (model.ChatMessage, bool) {
	chatMessage, ok := c.chatMessages[chatMessageId]
	return chatMessage, ok
}

func (c *chatMessageRepositoryImpl) DeleteById(chatMessageId string) bool {
	if _, ok := c.chatMessages[chatMessageId]; ok {
		delete(c.chatMessages, chatMessageId)
		return true
	}

	return false
}
