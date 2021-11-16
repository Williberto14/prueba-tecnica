# API REST de Productos

## prueba-tecnica

============================================

**Descripción:** Este proyecto tiene como objetivo desarrollar una api rest en el lenguaje Golang que exponga endpoints CRUD para gestionar productos y volúmenes.

=============================================

**Tecnologias usadas:**

- Golang/Go
- Gin Framework
- Sqlite Database
- GORM ORM
- Postman Test
- Ginkgo & gomega Test

=============================================

**Instalación:** Despues de clonar el repositorio ejecutar los siguiente comandos en la terminal

`go mod tidy`

`go run main.go`

**Dependencias**

- GORM y Driver SQLite

  `go get -u gorm.io/gorm`

  `go get -u gorm.io/driver/sqlite`

- Gin

  `go get -u github.com/gin-gonic/gin`

- Datatypes GORM

  `go get -u gorm.io/datatypes`

- Ginkgo y Gomega

  `go get github.com/onsi/ginkgo/ginkgo`

  `go get github.com/onsi/gomega/...`

=============================================

**Windows**
En windows 10 puede presentar problemas con el driver de SQlite
se puede resolver instalando el siguiente software y agregandolo al PATH.

https://jmeubank.github.io/tdm-gcc/

Fuente de la soluciÓn:

https://stackoverflow.com/questions/43580131/exec-gcc-executable-file-not-found-in-path-when-trying-go-build

=============================================

### La API se levantara en el puerto :9098

**Los endpoints expuestos al levantar la aplicacion son:**

- **POST** /products/create
- **GET** /products/list
- **GET** /products/find/:id
- **PUT** /products/update/:id
- **DELETE** /products/delete/:id

---

**Para acceder al siguiente endpoint debera autenticarse:**

    Usuario     Contraseña
    "user1":    "root",
    "user2":    "1234",

- **GET** /volumes

=============================================

## Estructura del proyecto

![Estructura del proyecto](https://i.ibb.co/cJS40vV/Estructura-del-proyecto.png)

================================================

## Test Endpoints y validaciones con Postman

### Crear producto 1

![Crear producto 1](https://i.ibb.co/RYSMW29/p1.png)

### Crear producto 2

![Crear producto 2](https://i.ibb.co/yNdStm1/p2.png)

### Crear producto id repetido

![Crear producto id repetido](https://i.ibb.co/L9LZ9n8/3.png)

### Actualizar producto

![Actualizar producto](https://i.ibb.co/pxBrLFm/4.png)

### Actualizar producto id errado

![Actualizar producto id errado](https://i.ibb.co/4PRWtyH/5.png)

### Listar todos los productos

![Listar todos los productos](https://i.ibb.co/Vt05R5j/6.png)

### Buscar producto por id

![Buscar producto por id](https://i.ibb.co/YbjFrGd/7.png)

### Buscar producto por id errado

![Buscar producto por id errado](https://i.ibb.co/0qJpsQd/8.png)

### Borrar producto

![Borrar producto](https://i.ibb.co/gFY0pJT/9.png)

### Borrar producto id errado

![Borrar producto id errado](https://i.ibb.co/3Tgh9Zy/10.png)
