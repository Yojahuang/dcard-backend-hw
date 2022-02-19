package model

import (
    "gorm.io/gorm"
)

type Url struct {
    gorm.Model
    OriginUrl string
    ExpireAt string
}
