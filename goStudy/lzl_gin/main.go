package main

import (
	"studyFiles/goStudy/lzl_gin/gin"
)

func main() {
	r := gin.New()

	r.GET("/print/method", func(ctx gin.Ctx) {
		ctx.Json(200, map[string]string{
			"method": ctx.Req.Method,
		})
	})

	err := r.Run(":12345")
	if err != nil {
		panic(err)
	}
}
