package post

import (
	"github.com/yamagn/qbox/model"
    "github.com/gin-gonic/gin"

    "github.com/yamagn/qbox/db"
)

type Service struct{}

type Post model.Post

func (s Service) GetAll() ([]Post, error) {
    db := db.GetDB()
    var p []Post

    if err := db.Find(&p).Error; err != nil {
        return nil, err
    }

    return p, nil
}

func (s Service) CreateModel(c *gin.Context) (Post, error) {
    db := db.GetDB()
    var p Post

    if err := c.BindJSON(&p); err != nil {
        return p, err
    }

    if err := db.Create(&p).Error; err != nil {
        return p, err
    }

    return p, nil
}

func (s Service) DeleteByID(id string) error {
    db := db.GetDB()
    var p Post

    if err := db.Where("id = ?", id).Delete(&p).Error; err != nil {
        return err
    }

    return nil
}