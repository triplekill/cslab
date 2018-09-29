package t1simple

import "testing"

func TestImpl_Opt(t *testing.T) {
	var client Api
	client = &Impl{}
	s := "xxxx"
	expect := ("Opt" + s)
	get := client.Opt(s)
	if get != expect {
		t.Errorf("%v != %v", get, expect)
	}
}
