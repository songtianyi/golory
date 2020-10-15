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
	"errors"
	"strings"
)

// Exported errors
var (
	// Error occurred when parse configuration
	ErrParseCfg = errors.New("parse cfg failed")
)

// Join error strings
func wrap(e ...error) error {
	errs := make([]string, 0)
	for _, v := range e {
		if v != nil {
			errs = append(errs, v.Error())
		}
	}
	if len(errs) < 1 {
		return nil
	}
	return errors.New(strings.Join(errs, ": "))
}
