package api

import "testing"

func Test_Ping(t *testing.T) {
	ping := ApiPing{}
	data, err := ping.Get()
	if err != nil {
		t.Error("ping error", err)
	}
	bs := data.([]byte)
	s := string(bs)
	t.Log("result is:", s)
}
