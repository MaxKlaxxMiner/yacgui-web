package main

import (
	"d3d-client/app"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gonutz/d3d9"
	"github.com/gonutz/w32/v2"
	"runtime"
	"syscall"
	"time"
)

var mouseX, mouseY int

func main() {
	runtime.LockOSThread()

	const className = "fullscreen_window_class"
	classNamePtr, _ := syscall.UTF16PtrFromString(className)
	w32.RegisterClassEx(&w32.WNDCLASSEX{
		Cursor: w32.LoadCursor(0, w32.MakeIntResource(w32.IDC_ARROW)),
		WndProc: syscall.NewCallback(func(window w32.HWND, msg uint32, w, l uintptr) uintptr {
			switch msg {
			case w32.WM_KEYDOWN:
				if w == w32.VK_ESCAPE {
					w32.SendMessage(window, w32.WM_CLOSE, 0, 0)
				}
				return 0
			case w32.WM_DESTROY:
				w32.PostQuitMessage(0)
				return 0
			case w32.WM_MOUSEMOVE:
				mouseX = int(l & 0xffff)
				mouseY = int(l >> 16)
				//fmt.Println(mouseX, mouseY)
				return 0
			default:
				return w32.DefWindowProc(window, msg, w, l)
			}
		}),
		ClassName: classNamePtr,
	})

	windowNamePtr, _ := syscall.UTF16PtrFromString("Static Triangle")
	windowHandle := w32.CreateWindow(
		classNamePtr,
		windowNamePtr,
		w32.WS_OVERLAPPEDWINDOW|w32.WS_VISIBLE,
		w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, 1792, 1008,
		0, 0, 0, nil,
	)

	d3d, err := d3d9.Create(d3d9.SDK_VERSION)
	check(err)
	defer d3d.Release()

	device, _, err := d3d.CreateDevice(
		d3d9.ADAPTER_DEFAULT,
		d3d9.DEVTYPE_HAL,
		d3d9.HWND(windowHandle),
		d3d9.CREATE_HARDWARE_VERTEXPROCESSING,
		d3d9.PRESENT_PARAMETERS{
			Windowed:             1,
			SwapEffect:           d3d9.SWAPEFFECT_DISCARD,
			HDeviceWindow:        d3d9.HWND(windowHandle),
			PresentationInterval: d3d9.PRESENT_INTERVAL_IMMEDIATE, // d3d9.PRESENT_INTERVAL_IMMEDIATE,
		},
	)
	check(err)
	defer device.Release()

	check(device.SetRenderState(d3d9.RS_CULLMODE, uint32(d3d9.CULL_NONE)))

	vs, err := device.CreateVertexShaderFromBytes(app.Vso)
	check(err)
	defer vs.Release()
	check(device.SetVertexShader(vs))

	ps, err := device.CreatePixelShaderFromBytes(app.Pso)
	check(err)
	defer ps.Release()
	check(device.SetPixelShader(ps))

	vertices := []float32{
		-0.5, -0.5,
		-0.5, 0.5,
		0.5, -0.5,
		0.5, 0.5,
	}
	vb, err := device.CreateVertexBuffer(uint(len(vertices)*4), d3d9.USAGE_WRITEONLY, 0, d3d9.POOL_DEFAULT, 0)
	check(err)
	defer vb.Release()
	data, err := vb.Lock(0, 0, d3d9.LOCK_DISCARD)
	check(err)
	data.SetFloat32s(0, vertices)
	check(vb.Unlock())
	check(device.SetStreamSource(0, vb, 0, 2*4))

	decl, err := device.CreateVertexDeclaration([]d3d9.VERTEXELEMENT{
		{0, 0, d3d9.DECLTYPE_FLOAT2, d3d9.DECLMETHOD_DEFAULT, d3d9.DECLUSAGE_POSITION, 0},
		d3d9.DeclEnd(),
	})
	check(err)
	defer decl.Release()
	check(device.SetVertexDeclaration(decl))

	rotation := float32(0)
	// create a timer that ticks every 100ms and register a callback for it
	w32.SetTimer(windowHandle, 1, 100, 0)
	var lastTick = time.Now()
	var msg w32.MSG
	for w32.GetMessage(&msg, 0, 0, 0) != 0 {
		w32.TranslateMessage(&msg)

		w32.DispatchMessage(&msg)

		check(device.Clear(nil, d3d9.CLEAR_TARGET, 0, 0, 0))

		ticks := time.Since(lastTick).Milliseconds()
		for i := int64(0); i < ticks; i++ {
			rotation += 0.002
		}
		lastTick = lastTick.Add(time.Duration(ticks) * time.Millisecond)

		//mvp := mgl32.Translate3D(float32(mouseX)*0.01, float32(mouseY)*0.01, 0)
		//mvp := mgl32.HomogRotate3DZ(rotation)
		mvp := mgl32.Ident4()
		mvp = mvp.Mul4(mgl32.Scale3D(0.1, 0.1, 1))
		mvp = mvp.Mul4(mgl32.HomogRotate3DZ(rotation))
		mvp = mvp.Mul4(mgl32.Translate3D((float32(mouseX)-1792.0/2.0)/(1792.0/2.0), -(float32(mouseY)-1008.0/2.0)/(1008.0/2.0), 0).Transpose())
		//mvp[3] += 0.1

		//device.SetViewport(d3d9.VIEWPORT{X: 0, Y: 0, Width: 1492, Height: 1008, MinZ: 0, MaxZ: 256})

		device.SetTransform(d3d9.TS_PROJECTION, d3d9.MATRIX(mgl32.Ident4()))
		check(device.SetVertexShaderConstantF(0, mvp[:]))
		check(device.BeginScene())
		check(device.DrawPrimitive(d3d9.PT_TRIANGLESTRIP, 0, 2))
		check(device.EndScene())

		check(device.Present(nil, nil, 0, nil))
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
