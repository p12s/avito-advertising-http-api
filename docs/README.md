# Swagger

Документация уже настроена и должна открываться по адресу:  
http://localhost:80/swagger/index.html

## Памятка для настраивания документации в проекте

### 1. Документируем методы

([официальный репозиторий](https://github.com/swaggo/swag))
```go
// @title Todo App API
// @version 1.0
// @description API Server for TodoList Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	...
}
```
```go
// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body todo.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	...
}
```

### 2. Генерируем папку с документацией  

Скачиваем пакет:  
```
go get -u github.com/swaggo/swag/cmd/swag
```
запускаем инициализацию:
```
swag init -g cmd/main.go
/GO-PATH/bin/bin/swag init -g ./cmd/main.go // иногда приходится указвать полный путь
```
в корне должна быдет появиться директория 'docs'   

### 3. Открываем документацию в браузере

В главном хендлере подключаем модули и добавляем адрес для открытия документации в браузере:
```go
package handler

import (
	...
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/p12s/avito-advertising-http-api/docs"
)
...

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	...
}
```
Документация должна будет открыться по адресу (порт подставить из конфига, в моем случае - 80):  
http://localhost:80/swagger/index.html
