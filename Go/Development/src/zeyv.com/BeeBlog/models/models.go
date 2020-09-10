package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/unknwon/com"
	"os"
	"path"
	"time"
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

// 添加class
func AddCategory (name string) error {
	o := orm.NewOrm()
	cate := &Category{Title: name}
	fmt.Println(cate, "one")
	qs := o.QueryTable("category") // 查询这个数据表
	err := qs.Filter("title",name).One(cate) // 查询数据库里有没有这个name
	if err == nil {
		return err
	}
	fmt.Println(cate, "Insert")
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}
// 获取classList
func GetAllCategory () ([]*Category, error) {
	o:=orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_,err := qs.All(&cates)
	fmt.Println(cates, "cates")
	return cates, err
}
//
//func DelCategory(id string) error {
//	cid,err := strconv.ParseInt(id,10,64) // string转换成int60位十进制
//	if err != nil {
//		return err
//	}
//	o := orm.NewOrm()
//	cate := &Category{Id: cid}
//	_,err=o.Delete(cate)
//	return err
//}


