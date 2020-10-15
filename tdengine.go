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
	"database/sql"
	"fmt"
	_ "github.com/taosdata/driver-go/taosSql" // tdengine sql interface implementation
)

// TDengineCfg wrapped tdengine options and some other golory options
type TDengineCfg struct {
	Username string // Database Username
	Password string // Database Password (requires User)
	Addr     string // Database Network address
	DBName   string // Database name
}

// TDengineClient wrapped go-redis TDengineClient
type TDengineClient struct {
	*sql.DB
}

// init redis connection from redis config
func (cfg *TDengineCfg) init() (*TDengineClient, error) {
	conn, err := sql.Open("taosSql", fmt.Sprintf("%s:%s@/tcp(%s)/%s?interpolateParams=true",
		cfg.Username, cfg.Password, cfg.Addr, cfg.DBName))
	if err != nil {
		return nil, fmt.Errorf("open taosql error: %s", err)
	}
	return &TDengineClient{
		conn,
	}, nil
}

func (td *TDengineClient) close() error {
	return td.Close()
}
