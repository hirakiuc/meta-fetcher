package model

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type FeedItem struct {
	Id          int       `orm:"auto"`
	Title       string    `orm:"size(255)"`
	Link        string    `orm:"size(255)"`
	Desc        string    `orm:"size(65535)"`
	DcDate      time.Time `orm:"column(dc_date);type(datetime)"`
	DcSubject   string    `orm:"column(dc_subject);size(255)"`
	HatebuCount int32     `orm:"column(hatebu_count)"`
	Feed        *Feed     `orm:"rel(fk)"`
	CreatedAt   time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt   time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(FeedItem))
}

func (feedItem *FeedItem) TableName() string {
	return "feed_items"
}
