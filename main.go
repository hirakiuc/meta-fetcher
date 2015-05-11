package main

import (
	"./common"
	_ "./fetcher"
	"./model"
	"flag"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/op/go-logging"
	"os"
)

var log *logging.Logger

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
}

func initDatabase(config *common.Config) {
	orm.RegisterDataBase("default", "mysql", config.ConnectString())
}

func invokeFeedFetcher() bool {
	feeds, err := model.RdfFeeds()
	if err != nil {
		log.Error("err: %v\n", err)
		return false
	}

	for _, feed := range feeds {
		log.Info("%s\n", feed.Url)
	}

	return true
}

func main() {
	log = common.GetLogger()

	path := flag.String("c", "./config.toml", "path to config.tml")
	flag.Parse()

	var config, err = common.LoadConfig(*path)
	if err != nil {
		os.Exit(1)
	}

	initDatabase(config)

	// invoke fetcher
	invokeFeedFetcher()
}
