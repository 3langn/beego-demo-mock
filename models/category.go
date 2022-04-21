package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type CategoryDto struct {
	Title string `json:"title"`
}

type Category struct {
	Base
	Name  string  `json:"name"`
	Foods []*Food `json:"foods" gorm:"many2many:food_categories"`
}

//func (c *Category) TableName() string {
//	return "categories"
//}

func (c *Category) Create(dto []*Category) ([]*Category, error) {
	var categories []*Category
	for _, v := range dto {
		categories = append(categories, &Category{
			Name: v.Name,
		})
	}

	if err := GetDB().Save(categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *Category) GetById(id string) error {
	if err := GetDB().First(&c, "id = ?", id).Error; err != nil {
		return err
	}
	GetDB().Debug().Model(&c).Association("Foods").Find(&c.Foods)
	return nil
}

func (c *Category) Update(id string, dto Category) error {
	if err := GetDB().Model(Category{}).Where("id = ?", id).Updates(dto).Error; err != nil {
		return err
	}

	return nil
}

func (c *Category) Delete(id string) error {
	if err := GetDB().Model(Category{}).Where("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (c *Category) GetAll() ([]Category, error) {
	var categories []Category
	if err := GetDB().Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *Category) GetAllCategory(sort string, offset int, limit int) ([]Category, error) {
	var categories []Category

	key := "categories" + ":all" + sort + fmt.Sprint(offset) + fmt.Sprint(limit)

	value, err := RedisClient.Get(key).Result()
	if err != nil {
		fmt.Println("Data Redis not found")
		GetDB().Order(sort).Offset(offset).Limit(limit).Find(&categories)

		//jsonData, _ := json.Marshal(categories)
		if err := RedisClient.Set(key, "123213", 6*24*time.Hour).Err(); err != nil {
			fmt.Println("Redis error: ", err)
		}
	} else {
		fmt.Println("Redis success: ", value)
		err = json.Unmarshal([]byte(value), &categories)
		return categories, err
	}
	fmt.Printf("value: %s\n", value)

	if err := GetDB().Paging(sort, offset, limit).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
