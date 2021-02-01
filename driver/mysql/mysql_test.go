package mysql

import "testing"

func TestNewDriver(t *testing.T) {
	testCases := []struct {
		arg interface{}
	}{
		{
			arg: "",
		},
		{
			arg: Config{
				HostName: "127.0.0.1",
				HostPort: "3306",
				DBName:   "test",
				Charset:  "utf8",
				UserName: "root",
			},
		},
		{
			arg: "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local",
		},
		{
			arg: Config{},
		},
	}
	for _, testCase := range testCases {
		d := New(testCase.arg)
		if d != nil {
			t.Logf("name: %s, dns: %s", d.Name(), d.DNS())
		} else {
			t.Error(testCase, "driver is nil")
		}
	}
}
