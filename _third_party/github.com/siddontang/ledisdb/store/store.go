package store

import (
	"fmt"
	"os"
	"path"

	"bosun.org/_third_party/github.com/siddontang/ledisdb/config"
	"bosun.org/_third_party/github.com/siddontang/ledisdb/store/driver"

	_ "bosun.org/_third_party/github.com/siddontang/ledisdb/store/boltdb"
	_ "bosun.org/_third_party/github.com/siddontang/ledisdb/store/goleveldb"
	_ "bosun.org/_third_party/github.com/siddontang/ledisdb/store/leveldb"
	_ "bosun.org/_third_party/github.com/siddontang/ledisdb/store/mdb"
	_ "bosun.org/_third_party/github.com/siddontang/ledisdb/store/rocksdb"
)

func getStorePath(cfg *config.Config) string {
	if len(cfg.DBPath) > 0 {
		return cfg.DBPath
	} else {
		return path.Join(cfg.DataDir, fmt.Sprintf("%s_data", cfg.DBName))
	}
}

func Open(cfg *config.Config) (*DB, error) {
	s, err := driver.GetStore(cfg)
	if err != nil {
		return nil, err
	}

	path := getStorePath(cfg)

	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, err
	}

	idb, err := s.Open(path, cfg)
	if err != nil {
		return nil, err
	}

	db := new(DB)
	db.db = idb
	db.name = s.String()
	db.st = &Stat{}
	db.cfg = cfg

	return db, nil
}

func Repair(cfg *config.Config) error {
	s, err := driver.GetStore(cfg)
	if err != nil {
		return err
	}

	path := getStorePath(cfg)

	return s.Repair(path, cfg)
}

func init() {
}
