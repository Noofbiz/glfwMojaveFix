// +build !darwin arm arm64

package glfwMojaveFix

import "github.com/go-gl/glfw/v3.2/glfw"

func UpdateNSGLContext(w glfw.Window) {}
