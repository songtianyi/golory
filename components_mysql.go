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
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB contains information for current db connection
type MySQLClient struct {
	*gorm.DB
	ConnectionErr error
}

//  CommonCfg  params
type MySQLCfg struct {
	Type        string                 // database type defult: mysql
	Username    string                 // database username
	Password    string                 // database password
	Name        string                 // database name
	Addr        string                 // database addr
	TablePrefix string                 // database table prefix
	Dsn         map[string]interface{} //  mysql suffix dsn params

	SingularTable bool
}

// Boot  Connection db return mysql.DB
func MySQLBoot(cfg MySQLCfg) *MySQLClient {
	if cfg.Type == "" {
		cfg.Type = "mysql"
	}
	if cfg.TablePrefix == "" {
		cfg.TablePrefix = "golory_"
	}
	var buf bytes.Buffer
	buf.WriteString(cfg.Username)
	buf.WriteString(":")
	buf.WriteString(cfg.Password)
	buf.WriteString("@tcp(")
	buf.WriteString(cfg.Addr)
	buf.WriteString(")/")
	buf.WriteString(cfg.Name)
	buf.WriteString("?")
	if len(cfg.Dsn) <= 0 {
		buf.WriteString("charset=utf8&parseTime=True&loc=Local")
	} else {
		for k, v := range cfg.Dsn {
			buf.WriteString(k)
			buf.WriteString("=")
			buf.WriteString(fmt.Sprintf("%s", v))
			buf.WriteString("&")
		}
	}
	db, err := gorm.Open(cfg.Type, buf.String())
	if err != nil {
		return &MySQLClient{nil, err}
	}
	// TODO table prefix config ?
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return cfg.TablePrefix + defaultTableName
	}
	// 全局禁用表名复数
	db.SingularTable(cfg.SingularTable)

	return &MySQLClient{db, nil}
}

// close db connection
func (db *MySQLClient) Close() {
	db.Close()
}
