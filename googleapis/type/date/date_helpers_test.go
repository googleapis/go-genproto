// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package date

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	for _, test := range []struct {
		name    string
		y, m, d int
	}{
		{"NormalDate", 2012, 4, 21},
		{"LongAgo", 1776, 7, 4},
		{"Future", 2032, 4, 21},
		{"FarFuture", 2062, 4, 21},
	} {
		t.Run(test.name, func(t *testing.T) {
			datePb := &Date{Year: int32(test.y), Month: int32(test.m), Day: int32(test.d)}
			times := map[string]time.Time{
				"local": datePb.ToLocalTime(),
				"utc":   datePb.ToUTCTime(),
			}
			for k, time := range times {
				t.Run(k, func(t *testing.T) {
					assertEqual(t, "year", time.Year(), test.y)
					assertEqual(t, "month", int(time.Month()), test.m)
					assertEqual(t, "day", time.Day(), test.d)
					dtPb := NewFromTime(time)
					t.Run("ToProto", func(t *testing.T) {
						assertEqual(t, "year", dtPb.GetYear(), int32(test.y))
						assertEqual(t, "month", dtPb.GetMonth(), int32(test.m))
						assertEqual(t, "day", dtPb.GetDay(), int32(test.d))
					})
				})
			}
		})
	}
}

func assertEqual(t *testing.T, name string, got, want interface{}) {
	t.Run(name, func(t *testing.T) {
		if got != want {
			t.Errorf("Got %v, expected %v.", got, want)
		}
	})
}
