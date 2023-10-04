package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Id                   string `json:"id" `
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	Onboarded            bool   `json:"Onboarded"`
	Fqdns                string `json:"Fqdns"`
	TrackingPageUrl      string `json:"TrackingPageUrl  "`
	ReplaceTrackingLinks *bool  `json:"ReplaceTrackingLinks "`
	CurrentBillingId     string `json:"CurrentBillingId  "`
	IsFreeTrialUsed      bool   `json:"IsFreeTrialUsed "`
	AuthProvider         string `json:"AuthProvider "`
	CountryCode          string `json:"CountryCode"`
	Currency             string `json:"Currency  "`
	Category             string `json:"Category   "`
	Verified             bool   `json:"Verified "`
	EmailOptIn           *bool  `json:"EmailOptIn  "`
	SmsOptIn             *bool  `json:"SmsOptIn   "`
	ShopifyAppScopes     string `json:"ShopifyAppScopes   "`
	CreatedAt            string `json:"CreatedAt"`
}

var todos = []todo{

	{Id: "1",
		Name:                 "Sarthak",
		Email:                "dragmaster019@gmail.com",
		Phone:                "7872129806",
		Onboarded:            false,
		Fqdns:                "Null",
		TrackingPageUrl:      "01",
		ReplaceTrackingLinks: new(bool),
		CurrentBillingId:     "01",
		IsFreeTrialUsed:      false,
		AuthProvider:         "shopify",
		CountryCode:          "+91",
		Currency:             "INR",
		Category:             "e commerce",
		Verified:             false,
		EmailOptIn:           new(bool),
		SmsOptIn:             new(bool),
		ShopifyAppScopes:     "01",
		CreatedAt:            "4/10/23"},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)

}
func addTodos(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.JSON(http.StatusCreated, newTodo)
}
func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodos)
	router.Run("localhost:9090")

}
