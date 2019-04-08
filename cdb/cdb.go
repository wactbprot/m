package cdb

import (
	couch "github.com/dustin/go-couch"
	log "github.com/inconshreveable/log15"
	//	redis "gopkg.in/redis.v5"
)
// ... new more simpler mpdef struct here ... 
func GetDef(db couch.Database, id string) (mp MpDef) {
	err := db.Retrieve(id, &mp)
	if err!= nil {
		log.Error(err.Error())
	}
	log.Debug(mp._id)
	return
}
