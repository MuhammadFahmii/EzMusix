package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogMiddleware struct {
	DBLog *mongo.Client
}

func (log *LogMiddleware) Log(next echo.HandlerFunc) echo.HandlerFunc {
	coll := log.DBLog.Database("ez_musix").Collection("logs")

	return func(c echo.Context) error {
		log := bson.M{
			"time":   time.Now(),
			"method": c.Request().Method,
			"path":   c.Path(),
		}
		response := next(c)
		log["response"] = c.Response().Status
		coll.InsertOne(c.Request().Context(), log)
		return response
	}
}
