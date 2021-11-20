package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"` //foreign key for user
	User      User      `json:"user" gorm:"foreignKey:UserID"`
}

func CreateBlog(blog *Blog) (*Blog, error) {
	db := db.Create(&blog)
	fmt.Println("Rows Added", db.RowsAffected)
	db.Preload("User").Find(&blog)
	return blog, db.Error
}

func GetAllBlogs(blogs *[]Blog) (*[]Blog, error) {
	db := db.Preload("User").Find(&blogs)
	fmt.Println("Rows Affected", db.RowsAffected)
	return blogs, db.Error
}

func GetBlogByID(id int, blog *Blog) (*Blog, error) {
	// db := db.Preload("User").Preload("User.Blogs").Find(&blog, "id = ?", id)
	db := db.Preload("User").Find(&blog, "id = ?", id)
	if db.Error != nil {
		return nil, db.Error
	}
	return blog, nil
}

func DeleteBlogByID(id int) *gorm.DB {
	db := db.Delete(&Blog{}, "id = ?", id)
	fmt.Println("Rows Deleted", db.RowsAffected)
	return db
}

func UpdateBlog(id int, fieldsToBeUpdated map[string]interface{}, blog *Blog) (*Blog, error) {
	blog, err := GetBlogByID(id, blog)
	if err != nil {
		return nil, err
	}
	fmt.Println("fields To be Updated", fieldsToBeUpdated)
	fmt.Println("Original Data", blog)
	db := db.Preload("User").Model(blog).Updates(fieldsToBeUpdated) //updates the fields in the database with the new values in the fieldsToBeUpdated map and returns the updated user object in the user pointer variable and error if any error occurs while updating the user object in the database table.
	fmt.Println("Rows Updated", db.RowsAffected)
	fmt.Println("New Data", blog)
	return blog, db.Error
}
