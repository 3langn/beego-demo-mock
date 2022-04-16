package models

import (
	"errors"
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
	FoodID   uuid.UUID  `json:"-"`
	Food     *Food      `json:"food" gorm:"foreignkey:FoodID"`
	Quantity int        `json:"count"`
	BasketID *uuid.UUID `json:"-"`
	OrderID  *uuid.UUID `json:"-"`
	Status   string     `json:"status"`
}

type BasketResponse struct {
	Basket Basket  `json:"basket,omitempty"`
	Total  float64 `json:"total"`
}

type BasketItemRequest struct {
	FoodID   string `json:"food_id"`
	Quantity int    `json:"quantity"`
}

// Util functions
func CalculateTotal(items []BasketItem) (total float64) {
	for _, item := range items {
		total += item.Food.Price * float64(item.Quantity)
	}
	return
}

func (b *Basket) GetBasket(user_id string) (*BasketResponse, error) {
	//TODO: Group by food name, count
	var basket Basket
	err := db.Debug().
		Model(&Basket{}).
		Where("user_id = ?", user_id).
		Preload("BasketItems").
		Preload("BasketItems.Food").
		First(&basket).Error

	if err != nil {
		return nil, err
	}

	total := CalculateTotal(basket.BasketItems)

	return &BasketResponse{
		Basket: basket,
		Total:  total,
	}, nil
}

func (b *Basket) UpdateBasketItem(user_id string, food_id string, quantity int) (err error) {
	var food Food
	err = db.First(&food, "id = ?", food_id).Error
	if err == gorm.ErrRecordNotFound {
		return errors.New("Food not found")
	}

	var basket Basket
	err = GetDB().First(&basket, "user_id = ?", user_id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			var basket_items []BasketItem
			basket_items = append(basket_items, BasketItem{FoodID: food.ID, Quantity: quantity})
			basket = Basket{
				UserID:      uuid.FromStringOrNil(user_id),
				BasketItems: basket_items,
			}
			err = db.Debug().Omit("UserID.*").Create(&basket).Error
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	var basket_item BasketItem
	err = db.Debug().
		Where("basket_id = ? AND food_id = ?", basket.ID, food.ID).
		First(&basket_item).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			basket_item = BasketItem{FoodID: food.ID, Quantity: quantity, BasketID: &basket.ID}
			db.Debug().Omit("UserID.*").Omit("Food.*").Create(&basket_item)
			return nil
		}
		return err
	}

	basket_item.Quantity = quantity
	db.Debug().Omit("UserID.*").Omit("Food.*").Save(&basket_item)

	return nil
}
