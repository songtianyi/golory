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
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driver
	"strings"
	"time"
)

// GormClient contains information for current db connection
type GormClient struct {
	*gorm.DB
	Err error
}

// GormCfg is sql config struct
type GormCfg struct {
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
func (cfg *GormCfg) init() *GormClient {
	if cfg.Engine == "" {
		// default engine
		cfg.Engine = "mysql"
	}

	pa := make([]string, 0)
	for k, v := range cfg.Params {
		pa = append(pa, fmt.Sprintf("%s=%s", k, v))
	}

	var buf bytes.Buffer
	switch cfg.Engine {
	case "mysql":
		// user:password@tcp(host)/dbname?charset=utf8&parseTime=True&loc=Local
		buf.WriteString(fmt.Sprintf("%s:%s@tcp(%s)/%s?",
			cfg.Username,
			cfg.Password,
			cfg.Addr,
			cfg.DBName))
		if len(pa) == 0 {
			// default params
			buf.WriteString("charset=utf8&parseTime=True&loc=Local")
		} else {
			// user defined params
			buf.WriteString(strings.Join(pa, "&"))
		}
		break
	case "postgres":
		// host=myhost port=myport user=gorm dbname=gorm password=mypassword
		hp := strings.Split(cfg.Addr, ":")
		if len(hp) < 1 {
			return &GormClient{nil, fmt.Errorf("addr %s invalid", cfg.Addr)}
		}
		buf.WriteString(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
			hp[0],
			hp[1],
			cfg.Username,
			cfg.Password,
			cfg.DBName))
		if len(pa) > 0 {
			buf.WriteString(" " + strings.Join(pa, " "))
		}
		break
	default:
		return &GormClient{
			nil,
			fmt.Errorf("engine %s not support", cfg.Engine),
		}
	}
	db, err := gorm.Open(cfg.Engine, buf.String())
	if err != nil {
		return &GormClient{nil, err}
	}
	db.LogMode(cfg.Debug)
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

	return &GormClient{db, nil}
}

// Close sql client
func (c *GormClient) Close() {
	c.Close()
}
