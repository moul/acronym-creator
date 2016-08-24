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

		acronyms := creator.CreateAcronyms()
		So(len(acronyms) > 0, ShouldBeTrue)
	})
}
