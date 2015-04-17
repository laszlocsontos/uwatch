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

package youtube

import (
	"fmt"
	"testing"

	"github.com/lcsontos/uwatch/catalog"
	"github.com/lcsontos/uwatch/config"
)

func TestSearchByID(t *testing.T) {
	service := newService(t)

	// Call with existing video ID

	service.doTestSearchByID("3M3iK_a-azM", false, t)

	// Call with existing video ID

	service.doTestSearchByID("a", true, t)
}

func TestSearchByTitle(t *testing.T) {

}

func (service *Service) doTestSearchByID(videoId string, wantNoSuchVideoError bool, t *testing.T) {
	videoRecord, err := service.SearchByID(videoId)

	nsve, ok := err.(*catalog.NoSuchVideoError)

	if !ok && err != nil {
		t.Fatal(err)
	}

	if (nsve != nil && !wantNoSuchVideoError) || (nsve == nil && wantNoSuchVideoError) {
		fmt.Printf(
			"err=%v, ok=%v, videoRecord=%v, wantNoSuchVideoError=%v\n",
			err, ok, videoRecord, wantNoSuchVideoError)

		t.Fatal(err)
	}
}

func newService(t *testing.T) *Service {
	config.Init("youtube_key.xml")

	developerKey = config.GetValue("YouTubeDevKey")

	fmt.Printf("Developer Key is: %s", developerKey)

	setDeveloperKey(developerKey)

	service, err := New()

	if err != nil {
		t.Fatal(err)
	}

	return service
}
