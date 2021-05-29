package gls

import (
	"unsafe"

	"github.com/modern-go/reflect2"
	"github.com/xmseekshine/eutil/gls/g"
)

// offset for go1.4
var goidOffset = getOffSet()

// Init
func getOffSet() uintptr {
	gType := reflect2.TypeByName("runtime.g").(reflect2.StructType)
	if gType == nil {
		panic("failed to get runtime.g type")
	}
	goidField := gType.FieldByName("goid")
	return goidField.Offset()
}

//GetGroutineID ... GoID returns the goroutine id of current goroutine
func GetGroutineID() int64 {
	gPointer := g.G()
	p := (*int64)(unsafe.Pointer(uintptr(gPointer) + goidOffset))

	return *p
}
