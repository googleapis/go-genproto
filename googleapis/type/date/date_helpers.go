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

import "time"

// ToLocalTime returns a new Time based on d, in the system's time zone.
// Hours, minues, seconds, and nanoseconds are set to 0.
func (d *Date) ToLocalTime() time.Time {
	return d.ToTime(time.Local)
}

// ToUTCTime returns a new Time based on d, in UTC.
// Hours, minutes, seconds, and nanoseconds are set to 0.
func (d *Date) ToUTCTime() time.Time {
	return d.ToTime(time.UTC)
}

// ToTime returns a new Time based on d and the provided *time.Location.
// Hours, minutes, seconds, and nanoseconds are set to 0.
func (d *Date) ToTime(l *time.Location) time.Time {
	return time.Date(int(d.GetYear()), time.Month(d.GetMonth()), int(d.GetDay()), 0, 0, 0, 0, l)
}

// NewFromTime returns a new Date based on the provided time.Time.
// The location is ignored, as is anything more precise than the day.
func NewFromTime(t time.Time) *Date {
	return &Date{
		Year:  int32(t.Year()),
		Month: int32(t.Month()),
		Day:   int32(t.Day()),
	}
}
