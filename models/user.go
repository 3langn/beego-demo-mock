package models

// Don't use adapter/orm => https://github.com/beego/beego/issues/4683
import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Base
	Username string  ` json:"username"`
	Password string  ` json:"-"`
	Token    []Token ` json:"token,omitempty"`
	Role     string  ` json:"role"`
	Basket   Basket  `json:"basket,omitempty"`
	Address  string  `json:"address,omitempty"`
}

func (u *User) TableName() string {
	return "users"
}
func (u *User) Create(username string, password string) error {
	*u = User{
		Username: username,
		Password: password,
	}
	temp := User{}

	if err := GetDB().First(&temp, "username = ?", username).Error; err != gorm.ErrRecordNotFound {
		return fmt.Errorf("username already exists")
	}

	h, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(h)
	if err := GetDB().Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) GetAll() ([]*User, error) {
	var users []*User
	if err := GetDB().Find(&users).Error; err != nil {
		panic(err)
		return nil, err
	}
	return users, nil
}

func (u *User) FindByUsername(username string) error {
	err := GetDB().Where("username = ?", username).First(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) FindById(id int64) error {
	err := GetDB().First(&u, "id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateRole(role string) error {
	if err := GetDB().Model(&u).Update("role", role).Error; err != nil {
		return err
	}
	return nil
}
