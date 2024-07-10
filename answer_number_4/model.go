package main

type User struct {
	Phone      string `json:"phone" gorm:"type:varchar(16);primary_key" validate:"required"`
	Email      string `json:"email" gorm:"type:varchar(255);uniqueIndex" validate:"required"`
	Name       string `json:"name" gorm:"type:varchar(255)" validate:"required"`
	Password   string `json:"password,omitempty" gorm:"type:varchar(255)" validate:"required,min=6"`
	CreatedAt  int64  `json:"-" gorm:"type:bigint"`
	UpdatedAt  int64  `json:"-" gorm:"type:bigint"`
	StatusLike string `json:"status_like"`
	Token      string `json:"token,omitempty" gorm:"-"`
}

type JsonResponse struct {
	Status   int         `json:"status"`
	Messages string      `json:"messages"`
	Data     interface{} `json:"data"`
}
