package main

import (
    "github.com/gin-gonic/gin"
    gomod "github.com/caseyfu/gomod"
)

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.GET("/v100/say", func(c *gin.Context) {
        msg := gomod.SayHiV100("Fu Kai")
        c.JSON(200, gin.H{
            "message": msg,
        })
    })
    r.GET("/v101/say", func(c *gin.Context) {
        msg := gomod.SayHiV101("Fu Kai")
        c.JSON(200, gin.H{
            "message": msg,
        })
    })
    // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
    r.Run() 
}