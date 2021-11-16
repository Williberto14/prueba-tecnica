package main

import (
	"fmt"

	"github.com/Williberto14/prueba-tecnica/conf"
	"github.com/Williberto14/prueba-tecnica/entity"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db = conf.ConfDatabase()

func main() {

	defer db.Close()

	r := gin.Default()
	r.GET("/products/", GetProducts)
	r.GET("/product/:id", GetProduct)
	r.POST("/products", CreateProduct)
	r.PUT("/products/:id", UpdateProduct)
	r.DELETE("/products/:id", DeleteProduct)
	r.Run(":9098")
}

func CreateProduct(c *gin.Context) {
	var product entity.Product
	c.BindJSON(&product)
	id := product.Product_id
	if found := db.Where("product_id = ?", id).First(&product).Error; found == nil {
		c.JSON(400, "Error! ya existe un producto con el ID: "+id)
	} else {
		db.Create(&product)
		c.JSON(200, product)
	}
}

func GetProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product entity.Product
	if err := db.Where("product_id = ?", id).First(&product).Error; err != nil {
		c.JSON(404, "No existe un producto con el ID: "+id)
		fmt.Println(err)
	} else {
		c.JSON(200, product)
	}
}

func GetProducts(c *gin.Context) {
	var products []entity.Product
	if err := db.Find(&products).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, products)
	}
}

func UpdateProduct(c *gin.Context) {
	var product entity.Product
	id := c.Params.ByName("id")
	if err := db.Where("product_id  = ?", id).First(&product).Error; err != nil {
		c.JSON(404, "Error! No se pudo actualizar el producto con ID: "+id)
		fmt.Println(err)
	} else {
		c.BindJSON(&product)
		db.Save(&product)
		c.JSON(200, product)
	}

}

func DeleteProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product entity.Product
	d := db.Where("product_id  = ?", id)
	if err := db.Where("product_id  = ?", id).First(&product).Error; err != nil {
		c.JSON(404, "No existe un producto con ID: "+id)
	} else {
		d.Delete(&product)
		c.JSON(200, gin.H{"id #" + id: "deleted"})
	}
}
