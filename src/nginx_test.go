package main

import (
	"testing"
)

var vhostsRootPath string

func init() {
	vhostsRootPath = "D:\\nginx\\vhosts"
}

func TestNginxCreate(t *testing.T) {

	ng := Nginx{}
	result := ng.Create("domain.com", "192.168.2.5", "80", false, true, vhostsRootPath)

	if !result.Success {
		t.Error(result.Message)
	} else {
		t.Log(result.Message)
	}
}

func TestNginxList(t *testing.T) {

	ng := Nginx{}
	result := ng.List(vhostsRootPath)

	if len(result.Vhosts) == 0 {
		t.Error("List is Empty!")
		return
	}

	for _, value := range result.Vhosts {
		t.Log(value)
	}
}
