package main

import (
    "github.com/gin-gonic/gin"
    "github.com/caseyfu/go/module/say"
)

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.GET("/v1/say", func(c *gin.Context) {
        msg := say.SayHi("Fu Kai")
        c.JSON(200, gin.H{
            "message": msg,
        })
    })
    // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
    r.Run() 
}