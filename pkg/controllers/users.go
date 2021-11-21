package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ShivanshVerma-coder/golang-socialApp/pkg/models"
	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	databytes, err := ioutil.ReadAll(r.Body) // read the body of the request from the client to the server and store it in the variable databytes
	// fmt.Println("databytes", databytes)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(databytes, &user) // convert json to struct
	fmt.Println(user)
	user, err = models.CreateUser(user) //doesnt changes datatype
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request, userID int) {
	var users *[]models.User
	fmt.Println(userID, "userID")
	users, err := models.GetAllUsers(users)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request, userID int) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	var user *models.User
	user, err = models.GetUserByID(id, user)
	if err != nil {
		panic(err)
	}
	if user.ID != 0 {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request, userID int) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	db := models.DeleteUser(id)
	if db.Error != nil {
		panic(err)
	}
	if db.RowsAffected == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User Deleted"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request, userID int) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	var fieldsToBeUpdated map[string]interface{}
	var user *models.User
	databytes, err := ioutil.ReadAll(r.Body) // read the body of the request from the client to the server and store it in the variable databytes
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(databytes, &fieldsToBeUpdated) // convert json to struct
	if err != nil {
		panic(err)
	}
	user, err = models.UpdateUser(id, fieldsToBeUpdated, user)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func GetProfile(w http.ResponseWriter, r *http.Request, userID int) {
	var user *models.User
	fmt.Println(userID)
	user, err := models.GetUserByID(userID, user)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}
