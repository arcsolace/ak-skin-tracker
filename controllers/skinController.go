package controllers

import (
    "context"
	// "fmt"
    // "log"
    // "math"
    "strconv"
    "time"

    f "github.com/gofiber/fiber/v2"
    "github.com/arcsolace/ak-skin-tracker/config"
    m "github.com/arcsolace/ak-skin-tracker/models"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllSkins(c *f.Ctx) error {
	skinCol := config.MI.DB.Collection("skins")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var skins []m.Skin

	filter := bson.M{}
	sort := bson.D{{"rarity", -1}}
	opts := options.Find().SetSort(sort)

	cursor, err := skinCol.Find(ctx, filter, opts)

	if err != nil {
		return c.Status(f.StatusNotFound).JSON(f.Map{
			"success":  false,
			"message":  "No skins found!",
			"error" : err,
		})
	}

	for cursor.Next(ctx) {
		var skin m.Skin
		cursor.Decode(&skin)
		skins = append(skins, skin)
	}

	defer cursor.Close(ctx)

	return c.Status(f.StatusOK).JSON(f.Map{
		"data":	skins,
		"total": len(skins),
	})
}

func GetSkinByID(c *f.Ctx) error {
	skinCol := config.MI.DB.Collection("skins")
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	skinID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(f.StatusNotFound).JSON(f.Map{
			"success": false,
			"message": "Invalid ID!",
		})
	}

	var result m.Skin

	filter := bson.M{"skin_id": skinID}
	ferr := skinCol.FindOne(ctx, filter).Decode(&result)

	if ferr != nil {
		if ferr == mongo.ErrNoDocuments {
			return c.Status(f.StatusNotFound).JSON(f.Map{
				"success": false,
				"message": "Skin not found!",
			})
		}
		return c.Status(f.StatusInternalServerError).JSON(f.Map{
			"success": false,
			"message": "Error retrieving skin!",
			"error":   err.Error(),
		})
	}

	return c.Status(f.StatusOK).JSON(f.Map{
		"success": true,
		"data":    result,
	})
}