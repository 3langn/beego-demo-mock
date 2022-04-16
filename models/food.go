package models

import "errors"

type Food struct {
	Base
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Price       float64       `json:"price"`
	Image       string        `json:"image"`
	Categories  []*Category   `json:"categories,omitempty" gorm:"many2many:food_categories"`
	BasketItems []*BasketItem `json:"basket_items,omitempty"`
}

type FoodCategory struct {
	FoodID     uint
	CategoryID uint
}

type FoodDto struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Price        float64  `json:"price"`
	Image        string   `json:"image"`
	CategoriesId []string `json:"categories_id"`
}

func (f *Food) FindOne(id string) error {
	return db.Where("id = ?", id).First(f).Error
}

func (f *Food) FindAll(category string) ([]Food, error) {
	var foods []Food
	var err error
	if category == "" {
		err = db.Find(&foods).Error
	} else {
		err = db.Where("categories.name = ?", category).
			Joins("JOIN food_categories ON food_categories.food_id = foods.id").
			Joins("JOIN categories ON categories.id = food_categories.category_id").
			Find(&foods).Error
	}
	return foods, err
}

func (f *Food) Create(dto *FoodDto) (Food, error) {
	var categories []*Category
	db.Model(Category{}).Where("id IN (?)", dto.CategoriesId).Find(&categories)
	if len(categories) == 0 {
		return Food{}, errors.New("No categories found")
	}
	var food Food
	food.Name = dto.Name
	food.Description = dto.Description
	food.Price = dto.Price
	food.Image = dto.Image
	food.Categories = categories

	err := db.Debug().Omit("Categories.*").Create(&food).Error
	return food, err
}

func (f *Food) Update(id string, dto *Food) error {
	return db.Model(Food{}).Where("id = ?", id).Updates(dto).Error
}

func (f *Food) Delete(id string) error {
	return db.Delete(f, "id = ?", id).Error
}
