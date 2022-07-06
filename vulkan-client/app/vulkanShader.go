package app

import (
	"fmt"
	vk "github.com/vulkan-go/vulkan"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"runtime"
	"unsafe"
)

func (a *App) createGraphicsPipeline() error {

	_, fileName, _, _ := runtime.Caller(1)
	fileName, err := filepath.Abs(fileName)
	if err != nil {
		return err
	}

	fragCode, err := ioutil.ReadFile(filepath.Join(filepath.Dir(fileName), "../shaders/frag.spv"))
	if err != nil {
		return err
	}

	vertCode, err := ioutil.ReadFile(filepath.Join(filepath.Dir(fileName), "../shaders/vert.spv"))
	if err != nil {
		return err
	}

	fragModule, err := a.createShaderModule(fragCode)
	if err != nil {
		return err
	}
	vertModule, err := a.createShaderModule(vertCode)
	if err != nil {
		return err
	}

	vertStageCreateInfo := vk.PipelineShaderStageCreateInfo{
		SType:  vk.StructureTypePipelineShaderStageCreateInfo,
		Stage:  vk.ShaderStageVertexBit,
		Module: vertModule,
		PName:  "main",
	}
	fragStageCreateInfo := vk.PipelineShaderStageCreateInfo{
		SType:  vk.StructureTypePipelineShaderStageCreateInfo,
		Stage:  vk.ShaderStageFragmentBit,
		Module: fragModule,
		PName:  "main",
	}

	shaderStages := []vk.PipelineShaderStageCreateInfo{vertStageCreateInfo, fragStageCreateInfo}

	if len(shaderStages) == 2 {
	}

	vk.DestroyShaderModule(a.logicalDevice, fragModule, nil)
	vk.DestroyShaderModule(a.logicalDevice, vertModule, nil)
	return nil
}

func (a *App) createShaderModule(code []byte) (vk.ShaderModule, error) {
	createInfo := vk.ShaderModuleCreateInfo{
		SType:    vk.StructureTypeShaderModuleCreateInfo,
		PNext:    nil,
		Flags:    0,
		CodeSize: uint(len(code)),
		PCode:    repackUint32(code),
	}

	var shaderModule vk.ShaderModule
	err := vk.Error(vk.CreateShaderModule(a.logicalDevice, &createInfo, nil, &shaderModule))
	if err != nil {
		return nil, fmt.Errorf("could not create shader module - " + err.Error())
	}

	return shaderModule, nil
}

func repackUint32(data []byte) []uint32 {
	buf := make([]uint32, len(data)/4)

	vk.Memcopy(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&buf)).Data), data)
	return buf
}
