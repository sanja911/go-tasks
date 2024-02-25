package userControllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"task-management-system/configs"
	"task-management-system/models"
	"task-management-system/responses"
	"task-management-system/security"
	util "task-management-system/utils"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func Register(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data: &fiber.Map{
				"data": err.Error(),
			}})
	}

	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			}})
	}

	user.Email = util.NormalizeEmail(user.Email)
	if !govalidator.IsEmail(user.Email) {
		return c.Status(http.StatusBadRequest).JSON(util.NewJSONrror(util.ErrInvalidEmail))
	}

	existingUsers := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&user)
	if existingUsers == nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data: &fiber.Map{
				"data": "This email is already used on another users",
			}})
	}

	user.Password = security.EncryptPassword(user.Password)
	userEntry := models.User{
		Id:       primitive.NewObjectID(),
		Name:     user.Name,
		Email:    user.Email,
		Location: user.Location,
		Password: user.Password,
	}

	result, err := userCollection.InsertOne(ctx, userEntry)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data: &fiber.Map{
				"data": err.Error(),
			}})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data: &fiber.Map{
			"data": result,
		}})

}

func Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var input models.User
	var user models.User
	defer cancel()

	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: "error",
			Data: &fiber.Map{
				"data": err.Error(),
			}})
	}

	input.Email = util.NormalizeEmail(input.Email)
	err = userCollection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&user)
	if err != nil {
		log.Printf("%s signin failed: %v\n", input.Email, err.Error())
		return c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{
			Status:  http.StatusUnauthorized,
			Message: "signin failed",
			Data: &fiber.Map{
				"data": err.Error(),
			}})
	}

	err = security.VerifyPassword(user.Password, input.Password)
	if err != nil {
		return c.Status(http.StatusUnauthorized).
			JSON(responses.UserResponse{
				Status:  http.StatusUnauthorized,
				Message: "Invalid Credentials",
				Data: &fiber.Map{
					"data": err.Error(),
				}})
	}

	token, err := security.NewToken(input.Id.Hex())
	if err != nil {
		return c.Status((http.StatusUnauthorized)).JSON(responses.UserResponse{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized Access",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	return c.Status(http.StatusOK).
		JSON(responses.UserResponse{
			Status:  http.StatusAccepted,
			Message: "Success",
			Data: &fiber.Map{
				"user":  user,
				"token": fmt.Sprintf("Bearer %s", token),
			}})
}
