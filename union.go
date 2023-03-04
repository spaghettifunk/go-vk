// Code generated by go-vk from vk.xml at 2023-03-03 21:43:42.7958204 -0600 CST m=+2.390740001. DO NOT EDIT.
package vk

import "unsafe"

// AccelerationStructureGeometryDataKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryDataKHR.html
type AccelerationStructureGeometryDataKHR struct {
	Triangles   AccelerationStructureGeometryTrianglesDataKHR
	asTriangles bool
	Aabbs       AccelerationStructureGeometryAabbsDataKHR
	asAabbs     bool
	Instances   AccelerationStructureGeometryInstancesDataKHR
	asInstances bool
}

func (u *AccelerationStructureGeometryDataKHR) AsTriangles(val AccelerationStructureGeometryTrianglesDataKHR) {
	u.Triangles = val
	u.asTriangles = true
	u.asAabbs = false
	u.asInstances = false
}

func (u *AccelerationStructureGeometryDataKHR) AsAabbs(val AccelerationStructureGeometryAabbsDataKHR) {
	u.Aabbs = val
	u.asTriangles = false
	u.asAabbs = true
	u.asInstances = false
}

func (u *AccelerationStructureGeometryDataKHR) AsInstances(val AccelerationStructureGeometryInstancesDataKHR) {
	u.Instances = val
	u.asTriangles = false
	u.asAabbs = false
	u.asInstances = true
}

type _vkAccelerationStructureGeometryDataKHR [unsafe.Sizeof(_vkAccelerationStructureGeometryTrianglesDataKHR{})]byte

func (u *AccelerationStructureGeometryDataKHR) Vulkanize() *_vkAccelerationStructureGeometryDataKHR {
	switch true {
	case u.asTriangles:
		return (*_vkAccelerationStructureGeometryDataKHR)(unsafe.Pointer(&u.Triangles))
	case u.asAabbs:
		return (*_vkAccelerationStructureGeometryDataKHR)(unsafe.Pointer(&u.Aabbs))
	case u.asInstances:
		return (*_vkAccelerationStructureGeometryDataKHR)(unsafe.Pointer(&u.Instances))
	default:
		return &_vkAccelerationStructureGeometryDataKHR{}
	}
}

// AccelerationStructureMotionInstanceDataNV: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInstanceDataNV.html
type AccelerationStructureMotionInstanceDataNV struct {
	StaticInstance         AccelerationStructureInstanceKHR
	asStaticInstance       bool
	MatrixMotionInstance   AccelerationStructureMatrixMotionInstanceNV
	asMatrixMotionInstance bool
	SrtMotionInstance      AccelerationStructureSRTMotionInstanceNV
	asSrtMotionInstance    bool
}

func (u *AccelerationStructureMotionInstanceDataNV) AsStaticInstance(val AccelerationStructureInstanceKHR) {
	u.StaticInstance = val
	u.asStaticInstance = true
	u.asMatrixMotionInstance = false
	u.asSrtMotionInstance = false
}

func (u *AccelerationStructureMotionInstanceDataNV) AsMatrixMotionInstance(val AccelerationStructureMatrixMotionInstanceNV) {
	u.MatrixMotionInstance = val
	u.asStaticInstance = false
	u.asMatrixMotionInstance = true
	u.asSrtMotionInstance = false
}

func (u *AccelerationStructureMotionInstanceDataNV) AsSrtMotionInstance(val AccelerationStructureSRTMotionInstanceNV) {
	u.SrtMotionInstance = val
	u.asStaticInstance = false
	u.asMatrixMotionInstance = false
	u.asSrtMotionInstance = true
}

type _vkAccelerationStructureMotionInstanceDataNV [unsafe.Sizeof(_vkAccelerationStructureMatrixMotionInstanceNV{})]byte

func (u *AccelerationStructureMotionInstanceDataNV) Vulkanize() *_vkAccelerationStructureMotionInstanceDataNV {
	switch true {
	case u.asStaticInstance:
		return (*_vkAccelerationStructureMotionInstanceDataNV)(unsafe.Pointer(&u.StaticInstance))
	case u.asMatrixMotionInstance:
		return (*_vkAccelerationStructureMotionInstanceDataNV)(unsafe.Pointer(&u.MatrixMotionInstance))
	case u.asSrtMotionInstance:
		return (*_vkAccelerationStructureMotionInstanceDataNV)(unsafe.Pointer(&u.SrtMotionInstance))
	default:
		return &_vkAccelerationStructureMotionInstanceDataNV{}
	}
}

// ClearColorValue: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkClearColorValue.html
type ClearColorValue struct {
	TypeFloat32   [4]float32
	asTypeFloat32 bool
	TypeInt32     [4]int32
	asTypeInt32   bool
	TypeUint32    [4]uint32
	asTypeUint32  bool
}

func (u *ClearColorValue) AsTypeFloat32(val [4]float32) {
	u.TypeFloat32 = val
	u.asTypeFloat32 = true
	u.asTypeInt32 = false
	u.asTypeUint32 = false
}

func (u *ClearColorValue) AsTypeInt32(val [4]int32) {
	u.TypeInt32 = val
	u.asTypeFloat32 = false
	u.asTypeInt32 = true
	u.asTypeUint32 = false
}

func (u *ClearColorValue) AsTypeUint32(val [4]uint32) {
	u.TypeUint32 = val
	u.asTypeFloat32 = false
	u.asTypeInt32 = false
	u.asTypeUint32 = true
}

type _vkClearColorValue [unsafe.Sizeof([4]float32{})]byte

func (u *ClearColorValue) Vulkanize() *_vkClearColorValue {
	switch true {
	case u.asTypeFloat32:
		return (*_vkClearColorValue)(unsafe.Pointer(&u.TypeFloat32))
	case u.asTypeInt32:
		return (*_vkClearColorValue)(unsafe.Pointer(&u.TypeInt32))
	case u.asTypeUint32:
		return (*_vkClearColorValue)(unsafe.Pointer(&u.TypeUint32))
	default:
		return &_vkClearColorValue{}
	}
}

// ClearValue: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkClearValue.html
type ClearValue struct {
	Color          ClearColorValue
	asColor        bool
	DepthStencil   ClearDepthStencilValue
	asDepthStencil bool
}

func (u *ClearValue) AsColor(val ClearColorValue) {
	u.Color = val
	u.asColor = true
	u.asDepthStencil = false
}

func (u *ClearValue) AsDepthStencil(val ClearDepthStencilValue) {
	u.DepthStencil = val
	u.asColor = false
	u.asDepthStencil = true
}

type _vkClearValue [unsafe.Sizeof(_vkClearColorValue{})]byte

func (u *ClearValue) Vulkanize() *_vkClearValue {
	switch true {
	case u.asColor:
		return (*_vkClearValue)(unsafe.Pointer(&u.Color))
	case u.asDepthStencil:
		return (*_vkClearValue)(unsafe.Pointer(&u.DepthStencil))
	default:
		return &_vkClearValue{}
	}
}

// DescriptorDataEXT: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDescriptorDataEXT.html
type DescriptorDataEXT struct {
	PSampler                *Sampler
	asPSampler              bool
	PCombinedImageSampler   *DescriptorImageInfo
	asPCombinedImageSampler bool
	PInputAttachmentImage   *DescriptorImageInfo
	asPInputAttachmentImage bool
	PSampledImage           *DescriptorImageInfo
	asPSampledImage         bool
	PStorageImage           *DescriptorImageInfo
	asPStorageImage         bool
	PUniformTexelBuffer     *DescriptorAddressInfoEXT
	asPUniformTexelBuffer   bool
	PStorageTexelBuffer     *DescriptorAddressInfoEXT
	asPStorageTexelBuffer   bool
	PUniformBuffer          *DescriptorAddressInfoEXT
	asPUniformBuffer        bool
	PStorageBuffer          *DescriptorAddressInfoEXT
	asPStorageBuffer        bool
	AccelerationStructure   DeviceAddress
	asAccelerationStructure bool
}

func (u *DescriptorDataEXT) AsPSampler(vals []*Sampler) {
	// copy(u.PSampler[:], vals) // TODO this is a known bug in vk-gen
	u.asPSampler = true
	u.asPCombinedImageSampler = false
	u.asPInputAttachmentImage = false
	u.asPSampledImage = false
	u.asPStorageImage = false
	u.asPUniformTexelBuffer = false
	u.asPStorageTexelBuffer = false
	u.asPUniformBuffer = false
	u.asPStorageBuffer = false
	u.asAccelerationStructure = false
}

func (u *DescriptorDataEXT) AsPCombinedImageSampler(vals []*DescriptorImageInfo) {
	// copy(u.PCombinedImageSampler[:], vals)
	u.asPSampler = false
	u.asPCombinedImageSampler = true
	u.asPInputAttachmentImage = false
	u.asPSampledImage = false
	u.asPStorageImage = false
	u.asPUniformTexelBuffer = false
	u.asPStorageTexelBuffer = false
	u.asPUniformBuffer = false
	u.asPStorageBuffer = false
	u.asAccelerationStructure = false
}

func (u *DescriptorDataEXT) AsPInputAttachmentImage(vals []*DescriptorImageInfo) {
	// copy(u.PInputAttachmentImage[:], vals)
	u.asPSampler = false
	u.asPCombinedImageSampler = false
	u.asPInputAttachmentImage = true
	u.asPSampledImage = false
	u.asPStorageImage = false
	u.asPUniformTexelBuffer = false
	u.asPStorageTexelBuffer = false
	u.asPUniformBuffer = false
	u.asPStorageBuffer = false
	u.asAccelerationStructure = false
}

func (u *DescriptorDataEXT) AsPSampledImage(vals []*DescriptorImageInfo) {
	// copy(u.PSampledImage[:], vals)
	u.asPSampler = false
	u.asPCombinedImageSampler = false
	u.asPInputAttachmentImage = false
	u.asPSampledImage = true
	u.asPStorageImage = false
	u.asPUniformTexelBuffer = false
	u.asPStorageTexelBuffer = false
	u.asPUniformBuffer = false
	u.asPStorageBuffer = false
	u.asAccelerationStructure = false
}

func (u *DescriptorDataEXT) AsPStorageImage(vals []*DescriptorImageInfo) {
	// copy(u.PStorageImage[:], vals)
	u.asPSampler = false
	u.asPCombinedImageSampler = false
	u.asPInputAttachmentImage = false
	u.asPSampledImage = false
	u.asPStorageImage = true
	u.asPUniformTexelBuffer = false
	u.asPStorageTexelBuffer = false
	u.asPUniformBuffer = false
	u.asPStorageBuffer = false
	u.asAccelerationStructure = false
}

func (u *DescriptorDataEXT) AsPUniformTexelBuffer(vals []*DescriptorAddressInfoEXT) {
	// copy(u.PUniformTexelBuffer[:], vals)
	u.asPSampler = false
	u.asPCombinedImageSampler = false
	u.asPInputAttachmentImage = false
	u.asPSampledImage = false
	u.asPStorageImage = false
	u.asPUniformTexelBuffer = true
	u.asPStorageTexelBuffer = false
	u.asPUniformBuffer = false
	u.asPStorageBuffer = false
	u.asAccelerationStructure = false
}

func (u *DescriptorDataEXT) AsPStorageTexelBuffer(vals []*DescriptorAddressInfoEXT) {
	// copy(u.PStorageTexelBuffer[:], vals)
	u.asPSampler = false
	u.asPCombinedImageSampler = false
	u.asPInputAttachmentImage = false
	u.asPSampledImage = false
	u.asPStorageImage = false
	u.asPUniformTexelBuffer = false
	u.asPStorageTexelBuffer = true
	u.asPUniformBuffer = false
	u.asPStorageBuffer = false
	u.asAccelerationStructure = false
}

func (u *DescriptorDataEXT) AsPUniformBuffer(vals []*DescriptorAddressInfoEXT) {
	// copy(u.PUniformBuffer[:], vals)
	u.asPSampler = false
	u.asPCombinedImageSampler = false
	u.asPInputAttachmentImage = false
	u.asPSampledImage = false
	u.asPStorageImage = false
	u.asPUniformTexelBuffer = false
	u.asPStorageTexelBuffer = false
	u.asPUniformBuffer = true
	u.asPStorageBuffer = false
	u.asAccelerationStructure = false
}

func (u *DescriptorDataEXT) AsPStorageBuffer(vals []*DescriptorAddressInfoEXT) {
	// copy(u.PStorageBuffer[:], vals)
	u.asPSampler = false
	u.asPCombinedImageSampler = false
	u.asPInputAttachmentImage = false
	u.asPSampledImage = false
	u.asPStorageImage = false
	u.asPUniformTexelBuffer = false
	u.asPStorageTexelBuffer = false
	u.asPUniformBuffer = false
	u.asPStorageBuffer = true
	u.asAccelerationStructure = false
}

func (u *DescriptorDataEXT) AsAccelerationStructure(val DeviceAddress) {
	u.AccelerationStructure = val
	u.asPSampler = false
	u.asPCombinedImageSampler = false
	u.asPInputAttachmentImage = false
	u.asPSampledImage = false
	u.asPStorageImage = false
	u.asPUniformTexelBuffer = false
	u.asPStorageTexelBuffer = false
	u.asPUniformBuffer = false
	u.asPStorageBuffer = false
	u.asAccelerationStructure = true
}

type _vkDescriptorDataEXT [unsafe.Sizeof(unsafe.Pointer(nil))]byte

func (u *DescriptorDataEXT) Vulkanize() *_vkDescriptorDataEXT {
	switch true {
	case u.asPSampler:
		return (*_vkDescriptorDataEXT)(unsafe.Pointer(&u.PSampler))
	case u.asPCombinedImageSampler:
		return (*_vkDescriptorDataEXT)(unsafe.Pointer(&u.PCombinedImageSampler))
	case u.asPInputAttachmentImage:
		return (*_vkDescriptorDataEXT)(unsafe.Pointer(&u.PInputAttachmentImage))
	case u.asPSampledImage:
		return (*_vkDescriptorDataEXT)(unsafe.Pointer(&u.PSampledImage))
	case u.asPStorageImage:
		return (*_vkDescriptorDataEXT)(unsafe.Pointer(&u.PStorageImage))
	case u.asPUniformTexelBuffer:
		return (*_vkDescriptorDataEXT)(unsafe.Pointer(&u.PUniformTexelBuffer))
	case u.asPStorageTexelBuffer:
		return (*_vkDescriptorDataEXT)(unsafe.Pointer(&u.PStorageTexelBuffer))
	case u.asPUniformBuffer:
		return (*_vkDescriptorDataEXT)(unsafe.Pointer(&u.PUniformBuffer))
	case u.asPStorageBuffer:
		return (*_vkDescriptorDataEXT)(unsafe.Pointer(&u.PStorageBuffer))
	case u.asAccelerationStructure:
		return (*_vkDescriptorDataEXT)(unsafe.Pointer(&u.AccelerationStructure))
	default:
		return &_vkDescriptorDataEXT{}
	}
}

// DeviceOrHostAddressConstKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressConstKHR.html
type DeviceOrHostAddressConstKHR struct {
	DeviceAddress   DeviceAddress
	asDeviceAddress bool
	HostAddress     unsafe.Pointer
	asHostAddress   bool
}

func (u *DeviceOrHostAddressConstKHR) AsDeviceAddress(val DeviceAddress) {
	u.DeviceAddress = val
	u.asDeviceAddress = true
	u.asHostAddress = false
}

func (u *DeviceOrHostAddressConstKHR) AsHostAddress(val unsafe.Pointer) {
	u.HostAddress = val
	u.asDeviceAddress = false
	u.asHostAddress = true
}

type _vkDeviceOrHostAddressConstKHR [8]byte

func (u *DeviceOrHostAddressConstKHR) Vulkanize() *_vkDeviceOrHostAddressConstKHR {
	switch true {
	case u.asDeviceAddress:
		return (*_vkDeviceOrHostAddressConstKHR)(unsafe.Pointer(&u.DeviceAddress))
	case u.asHostAddress:
		return (*_vkDeviceOrHostAddressConstKHR)(unsafe.Pointer(&u.HostAddress))
	default:
		return &_vkDeviceOrHostAddressConstKHR{}
	}
}

// DeviceOrHostAddressKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressKHR.html
type DeviceOrHostAddressKHR struct {
	DeviceAddress   DeviceAddress
	asDeviceAddress bool
	HostAddress     unsafe.Pointer
	asHostAddress   bool
}

func (u *DeviceOrHostAddressKHR) AsDeviceAddress(val DeviceAddress) {
	u.DeviceAddress = val
	u.asDeviceAddress = true
	u.asHostAddress = false
}

func (u *DeviceOrHostAddressKHR) AsHostAddress(val unsafe.Pointer) {
	u.HostAddress = val
	u.asDeviceAddress = false
	u.asHostAddress = true
}

type _vkDeviceOrHostAddressKHR [8]byte

func (u *DeviceOrHostAddressKHR) Vulkanize() *_vkDeviceOrHostAddressKHR {
	switch true {
	case u.asDeviceAddress:
		return (*_vkDeviceOrHostAddressKHR)(unsafe.Pointer(&u.DeviceAddress))
	case u.asHostAddress:
		return (*_vkDeviceOrHostAddressKHR)(unsafe.Pointer(&u.HostAddress))
	default:
		return &_vkDeviceOrHostAddressKHR{}
	}
}

// PerformanceCounterResultKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterResultKHR.html
type PerformanceCounterResultKHR struct {
	TypeInt32     int32
	asTypeInt32   bool
	Int64         int64
	asInt64       bool
	TypeUint32    uint32
	asTypeUint32  bool
	Uint64        uint64
	asUint64      bool
	TypeFloat32   float32
	asTypeFloat32 bool
	Float64       float64
	asFloat64     bool
}

func (u *PerformanceCounterResultKHR) AsTypeInt32(val int32) {
	u.TypeInt32 = val
	u.asTypeInt32 = true
	u.asInt64 = false
	u.asTypeUint32 = false
	u.asUint64 = false
	u.asTypeFloat32 = false
	u.asFloat64 = false
}

func (u *PerformanceCounterResultKHR) AsInt64(val int64) {
	u.Int64 = val
	u.asTypeInt32 = false
	u.asInt64 = true
	u.asTypeUint32 = false
	u.asUint64 = false
	u.asTypeFloat32 = false
	u.asFloat64 = false
}

func (u *PerformanceCounterResultKHR) AsTypeUint32(val uint32) {
	u.TypeUint32 = val
	u.asTypeInt32 = false
	u.asInt64 = false
	u.asTypeUint32 = true
	u.asUint64 = false
	u.asTypeFloat32 = false
	u.asFloat64 = false
}

func (u *PerformanceCounterResultKHR) AsUint64(val uint64) {
	u.Uint64 = val
	u.asTypeInt32 = false
	u.asInt64 = false
	u.asTypeUint32 = false
	u.asUint64 = true
	u.asTypeFloat32 = false
	u.asFloat64 = false
}

func (u *PerformanceCounterResultKHR) AsTypeFloat32(val float32) {
	u.TypeFloat32 = val
	u.asTypeInt32 = false
	u.asInt64 = false
	u.asTypeUint32 = false
	u.asUint64 = false
	u.asTypeFloat32 = true
	u.asFloat64 = false
}

func (u *PerformanceCounterResultKHR) AsFloat64(val float64) {
	u.Float64 = val
	u.asTypeInt32 = false
	u.asInt64 = false
	u.asTypeUint32 = false
	u.asUint64 = false
	u.asTypeFloat32 = false
	u.asFloat64 = true
}

type _vkPerformanceCounterResultKHR [8]byte

func (u *PerformanceCounterResultKHR) Vulkanize() *_vkPerformanceCounterResultKHR {
	switch true {
	case u.asTypeInt32:
		return (*_vkPerformanceCounterResultKHR)(unsafe.Pointer(&u.TypeInt32))
	case u.asInt64:
		return (*_vkPerformanceCounterResultKHR)(unsafe.Pointer(&u.Int64))
	case u.asTypeUint32:
		return (*_vkPerformanceCounterResultKHR)(unsafe.Pointer(&u.TypeUint32))
	case u.asUint64:
		return (*_vkPerformanceCounterResultKHR)(unsafe.Pointer(&u.Uint64))
	case u.asTypeFloat32:
		return (*_vkPerformanceCounterResultKHR)(unsafe.Pointer(&u.TypeFloat32))
	case u.asFloat64:
		return (*_vkPerformanceCounterResultKHR)(unsafe.Pointer(&u.Float64))
	default:
		return &_vkPerformanceCounterResultKHR{}
	}
}

// PerformanceValueDataINTEL: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPerformanceValueDataINTEL.html
type PerformanceValueDataINTEL struct {
	Value32       uint32
	asValue32     bool
	Value64       uint64
	asValue64     bool
	ValueFloat    float32
	asValueFloat  bool
	ValueBool     bool
	asValueBool   bool
	ValueString   string
	asValueString bool
}

func (u *PerformanceValueDataINTEL) AsValue32(val uint32) {
	u.Value32 = val
	u.asValue32 = true
	u.asValue64 = false
	u.asValueFloat = false
	u.asValueBool = false
	u.asValueString = false
}

func (u *PerformanceValueDataINTEL) AsValue64(val uint64) {
	u.Value64 = val
	u.asValue32 = false
	u.asValue64 = true
	u.asValueFloat = false
	u.asValueBool = false
	u.asValueString = false
}

func (u *PerformanceValueDataINTEL) AsValueFloat(val float32) {
	u.ValueFloat = val
	u.asValue32 = false
	u.asValue64 = false
	u.asValueFloat = true
	u.asValueBool = false
	u.asValueString = false
}

func (u *PerformanceValueDataINTEL) AsValueBool(val bool) {
	u.ValueBool = val
	u.asValue32 = false
	u.asValue64 = false
	u.asValueFloat = false
	u.asValueBool = true
	u.asValueString = false
}

func (u *PerformanceValueDataINTEL) AsValueString(val string) {
	u.ValueString = val
	u.asValue32 = false
	u.asValue64 = false
	u.asValueFloat = false
	u.asValueBool = false
	u.asValueString = true
}

type _vkPerformanceValueDataINTEL [8]byte

func (u *PerformanceValueDataINTEL) Vulkanize() *_vkPerformanceValueDataINTEL {
	switch true {
	case u.asValue32:
		return (*_vkPerformanceValueDataINTEL)(unsafe.Pointer(&u.Value32))
	case u.asValue64:
		return (*_vkPerformanceValueDataINTEL)(unsafe.Pointer(&u.Value64))
	case u.asValueFloat:
		return (*_vkPerformanceValueDataINTEL)(unsafe.Pointer(&u.ValueFloat))
	case u.asValueBool:
		return (*_vkPerformanceValueDataINTEL)(unsafe.Pointer(&u.ValueBool))
	case u.asValueString:
		return (*_vkPerformanceValueDataINTEL)(unsafe.Pointer(&u.ValueString))
	default:
		return &_vkPerformanceValueDataINTEL{}
	}
}

// PipelineExecutableStatisticValueKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPipelineExecutableStatisticValueKHR.html
type PipelineExecutableStatisticValueKHR struct {
	B32   bool
	asB32 bool
	I64   int64
	asI64 bool
	U64   uint64
	asU64 bool
	F64   float64
	asF64 bool
}

func (u *PipelineExecutableStatisticValueKHR) AsB32(val bool) {
	u.B32 = val
	u.asB32 = true
	u.asI64 = false
	u.asU64 = false
	u.asF64 = false
}

func (u *PipelineExecutableStatisticValueKHR) AsI64(val int64) {
	u.I64 = val
	u.asB32 = false
	u.asI64 = true
	u.asU64 = false
	u.asF64 = false
}

func (u *PipelineExecutableStatisticValueKHR) AsU64(val uint64) {
	u.U64 = val
	u.asB32 = false
	u.asI64 = false
	u.asU64 = true
	u.asF64 = false
}

func (u *PipelineExecutableStatisticValueKHR) AsF64(val float64) {
	u.F64 = val
	u.asB32 = false
	u.asI64 = false
	u.asU64 = false
	u.asF64 = true
}

// WARNING - union PipelineExecutableStatisticValueKHR is returned only, which is not yet handled in the binding
type _vkPipelineExecutableStatisticValueKHR [8]byte

func (u *PipelineExecutableStatisticValueKHR) Vulkanize() *_vkPipelineExecutableStatisticValueKHR {
	switch true {
	case u.asB32:
		return (*_vkPipelineExecutableStatisticValueKHR)(unsafe.Pointer(&u.B32))
	case u.asI64:
		return (*_vkPipelineExecutableStatisticValueKHR)(unsafe.Pointer(&u.I64))
	case u.asU64:
		return (*_vkPipelineExecutableStatisticValueKHR)(unsafe.Pointer(&u.U64))
	case u.asF64:
		return (*_vkPipelineExecutableStatisticValueKHR)(unsafe.Pointer(&u.F64))
	default:
		return &_vkPipelineExecutableStatisticValueKHR{}
	}
}
