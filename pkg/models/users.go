package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password"`
	Bio       string    `json:"bio"`
	Blogs     []Blog    `json:"blogs" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` //foreign key for blogs table will be UserID
}

func CreateUser(user *User) (*User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	db := db.Create(&user)
	fmt.Println("Rows Added", db.RowsAffected)
	return user, db.Error
}

func GetAllUsers(users *[]User) (*[]User, error) {
	db := db.Preload("Blogs").Find(&users)
	fmt.Println("Rows Affected", db.RowsAffected)
	return users, db.Error
}

func GetUserByID(id int, user *User) (*User, error) {
	db := db.Preload("Blogs").Find(&user, "id = ?", id)
	if db.Error != nil {
		return nil, db.Error
	}
	return user, nil
}

func DeleteUser(id int) *gorm.DB {
	db := db.Delete(&User{}, "id = ?", id)
	fmt.Println("Rows Deleted", db.RowsAffected)
	return db
}

func UpdateUser(id int, fieldsToBeUpdated map[string]interface{}, user *User) (*User, error) {
	user, err := GetUserByID(id, user)
	if err != nil {
		return nil, err
	}
	fmt.Println("fields To be Updated", fieldsToBeUpdated)
	fmt.Println("Original Data", user)
	db := db.Preload("Blogs").Model(&user).Updates(fieldsToBeUpdated) //updates the fields in the database with the new values in the fieldsToBeUpdated map and returns the updated user object in the user pointer variable and error if any error occurs while updating the user object in the database table.
	fmt.Println("Rows Updated", db.RowsAffected)
	fmt.Println("New Data", user)
	return user, db.Error
}

func AuthenticateUser(email, password string) (bool, *User) {
	var user *User
	db := db.Where("email = ?", email).First(&user)
	if db.Error != nil {
		fmt.Println(db.Error)
		return false, user
	}
	if user.ID == 0 {
		return false, user
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println(err.Error())
		return false, user
	}
	return true, user
}
