package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product // hasMany
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category     // belongsTo
	SerialNumber SerialNumber // hasOne
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// create category
	// category := Category{Name: "Cozinha"}
	// db.Create(&category)

	// create product
	// product := Product{
	// 	Name:       "Panela",
	// 	Price:      99.00,
	// 	CategoryID: 4,
	// }
	// db.Create(&product)

	// create serial number
	// db.Create(&SerialNumber{
	// 	Number:    "123456",
	// 	ProductID: product.ID,
	// })

	// var products []Product
	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	// }

	// Has Many
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name + ":")
		for _, product := range category.Products {
			fmt.Println("-", product.Name, "- Serial Number:", product.SerialNumber.Number)
		}
	}
}
