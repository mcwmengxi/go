package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
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
		Logger: logger.Default.LogMode(logger.Error),

	})
	handler = db
	if err != nil {
		log.Println(err)
	}
}

func sqlPolicy() {
	adapter, _ := gormadapter.NewAdapterByDB(handler)
	const text = `[request_definition]
	r = sub, obj, act
	
	[policy_definition]
	p = sub, obj, act
	
	[role_definition]
	g = _, _
	
	[policy_effect]
	e = some(where (p.eft == allow))
	
	[matchers]
	m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act`
	m, err := model.NewModelFromString(text)
	// m, err := model.NewModelFromFile("./.model.conf")
	if err != nil {
		zap.L().Error("字符串加载模型失败!", zap.Error(err))
	}
	e, _ := casbin.NewCachedEnforcer(m, adapter)
	e.SetExpireTime(60*60)
	e.LoadPolicy()

	// sub := "alice"                                      // 想要访问资源的用户。
	// obj := "data1"                                      // 将被访问的资源。
	// act := "read"                                       // 用户对资源执行的操作。
	// added, err := e.AddPolicy("alice", "data1", "read") // 给数据库添加数据
	// fmt.Println(added)                                  // 第一次添加成功就为true，第二次就为false，因为已经存在了这个数据
	// check(e, sub, obj, act)
	e.AddPolicy("admin", "/api/users", "GET")
	e.AddRoleForUser("zhangsan", "admin")
	check(e, "zhangsan", "/api/users", "GET")
	e.RemoveGroupingPolicy("zhangsan", "admin")
	e.RemovePolicy("admin", "/api/users", "GET")
	check(e, "zhangsan", "/api/users", "GET")
	e.SavePolicy()
	check(e, "zhangsan", "/api/users", "GET")
}

func check(e *casbin.CachedEnforcer, sub, obj, act string) {
  ok, _ := e.Enforce(sub, obj, act)
  if ok {
    fmt.Printf("%s CAN %s %s\n", sub, act, obj)
  } else {
    fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
  }
}

func filePolicy() {
	enforcer, err := casbin.NewCachedEnforcer("./model.pml", "./policy.csv")
	if err != nil {
		zap.L().Error("new enforcer failed", zap.Error(err))
	}

	check(enforcer, "zhangsan", "/index", "POST")
	check(enforcer, "zhangsan", "/home", "GET")
  check(enforcer, "zhangsan", "/users", "POST")
  check(enforcer, "wangwu", "/users", "POST")
	enforcer.AddPolicy("wangwu", "/users", "POST")
	enforcer.SavePolicy()
	enforcer.RemovePolicy("wangwu", "/users", "POST")
	enforcer.SavePolicy()
  check(enforcer, "wangwu", "/users", "POST")
}
func main() {
	sqlPolicy()
}
