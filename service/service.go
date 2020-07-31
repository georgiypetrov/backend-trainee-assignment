package service

import (
	"encoding/json"
	"github.com/allerria/backend-trainee-assignment/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Service struct {
	DB     *models.DB
	Server http.Server
}

type CreateUserRequestBody struct {
	Username string `json:"username"`
}

type CreateChatRequestBody struct {
	Name  string   `json:"name"`
	Users []string `json:"users"`
}

type CreateMessageRequestBody struct {
	Chat   string `json:"chat"`
	Author string `json:"author"`
	Text   string `json:"text"`
}

type GetUserChatsRequestBody struct {
	ID string `json:"user"`
}

type GetChatMessagesRequestBody struct {
	Chat string `json:"chat"`
}

func CreateRouter(s *Service) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users/add", s.createUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/chats/add", s.creatChatHandler).Methods(http.MethodPost)
	r.HandleFunc("/chats/get", s.getUserChatsHandler).Methods(http.MethodPost)
	r.HandleFunc("/messages/add", s.createMessageHandler).Methods(http.MethodPost)
	r.HandleFunc("/messages/get", s.getChatMessagesHandler).Methods(http.MethodPost)
	return r
}

func (s *Service) createUserHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := CreateUserRequestBody{}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id, err := s.DB.CreateUser(data.Username)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	msg, err := json.Marshal(map[string]string{"id": id})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(msg)
}

func (s *Service) creatChatHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := CreateChatRequestBody{}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id, err := s.DB.CreateChat(data.Name, data.Users)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	msg, err := json.Marshal(map[string]uint64{"id": id})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(msg)
}

func (s *Service) createMessageHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := CreateMessageRequestBody{}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var chatID int
	chatID, err = strconv.Atoi(data.Chat)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id, err := s.DB.CreateMessage(uint64(chatID), data.Author, data.Text)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	msg, err := json.Marshal(map[string]uint64{"id": id})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(msg)
}

func (s *Service) getUserChatsHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := GetUserChatsRequestBody{}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	chats, err := s.DB.GetUserChats(data.ID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	msg, err := json.Marshal(chats)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(msg)
}

func (s *Service) getChatMessagesHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := GetChatMessagesRequestBody{}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var chatID int
	chatID, err = strconv.Atoi(data.Chat)
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	chats, err := s.DB.GetChatMessages(uint64(chatID))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	msg, err := json.Marshal(chats)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(msg)
}
