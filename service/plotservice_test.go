package service

import (
	"testing"
)

func TestPlotBar(t *testing.T) {
	_, body, _ := GetCertainRepo("X-lab2017/open-digger", "openrank")
	PlotBar("openrank", body)
}