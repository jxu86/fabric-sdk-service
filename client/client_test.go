package client

import (
	"testing"
)

func TestGetConfigByUrl(t *testing.T) {

	resp, err := getConfigByUrl("http://10.10.7.30:30628/api/cfg/conn")
	if err != nil {
		t.Error(err)
	}
	//defer resp.Close()
	t.Log(string(resp))
}
