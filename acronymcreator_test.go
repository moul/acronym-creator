package actor

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreator_CreateAcronyms(t *testing.T) {
	Convey("Testing Creator.CreateAcronyms()", t, func() {
		columns := [][]string{
			{"acronym"},
			{"creator"},
		}
		creator := New(columns)

		acronyms, err := creator.CreateAcronyms()
		So(err, ShouldBeNil)
		So(len(acronyms) > 0, ShouldBeTrue)
		So(len(acronyms["ac"]), ShouldEqual, 1)
		So(len(creator.getCombinations()), ShouldEqual, 1)
		So(acronyms["ac"][0].Combination, ShouldResemble, []string{"acronym", "creator"})
	})
}
