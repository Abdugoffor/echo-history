#Yuklash
```
go get github.com/Abdugoffor/echo-history/history
```
#Doc
```
package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/username/history"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Category struct {
	ID       int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// migrate
	db.AutoMigrate(&Category{}, &history.History{})

	// hooklarni ulash
	history.RegisterHooks(db)

	e := echo.New()
	e.Use(history.RequestContext)

	// demo
	e.POST("/cat", func(c echo.Context) error {
		cat := Category{Name: "Books", IsActive: true}
		if err := db.Create(&cat).Error; err != nil {
			return err
		}
		return c.JSON(200, cat)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

```


