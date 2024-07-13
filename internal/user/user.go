package user

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	v "github.com/core-go/core/v10"
	"github.com/core-go/mongo/adapter"
	"github.com/core-go/search"

	"go-service/internal/user/handler"
	"go-service/internal/user/model"
	"go-service/internal/user/repository/query"
	"go-service/internal/user/service"
)

type UserTransport interface {
	All(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(db *mongo.Database, logError func(context.Context, string, ...map[string]interface{})) (UserTransport, error) {
	validator, err := v.NewValidator()
	if err != nil {
		return nil, err
	}

	// userRepository := adapter.NewUserAdapter(db, query.BuildQuery)
	userRepository := adapter.NewSearchAdapter[model.User, string, *model.UserFilter](db, "users", query.BuildQuery, search.GetSort)
	userService := service.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userService, validator.Validate, logError)
	return userHandler, nil
}
