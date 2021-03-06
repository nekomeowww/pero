package database

import (
	"github.com/xujiajun/nutsdb"

	"github.com/nekomeowww/pero/config"
	"github.com/nekomeowww/pero/logger"
)

var db *nutsdb.DB

// NutsDB 数据库
type NutsDB struct {
	*nutsdb.DB
	Options nutsdb.Options
}

// InitNutsDB 初始化 nutsdb
func InitNutsDB(conf config.Config) *NutsDB {
	opt := nutsdb.DefaultOptions
	opt.Dir = conf.NutsDB.Dir
	db, err := nutsdb.Open(opt)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	ndb := new(NutsDB)
	ndb.Options = opt
	return ndb
}

// Set 设定键值
func (ndb *NutsDB) Set(bucket string, key, val []byte, ttl uint32) error {
	db, err := nutsdb.Open(ndb.Options)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Update(func(tx *nutsdb.Tx) error {
		if err := tx.Put(bucket, key, val, ttl); err != nil {
			return err
		}
		return nil
	})
	return err
}

// Get 获取键值
func (ndb *NutsDB) Get(bucket string, key []byte) ([]byte, error) {
	b := make([]byte, 0)

	db, err := nutsdb.Open(ndb.Options)
	if err != nil {
		return b, err
	}
	defer db.Close()

	err = db.View(func(tx *nutsdb.Tx) error {
		e, err := tx.Get(bucket, key)
		if err != nil {
			return err
		}
		b = e.Value
		return nil
	})
	return b, err
}

// Del 删除键值
func (ndb *NutsDB) Del(bucket string, key []byte) error {
	db, err := nutsdb.Open(ndb.Options)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Update(func(tx *nutsdb.Tx) error {
		if err := tx.Delete(bucket, key); err != nil {
			return err
		}
		return nil
	})
	return err
}
