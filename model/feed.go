package model

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Feed struct {
	Id           int       `orm:"auto"`
	Url          string    `orm:"size(255)"`
	Type         string    `orm:"column(feed_type);size(255)"`
	Label        string    `orm:"column(feed_label);size(255)"`
	LastLoadedAt time.Time `orm:"column(last_loaded_at);type(datetime)"`
	Enabled      bool      `orm:""`
	CreatedAt    time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt    time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Feed))
}

func (feed *Feed) TableName() string {
	return "feeds"
}

func RdfFeeds() ([]*Feed, error) {
	o := orm.NewOrm()

	var feeds []*Feed

	_, err := o.QueryTable("feeds").
		Filter("enabled", 1).
		Filter("feed_type", "Hatebu::Rdf").
		OrderBy("url").
		All(&feeds)

	if err != nil {
		return []*Feed{}, err
	} else {
		return feeds, err
	}
}
