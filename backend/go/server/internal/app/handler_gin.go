package app

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gushikem01/users-go/server/internal/users"
)

// AppGin
type AppGin struct {
	uh *users.UserHandler
}

var ginLambda *ginadapter.GinLambda

// NewAppGin 作成
func NewAppGin(uh *users.UserHandler) *AppGin {
	return &AppGin{uh}
}

// Run 実行
func (app *AppGin) Run() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://host.docker.internal:3000",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			"x-api-key",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
	v1 := r.Group("/v1")
	{
		app.uh.AddRoute(v1)
	}
	ginLambda = ginadapter.New(r)
	// return r.Run()
	lambda.Start(Handler)
}

// Handler ハンドラ
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}
