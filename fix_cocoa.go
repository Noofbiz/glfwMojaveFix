// +build darwin,!arm,!arm64

package glfwMojaveFix

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#include <Cocoa/Cocoa.h>
void cocoa_update_nsgl_context(void* id) {
    NSOpenGLContext *ctx = id;
    [ctx update];
}
*/
import "C"

import (
	"unsafe"

	"github.com/go-gl/glfw/v3.2/glfw"
)

var numUpdates int

func UpdateNSGLContext(w glfw.Window) {
	if numUpdates < 2 {
		ctx := w.GetNSGLContext()
		C.cocoa_update_nsgl_context(unsafe.Pointer(ctx))
		numUpdates++
	}
}
