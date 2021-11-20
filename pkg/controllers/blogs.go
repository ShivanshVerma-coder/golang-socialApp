package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ShivanshVerma-coder/golang-socialApp/pkg/models"
	"github.com/gorilla/mux"
)

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	databytes, err := ioutil.ReadAll(r.Body) // read the body of the request from the client to the server and store it in the variable databytes
	if err != nil {
		panic(err)
	}
	var blog *models.Blog
	json.Unmarshal(databytes, &blog)
	blog, err = models.CreateBlog(blog)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(blog)
}

func GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	var blogs *[]models.Blog
	blogs, err := models.GetAllBlogs(blogs)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}

func GetBlog(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	var blog *models.Blog
	blog, err = models.GetBlogByID(id, blog)
	if err != nil {
		panic(err)
	}
	if blog.ID != 0 {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(blog)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Blog not found"))
	}
}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	db := models.DeleteBlogByID(id)
	if db.RowsAffected == 1 {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Blog deleted"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Blog not found"))
	}
}

func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	databytes, err := ioutil.ReadAll(r.Body) // read the body of the request from the client to the server and store it in the variable databytes
	if err != nil {
		panic(err)
	}
	var blog *models.Blog
	var dataToBeUpdated map[string]interface{}
	json.Unmarshal(databytes, &dataToBeUpdated)
	blog, err = models.UpdateBlog(id, dataToBeUpdated, blog)
	if err != nil {
		panic(err)
	}
	if blog.ID != 0 {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(blog)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Blog not found"))
	}
}
