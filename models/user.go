package models 

import ("gorm.io/gorm"
"time"
)

type User struct {
	gorm.Model
	Name  string `json:"name" form:"name"`
	Age  int `json:"age" form:age`
	Power int `json:"power" form:power`
	Tier string `json:"tier" form:tier`
	Abilities string `json:"abilities" form:abilities`
	Image string `json:"image" form:image`
	CreatedAt time.Time
	UpdatedAt time.Time 
}
