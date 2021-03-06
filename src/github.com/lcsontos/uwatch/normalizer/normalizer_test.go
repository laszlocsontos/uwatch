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

package normalizer

import (
	"testing"
)

type testCase struct {
	want string
	text string
}

var testCaseData = []testCase{
	{
		"zen---musica-de-relajacion-y-balance-espiritual",
		"ZEN - MÚSICA DE RELAJACIÓN Y BALANCE ESPIRITUAL",
	},
	{
		"meditations-with-sri-chinmoy-vol--1",
		"Meditations with Sri Chinmoy Vol. 1",
	},
	{
		"ouooueaui",
		"öüóőúéáűí",
	},
}

func TestNormalize(t *testing.T) {
	for _, data := range testCaseData {
		text := data.text

		got, want := Normalize(text), data.want

		if got != want {
			t.Fatalf("Normalize(%s) = %s, but wanted %s", text, got, want)
		}
	}
}
