package str

import (
	"git.qutoutiao.net/xialu_test/utdemo/testutils"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestContainString(t *testing.T) {
	Convey("测试containString方法", t, func() {
		testutils.Info(t)
		strSlice := []string{"abc", "bcd"}
		So(containString(strSlice, "abc"), ShouldBeTrue)
		So(containString(strSlice, "ab"), ShouldBeFalse)
	})
}
