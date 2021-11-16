package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Williberto14/prueba-tecnica/conf"
	"github.com/Williberto14/prueba-tecnica/model"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db = conf.ConfDatabase()

func main() {
	defer db.Close()

	r := gin.Default()

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"user1": "root",
		"user2": "1234",
	}))

	authorized.GET("/volumes", GetVolumes)

	r.POST("/products/create", CreateProduct)
	r.GET("/products/list", GetProducts)
	r.GET("/products/find/:id", GetProduct)
	r.PUT("/products/update/:id", UpdateProduct)
	r.DELETE("/products/delete/:id", DeleteProduct)

	r.Run(":9098")
}

//=====================================================================================

func CreateProduct(c *gin.Context) {
	var product model.Product
	c.BindJSON(&product)
	id := product.Product_id
	if found := db.Where("product_id = ?", id).First(&product).Error; found == nil {
		c.JSON(400, "Error! ya existe un producto con id: "+id)
	} else {
		db.Create(&product)
		c.JSON(200, product)
	}
}

func GetProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product model.Product
	if err := db.Where("product_id = ?", id).First(&product).Error; err != nil {
		c.JSON(404, "No existe un producto con id: "+id)
		fmt.Println(err)
	} else {
		c.JSON(200, product)
	}
}

func GetProducts(c *gin.Context) {
	var products []model.Product
	if err := db.Find(&products).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, products)
	}
}

func UpdateProduct(c *gin.Context) {
	var product model.Product
	id := c.Params.ByName("id")
	if err := db.Where("product_id  = ?", id).First(&product).Error; err != nil {
		c.JSON(404, "Error! No se pudo actualizar el producto con id: "+id)
		fmt.Println(err)
	} else {
		c.BindJSON(&product)
		db.Save(&product)
		c.JSON(200, product)
	}
}

func DeleteProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product model.Product
	d := db.Where("product_id  = ?", id)
	if err := db.Where("product_id  = ?", id).First(&product).Error; err != nil {
		c.JSON(404, "No existe un producto con id: "+id)
	} else {
		d.Delete(&product)
		c.JSON(200, gin.H{"Producto con id #" + id: "eliminado con exito"})
	}
}

func GetVolumes(c *gin.Context) {
	fmt.Println("Estamos vivos")
	jsonFile, _ := ioutil.ReadFile("./volumes/volumen_list.json")
	var data interface{}
	json.Unmarshal(jsonFile, &data)
	c.JSON(http.StatusOK, data)
}
