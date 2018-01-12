package main

import (
	"github.com/kataras/iris"
	"github.com/nfrush/healthcare-graphql/models"
	"github.com/nfrush/healthcare-graphql/database"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/middleware/logger"
	"log"
	r "gopkg.in/gorethink/gorethink.v4"
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

		if err := r.Table("providers").Insert(p).Exec(providerSession); err != nil {
			log.Printf("[ERROR] ", err.Error())
			errMsg := models.ErrorMessage{err.Error()}
			ctx.StatusCode(iris.StatusConflict)
			ctx.JSON(errMsg)
			return
		}

		ctx.StatusCode(iris.StatusOK)
	})

	app.Handle("POST", "/users", func(ctx iris.Context) {
		u := &models.User{}

		if err := ctx.ReadJSON(u); err != nil {
			log.Printf("[ERROR] ", err.Error())
			errMsg := models.ErrorMessage{err.Error()}
			ctx.StatusCode(iris.StatusConflict)
			ctx.JSON(errMsg)
			return
		}

		userSession := database.GetSession()

		if err := r.Table("users").Insert(u).Exec(userSession); err != nil {
			log.Printf("[ERROR] ", err.Error())
			errMsg := models.ErrorMessage{err.Error()}
			ctx.StatusCode(iris.StatusConflict)
			ctx.JSON(errMsg)
			return
		}

		ctx.StatusCode(iris.StatusOK)
	})

	app.Handle("GET", "/providers", func(ctx iris.Context) {
		p := []models.Provider{}

		providerSession := database.GetSession()

		res, err := r.Table("providers").Run(providerSession);
		if err != nil {
			log.Printf("[ERROR] ", err.Error())
			errMsg := models.ErrorMessage{err.Error()}
			ctx.StatusCode(iris.StatusConflict)
			ctx.JSON(errMsg)
			return
		}

		res.All(&p);
		res.Close()

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(p)
	})

	app.Handle("GET", "/users", func(ctx iris.Context) {
		u := []models.User{}

		userSession := database.GetSession()

		res, err := r.Table("users").Run(userSession);
		if err != nil {
			log.Printf("[ERROR] ", err.Error())
			errMsg := models.ErrorMessage{err.Error()}
			ctx.StatusCode(iris.StatusConflict)
			ctx.JSON(errMsg)
			return
		}

		res.All(&u);
		res.Close()

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(u)
	})

	app.Handle("GET", "/providers/{id:string}", func(ctx iris.Context) {
		p := []models.Provider{}

		requestedId := ctx.Params().Get("id")

		providerSession := database.GetSession()

		res, err := r.Table("providers").Get(requestedId).Run(providerSession);
		if err != nil {
			log.Printf("[ERROR] ", err.Error())
			errMsg := models.ErrorMessage{err.Error()}
			ctx.StatusCode(iris.StatusConflict)
			ctx.JSON(errMsg)
			return
		}

		res.All(&p);
		res.Close()

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(p)
	})

	app.Run(iris.Addr(":9999"))
}