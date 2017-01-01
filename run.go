package main

import (
	"github.com/wactbprot/m/ini"
	"github.com/wactbprot/m/cdb"
)

func main() {

	config := ini.ParseConfig(ini.GetConfig("config.json"))

	ds := ini.GetDataClient(config)
	ini.TestDataClient(ds)
	ini.UploadConfig(ds, config)

	db := ini.GetDbClient(config)
	ini.TestDbClient(db)
	cdb.GetDef(db, "mpd-ce3-calib")
}
