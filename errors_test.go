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
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestErros_warp(t *testing.T) {
	e1 := errors.New("err1")
	e2 := errors.New("err2")
	Convey("errors wrap:", t, func() {
		So(
			wrap(nil),
			ShouldBeNil,
		)
		So(
			wrap(e1),
			ShouldNotBeNil,
		)
		So(
			wrap(e1, nil),
			ShouldNotBeNil,
		)
		So(
			wrap(e1, nil, nil),
			ShouldNotBeNil,
		)
		So(
			wrap(e1, e2),
			ShouldNotBeNil,
		)
		So(
			wrap(e1, nil, e2),
			ShouldNotBeNil,
		)
		So(
			wrap(nil, nil, e2),
			ShouldNotBeNil,
		)
	})
}
