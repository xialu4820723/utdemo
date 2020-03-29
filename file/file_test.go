package file

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoadLines(t *testing.T) {
	Convey("测试LoadLines方法", t, func() {
		content := loadLines("testdata/test.txt")
		So(content, ShouldResemble, []string{"123", "一二三"})
	})
}
