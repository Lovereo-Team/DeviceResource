package routers

import (
	"DeviceResource/core"
	"DeviceResource/generator/routers/gen"
)

var InitRouters = []*core.GroupBase{
	// gen
	gen.GenGroup,
}
