package main

import (
	"testing"
	"time"

	"github.com/Williberto14/prueba-tecnica/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/datatypes"
)

func TestPruebaTecnica(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PruebaTecnica Suite")
}

var productoDePrueba = model.Product{
	Product_id:       "s1",
	Name:             "producto1",
	Description:      "producto de prueba",
	Status:           "disponible",
	Creation_date:    time.Now(),
	Update_date:      time.Now(),
	Account_id:       "xlr",
	Format_product:   datatypes.JSON([]byte(`{"formato": "formato de prueba"}`)),
	Value_unit:       150.00,
	Unit_name:        "nombre unitario",
	Unit_description: "descripcion unitaria",
	Stock:            1,
}

var _ = Describe("Pruebas funcionalidad CRUD productos", func() {
	It("Crear producto", func() {

		// gin.SetMode(gin.TestMode)

		// productoJson, _ := json.Marshal(productoDePrueba)

		// w := httptest.NewRecorder()
		// c, _ := gin.CreateTestContext(w)

		// // c.Params = []gin.Param{{Key: "product_id", Value: "p25"}}
		// // c.Params = []gin.Param{{Key: "name", Value: "producto1"}}
		// // c.Params = []gin.Param{{Key: "description", Value: "product de prueba"}}
		// // c.Params = []gin.Param{{Key: "status", Value: "disponible"}}
		// // c.Params = []gin.Param{{Key: "creation_date", Value: "2020-08-12T12:20:14.4469114-05:00"}}
		// // c.Params = []gin.Param{{Key: "update_date", Value: "2020-08-12T12:20:14.4469114-05:00"}}
		// // c.Params = []gin.Param{{Key: "account_id", Value: "xlr"}}
		// // c.Params = []gin.Param{{Key: "format_product", Value: `{'formato' : 'formato de prueba'}`}}
		// // c.Params = []gin.Param{{Key: "value_unit", Value: "150"}}
		// // c.Params = []gin.Param{{Key: "unit_name", Value: "nombre unitario"}}
		// // c.Params = []gin.Param{{Key: "unit_description", Value: "descripcion unitaria"}}
		// // c.Params = []gin.Param{{Key: "stock", Value: "5"}}

		// CreateProduct(c)

		// Expect(w.Code).To(Equal(200))
	})
})
