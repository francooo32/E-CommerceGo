package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	productCollection *mongo.Collection
	userCollection    *mongo.Collection

	func NewApplication(productCollection, userCollection *mongo.Collection) *Application{
		return &Application{
			productCollection: productCollection,
			userCollection: userCollection
		}
	}
}

func AddToCart() gin.Handler {

}

func RemoteItem() gin.HandlerFunc {

}

func GetItemFromCart() gin.HandlerFunc {

}

func BuyFromCart() gin.HandlerFunc {

}

func InstantBuy() gin.HandlerFunc {

}
