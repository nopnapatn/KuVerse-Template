package controllers

import (
	"context"
	"kuverse/config"
	"kuverse/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PayloadReport struct {
	UserUid  string                 `json:"user_uid"`
	VerseUid string                 `json:"verse_uid"`
	Type     string                 `json:"type"`
	Note     string                 `json:"note"`
	Data     map[string]interface{} `json:"data"`
}
type CountReport struct {
	Type  string `json:"type"`
	Count int64  `json:"count"`
}

type DataReport struct {
	Pages        int             `json:"pages"`
	CurrentPage  int             `json:"current_page"`
	CurrentCount int             `json:"current_count"`
	TotalCount   int64           `json:"total_count"`
	Result       []models.Report `json:"result"`
}

var client *mongo.Client

func InitMongoDB() {
	clientOptions := options.Client().ApplyURI(config.GetMongoURI())
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func GetAllReport(c *fiber.Ctx) error {

	InitMongoDB()
	if client == nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to connect to MongoDB")
	}

	collection := client.Database("staging-user-report-api").Collection("report")
	cursor, err := collection.Find(context.Background(), bson.M{})
	defer client.Disconnect(context.Background())
	var reports []models.Report

	for cursor.Next(context.Background()) {
		var report models.Report
		if err := cursor.Decode(&report); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		reports = append(reports, report)
	}

	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	return c.JSON(fiber.Map{
		"status": "ok",
		"result": reports,
	})

}
