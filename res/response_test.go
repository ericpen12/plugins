package res

import (
	"errors"
	"github.com/quzhen12/plugins/error"
	"testing"
)

func Test_mgs(t *testing.T) {
	tests := []struct {
		data  interface{}
		error interface{}
	}{
		{nil, errors.New("test")},
		{nil, error.Err{Code: 1002, Msg: "登录失败"}},
		{nil, error.Err{Code: 1001, Msg: "请求失败"}},
		{"hello world", nil},
		{map[string]interface{}{"key": "value"}, nil},
		{nil, nil},
		{100, nil},
	}
	for _, test := range tests {
		m := msg(test.data, test.error)
		t.Logf("msg: %+v", m)
	}
}
