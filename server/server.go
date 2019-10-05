package server

import (
    "github.com/gin-gonic/gin"

    "github.com/yamagn/night-challenge/controller"
)

func Init() {
    r := router()
    r.Run()
}

func router() *gin.Engine {
    r := gin.Default()

    u := r.Group("/posts")
    {
        ctrl := post.Controller{}
        u.GET("", ctrl.Index)
        u.POST("", ctrl.Create)
        u.DELETE("/:id", ctrl.Delete)
    }

    return r
}