/*
Copyright 2015 All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	inMemoryCache = cache.New(time.Duration(10)*time.Minute, time.Duration(1)*time.Minute)
)

type cacheFunc func() (interface{}, error)

func getCached(key string, fn cacheFunc) (interface{}, error) {
	// step: check if the key is in the cache
	item, found := inMemoryCache.Get(key)
	if found {
		return item, nil
	}
	// step: call the caching function
	data, err := fn()
	if err != nil {
		return data, err
	}
	// step: inject into the cache
	inMemoryCache.Set(key, data, 0)

	return data, nil
}
