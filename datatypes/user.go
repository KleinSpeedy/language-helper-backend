package datatypes

// User model
type User struct {
	UserId   uint   `gorm:"column:User_ID;primarykey"`
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Created  int64  // unix timestamp
}
