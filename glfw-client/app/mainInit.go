package app

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
	"unsafe"
)

var vertexShaderSource = `
#version 410 core
layout (location = 0) in vec3 vp;

uniform mat4 proj;
uniform vec2 ppos;

void main()
{
   gl_Position = proj * vec4(vp.x + ppos.x, vp.y + ppos.y, vp.z, 1.0);
}
`

var fragmentShaderSource = `
#version 410 core
out vec4 color;
void main()
{
    color = vec4(1.0f, 0.5f, 0.2f, 1.0f);
}
`

type getGlParam func(uint32, uint32, *int32)
type getInfoLog func(uint32, int32, *int32, *uint8)

func checkGlError(glObject uint32, errorParam uint32, getParamFn getGlParam,
	getInfoLogFn getInfoLog, failMsg string) {

	var success int32
	getParamFn(glObject, errorParam, &success)
	if success != 1 {
		var infoLog [512]byte
		getInfoLogFn(glObject, 512, nil, (*uint8)(unsafe.Pointer(&infoLog)))
		log.Fatalln(failMsg, "\n", string(infoLog[:512]))
	}
}

func checkShaderCompileErrors(shader uint32) {
	checkGlError(shader, gl.COMPILE_STATUS, gl.GetShaderiv, gl.GetShaderInfoLog,
		"ERROR::SHADER::COMPILE_FAILURE")
}

func checkProgramLinkErrors(program uint32) {
	checkGlError(program, gl.LINK_STATUS, gl.GetProgramiv, gl.GetProgramInfoLog,
		"ERROR::PROGRAM::LINKING_FAILURE")
}

func compileShaders() []uint32 {
	// create the vertex shader
	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	shaderSourceChars, freeVertexShaderFunc := gl.Strs(vertexShaderSource)
	gl.ShaderSource(vertexShader, 1, shaderSourceChars, nil)
	gl.CompileShader(vertexShader)
	checkShaderCompileErrors(vertexShader)

	// create the fragment shader
	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	shaderSourceChars, freeFragmentShaderFunc := gl.Strs(fragmentShaderSource)
	gl.ShaderSource(fragmentShader, 1, shaderSourceChars, nil)
	gl.CompileShader(fragmentShader)
	checkShaderCompileErrors(fragmentShader)

	defer freeFragmentShaderFunc()
	defer freeVertexShaderFunc()

	return []uint32{vertexShader, fragmentShader}
}

/*
 * Link the provided shaders in the order they were given and return the linked program.
 */
func linkShaders(shaders []uint32) uint32 {
	program := gl.CreateProgram()
	for _, shader := range shaders {
		gl.AttachShader(program, shader)
	}
	gl.LinkProgram(program)
	checkProgramLinkErrors(program)

	// shader objects are not needed after they are linked into a program object
	for _, shader := range shaders {
		gl.DeleteShader(shader)
	}

	return program
}

/*
 * Creates the Vertex Array Object for a triangle.
 */
func createTriangleVAO() uint32 {
	vertices := []float32{
		0, 0, 0,
		100, 0, 0,
		100, 100, 0,
		0, 100, 0,
	}

	var VAO uint32
	gl.GenVertexArrays(1, &VAO)

	var VBO uint32
	gl.GenBuffers(1, &VBO)

	// Bind the Vertex Array Object first, then bind and set vertex buffer(s) and attribute pointers()
	gl.BindVertexArray(VAO)

	// copy vertices data into VBO (it needs to be bound first)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// specify the format of our vertex input
	// (shader) input 0
	// vertex has size 3
	// vertex items are of type FLOAT
	// do not normalize (already done)
	// stride of 3 * sizeof(float) (separation of vertices)
	// offset of where the position data starts (0 for the beginning)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)

	// unbind the VAO (safe practice so we don't accidentally (mis)configure it later)
	gl.BindVertexArray(0)

	return VAO
}

var shaders []uint32
var shaderProgram uint32
var vao uint32

func MainInit(window *glfw.Window) {
	window.SetKeyCallback(KeyboardCallBack)

	glfw.SwapInterval(0)

	shaders = compileShaders()
	shaderProgram = linkShaders(shaders)

	// VAO contains all the information about the data to be rendered
	vao = createTriangleVAO()
}
