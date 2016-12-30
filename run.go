package main
import "github.com/wactbprot/m/ini"

func main() {

	config := ini.ParseConfig(ini.GetConfig("config.json"))
	ds := ini.GetDataClient(config)

	ini.TestDataClient(ds)
	ini.UploadConfig(ds, config)
 
	db := ini.GetDbClient(config)
	ini.TestDbClient(db)
}
