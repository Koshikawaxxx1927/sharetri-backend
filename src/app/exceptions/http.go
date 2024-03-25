package exceptions

import (
	"gorm.io/gorm"
)

var NotFound = gorm.ErrRecordNotFound