package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chiachun0920/platform-api/pkg/external/mongodb"
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
}
