package controllers

import (
    "context"
    "time"

    f "github.com/gofiber/fiber/v2"
    "github.com/arcsolace/ak-skin-tracker/config"
    m "github.com/arcsolace/ak-skin-tracker/models"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSharedSkins(c *f.Ctx) error {
	userCol := config.MI.DB.Collection("users")
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	shareCode := c.Params("share_code")
	filter := bson.M{"share_code": shareCode}

	var sharedUser m.SharedUser

	err := userCol.FindOne(ctx, filter).Decode(&sharedUser)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(f.StatusNotFound).JSON(f.Map{
				"success": false,
				"message": "User not found!",
			})
		}
		return c.Status(f.StatusInternalServerError).JSON(f.Map{
			"success": false,
			"message": "Error retrieving user!",
			"error":   err.Error(),
		})
	}

	skinCol := config.MI.DB.Collection("skins")
	skinFilter := bson.M{"skin_id": bson.M{"$in": sharedUser.Skins}}

	cursor, err := skinCol.Find(ctx, skinFilter)
	if err != nil {
		return c.Status(f.StatusInternalServerError).JSON(f.Map{
			"success": false,
			"message": "Error retrieving skins!",
			"error":   err.Error(),
		})
	}
	defer cursor.Close(ctx)

	var skins []m.Skin
	for cursor.Next(ctx) {
		var skin m.Skin
		if err := cursor.Decode(&skin); err != nil {
			return c.Status(f.StatusInternalServerError).JSON(f.Map{
				"success": false,
				"message": "Error decoding skin!",
				"error":   err.Error(),
			})
		}
		skins = append(skins, skin)
	}

	if err := cursor.Err(); err != nil {
		return c.Status(f.StatusInternalServerError).JSON(f.Map{
			"success": false,
			"message": "Error iterating over skins!",
			"error":   err.Error(),
		})
	}

	return c.Status(f.StatusOK).JSON(f.Map{
		"success": true,
		"data": f.Map{
			"skins":     skins,
			"user_name": sharedUser.UserName,
		},
	})
}