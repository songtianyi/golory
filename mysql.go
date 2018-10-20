// Copyright 2018 golory Authors @1pb.club. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package golory

import (
	"bytes"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // orm need this
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

// MySQLClient contains information for current db connection
type MySQLClient struct {
	*gorm.DB
	ConnectionErr error
}

// MySQLCfg is sql config struct
type MySQLCfg struct {
	Debug    bool
	Engine   string                 // Database type, default: mysql
	Username string                 // Database Username
	Password string                 // Database Password (requires User)
	Net      string                 // Database Network type
	Addr     string                 // Database Network address (requires Net)
	DBName   string                 // Database name
	Params   map[string]interface{} //  mysql suffix params ，See in detail：https://github.com/go-sql-driver/mysql#parameters

	TablePrefix string // database table prefix

	SingularTable bool          // Set true to disable table name's pluralization globally
	MaxOpenConn   int           // Maximum connection of database
	MaxIdleConn   int           // Maximum idle connection number of database
	MaxLifetime   time.Duration // The maximum time of a single connection in a database
}

// init db connection from sql config
func (cfg *MySQLCfg) init() *MySQLClient {
	if cfg.Engine == "" {
		cfg.Engine = "mysql"
	}

	var buf bytes.Buffer
	buf.WriteString(cfg.Username)
	buf.WriteString(":")
	buf.WriteString(cfg.Password)
	buf.WriteString("@tcp(")
	buf.WriteString(cfg.Addr)
	buf.WriteString(")/")
	buf.WriteString(cfg.DBName)
	buf.WriteString("?")
	if len(cfg.Params) <= 0 {
		buf.WriteString("charset=utf8&parseTime=True&loc=Local")
	} else {
		a := make([]string, 0)
		for k, v := range cfg.Params {
			a = append(a, fmt.Sprintf("%s=%s", k, v))
		}
		buf.WriteString(strings.Join(a, "&"))
	}
	db, err := gorm.Open(cfg.Engine, buf.String())
	if err != nil {
		return &MySQLClient{nil, err}
	}
	db.LogMode(cfg.Debug)
	// TODO table prefix config ?
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return cfg.TablePrefix + defaultTableName
	}
	db.SingularTable(cfg.SingularTable)

	if cfg.MaxLifetime != 0 {
		db.DB().SetConnMaxLifetime(cfg.MaxLifetime)
	}
	if cfg.MaxOpenConn != 0 {
		db.DB().SetMaxOpenConns(cfg.MaxOpenConn)
	}
	if cfg.MaxIdleConn != 0 {
		db.DB().SetMaxIdleConns(cfg.MaxIdleConn)
	}

	return &MySQLClient{db, nil}
}

// Close sql client
func (c *MySQLClient) Close() {
	c.Close()
}
