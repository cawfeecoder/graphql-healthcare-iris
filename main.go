package main

import (
	"github.com/kataras/iris"
	"github.com/nfrush/healthcare-graphql/models"
	"github.com/nfrush/healthcare-graphql/database"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/middleware/logger"
	"log"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	app.Use(recover.New())

	app.Use(logger.New())

	app.Handle("POST", "/providers", func(ctx iris.Context) {
		p := &models.Provider{}

		if err := ctx.ReadJSON(p); err != nil {
			log.Printf("[ERROR] ", err.Error())
			errMsg := models.ErrorMessage{err.Error()}
			ctx.StatusCode(iris.StatusConflict)
			ctx.JSON(errMsg)
			return
		}

		providerSession := database.GetSession()

		if err := providerSession.DB("").C("providers").Insert(p); err != nil {
			log.Printf("[ERROR] ", err.Error())
			errMsg := models.ErrorMessage{err.Error()}
			ctx.StatusCode(iris.StatusConflict)
			ctx.JSON(errMsg)
			return
		}

		defer providerSession.Close()

		ctx.StatusCode(iris.StatusOK)
	})

	app.Handle("GET", "/providers", func(ctx iris.Context) {
		p := []models.Provider{}

		providerSession := database.GetSession()

		if err := providerSession.DB("").C("providers").Find(bson.M{}).All(&p); err != nil {
			log.Printf("[ERROR] ", err.Error())
			errMsg := models.ErrorMessage{err.Error()}
			ctx.StatusCode(iris.StatusConflict)
			ctx.JSON(errMsg)
			return
		}

		defer providerSession.Close()

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(p)
	})

	app.Run(iris.Addr(":9999"))
}