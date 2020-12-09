package service

import (
	"testing"
)

func TestCodeController(t *testing.T) {
	datas := make([]string, 1)
	datas[0] = "Name           string    `json:\"name\" form:\"name\"`"
	AutoGenerationMod("series", "系列管理", datas)
}
