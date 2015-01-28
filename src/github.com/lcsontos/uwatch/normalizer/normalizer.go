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
	"strings"
)

func Normalize(text string) string {
	if text == "" {
		return text
	}

	var normalized = make([]rune, len(text))

	text = strings.ToLower(text)

	for index, runeValue := range text {
		char := runeValue

		switch {
		case runeValue >= '0' && runeValue <= '9':
			normalized[index] = char
		case runeValue >= 'a' && runeValue <= 'z':
			normalized[index] = char
		case runeValue >= 224 && runeValue <= 229:
			normalized[index] = 'a'
		case runeValue >= 232 && runeValue <= 235:
			normalized[index] = 'e'
		case runeValue >= 236 && runeValue <= 239:
			normalized[index] = 'i'
		case runeValue == 241:
			normalized[index] = 'n'
		case (runeValue >= 242 && runeValue <= 246) || (runeValue == 337):
			normalized[index] = 'o'
		case (runeValue >= 249 && runeValue <= 252) || (runeValue == 369):
			normalized[index] = 'u'
		case runeValue == 253:
			normalized[index] = 'y'
		case runeValue == 255:
			normalized[index] = 'y'
		default:
			normalized[index] = '-'
		}

	}

	return string(normalized)
}
