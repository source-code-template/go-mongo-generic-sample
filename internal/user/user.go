package user

import (
	"net/http"

	repo "github.com/core-go/mongo/repository"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/core-go/core"
	v "github.com/core-go/core/validator"
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

func NewUserHandler(db *mongo.Database, logError core.Log, action *core.ActionConfig) (UserTransport, error) {
	validator, err := v.NewValidator[*model.User]()
	if err != nil {
		return nil, err
	}

	// userRepository := adapter.NewUserAdapter(db, query.BuildQuery)
	userRepository := repo.NewSearchRepository[model.User, string, *model.UserFilter](db, "users", query.BuildQuery, search.GetSort)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService, logError, validator.Validate, action)
	return userHandler, nil
}
