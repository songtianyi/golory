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

package log_test

import (
	"fmt"
	"github.com/1pb-club/golory/components/log"
)

func ExampleBoot() {
	logger := log.Boot(log.CommonCfg{
		Debug: false,
		Level: "INFO",
		Path:  "golory.log",
	})
	logger.Info("test debug log")
	if logger != nil {
		fmt.Println("OK")
	}

	// Output:
	// OK
}
