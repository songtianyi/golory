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
	"github.com/1pb-club/golory/components/log"
	"github.com/1pb-club/golory/components/redis"
)

type handler struct {
	loggers      map[string]*log.Logger
	redisClients map[string]*redis.Client
}

func newHandler() *handler {
	return &handler{
		loggers: make(map[string]*log.Logger),
	}
}
func (s *handler) setLogger(k string, v *log.Logger) {
	s.loggers[k] = v
}

func (s *handler) getLogger(k string) *log.Logger {
	l := s.loggers[k]
	return l
}

func (s *handler) setRedis(k string, v *redis.Client) {
	s.redisClients[k] = v
}

func (s *handler) getRedis(k string) *redis.Client {
	c := s.redisClients[k]
	return c
}
