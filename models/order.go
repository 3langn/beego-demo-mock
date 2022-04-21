package models

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"mock/utils"
)

type Order struct {
	Base
	UserID      uuid.UUID    `json:"-"`
	BasketItems []BasketItem `json:"order_items"`
}

func (o *Order) TableName() string {
	return "orders"
}

func (o *Order) GetOrder(user_id string) (Order, error) {
	fmt.Println("User ID: ", user_id)
	var order Order
	var result interface{}
	err := db.Model(&Order{}).
		Preload("BasketItems.Food").
		Where("user_id = ?", user_id).
		Take(&order, "user_id = ?", user_id).
		Error
	fmt.Printf("%+v\n", result)
	//TODO: add total price
	return order, err
}

func (o *Order) CreateOrder(basketItemId string) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		basketItem := BasketItem{}
		err := db.Preload("Food").First(&basketItem, "id = ?", basketItemId).Error
		if err != nil {
			return err
		}
		fmt.Printf("%+v\n", basketItem)
		if basketItem.OrderID != nil {
			return errors.New("Basket item already has an order")
		}

		basket := Basket{}
		err = db.First(&basket, "id = ?", basketItem.BasketID).Error
		if err != nil {
			return err
		}
		fmt.Printf("%+v\n", basket)

		basketItem.BasketID = nil
		basketItem.Status = utils.OrderStatusPending

		order := Order{}
		err = db.First(&order, "user_id = ?", basket.UserID).Error
		if err != nil {
			return err
		}

		user := User{}
		err = db.First(&user, "id = ?", basket.UserID).Error
		if err != nil {
			return err
		}

		balance := user.Balance - (basketItem.Food.Price * float64(basketItem.Quantity))
		basketItem.Status = utils.OrderStatusPending
		if balance < 0 {
			return errors.New("Not enough balance")
		}

		db.Model(&user).Update("balance", balance)
		if err != nil {
			if err == gorm.ErrRecordNotFound {

				order.UserID = basket.UserID
				err = db.Debug().Save(&order).Error
				basketItem.OrderID = &order.ID
				err = db.Debug().Save(&basketItem).Error
			}
		} else {
			basketItem.OrderID = &order.ID
			err = db.Debug().Save(&basketItem).Error
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) UpdateOrderStatus(itemId, user_id, status string) error {
	db.Transaction(func(tx *gorm.DB) error {
		basketItem := BasketItem{}
		err := db.Preload("Food").First(&basketItem, "id = ?", itemId).Error
		if err != nil {
			return err
		}

		if status == basketItem.Status {
			return errors.New("Status is already " + status)
		}

		if basketItem.Status == utils.OrderStatusCancelled {
			return errors.New("Order is cancelled")
		}

		if status == utils.OrderStatusCancelled {
			user := User{}
			err = db.First(&user, "id = ?", user_id).Error
			balance := user.Balance + basketItem.Food.Price*float64(basketItem.Quantity)

			err = db.Model(&user).Update("balance", balance).Error
			if err != nil {
				return err
			}
		}

		return db.Model(&basketItem).Update("status", status).Error
	})
	return nil
}

func (o *Order) CancelOrder(itemId string) error {
	basketItem := BasketItem{}
	err := db.First(&basketItem, "id = ?", itemId).Error
	if err != nil {
		return err
	}
	return db.Model(&basketItem).Update("status", utils.OrderStatusCancelled).Error
}

// Order utils functions
