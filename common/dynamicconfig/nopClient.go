// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package dynamicconfig

import (
	"errors"
	"time"

	"github.com/uber/cadence/common/log"
	"github.com/uber/cadence/common/types"
)

// nopClient is a dummy implements of dynamicconfig Client interface, all operations will always return default values.
type nopClient struct{}

func (mc *nopClient) GetValue(name Key, defaultValue interface{}) (interface{}, error) {
	return nil, NotFoundError
}

func (mc *nopClient) GetValueWithFilters(
	name Key, filters map[Filter]interface{}, defaultValue interface{},
) (interface{}, error) {
	return nil, NotFoundError
}

func (mc *nopClient) GetIntValue(name Key, filters map[Filter]interface{}, defaultValue int) (int, error) {
	return defaultValue, NotFoundError
}

func (mc *nopClient) GetFloatValue(name Key, filters map[Filter]interface{}, defaultValue float64) (float64, error) {
	return defaultValue, NotFoundError
}

func (mc *nopClient) GetBoolValue(name Key, filters map[Filter]interface{}, defaultValue bool) (bool, error) {
	if filters[DomainName] == "TestRawHistoryDomain" {
		return true, NotFoundError
	}
	return defaultValue, NotFoundError
}

func (mc *nopClient) GetStringValue(name Key, filters map[Filter]interface{}, defaultValue string) (string, error) {
	return defaultValue, NotFoundError
}

func (mc *nopClient) GetMapValue(
	name Key, filters map[Filter]interface{}, defaultValue map[string]interface{},
) (map[string]interface{}, error) {
	return defaultValue, NotFoundError
}

func (mc *nopClient) GetDurationValue(
	name Key, filters map[Filter]interface{}, defaultValue time.Duration,
) (time.Duration, error) {
	return defaultValue, NotFoundError
}

func (mc *nopClient) UpdateValue(name Key, value interface{}) error {
	return errors.New("unable to update key")
}

func (mc *nopClient) RestoreValue(name Key, filters map[Filter]interface{}) error {
	return errors.New("not supported for file based client")
}

func (mc *nopClient) ListValue(name Key) ([]*types.DynamicConfigEntry, error) {
	return nil, errors.New("not supported for file based client")
}

// NewNopClient creates a nop client
func NewNopClient() Client {
	return &nopClient{}
}

// NewNopCollection creates a new nop collection
func NewNopCollection() *Collection {
	return NewCollection(&nopClient{}, log.NewNoop())
}
