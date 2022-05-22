package main

import (
	"fmt"

	"github.com/gushikem01/users-go/pkg/udynamodb"
	"github.com/gushikem01/users-go/pkg/ulog"
	"github.com/gushikem01/users-go/pkg/uusers"
	"github.com/gushikem01/users-go/server/internal/app"
	"github.com/gushikem01/users-go/server/internal/users"
)

func main() {
	// TODO: wire
	l, err := ulog.NewZap()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ddb := udynamodb.NewDynamodb(l)
	uRepo := uusers.NewRepository(ddb)
	us := uusers.NewService(uRepo, l)
	uh := users.NewUsers(l, us)
	appGin := app.NewAppGin(uh)
	appGin.Run()

	// l, err := ulog.NewZap()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// db := udynamodb.NewDynamodb(l)
	// m := map[string]*string{
	// 	"#name": aws.String("Name"),
	// }
	// v := map[string]*dynamodb.AttributeValue{
	// 	":name": {
	// 		S: aws.String("Bob"),
	// 	}}

	// res, err := db.QueryInput("Users", m, v, "#name = :name")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(res)
}
