package remote_package

import (
	uuid "github.com/satori/go.uuid"
	"testing"
)

// 测试导入其他包
// go get 获取github不能有文件目录 得直接文件展示
func TestName(t *testing.T) {
	u1 := uuid.Must(uuid.NewV4(), nil)
	t.Log(u1)
}
