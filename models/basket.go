package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Basket struct {
	Base
	BasketItems []BasketItem `json:"basket_items"`
	UserID      uuid.UUID    `json:"-"`
}

//TODO: Check relationship
type BasketItem struct {
	Base
	FoodID   uuid.UUID `json:"-"`
	Food     *Food     `json:"food" gorm:"foreignkey:FoodID"`
	Count    int       `json:"count"`
	BasketID uuid.UUID `json:"-"`
}

type BasketResponse struct {
	Basket Basket                `json:"basket"`
	Items  []*BasketItemResponse `json:"items"`
}

type BasketItemResponse struct {
	Food  Food `json:"food"`
	Count int  `json:"count"`
}

func (b *Basket) GetBasket(user_id string) (basket Basket, err error) {
	//TODO: Group by food name, count

	err = db.Debug().
		Model(&Basket{}).
		Where("user_id = ?", user_id).
		Preload("BasketItems").
		Preload("BasketItems.Food").
		Find(&basket).Error

	return basket, nil
}

func (b *Basket) AddFood(user_id string, food_id string) (err error) {
	var food Food
	err = db.First(&food, "id = ?", food_id).Error
	if err != nil {
		return err
	}

	var basket Basket
	err = GetDB().First(&basket, "user_id = ?", user_id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			var basket_items []BasketItem
			basket_items = append(basket_items, BasketItem{FoodID: food.ID, Count: 1})
			basket = Basket{
				UserID:      uuid.FromStringOrNil(user_id),
				BasketItems: basket_items,
			}
			db.Debug().Omit("UserID.*").Create(&basket)
			return nil
		}
		return err
	}

	var basket_item BasketItem
	// If basket item exists, increase count or create new basket item
	err = db.Debug().
		Where("basket_id = ? AND food_id = ?", basket.ID, food.ID).
		First(&basket_item).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			basket_item = BasketItem{FoodID: food.ID, Count: 1, BasketID: basket.ID}
			db.Debug().Omit("UserID.*").Omit("Food.*").Create(&basket_item)
			return nil
		}
		return err
	}

	// Increase count
	basket_item.Count++
	db.Debug().Omit("UserID.*").Omit("Food.*").Save(&basket_item)

	return nil
}
