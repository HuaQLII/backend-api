package modles

// User 用户表模型
type User struct {
	ID int64 // 主键
	//通过在字段后面的标签说明，定义golang字段和表字段的关系
	//例如 `gorm:"column:username"` 标签说明含义是: Mysql表的列名（字段名)为username
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	//创建时间，时间戳
}

// TableName 为User绑定表名
func (u User) TableName() string {
	//绑定MYSQL表名为users
	return "sys_users"
}
