package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var handler *gorm.DB

func init() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:123456@tcp(127.0.0.1:3306)/casbin_test?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         191,                                                                                    // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                   // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                   // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                   // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                  // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		// 不使用默认的事务
		SkipDefaultTransaction: false,
		// 不建立实际的外键约束,约束靠代码实现
		DisableForeignKeyConstraintWhenMigrating: true,
		// 表名迁移为单数
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		// 日志
		Logger: logger.Default.LogMode(logger.Info),
	})
	handler = db
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(db, 123)
	}
}
func main() {
	adapter, err := gormadapter.NewAdapterByDB(handler)
	e := casbin.NewEnforcer("./model.conf", adapter)
	if err != nil {
		log.Println(err, 23)
	} else {
		fmt.Println(adapter, e, 34)
	}

	sub := "alice"                                 // 想要访问资源的用户。
	obj := "data1"                                 // 将被访问的资源。
	act := "read"                                  // 用户对资源执行的操作。
	added := e.AddPolicy("alice", "data1", "read") // 给数据库添加数据
	fmt.Println(added)                             // 第一次添加成功就为true，第二次就为false，因为已经存在了这个数据

	ok := e.Enforce(sub, obj, act)
	e.LoadPolicy()
	if ok == true {
		// 允许alice读取data1
		fmt.Println("access")
	} else {
		// 拒绝请求，抛出异常
		fmt.Println("not access")
	}

}
