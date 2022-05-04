package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/murphd40/go-microservice-template/internal/dao/model"
	log "github.com/murphd40/go-microservice-template/internal/logging"
	types "github.com/murphd40/go-microservice-template/internal/server/datatypes"
	"github.com/murphd40/go-microservice-template/internal/service"
	"github.com/murphd40/go-microservice-template/internal/utils"
)

type ChatMessageHandler struct {
	chatMessageService service.ChatMessageService
}

func NewChatMessageHandler(chatMessageService service.ChatMessageService) *ChatMessageHandler {
	return &ChatMessageHandler{
		chatMessageService: chatMessageService,
	}
}

func (h *ChatMessageHandler) CreateChatMessage(w http.ResponseWriter, r *http.Request) {
	bs, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Error("Failed to read request body.", err)
	}

	data := types.ChatMessageRequest{}
	json.Unmarshal(bs, &data)

	chatMessage := model.ChatMessage{
		Content: data.Content,
		CreatedBy: "admin",
	}

	chatMessage = h.chatMessageService.CreateChatMessage(chatMessage)

	response := types.ChatMessageResponse{}
	utils.Convert(chatMessage, &response)

	jsonResponse(w, response)
}

func (h *ChatMessageHandler) GetChatMessageById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["chatMessageId"]

	chatMessage, ok := h.chatMessageService.GetChatMessageById(id)

	if !ok {
		w.WriteHeader(404)
		return
	}

	response := types.ChatMessageResponse{}
	utils.Convert(chatMessage, &response)
	jsonResponse(w, response)
}

func jsonResponse(w http.ResponseWriter, data any) {
	bs, _ := json.MarshalIndent(data, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(bs)
}