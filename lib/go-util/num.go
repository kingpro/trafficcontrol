package util

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import (
	"errors"
	"strconv"
)

const MSPerNS = int64(1000000)

// ToNumeric returns a float for any numeric type, and false if the interface does not hold a numeric type.
// This allows converting unknown numeric types (for example, from JSON) in a single line
// TODO try to parse string stats as numbers?
func ToNumeric(v interface{}) (float64, bool) {
	switch i := v.(type) {
	case uint8:
		return float64(i), true
	case uint16:
		return float64(i), true
	case uint32:
		return float64(i), true
	case uint64:
		return float64(i), true
	case int8:
		return float64(i), true
	case int16:
		return float64(i), true
	case int32:
		return float64(i), true
	case int64:
		return float64(i), true
	case float32:
		return float64(i), true
	case float64:
		return i, true
	case int:
		return float64(i), true
	case uint:
		return float64(i), true
	default:
		return 0.0, false
	}
}

// JSONIntStr unmarshals JSON strings or numbers into an int.
// This is designed to handle backwards-compatibility for old Perl endpoints which accept both. Please do not use this for new endpoints or new APIs, APIs should be well-typed.
type JSONIntStr int64

func (i *JSONIntStr) UnmarshalJSON(d []byte) error {
	if len(d) == 0 {
		return errors.New("empty object")
	}
	if d[0] == '"' {
		d = d[1 : len(d)-1] // strip JSON quotes
	}
	err := error(nil)
	di, err := strconv.ParseInt(string(d), 10, 64)
	if err != nil {
		return errors.New("not an integer")
	}
	*i = JSONIntStr(di)
	return nil
}
