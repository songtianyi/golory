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

type handler struct {
	loggers      map[string]*LoggerClient
	redisClients map[string]*RedisClient
	mySQLDb      map[string]*MySQLClient
}

func newHandler() *handler {
	return &handler{
		loggers:      make(map[string]*LoggerClient),
		redisClients: make(map[string]*RedisClient),
		mySQLDb:      make(map[string]*MySQLClient),
	}
}
func (s *handler) setLogger(k string, v *LoggerClient) {
	s.loggers[k] = v
}

func (s *handler) getLogger(k string) *LoggerClient {
	l := s.loggers[k]
	return l
}

func (s *handler) setRedis(k string, v *RedisClient) {
	s.redisClients[k] = v
}

func (s *handler) getRedis(k string) *RedisClient {
	c := s.redisClients[k]
	return c
}

func (s *handler) setMySQL(k string, v *MySQLClient) {
	s.mySQLDb[k] = v
}

func (s *handler) getMySQL(k string) *MySQLClient {
	c := s.mySQLDb[k]
	return c
}
