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

// Logger return a log.Logger by name
func Logger(name string) *LoggerClient {
	return gly.components.getLogger(name)
}

// Redis return a redis.Client by name
func Redis(name string) *RedisClient {
	return gly.components.getRedis(name)
}

// Gorm return a *GromClient by name
func Gorm(name string) *GormClient {
	return gly.components.getGorm(name)
}

// TDengine return *TDengineClient by name
func TDengine(name string) *TDengineClient {
	return gly.components.getTDengine(name)
}
