//
// Copyright (C) 2015-present  Laszlo Csontos
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>. */
//

package config

import (
	"testing"
)

type testCase struct {
	configString string
	configXml    string
}

var testCaseData = []testCase{
	{
		"name1=value1",
		"<config><parameter><name>name1</name><value>value1</value></parameter></config>",
	},
}

func TestParseConfig(t *testing.T) {
	for _, data := range testCaseData {
		bytes := []byte(data.configXml)

		config := parseConfig(bytes)

		got, want := config.String(), data.configString

		if got != want {
			t.Fatalf("parseConfig(%s) = %s, but wanted %s", bytes, got, want)
		}
	}
}
