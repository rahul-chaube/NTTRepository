package handler

import (
	"NTTHomeTestDemo/middleware"
	"NTTHomeTestDemo/model"
	"NTTHomeTestDemo/utility"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
}

func (uh UserHandler) NewUserHandler(handler *mux.Router) {

	handler.Handle("/login", middleware.Logging(http.HandlerFunc(uh.Login))).Methods("Post")
}

func (uh UserHandler) Login(res http.ResponseWriter, req *http.Request) {

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		userResponse := model.UserResponse{
			Message: "Error reading request body",
		}
		res.Header().Set(utility.ContentType, utility.ApplicationJSON)

		res.WriteHeader(http.StatusBadRequest)

		resp, _ := json.Marshal(userResponse)
		res.Write(resp)
		return
	}
	defer req.Body.Close()
	user := model.User{}
	err = json.Unmarshal(data, &user)
	if err != nil {
		userResponse := model.UserResponse{
			Message: "Error reading request body",
		}
		res.Header().Set(utility.ContentType, utility.ApplicationJSON)

		res.WriteHeader(http.StatusInternalServerError)

		resp, _ := json.Marshal(userResponse)
		res.Write(resp)
		return
	}
	password, isExists := model.AllowUser[user.Username]
	if !isExists {
		userResponse := model.UserResponse{
			Message: "User is not found " + user.Username,
		}
		res.Header().Set(utility.ContentType, utility.ApplicationJSON)

		res.WriteHeader(http.StatusNotFound)

		resp, _ := json.Marshal(userResponse)
		res.Write(resp)
		return
	}
	if password != user.Password {
		userResponse := model.UserResponse{
			Message: "Invalid username/password",
		}
		res.Header().Set(utility.ContentType, utility.ApplicationJSON)

		res.WriteHeader(http.StatusUnauthorized)

		resp, _ := json.Marshal(userResponse)
		res.Write(resp)
		return
	}
	token, err := utility.CreateToken(user.Username)
	if err != nil {
		userResponse := model.UserResponse{
			Message: "Error while creating token" + err.Error(),
		}
		res.Header().Set(utility.ContentType, utility.ApplicationJSON)

		res.WriteHeader(http.StatusInternalServerError)

		resp, _ := json.Marshal(userResponse)
		res.Write(resp)
		return
	}

	userResponse := model.UserResponse{
		Message: "Success",
		Data:    struct{ Token string }{Token: token},
	}
	res.Header().Set(utility.ContentType, utility.ApplicationJSON)

	res.WriteHeader(http.StatusInternalServerError)

	resp, _ := json.Marshal(userResponse)
	res.Write(resp)
}
