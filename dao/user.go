package dao

import (
	"time"
	"github.com/didi/gendry/manager"
	"github.com/didi/gendry/scanner"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

//用户名
const userName string = "root"

//密码
const password = "123456"

//host和端口
const (
	host     string = "127.0.0.1"
	port            = 3306
	database        = "test"
)

var dbs *sql.DB
var err error

//导入包完成后自动初始化
func init() {
	dbs, err = manager.New(database, userName, password, host).Set(
		manager.SetCharset("utf8"),
		manager.SetAllowCleartextPasswords(true),
		manager.SetInterpolateParams(true),
		manager.SetTimeout(1*time.Second),
		manager.SetReadTimeout(1*time.Second),
	).Port(port).Open(true)
	if err != nil {
		log.Fatalf("connect database fail: %v", err)
	}
}

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}
func (this *UserDao) GetUsers() ([]map[string]interface{}, error) {
	rows, _ := dbs.Query("select * from user")
	fmt.Println(rows)
	defer rows.Close()
	//err := scanner.Scan(rows, &users)
	//fmt.Println(err)
	res, err := scanner.ScanMap(rows)
	fmt.Println(res)
	return res, err
}
