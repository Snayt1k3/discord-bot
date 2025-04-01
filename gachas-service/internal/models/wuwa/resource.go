package wuwa

import "gorm.io/gorm"

type Material struct {
    gorm.Model
    Name   string `json:"name" gorm:"unique;not null"`
    Type   string `json:"type" gorm:"not null"`
    Source string `json:"source"`
}