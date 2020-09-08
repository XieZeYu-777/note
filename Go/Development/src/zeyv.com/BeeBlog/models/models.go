package models

import (
	"path"
	"time"
	"os"
	"github.com/unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

// 数据库信息 驱动的名称
const (
	_DB_NAME = "data/beeblog.db" // 数据库路径
	_SQLITE3_DEIVER = "sqlite3" // 驱动
)

// 分类结构
type Category struct {
	Id int64
	Title string
	Created time.Time `orm:"index"`
	Views int64 `orm:"index"`
	TopicTime time.Time `orm:"index"`
	TopicCount int64
	TopicLastUserId int64
}

// 评论
type Topic struct {
	Id int64
	Uid int64
	Title string
	Content string `orm:"size(5000)"`
	Attachment string
	Created time.Time `orm:"index"`
	Updated time.Time `orm:"index"`
	Views int64 `orm:"index"`
	Author string
	ReplyTime time.Time `orm:"index"`
	ReplyCount int64
	ReplyLastUserId int64
}

func RegisterDB()  {
	if !com.IsExist(_DB_NAME) { // todo 是否存在这个文件
		// todo 不存在改数据库就 创建数据库目录
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME) // todo 创建
	}
	// 注册结构体
	orm.RegisterModel(new(Category), new(Topic))
	// 注册驱动
	orm.RegisterDriver(_SQLITE3_DEIVER,orm.DRSqlite)
	// todo 连接数据库 数据库名称 驱动名称 数据库的路径 最大连接数
	orm.RegisterDataBase("default", _SQLITE3_DEIVER,_DB_NAME, 10)
}