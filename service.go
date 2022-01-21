package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type GetPresenceBody struct {
	authToken string
	userId    string
}

type NotifyPresenceBody struct {
	authToken string `form:"authToken" json:"authToken" xml:"authToken"  binding:"required"`
	userId    string `form:"userId" json:"userId" xml:"userId" binding:"required"`
}

type UserPresenceService struct {
	redisClient *RedisClient
	defaultTTL  time.Duration
}

func CreateUserPresenceService(redisClient *RedisClient, defaultTTL time.Duration) *UserPresenceService {
	return &UserPresenceService{redisClient: redisClient, defaultTTL: defaultTTL}
}

func (service *UserPresenceService) updateUserPresence(userId string) bool {
	_, err := service.redisClient.SetKey(userId, "online", service.defaultTTL)
	return err == nil
}

func (service *UserPresenceService) isUserOnline(userId string) bool {
	_, err := service.redisClient.GetKey(userId)
	return err == nil
}

func (service UserPresenceService) HandlerNotifyUserPresence(rw http.ResponseWriter, r *http.Request) {
	// FIX ME
	var status = service.updateUserPresence("testUser")
	var notifyPresenseBody NotifyPresenceBody
	log.Println("Request Body")

	err := json.NewDecoder(r.Body).Decode(&notifyPresenseBody)
	if err != nil {
		log.Println(err)
	}
	log.Println("Request body", notifyPresenseBody)
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"updated": status,
	})
}

func (service UserPresenceService) HandlerGetUserPresence(rw http.ResponseWriter, r *http.Request) {
	// FIX ME
	var status = service.isUserOnline("testUser")
	var getPresenceBody GetPresenceBody
	err := json.NewDecoder(r.Body).Decode(&getPresenceBody)
	if err != nil {
		log.Println(err)
	}
	log.Println("Request body", getPresenceBody)
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"isOnline": status == true,
	})
}
