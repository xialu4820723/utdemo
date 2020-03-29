package num

import (
	"git.qutoutiao.net/xialu_test/utdemo/testutils"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTransToStr(t *testing.T) {
	Convey("测试transToStr方法", t, func() {
		testutils.Info(t)
		So(transToStr(1234), ShouldEqual, "一二三四")
		So(transToStr(390), ShouldEqual, "三九零")
	})
}
