package app

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

var winWidth int
var winHeight int
var posUniform int32

func MainDraw(window *glfw.Window) {
	gl.ClearColor(0, 0.25, 0.5, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	width, height := window.GetFramebufferSize()
	if width != winWidth || height != winHeight {
		gl.Viewport(0, 0, int32(width), int32(height))
		winWidth = width
		winHeight = height

		gl.UseProgram(shaderProgram) // ensure the right shader program is being used
		projection := mgl32.Ortho2D(0, float32(width), float32(height), 0)
		projectionUniform := gl.GetUniformLocation(shaderProgram, gl.Str("proj\x00"))
		gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

		posUniform = gl.GetUniformLocation(shaderProgram, gl.Str("ppos\x00"))
	}

	//double xpos, ypos;
	//glfwGetCursorPos(window, &xpos, &ypos);
	mx, my := window.GetCursorPos()

	gl.Uniform2f(posUniform, float32(mx), float32(my))

	gl.BindVertexArray(vao)              // bind data
	gl.DrawArrays(gl.TRIANGLE_FAN, 0, 4) // perform draw call
	gl.BindVertexArray(0)                // unbind data (so we don't mistakenly use/modify it)
}
