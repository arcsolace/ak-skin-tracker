package controllers

import (
    "context"
	// "fmt"
    // "log"
    // "math"
    // "strconv"
    "time"
	// "encoding/json"

    f "github.com/gofiber/fiber/v2"
    "github.com/arcsolace/ak-skin-tracker/config"
    m "github.com/arcsolace/ak-skin-tracker/models"
	u "github.com/arcsolace/ak-skin-tracker/utils"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/bson/primitive"
    // "go.mongodb.org/mongo-driver/mongo/options"
)

func GetUserSkins(c *f.Ctx) error {
	userCol := config.MI.DB.Collection("users")
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	userCode := c.Params("user_code")
	filter := bson.M{"user_code": userCode}

	var user m.User

	err := userCol.FindOne(ctx, filter).Decode(&user)

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
	skinFilter := bson.M{"skin_id": bson.M{"$in": user.Skins}}

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
		"data":    skins,
	})
}

func UpdateUserSkins(c *f.Ctx) error {
	userCol := config.MI.DB.Collection("users")
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	userCode := c.Params("user_code")

	var req struct {
		SkinIDs []int `json:"skin_ids"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(f.StatusBadRequest).JSON(f.Map{
			"success": false,
			"message": "Invalid request body!",
		})
	}

	filter := bson.M{"user_code": userCode}
	update := bson.M{
		"$addToSet": bson.M{"skins": bson.M{"$each": req.SkinIDs}},
	}

	updateResult, err := userCol.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.Status(f.StatusInternalServerError).JSON(f.Map{
			"success": false,
			"message": "Error updating user skins!",
			"error":   err.Error(),
		})
	}

	if updateResult.MatchedCount == 0 {
		return c.Status(f.StatusNotFound).JSON(f.Map{
			"success": false,
			"message": "User not found!",
		})
	}

	return c.Status(f.StatusOK).JSON(f.Map{
		"success": true,
		"message": "User skins updated successfully!",
	})
}

func CreateUser(c *f.Ctx) error {
    userCol := config.MI.DB.Collection("users")
    ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
    defer cancel()

    var req struct {
        UserName string `json:"user_name"`
    }
    if err := c.BodyParser(&req); err != nil {
        return c.Status(f.StatusBadRequest).JSON(f.Map{
            "success": false,
            "message": "Invalid request body!",
        })
    }

    userCode := u.GenerateRandomUserCode(6)
    newUser := m.User{
        UserCode: userCode,
        UserName: req.UserName,
        Skins:    []int{},
    }

    _, err := userCol.InsertOne(ctx, newUser)
    if err != nil {
        return c.Status(f.StatusInternalServerError).JSON(f.Map{
            "success": false,
            "message": "Error creating user!",
            "error":   err.Error(),
        })
    }

    return c.Status(f.StatusOK).JSON(f.Map{
        "success":    true,
        "message":    "User created successfully!",
        "user_code":  userCode,
    })
}