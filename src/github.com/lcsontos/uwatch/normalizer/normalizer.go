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
	"unicode/utf8"
)

func Normalize(text string) string {
	if text == "" {
		return text
	}

	text = strings.ToLower(text)

	var normalized = make([]rune, 0)

	for index, width := 0, 0; index < len(text); index += width {
		var newRuneValue, oldRuneValue rune

		oldRuneValue, width = utf8.DecodeRuneInString(text[index:])

		switch {
		case oldRuneValue >= '0' && oldRuneValue <= '9':
			newRuneValue = oldRuneValue
		case oldRuneValue >= 'a' && oldRuneValue <= 'z':
			newRuneValue = oldRuneValue
		case oldRuneValue >= 224 && oldRuneValue <= 229:
			newRuneValue = 'a'
		case oldRuneValue >= 232 && oldRuneValue <= 235:
			newRuneValue = 'e'
		case oldRuneValue >= 236 && oldRuneValue <= 239:
			newRuneValue = 'i'
		case oldRuneValue == 241:
			newRuneValue = 'n'
		case (oldRuneValue >= 242 && oldRuneValue <= 246) || (oldRuneValue == 337):
			newRuneValue = 'o'
		case (oldRuneValue >= 249 && oldRuneValue <= 252) || (oldRuneValue == 369):
			newRuneValue = 'u'
		case oldRuneValue == 253:
			newRuneValue = 'y'
		case oldRuneValue == 255:
			newRuneValue = 'y'
		default:
			newRuneValue = '-'
		}

		normalized = append(normalized, newRuneValue)
	}

	return string(normalized)
}
