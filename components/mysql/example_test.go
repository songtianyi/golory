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

package mysql

import (
	"fmt"
	"testing"
)

type User struct {
	ID       int
	Username string
	Password string
}

func TestBoot(t *testing.T) {
	cfg := CommonCfg{
		UserName:      "root",
		PassWord:      "",
		Addr:          "127.0.0.1:3306",
		Name:          "golory",
		SingularTable: true,
	}
	db := Boot(cfg)
	if db.ConnectionErr != nil {
		t.Errorf("Connection to database failed，error : %v \n", db.ConnectionErr)
		return
	}
	fmt.Println("Successful connection to database.")
}
