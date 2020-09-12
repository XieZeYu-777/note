package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/unknwon/com"
	"os"
	"path"
	"strconv"
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

func AddCategory (name string) error {
	// orm连接建立
	o := orm.NewOrm()
	// 初始化结构体
	cate := &Category{Title:name}
	// 查询表category
	qs := o.QueryTable("category")
	// 查询这个表里是否有name这个
	err := qs.Filter("title", name).One(cate)
	if err ==  nil {
		return err
	}
	_, err = o.Insert(cate) // 插入数据
	if err != nil {
		return err
	}
	return nil
}

// 返回所有分类
func GetCategoryAll() ([] *Category, error)  {
	// cates make下
	cates := make([] *Category, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	//传递指针数据
	_, err := qs.All(&cates)
	return cates, err
}

// 删除分类
func DelCategory(id string) error {
	// string转int64
	cid, err := strconv.ParseInt(id,10,64)
	// 如果err ！= nil 报错
	if err != nil {
		return  err
	}
	// 连接建立
	o := orm.NewOrm()
	// 初始化结构
	cate := &Category{Id: cid}
	// 删除该数据
	_, err = o.Delete(cate)
	return err
}
