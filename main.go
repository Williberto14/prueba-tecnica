package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func main() {

 db, _ = gorm.Open("sqlite3", "./database/products.db")

 if err != nil {
    fmt.Println(err)
 }

 defer db.Close()

 db.AutoMigrate(&Product{})

 r := gin.Default()
 r.GET("/products/", GetProducts)
 r.GET("/product/:id", GetProduct)
 r.POST("/products", CreateProduct)
 r.PUT("/products/:id", UpdateProduct)
 r.DELETE("/products/:id", DeleteProduct)
 r.Run(":9098")
}

func DeleteProduct(c *gin.Context) {
 id := c.Params.ByName("id")
 var product Product
 d := db.Where("id = ?", id).Delete(&product)
 fmt.Println(d)
 c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdateProduct(c *gin.Context) {
var product Product
 id := c.Params.ByName("id")
 if err := db.Where("id = ?", id).First(&product).Error; err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
 }
 c.BindJSON(&product)
 db.Save(&product)
 c.JSON(200, product)
}

func CreateProduct(c *gin.Context) {
var product Product
 c.BindJSON(&product)
 db.Create(&product)
 c.JSON(200, product)
}

func GetProduct(c *gin.Context) {
 id := c.Params.ByName("id")
 var product Product
 if err := db.Where("id = ?", id).First(&product).Error; err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
 } else {
    c.JSON(200, product)
 }
}

func GetProducts(c *gin.Context) {
 var products []Product
 if err := db.Find(&products).Error; err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
 } else {
    c.JSON(200, products)
 }
}