package v1

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/c-beel/userman/src/configman"
)

type UsermanServer struct {
	DB           *gorm.DB
	GOAuthConfig configman.GoogleOAuthConfig
}
