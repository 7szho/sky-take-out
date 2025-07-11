package response

import (
	"take-out/internal/model"
	"time"
)

type SetMealPageQueryVo struct {
	Id           uint64    `json:"id"`
	CategoryId   uint64    `json:"categoryId"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	Status       int       `json:"status"`
	Description  string    `json:"description"`
	Image        string    `json:"image"`
	CreateTime   time.Time `json:"crateTime"`
	UpdateTime   time.Time `json:"updateTime"`
	CreateUser   uint64    `json:"createUser"`
	UpdateUser   uint64    `json:"updateUser"`
	CategoryName string    `json:"categoryName"`
}

type SetMealWithDishByIdVo struct {
	Id            uint64              `json:"id"`
	CategoryId    uint64              `json:"categoryId"`
	CategoryName  string              `json:"categoryName"`
	Description   string              `json:"description"`
	Image         string              `json:"image"`
	Name          string              `json:"name"`
	Price         float64             `json:"price"`
	SetmealDishes []model.SetMealDish `json:"setmealDishes"`
	Status        int                 `json:"status"`
	UpdateTime    time.Time           `json:"updateTime"`
}
