package service

import (
	"testing"
)

func TestCodeController(t *testing.T) {
	datas := make([]string, 4)
	datas[0] = "Name string `json:\"name\" form:\"name\"` //名称"
	datas[1] = "BrandId string `json:\"brandId\" form:\"brandId\"` //品牌"
	datas[2] = "SeriesId string `json:\"seriesId\" form:\"seriesId\"` //系列"
	datas[3] = "Concentration float32 `json:\"seriesId\" form:\"seriesId\"` //酒精浓度"
	AutoGenerationMod("alcohol", "酒管理", datas)
}
