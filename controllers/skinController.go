package controllers

import (
    "context"
    "log"
    "math"
    "strconv"
    "time"

    f "github.com/gofiber/fiber/v2"
    "github.com/arcsolace/ak-skin-tracker/config"
    m "github.com/arcsolace/ak-skin-tracker/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllSkins(c *f.Ctx) error {
	skinCol := config.MI.DB.Collection("skins")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var skins []m.SkinBrief

	filter := bson.M{}
    findOptions := options.Find()

	cursor, err := skinCol.Find(ctx, filter, findOptions)
	defer cursor.Close(ctx)

	if err != nil {
		return c.Status(f.StatusNotFound).JSON(f.Map{
			"success":  false,
			"message":  "No skins found!",
			"error": 	err,
		})
	}

	for cursor.Next(ctx) {
		var skin m.SkinBrief
		cursor.Decode(&skin)
		skins = append(skins, skin)
	}

	return c.Status(f.StatusOK).JSON(f.Map{
		"data":	skins,
		"total": len(skins),
	})
}