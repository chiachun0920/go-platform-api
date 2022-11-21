package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/chiachun0920/platform-api/pkg/controller"
	"github.com/chiachun0920/platform-api/pkg/external/lineapi"
	"github.com/chiachun0920/platform-api/pkg/external/mongodb"
	"github.com/chiachun0920/platform-api/pkg/repository/dbrepo"
	"github.com/gin-gonic/gin"
)

func main() {
	vp, err := readConfig()
	if err != nil {
		log.Fatalf("error/init-config - %s", err)
	}

	conn := fmt.Sprintf(
		"%s://%s:%s",
		vp.GetString("db.protocol"),
		vp.GetString("db.uri"),
		vp.GetString("db.port"),
	)

	db, err := mongodb.NewDB(conn)
	if err != nil {
		log.Fatalf("error/connect-db-fail - %s", err)
	}
	defer func() {
		if err := db.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	msgRepo := dbrepo.NewMessageDBRepo(db, vp.GetString("db.dbName"))
	customerRepo := dbrepo.NewCustomerDBRepo(db, vp.GetString("db.dbName"))

	messaging := lineapi.NewLineAPI(
		vp.GetString("line.secret"),
		vp.GetString("line.token"),
	)
	msgController := controller.NewMessageController(msgRepo, customerRepo, messaging)

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.POST("/webhook/line", msgController.WebhookLine)
	router.POST("/messaging/line", msgController.SendMessage)
	router.GET("/message/customer/:customerId", msgController.ListCustomerMessage)

	router.Run()
}
