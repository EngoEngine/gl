// Copyright 2014 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build js
// +build !nogl

package gl

import (
	"errors"
	"image"
	"log"

	"github.com/gopherjs/gopherwasm/js"
)

type Texture struct{ js.Value }
type Buffer struct{ js.Value }
type FrameBuffer struct{ js.Value }
type RenderBuffer struct{ js.Value }
type Program struct{ js.Value }
type UniformLocation struct{ js.Value }
type Shader struct{ js.Value }

type ContextAttributes struct {
	// If Alpha is true, the drawing buffer has an alpha channel for
	// the purposes of performing OpenGL destination alpha operations
	// and compositing with the page.
	Alpha bool

	// If Depth is true, the drawing buffer has a depth buffer of at least 16 bits.
	Depth bool

	// If Stencil is true, the drawing buffer has a stencil buffer of at least 8 bits.
	Stencil bool

	// If Antialias is true and the implementation supports antialiasing
	// the drawing buffer will perform antialiasing using its choice of
	// technique (multisample/supersample) and quality.
	Antialias bool

	// If PremultipliedAlpha is true the page compositor will assume the
	// drawing buffer contains colors with premultiplied alpha.
	// This flag is ignored if the alpha flag is false.
	PremultipliedAlpha bool

	// If the value is true the buffers will not be cleared and will preserve
	// their values until cleared or overwritten by the author.
	PreserveDrawingBuffer bool
}

// Returns a copy of the default WebGL context attributes.
func DefaultAttributes() *ContextAttributes {
	return &ContextAttributes{true, true, false, true, true, false}
}

type Context struct {
	js.Value
	ARRAY_BUFFER                                 int
	ARRAY_BUFFER_BINDING                         int
	ATTACHED_SHADERS                             int
	BACK                                         int
	BLEND                                        int
	BLEND_COLOR                                  int
	BLEND_DST_ALPHA                              int
	BLEND_DST_RGB                                int
	BLEND_EQUATION                               int
	BLEND_EQUATION_ALPHA                         int
	BLEND_EQUATION_RGB                           int
	BLEND_SRC_ALPHA                              int
	BLEND_SRC_RGB                                int
	BLUE_BITS                                    int
	BOOL                                         int
	BOOL_VEC2                                    int
	BOOL_VEC3                                    int
	BOOL_VEC4                                    int
	BROWSER_DEFAULT_WEBGL                        int
	BUFFER_SIZE                                  int
	BUFFER_USAGE                                 int
	BYTE                                         int
	CCW                                          int
	CLAMP_TO_EDGE                                int
	CLAMP_TO_BORDER                              int // not supported! Defaults to CLAMP_TO_EDGE!
	COLOR_ATTACHMENT0                            int
	COLOR_BUFFER_BIT                             int
	COLOR_CLEAR_VALUE                            int
	COLOR_WRITEMASK                              int
	COMPILE_STATUS                               uint32
	COMPRESSED_TEXTURE_FORMATS                   int
	CONSTANT_ALPHA                               int
	CONSTANT_COLOR                               int
	CONTEXT_LOST_WEBGL                           int
	CULL_FACE                                    int
	CULL_FACE_MODE                               int
	CURRENT_PROGRAM                              int
	CURRENT_VERTEX_ATTRIB                        int
	CW                                           int
	DECR                                         int
	DECR_WRAP                                    int
	DELETE_STATUS                                int
	DEPTH_ATTACHMENT                             int
	DEPTH_BITS                                   int
	DEPTH_BUFFER_BIT                             int
	DEPTH_CLEAR_VALUE                            int
	DEPTH_COMPONENT                              int
	DEPTH_COMPONENT16                            int
	DEPTH_FUNC                                   int
	DEPTH_RANGE                                  int
	DEPTH_STENCIL                                int
	DEPTH_STENCIL_ATTACHMENT                     int
	DEPTH_TEST                                   int
	DEPTH_WRITEMASK                              int
	DITHER                                       int
	DONT_CARE                                    int
	DST_ALPHA                                    int
	DST_COLOR                                    int
	DYNAMIC_DRAW                                 int
	ELEMENT_ARRAY_BUFFER                         int
	ELEMENT_ARRAY_BUFFER_BINDING                 int
	EQUAL                                        int
	FASTEST                                      int
	FLOAT                                        int
	FLOAT_MAT2                                   int
	FLOAT_MAT3                                   int
	FLOAT_MAT4                                   int
	FLOAT_VEC2                                   int
	FLOAT_VEC3                                   int
	FLOAT_VEC4                                   int
	FRAGMENT_SHADER                              int
	FRAMEBUFFER                                  int
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME           int
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE           int
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE int
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL         int
	FRAMEBUFFER_BINDING                          int
	FRAMEBUFFER_COMPLETE                         int
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT            int
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS            int
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT    int
	FRAMEBUFFER_UNSUPPORTED                      int
	FRONT                                        int
	FRONT_AND_BACK                               int
	FRONT_FACE                                   int
	FUNC_ADD                                     int
	FUNC_REVERSE_SUBTRACT                        int
	FUNC_SUBTRACT                                int
	GENERATE_MIPMAP_HINT                         int
	GEQUAL                                       int
	GREATER                                      int
	GREEN_BITS                                   int
	HIGH_FLOAT                                   int
	HIGH_INT                                     int
	INCR                                         int
	INCR_WRAP                                    int
	INFO_LOG_LENGTH                              uint32 //not supported!
	INT                                          int
	INT_VEC2                                     int
	INT_VEC3                                     int
	INT_VEC4                                     int
	INVALID_ENUM                                 int
	INVALID_FRAMEBUFFER_OPERATION                int
	INVALID_OPERATION                            int
	INVALID_VALUE                                int
	INVERT                                       int
	KEEP                                         int
	LEQUAL                                       int
	LESS                                         int
	LINEAR                                       int
	LINEAR_MIPMAP_LINEAR                         int
	LINEAR_MIPMAP_NEAREST                        int
	LINES                                        int
	LINE_LOOP                                    int
	LINE_STRIP                                   int
	LINE_WIDTH                                   int
	LINK_STATUS                                  int
	LOW_FLOAT                                    int
	LOW_INT                                      int
	LUMINANCE                                    int
	LUMINANCE_ALPHA                              int
	MAX_COMBINED_TEXTURE_IMAGE_UNITS             int
	MAX_CUBE_MAP_TEXTURE_SIZE                    int
	MAX_FRAGMENT_UNIFORM_VECTORS                 int
	MAX_RENDERBUFFER_SIZE                        int
	MAX_TEXTURE_IMAGE_UNITS                      int
	MAX_TEXTURE_SIZE                             int
	MAX_VARYING_VECTORS                          int
	MAX_VERTEX_ATTRIBS                           int
	MAX_VERTEX_TEXTURE_IMAGE_UNITS               int
	MAX_VERTEX_UNIFORM_VECTORS                   int
	MAX_VIEWPORT_DIMS                            int
	MEDIUM_FLOAT                                 int
	MEDIUM_INT                                   int
	MIRRORED_REPEAT                              int
	MULTISAMPLE                                  int // not supported!
	NEAREST                                      int
	NEAREST_MIPMAP_LINEAR                        int
	NEAREST_MIPMAP_NEAREST                       int
	NEVER                                        int
	NICEST                                       int
	NONE                                         int
	NOTEQUAL                                     int
	NO_ERROR                                     int
	NUM_COMPRESSED_TEXTURE_FORMATS               int // not supported!
	ONE                                          int
	ONE_MINUS_CONSTANT_ALPHA                     int
	ONE_MINUS_CONSTANT_COLOR                     int
	ONE_MINUS_DST_ALPHA                          int
	ONE_MINUS_DST_COLOR                          int
	ONE_MINUS_SRC_ALPHA                          int
	ONE_MINUS_SRC_COLOR                          int
	OUT_OF_MEMORY                                int
	PACK_ALIGNMENT                               int
	POINTS                                       int
	POLYGON_OFFSET_FACTOR                        int
	POLYGON_OFFSET_FILL                          int
	POLYGON_OFFSET_UNITS                         int
	RED_BITS                                     int
	RENDERBUFFER                                 int
	RENDERBUFFER_ALPHA_SIZE                      int
	RENDERBUFFER_BINDING                         int
	RENDERBUFFER_BLUE_SIZE                       int
	RENDERBUFFER_DEPTH_SIZE                      int
	RENDERBUFFER_GREEN_SIZE                      int
	RENDERBUFFER_HEIGHT                          int
	RENDERBUFFER_INTERNAL_FORMAT                 int
	RENDERBUFFER_RED_SIZE                        int
	RENDERBUFFER_STENCIL_SIZE                    int
	RENDERBUFFER_WIDTH                           int
	RENDERER                                     int
	REPEAT                                       int
	REPLACE                                      int
	RGB                                          int
	RGB5_A1                                      int
	RGB565                                       int
	RGBA                                         int
	RGBA4                                        int
	SAMPLER_2D                                   int
	SAMPLER_CUBE                                 int
	SAMPLES                                      int
	SAMPLE_ALPHA_TO_COVERAGE                     int
	SAMPLE_BUFFERS                               int
	SAMPLE_COVERAGE                              int
	SAMPLE_COVERAGE_INVERT                       int
	SAMPLE_COVERAGE_VALUE                        int
	SCISSOR_BOX                                  int
	SCISSOR_TEST                                 int
	SHADER_COMPILER                              int // not supported!
	SHADER_SOURCE_LENGTH                         int // not supported!
	SHADER_TYPE                                  int
	SHADING_LANGUAGE_VERSION                     int
	SHORT                                        int
	SRC_ALPHA                                    int
	SRC_ALPHA_SATURATE                           int
	SRC_COLOR                                    int
	STATIC_DRAW                                  int
	STENCIL_ATTACHMENT                           int
	STENCIL_BACK_FAIL                            int
	STENCIL_BACK_FUNC                            int
	STENCIL_BACK_PASS_DEPTH_FAIL                 int
	STENCIL_BACK_PASS_DEPTH_PASS                 int
	STENCIL_BACK_REF                             int
	STENCIL_BACK_VALUE_MASK                      int
	STENCIL_BACK_WRITEMASK                       int
	STENCIL_BITS                                 int
	STENCIL_BUFFER_BIT                           int
	STENCIL_CLEAR_VALUE                          int
	STENCIL_FAIL                                 int
	STENCIL_FUNC                                 int
	STENCIL_INDEX                                int
	STENCIL_INDEX8                               int
	STENCIL_PASS_DEPTH_FAIL                      int
	STENCIL_PASS_DEPTH_PASS                      int
	STENCIL_REF                                  int
	STENCIL_TEST                                 int
	STENCIL_VALUE_MASK                           int
	STENCIL_WRITEMASK                            int
	STREAM_DRAW                                  int
	SUBPIXEL_BITS                                int
	TEXTURE                                      int
	TEXTURE0                                     int
	TEXTURE1                                     int
	TEXTURE2                                     int
	TEXTURE3                                     int
	TEXTURE4                                     int
	TEXTURE5                                     int
	TEXTURE6                                     int
	TEXTURE7                                     int
	TEXTURE8                                     int
	TEXTURE9                                     int
	TEXTURE10                                    int
	TEXTURE11                                    int
	TEXTURE12                                    int
	TEXTURE13                                    int
	TEXTURE14                                    int
	TEXTURE15                                    int
	TEXTURE16                                    int
	TEXTURE17                                    int
	TEXTURE18                                    int
	TEXTURE19                                    int
	TEXTURE20                                    int
	TEXTURE21                                    int
	TEXTURE22                                    int
	TEXTURE23                                    int
	TEXTURE24                                    int
	TEXTURE25                                    int
	TEXTURE26                                    int
	TEXTURE27                                    int
	TEXTURE28                                    int
	TEXTURE29                                    int
	TEXTURE30                                    int
	TEXTURE31                                    int
	TEXTURE_2D                                   int
	TEXTURE_BINDING_2D                           int
	TEXTURE_BINDING_CUBE_MAP                     int
	TEXTURE_CUBE_MAP                             int
	TEXTURE_CUBE_MAP_NEGATIVE_X                  int
	TEXTURE_CUBE_MAP_NEGATIVE_Y                  int
	TEXTURE_CUBE_MAP_NEGATIVE_Z                  int
	TEXTURE_CUBE_MAP_POSITIVE_X                  int
	TEXTURE_CUBE_MAP_POSITIVE_Y                  int
	TEXTURE_CUBE_MAP_POSITIVE_Z                  int
	TEXTURE_MAG_FILTER                           int
	TEXTURE_MIN_FILTER                           int
	TEXTURE_WRAP_S                               int
	TEXTURE_WRAP_T                               int
	TRIANGLES                                    int
	TRIANGLE_FAN                                 int
	TRIANGLE_STRIP                               int
	UNPACK_ALIGNMENT                             int
	UNPACK_COLORSPACE_CONVERSION_WEBGL           int
	UNPACK_FLIP_Y_WEBGL                          int
	UNPACK_PREMULTIPLY_ALPHA_WEBGL               int
	UNSIGNED_BYTE                                int
	UNSIGNED_INT                                 int
	UNSIGNED_SHORT                               int
	UNSIGNED_SHORT_4_4_4_4                       int
	UNSIGNED_SHORT_5_5_5_1                       int
	UNSIGNED_SHORT_5_6_5                         int
	VALIDATE_STATUS                              int
	VENDOR                                       int // not supported!
	VERSION                                      int
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING           int
	VERTEX_ATTRIB_ARRAY_ENABLED                  int
	VERTEX_ATTRIB_ARRAY_NORMALIZED               int
	VERTEX_ATTRIB_ARRAY_POINTER                  int
	VERTEX_ATTRIB_ARRAY_SIZE                     int
	VERTEX_ATTRIB_ARRAY_STRIDE                   int
	VERTEX_ATTRIB_ARRAY_TYPE                     int
	VERTEX_SHADER                                int
	VIEWPORT                                     int
	ZERO                                         int
	TRUE                                         int
}

// NewContext takes an HTML5 canvas object and optional context attributes.
// If an error is returned it means you won't have access to WebGL
// functionality.
func NewContext(canvas js.Value, ca *ContextAttributes) (*Context, error) {
	if js.Global().Get("WebGLRenderingContext").Type() == js.TypeUndefined {
		return nil, errors.New("Your browser doesn't appear to support webgl.")
	}

	ctx := new(Context)
	ctx.InitialContextValues()

	if ca == nil {
		ca = DefaultAttributes()
	}

	attrs := js.Global().Get("Object").New()
	attrs.Set("alpha", ca.Alpha)
	attrs.Set("depth", ca.Depth)
	attrs.Set("stencil", ca.Stencil)
	attrs.Set("antialias", ca.Antialias)
	attrs.Set("premultipliedAlpha", ca.PremultipliedAlpha)
	attrs.Set("preserveDrawingBuffer", ca.PreserveDrawingBuffer)
	gl := canvas.Call("getContext", "webgl", attrs)
	if gl.Type() == js.TypeNull {
		gl = canvas.Call("getContext", "experimental-webgl", attrs)
		if gl.Type() == js.TypeNull {
			return nil, errors.New("Creating a webgl context has failed.")
		}
	}
	ctx.Value = gl
	return ctx, nil
}

// InitialContextValues sets up the context by retrieving the values from the
// webgl context
func (c *Context) InitialContextValues() {
	webCtx := js.Global().Get("WebGLRenderingContext").Get("prototype")
	c.ARRAY_BUFFER = webCtx.Get("ARRAY_BUFFER").Int()
	c.ARRAY_BUFFER_BINDING = webCtx.Get("ARRAY_BUFFER_BINDING").Int()
	c.ATTACHED_SHADERS = webCtx.Get("ATTACHED_SHADERS").Int()
	c.BACK = webCtx.Get("BACK").Int()
	c.BLEND = webCtx.Get("BLEND").Int()
	c.BLEND_COLOR = webCtx.Get("BLEND_COLOR").Int()
	c.BLEND_DST_ALPHA = webCtx.Get("BLEND_DST_ALPHA").Int()
	c.BLEND_DST_RGB = webCtx.Get("BLEND_DST_RGB").Int()
	c.BLEND_EQUATION = webCtx.Get("BLEND_EQUATION").Int()
	c.BLEND_EQUATION_ALPHA = webCtx.Get("BLEND_EQUATION_ALPHA").Int()
	c.BLEND_EQUATION_RGB = webCtx.Get("BLEND_EQUATION_RGB").Int()
	c.BLEND_SRC_ALPHA = webCtx.Get("BLEND_SRC_ALPHA").Int()
	c.BLEND_SRC_RGB = webCtx.Get("BLEND_SRC_RGB").Int()
	c.BLUE_BITS = webCtx.Get("BLUE_BITS").Int()
	c.BOOL = webCtx.Get("BOOL").Int()
	c.BOOL_VEC2 = webCtx.Get("BOOL_VEC2").Int()
	c.BOOL_VEC3 = webCtx.Get("BOOL_VEC3").Int()
	c.BOOL_VEC4 = webCtx.Get("BOOL_VEC4").Int()
	c.BROWSER_DEFAULT_WEBGL = webCtx.Get("BROWSER_DEFAULT_WEBGL").Int()
	c.BUFFER_SIZE = webCtx.Get("BUFFER_SIZE").Int()
	c.BUFFER_USAGE = webCtx.Get("BUFFER_USAGE").Int()
	c.BYTE = webCtx.Get("BYTE").Int()
	c.CCW = webCtx.Get("CCW").Int()
	c.CLAMP_TO_EDGE = webCtx.Get("CLAMP_TO_EDGE").Int()
	c.CLAMP_TO_BORDER = webCtx.Get("CLAMP_TO_EDGE").Int()
	c.COLOR_ATTACHMENT0 = webCtx.Get("COLOR_ATTACHMENT0").Int()
	c.COLOR_BUFFER_BIT = webCtx.Get("COLOR_BUFFER_BIT").Int()
	c.COLOR_CLEAR_VALUE = webCtx.Get("COLOR_CLEAR_VALUE").Int()
	c.COLOR_WRITEMASK = webCtx.Get("COLOR_WRITEMASK").Int()
	c.COMPILE_STATUS = uint32(webCtx.Get("COMPILE_STATUS").Int())
	c.COMPRESSED_TEXTURE_FORMATS = webCtx.Get("COMPRESSED_TEXTURE_FORMATS").Int()
	c.CONSTANT_ALPHA = webCtx.Get("CONSTANT_ALPHA").Int()
	c.CONSTANT_COLOR = webCtx.Get("CONSTANT_COLOR").Int()
	c.CONTEXT_LOST_WEBGL = webCtx.Get("CONTEXT_LOST_WEBGL").Int()
	c.CULL_FACE = webCtx.Get("CULL_FACE").Int()
	c.CULL_FACE_MODE = webCtx.Get("CULL_FACE_MODE").Int()
	c.CURRENT_PROGRAM = webCtx.Get("CURRENT_PROGRAM").Int()
	c.CURRENT_VERTEX_ATTRIB = webCtx.Get("CURRENT_VERTEX_ATTRIB").Int()
	c.CW = webCtx.Get("CW").Int()
	c.DECR = webCtx.Get("DECR").Int()
	c.DECR_WRAP = webCtx.Get("DECR_WRAP").Int()
	c.DELETE_STATUS = webCtx.Get("DELETE_STATUS").Int()
	c.DEPTH_ATTACHMENT = webCtx.Get("DEPTH_ATTACHMENT").Int()
	c.DEPTH_BITS = webCtx.Get("DEPTH_BITS").Int()
	c.DEPTH_BUFFER_BIT = webCtx.Get("DEPTH_BUFFER_BIT").Int()
	c.DEPTH_CLEAR_VALUE = webCtx.Get("DEPTH_CLEAR_VALUE").Int()
	c.DEPTH_COMPONENT = webCtx.Get("DEPTH_COMPONENT").Int()
	c.DEPTH_COMPONENT16 = webCtx.Get("DEPTH_COMPONENT16").Int()
	c.DEPTH_FUNC = webCtx.Get("DEPTH_FUNC").Int()
	c.DEPTH_RANGE = webCtx.Get("DEPTH_RANGE").Int()
	c.DEPTH_STENCIL = webCtx.Get("DEPTH_STENCIL").Int()
	c.DEPTH_STENCIL_ATTACHMENT = webCtx.Get("DEPTH_STENCIL_ATTACHMENT").Int()
	c.DEPTH_TEST = webCtx.Get("DEPTH_TEST").Int()
	c.DEPTH_WRITEMASK = webCtx.Get("DEPTH_WRITEMASK").Int()
	c.DITHER = webCtx.Get("DITHER").Int()
	c.DONT_CARE = webCtx.Get("DONT_CARE").Int()
	c.DST_ALPHA = webCtx.Get("DST_ALPHA").Int()
	c.DST_COLOR = webCtx.Get("DST_COLOR").Int()
	c.DYNAMIC_DRAW = webCtx.Get("DYNAMIC_DRAW").Int()
	c.ELEMENT_ARRAY_BUFFER = webCtx.Get("ELEMENT_ARRAY_BUFFER").Int()
	c.ELEMENT_ARRAY_BUFFER_BINDING = webCtx.Get("ELEMENT_ARRAY_BUFFER_BINDING").Int()
	c.EQUAL = webCtx.Get("EQUAL").Int()
	c.FASTEST = webCtx.Get("FASTEST").Int()
	c.FLOAT = webCtx.Get("FLOAT").Int()
	c.FLOAT_MAT2 = webCtx.Get("FLOAT_MAT2").Int()
	c.FLOAT_MAT3 = webCtx.Get("FLOAT_MAT3").Int()
	c.FLOAT_MAT4 = webCtx.Get("FLOAT_MAT4").Int()
	c.FLOAT_VEC2 = webCtx.Get("FLOAT_VEC2").Int()
	c.FLOAT_VEC3 = webCtx.Get("FLOAT_VEC3").Int()
	c.FLOAT_VEC4 = webCtx.Get("FLOAT_VEC4").Int()
	c.FRAGMENT_SHADER = webCtx.Get("FRAGMENT_SHADER").Int()
	c.FRAMEBUFFER = webCtx.Get("FRAMEBUFFER").Int()
	c.FRAMEBUFFER_ATTACHMENT_OBJECT_NAME = webCtx.Get("FRAMEBUFFER_ATTACHMENT_OBJECT_NAME").Int()
	c.FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE = webCtx.Get("FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE").Int()
	c.FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE = webCtx.Get("FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE").Int()
	c.FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL = webCtx.Get("FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL").Int()
	c.FRAMEBUFFER_BINDING = webCtx.Get("FRAMEBUFFER_BINDING").Int()
	c.FRAMEBUFFER_COMPLETE = webCtx.Get("FRAMEBUFFER_COMPLETE").Int()
	c.FRAMEBUFFER_INCOMPLETE_ATTACHMENT = webCtx.Get("FRAMEBUFFER_INCOMPLETE_ATTACHMENT").Int()
	c.FRAMEBUFFER_INCOMPLETE_DIMENSIONS = webCtx.Get("FRAMEBUFFER_INCOMPLETE_DIMENSIONS").Int()
	c.FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT = webCtx.Get("FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT").Int()
	c.FRAMEBUFFER_UNSUPPORTED = webCtx.Get("FRAMEBUFFER_UNSUPPORTED").Int()
	c.FRONT = webCtx.Get("FRONT").Int()
	c.FRONT_AND_BACK = webCtx.Get("FRONT_AND_BACK").Int()
	c.FRONT_FACE = webCtx.Get("FRONT_FACE").Int()
	c.FUNC_ADD = webCtx.Get("FUNC_ADD").Int()
	c.FUNC_REVERSE_SUBTRACT = webCtx.Get("FUNC_REVERSE_SUBTRACT").Int()
	c.FUNC_SUBTRACT = webCtx.Get("FUNC_SUBTRACT").Int()
	c.GENERATE_MIPMAP_HINT = webCtx.Get("GENERATE_MIPMAP_HINT").Int()
	c.GEQUAL = webCtx.Get("GEQUAL").Int()
	c.GREATER = webCtx.Get("GREATER").Int()
	c.GREEN_BITS = webCtx.Get("GREEN_BITS").Int()
	c.HIGH_FLOAT = webCtx.Get("HIGH_FLOAT").Int()
	c.HIGH_INT = webCtx.Get("HIGH_INT").Int()
	c.INCR = webCtx.Get("INCR").Int()
	c.INCR_WRAP = webCtx.Get("INCR_WRAP").Int()
	c.INFO_LOG_LENGTH = uint32(0)
	c.INT = webCtx.Get("INT").Int()
	c.INT_VEC2 = webCtx.Get("INT_VEC2").Int()
	c.INT_VEC3 = webCtx.Get("INT_VEC3").Int()
	c.INT_VEC4 = webCtx.Get("INT_VEC4").Int()
	c.INVALID_ENUM = webCtx.Get("INVALID_ENUM").Int()
	c.INVALID_FRAMEBUFFER_OPERATION = webCtx.Get("INVALID_FRAMEBUFFER_OPERATION").Int()
	c.INVALID_OPERATION = webCtx.Get("INVALID_OPERATION").Int()
	c.INVALID_VALUE = webCtx.Get("INVALID_VALUE").Int()
	c.INVERT = webCtx.Get("INVERT").Int()
	c.KEEP = webCtx.Get("KEEP").Int()
	c.LEQUAL = webCtx.Get("LEQUAL").Int()
	c.LESS = webCtx.Get("LESS").Int()
	c.LINEAR = webCtx.Get("LINEAR").Int()
	c.LINEAR_MIPMAP_LINEAR = webCtx.Get("LINEAR_MIPMAP_LINEAR").Int()
	c.LINEAR_MIPMAP_NEAREST = webCtx.Get("LINEAR_MIPMAP_NEAREST").Int()
	c.LINES = webCtx.Get("LINES").Int()
	c.LINE_LOOP = webCtx.Get("LINE_LOOP").Int()
	c.LINE_STRIP = webCtx.Get("LINE_STRIP").Int()
	c.LINE_WIDTH = webCtx.Get("LINE_WIDTH").Int()
	c.LINK_STATUS = webCtx.Get("LINK_STATUS").Int()
	c.LOW_FLOAT = webCtx.Get("LOW_FLOAT").Int()
	c.LOW_INT = webCtx.Get("LOW_INT").Int()
	c.LUMINANCE = webCtx.Get("LUMINANCE").Int()
	c.LUMINANCE_ALPHA = webCtx.Get("LUMINANCE_ALPHA").Int()
	c.MAX_COMBINED_TEXTURE_IMAGE_UNITS = webCtx.Get("MAX_COMBINED_TEXTURE_IMAGE_UNITS").Int()
	c.MAX_CUBE_MAP_TEXTURE_SIZE = webCtx.Get("MAX_CUBE_MAP_TEXTURE_SIZE").Int()
	c.MAX_FRAGMENT_UNIFORM_VECTORS = webCtx.Get("MAX_FRAGMENT_UNIFORM_VECTORS").Int()
	c.MAX_RENDERBUFFER_SIZE = webCtx.Get("MAX_RENDERBUFFER_SIZE").Int()
	c.MAX_TEXTURE_IMAGE_UNITS = webCtx.Get("MAX_TEXTURE_IMAGE_UNITS").Int()
	c.MAX_TEXTURE_SIZE = webCtx.Get("MAX_TEXTURE_SIZE").Int()
	c.MAX_VARYING_VECTORS = webCtx.Get("MAX_VARYING_VECTORS").Int()
	c.MAX_VERTEX_ATTRIBS = webCtx.Get("MAX_VERTEX_ATTRIBS").Int()
	c.MAX_VERTEX_TEXTURE_IMAGE_UNITS = webCtx.Get("MAX_VERTEX_TEXTURE_IMAGE_UNITS").Int()
	c.MAX_VERTEX_UNIFORM_VECTORS = webCtx.Get("MAX_VERTEX_UNIFORM_VECTORS").Int()
	c.MAX_VIEWPORT_DIMS = webCtx.Get("MAX_VIEWPORT_DIMS").Int()
	c.MEDIUM_FLOAT = webCtx.Get("MEDIUM_FLOAT").Int()
	c.MEDIUM_INT = webCtx.Get("MEDIUM_INT").Int()
	c.MIRRORED_REPEAT = webCtx.Get("MIRRORED_REPEAT").Int()
	c.MULTISAMPLE = 0
	c.NEAREST = webCtx.Get("NEAREST").Int()
	c.NEAREST_MIPMAP_LINEAR = webCtx.Get("NEAREST_MIPMAP_LINEAR").Int()
	c.NEAREST_MIPMAP_NEAREST = webCtx.Get("NEAREST_MIPMAP_NEAREST").Int()
	c.NEVER = webCtx.Get("NEVER").Int()
	c.NICEST = webCtx.Get("NICEST").Int()
	c.NONE = webCtx.Get("NONE").Int()
	c.NOTEQUAL = webCtx.Get("NOTEQUAL").Int()
	c.NO_ERROR = webCtx.Get("NO_ERROR").Int()
	c.NUM_COMPRESSED_TEXTURE_FORMATS = 0
	c.ONE = webCtx.Get("ONE").Int()
	c.ONE_MINUS_CONSTANT_ALPHA = webCtx.Get("ONE_MINUS_CONSTANT_ALPHA").Int()
	c.ONE_MINUS_CONSTANT_COLOR = webCtx.Get("ONE_MINUS_CONSTANT_COLOR").Int()
	c.ONE_MINUS_DST_ALPHA = webCtx.Get("ONE_MINUS_DST_ALPHA").Int()
	c.ONE_MINUS_DST_COLOR = webCtx.Get("ONE_MINUS_DST_COLOR").Int()
	c.ONE_MINUS_SRC_ALPHA = webCtx.Get("ONE_MINUS_SRC_ALPHA").Int()
	c.ONE_MINUS_SRC_COLOR = webCtx.Get("ONE_MINUS_SRC_COLOR").Int()
	c.OUT_OF_MEMORY = webCtx.Get("OUT_OF_MEMORY").Int()
	c.PACK_ALIGNMENT = webCtx.Get("PACK_ALIGNMENT").Int()
	c.POINTS = webCtx.Get("POINTS").Int()
	c.POLYGON_OFFSET_FACTOR = webCtx.Get("POLYGON_OFFSET_FACTOR").Int()
	c.POLYGON_OFFSET_FILL = webCtx.Get("POLYGON_OFFSET_FILL").Int()
	c.POLYGON_OFFSET_UNITS = webCtx.Get("POLYGON_OFFSET_UNITS").Int()
	c.RED_BITS = webCtx.Get("RED_BITS").Int()
	c.RENDERBUFFER = webCtx.Get("RENDERBUFFER").Int()
	c.RENDERBUFFER_ALPHA_SIZE = webCtx.Get("RENDERBUFFER_ALPHA_SIZE").Int()
	c.RENDERBUFFER_BINDING = webCtx.Get("RENDERBUFFER_BINDING").Int()
	c.RENDERBUFFER_BLUE_SIZE = webCtx.Get("RENDERBUFFER_BLUE_SIZE").Int()
	c.RENDERBUFFER_DEPTH_SIZE = webCtx.Get("RENDERBUFFER_DEPTH_SIZE").Int()
	c.RENDERBUFFER_GREEN_SIZE = webCtx.Get("RENDERBUFFER_GREEN_SIZE").Int()
	c.RENDERBUFFER_HEIGHT = webCtx.Get("RENDERBUFFER_HEIGHT").Int()
	c.RENDERBUFFER_INTERNAL_FORMAT = webCtx.Get("RENDERBUFFER_INTERNAL_FORMAT").Int()
	c.RENDERBUFFER_RED_SIZE = webCtx.Get("RENDERBUFFER_RED_SIZE").Int()
	c.RENDERBUFFER_STENCIL_SIZE = webCtx.Get("RENDERBUFFER_STENCIL_SIZE").Int()
	c.RENDERBUFFER_WIDTH = webCtx.Get("RENDERBUFFER_WIDTH").Int()
	c.RENDERER = webCtx.Get("RENDERER").Int()
	c.REPEAT = webCtx.Get("REPEAT").Int()
	c.REPLACE = webCtx.Get("REPLACE").Int()
	c.RGB = webCtx.Get("RGB").Int()
	c.RGB5_A1 = webCtx.Get("RGB5_A1").Int()
	c.RGB565 = webCtx.Get("RGB565").Int()
	c.RGBA = webCtx.Get("RGBA").Int()
	c.RGBA4 = webCtx.Get("RGBA4").Int()
	c.SAMPLER_2D = webCtx.Get("SAMPLER_2D").Int()
	c.SAMPLER_CUBE = webCtx.Get("SAMPLER_CUBE").Int()
	c.SAMPLES = webCtx.Get("SAMPLES").Int()
	c.SAMPLE_ALPHA_TO_COVERAGE = webCtx.Get("SAMPLE_ALPHA_TO_COVERAGE").Int()
	c.SAMPLE_BUFFERS = webCtx.Get("SAMPLE_BUFFERS").Int()
	c.SAMPLE_COVERAGE = webCtx.Get("SAMPLE_COVERAGE").Int()
	c.SAMPLE_COVERAGE_INVERT = webCtx.Get("SAMPLE_COVERAGE_INVERT").Int()
	c.SAMPLE_COVERAGE_VALUE = webCtx.Get("SAMPLE_COVERAGE_VALUE").Int()
	c.SCISSOR_BOX = webCtx.Get("SCISSOR_BOX").Int()
	c.SCISSOR_TEST = webCtx.Get("SCISSOR_TEST").Int()
	c.SHADER_TYPE = webCtx.Get("SHADER_TYPE").Int()
	c.SHADING_LANGUAGE_VERSION = webCtx.Get("SHADING_LANGUAGE_VERSION").Int()
	c.SHORT = webCtx.Get("SHORT").Int()
	c.SRC_ALPHA = webCtx.Get("SRC_ALPHA").Int()
	c.SRC_ALPHA_SATURATE = webCtx.Get("SRC_ALPHA_SATURATE").Int()
	c.SRC_COLOR = webCtx.Get("SRC_COLOR").Int()
	c.STATIC_DRAW = webCtx.Get("STATIC_DRAW").Int()
	c.STENCIL_ATTACHMENT = webCtx.Get("STENCIL_ATTACHMENT").Int()
	c.STENCIL_BACK_FAIL = webCtx.Get("STENCIL_BACK_FAIL").Int()
	c.STENCIL_BACK_FUNC = webCtx.Get("STENCIL_BACK_FUNC").Int()
	c.STENCIL_BACK_PASS_DEPTH_FAIL = webCtx.Get("STENCIL_BACK_PASS_DEPTH_FAIL").Int()
	c.STENCIL_BACK_PASS_DEPTH_PASS = webCtx.Get("STENCIL_BACK_PASS_DEPTH_PASS").Int()
	c.STENCIL_BACK_REF = webCtx.Get("STENCIL_BACK_REF").Int()
	c.STENCIL_BACK_VALUE_MASK = webCtx.Get("STENCIL_BACK_VALUE_MASK").Int()
	c.STENCIL_BACK_WRITEMASK = webCtx.Get("STENCIL_BACK_WRITEMASK").Int()
	c.STENCIL_BITS = webCtx.Get("STENCIL_BITS").Int()
	c.STENCIL_BUFFER_BIT = webCtx.Get("STENCIL_BUFFER_BIT").Int()
	c.STENCIL_CLEAR_VALUE = webCtx.Get("STENCIL_CLEAR_VALUE").Int()
	c.STENCIL_FAIL = webCtx.Get("STENCIL_FAIL").Int()
	c.STENCIL_FUNC = webCtx.Get("STENCIL_FUNC").Int()
	c.STENCIL_INDEX = webCtx.Get("STENCIL_INDEX").Int()
	c.STENCIL_INDEX8 = webCtx.Get("STENCIL_INDEX8").Int()
	c.STENCIL_PASS_DEPTH_FAIL = webCtx.Get("STENCIL_PASS_DEPTH_FAIL").Int()
	c.STENCIL_PASS_DEPTH_PASS = webCtx.Get("STENCIL_PASS_DEPTH_PASS").Int()
	c.STENCIL_REF = webCtx.Get("STENCIL_REF").Int()
	c.STENCIL_TEST = webCtx.Get("STENCIL_TEST").Int()
	c.STENCIL_VALUE_MASK = webCtx.Get("STENCIL_VALUE_MASK").Int()
	c.STENCIL_WRITEMASK = webCtx.Get("STENCIL_WRITEMASK").Int()
	c.STREAM_DRAW = webCtx.Get("STREAM_DRAW").Int()
	c.SUBPIXEL_BITS = webCtx.Get("SUBPIXEL_BITS").Int()
	c.TEXTURE = webCtx.Get("TEXTURE").Int()
	c.TEXTURE0 = webCtx.Get("TEXTURE0").Int()
	c.TEXTURE1 = webCtx.Get("TEXTURE1").Int()
	c.TEXTURE2 = webCtx.Get("TEXTURE2").Int()
	c.TEXTURE3 = webCtx.Get("TEXTURE3").Int()
	c.TEXTURE4 = webCtx.Get("TEXTURE4").Int()
	c.TEXTURE5 = webCtx.Get("TEXTURE5").Int()
	c.TEXTURE6 = webCtx.Get("TEXTURE6").Int()
	c.TEXTURE7 = webCtx.Get("TEXTURE7").Int()
	c.TEXTURE8 = webCtx.Get("TEXTURE8").Int()
	c.TEXTURE9 = webCtx.Get("TEXTURE9").Int()
	c.TEXTURE10 = webCtx.Get("TEXTURE10").Int()
	c.TEXTURE11 = webCtx.Get("TEXTURE11").Int()
	c.TEXTURE12 = webCtx.Get("TEXTURE12").Int()
	c.TEXTURE13 = webCtx.Get("TEXTURE13").Int()
	c.TEXTURE14 = webCtx.Get("TEXTURE14").Int()
	c.TEXTURE15 = webCtx.Get("TEXTURE15").Int()
	c.TEXTURE16 = webCtx.Get("TEXTURE16").Int()
	c.TEXTURE17 = webCtx.Get("TEXTURE17").Int()
	c.TEXTURE18 = webCtx.Get("TEXTURE18").Int()
	c.TEXTURE19 = webCtx.Get("TEXTURE19").Int()
	c.TEXTURE20 = webCtx.Get("TEXTURE20").Int()
	c.TEXTURE21 = webCtx.Get("TEXTURE21").Int()
	c.TEXTURE22 = webCtx.Get("TEXTURE22").Int()
	c.TEXTURE23 = webCtx.Get("TEXTURE23").Int()
	c.TEXTURE24 = webCtx.Get("TEXTURE24").Int()
	c.TEXTURE25 = webCtx.Get("TEXTURE25").Int()
	c.TEXTURE26 = webCtx.Get("TEXTURE26").Int()
	c.TEXTURE27 = webCtx.Get("TEXTURE27").Int()
	c.TEXTURE28 = webCtx.Get("TEXTURE28").Int()
	c.TEXTURE29 = webCtx.Get("TEXTURE29").Int()
	c.TEXTURE30 = webCtx.Get("TEXTURE30").Int()
	c.TEXTURE31 = webCtx.Get("TEXTURE31").Int()
	c.TEXTURE_2D = webCtx.Get("TEXTURE_2D").Int()
	c.TEXTURE_BINDING_2D = webCtx.Get("TEXTURE_BINDING_2D").Int()
	c.TEXTURE_BINDING_CUBE_MAP = webCtx.Get("TEXTURE_BINDING_CUBE_MAP").Int()
	c.TEXTURE_CUBE_MAP = webCtx.Get("TEXTURE_CUBE_MAP").Int()
	c.TEXTURE_CUBE_MAP_NEGATIVE_X = webCtx.Get("TEXTURE_CUBE_MAP_NEGATIVE_X").Int()
	c.TEXTURE_CUBE_MAP_NEGATIVE_Y = webCtx.Get("TEXTURE_CUBE_MAP_NEGATIVE_Y").Int()
	c.TEXTURE_CUBE_MAP_NEGATIVE_Z = webCtx.Get("TEXTURE_CUBE_MAP_NEGATIVE_Z").Int()
	c.TEXTURE_CUBE_MAP_POSITIVE_X = webCtx.Get("TEXTURE_CUBE_MAP_POSITIVE_X").Int()
	c.TEXTURE_CUBE_MAP_POSITIVE_Y = webCtx.Get("TEXTURE_CUBE_MAP_POSITIVE_Y").Int()
	c.TEXTURE_CUBE_MAP_POSITIVE_Z = webCtx.Get("TEXTURE_CUBE_MAP_POSITIVE_Z").Int()
	c.TEXTURE_MAG_FILTER = webCtx.Get("TEXTURE_MAG_FILTER").Int()
	c.TEXTURE_MIN_FILTER = webCtx.Get("TEXTURE_MIN_FILTER").Int()
	c.TEXTURE_WRAP_S = webCtx.Get("TEXTURE_WRAP_S").Int()
	c.TEXTURE_WRAP_T = webCtx.Get("TEXTURE_WRAP_T").Int()
	c.TRIANGLES = webCtx.Get("TRIANGLES").Int()
	c.TRIANGLE_FAN = webCtx.Get("TRIANGLE_FAN").Int()
	c.TRIANGLE_STRIP = webCtx.Get("TRIANGLE_STRIP").Int()
	c.UNPACK_ALIGNMENT = webCtx.Get("UNPACK_ALIGNMENT").Int()
	c.UNPACK_COLORSPACE_CONVERSION_WEBGL = webCtx.Get("UNPACK_COLORSPACE_CONVERSION_WEBGL").Int()
	c.UNPACK_FLIP_Y_WEBGL = webCtx.Get("UNPACK_FLIP_Y_WEBGL").Int()
	c.UNPACK_PREMULTIPLY_ALPHA_WEBGL = webCtx.Get("UNPACK_PREMULTIPLY_ALPHA_WEBGL").Int()
	c.UNSIGNED_BYTE = webCtx.Get("UNSIGNED_BYTE").Int()
	c.UNSIGNED_INT = webCtx.Get("UNSIGNED_INT").Int()
	c.UNSIGNED_SHORT = webCtx.Get("UNSIGNED_SHORT").Int()
	c.UNSIGNED_SHORT_4_4_4_4 = webCtx.Get("UNSIGNED_SHORT_4_4_4_4").Int()
	c.UNSIGNED_SHORT_5_5_5_1 = webCtx.Get("UNSIGNED_SHORT_5_5_5_1").Int()
	c.UNSIGNED_SHORT_5_6_5 = webCtx.Get("UNSIGNED_SHORT_5_6_5").Int()
	c.VALIDATE_STATUS = webCtx.Get("VALIDATE_STATUS").Int()
	c.VERSION = webCtx.Get("VERSION").Int()
	c.VERTEX_ATTRIB_ARRAY_BUFFER_BINDING = webCtx.Get("VERTEX_ATTRIB_ARRAY_BUFFER_BINDING").Int()
	c.VERTEX_ATTRIB_ARRAY_ENABLED = webCtx.Get("VERTEX_ATTRIB_ARRAY_ENABLED").Int()
	c.VERTEX_ATTRIB_ARRAY_NORMALIZED = webCtx.Get("VERTEX_ATTRIB_ARRAY_NORMALIZED").Int()
	c.VERTEX_ATTRIB_ARRAY_POINTER = webCtx.Get("VERTEX_ATTRIB_ARRAY_POINTER").Int()
	c.VERTEX_ATTRIB_ARRAY_SIZE = webCtx.Get("VERTEX_ATTRIB_ARRAY_SIZE").Int()
	c.VERTEX_ATTRIB_ARRAY_STRIDE = webCtx.Get("VERTEX_ATTRIB_ARRAY_STRIDE").Int()
	c.VERTEX_ATTRIB_ARRAY_TYPE = webCtx.Get("VERTEX_ATTRIB_ARRAY_TYPE").Int()
	c.VERTEX_SHADER = webCtx.Get("VERTEX_SHADER").Int()
	c.VIEWPORT = webCtx.Get("VIEWPORT").Int()
	c.ZERO = webCtx.Get("ZERO").Int()
	c.TRUE = 1
}

// Returns the context attributes active on the context. These values might
// be different than what was requested on context creation if the
// browser's implementation doesn't support a feature.
func (c *Context) GetContextAttributes() ContextAttributes {
	ca := c.Call("getContextAttributes")
	return ContextAttributes{
		ca.Get("alpha").Bool(),
		ca.Get("depth").Bool(),
		ca.Get("stencil").Bool(),
		ca.Get("antialias").Bool(),
		ca.Get("premultipliedAlpha").Bool(),
		ca.Get("preservedDrawingBuffer").Bool(),
	}
}

// PerFragment ---------------------------------------------------------------

// The GL_BLEND_COLOR may be used to calculate the source and destination blending factors.
func (c *Context) BlendColor(r, g, b, a float64) {
	c.Call("blendColor", r, g, b, a)
}

// Sets the equation used to blend RGB and Alpha values of an incoming source
// fragment with a destination values as stored in the fragment's frame buffer.
func (c *Context) BlendEquation(mode int) {
	c.Call("blendEquation", mode)
}

// Controls the blending of an incoming source fragment's R, G, B, and A values
// with a destination R, G, B, and A values as stored in the fragment's WebGLFramebuffer.
func (c *Context) BlendEquationSeparate(modeRGB, modeAlpha int) {
	c.Call("blendEquationSeparate", modeRGB, modeAlpha)
}

// Sets the blending factors used to combine source and destination pixels.
func (c *Context) BlendFunc(sfactor, dfactor int) {
	c.Call("blendFunc", sfactor, dfactor)
}

// Sets the weighting factors that are used by blendEquationSeparate.
func (c *Context) BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha int) {
	c.Call("blendFuncSeparate", srcRGB, dstRGB, srcAlpha, dstAlpha)
}

// Sets a function to use to compare incoming pixel depth to the
// current depth buffer value.
func (c *Context) DepthFunc(fun int) {
	c.Call("depthFunc", fun)
}

func (c *Context) SampleCoverage(value float64, invert bool) {
	c.Call("sampleCoverage", value, invert)
}

func (c *Context) StencilFunc(function, ref, mask int) {
	c.Call("stencilFunc", function, ref, mask)
}

func (c *Context) StencilFuncSeparate(face, function, ref, mask int) {
	c.Call("stencilFuncSeparate", face, function, ref, mask)
}

// public function stencilOp(fail:GLenum, zfail:GLenum, zpass:GLenum) : Void;
// public function stencilOpSeparate(face:GLenum, fail:GLenum, zfail:GLenum, zpass:GLenum) : Void;

// FrameBuffer

// ---------------------------------------------------------------------------

// Specifies the active texture unit.
func (c *Context) ActiveTexture(texture int) {
	c.Call("activeTexture", texture)
}

// Attaches a WebGLShader object to a WebGLProgram object.
func (c *Context) AttachShader(program *Program, shader *Shader) {
	c.Call("attachShader", program.Value, shader.Value)
}

// Binds a generic vertex index to a user-defined attribute variable.
func (c *Context) BindAttribLocation(program *Program, index int, name string) {
	c.Call("bindAttribLocation", program.Value, index, name)
}

// Associates a buffer with a buffer target.
func (c *Context) BindBuffer(target int, buffer *Buffer) {
	if buffer == nil {
		c.Call("bindBuffer", target, nil)
		return
	}
	c.Call("bindBuffer", target, buffer.Value)
}

// Associates a WebGLFramebuffer object with the FRAMEBUFFER bind target.
func (c *Context) BindFramebuffer(target int, framebuffer *FrameBuffer) {
	if framebuffer == nil {
		c.Call("bindFramebuffer", target, nil)
		return
	}
	c.Call("bindFramebuffer", target, framebuffer.Value)
}

// Binds a WebGLRenderbuffer object to be used for rendering.
func (c *Context) BindRenderbuffer(target int, renderbuffer *RenderBuffer) {
	if renderbuffer == nil {
		c.Call("bindRenderBuffer", target, nil)
		return
	}
	c.Call("bindRenderbuffer", target, renderbuffer)
}

// Binds a named texture object to a target.
func (c *Context) BindTexture(target int, texture *Texture) {
	if texture == nil {
		c.Call("bindTexture", target, nil)
		return
	}
	c.Call("bindTexture", target, texture.Value)
}

// Creates a buffer in memory and initializes it with array data.
// If no array is provided, the contents of the buffer is initialized to 0.
func (c *Context) BufferData(target int, data interface{}, usage int) {
	buf, ok := data.([]uint8)
	d := js.TypedArrayOf([]uint8{})
	if ok {
		d = js.TypedArrayOf(buf)
	}
	buf2, ok := data.([]uint16)
	if ok {
		d = js.TypedArrayOf(buf2)
	}
	buf3, ok := data.([]uint32)
	if ok {
		d = js.TypedArrayOf(buf3)
	}
	buff32, ok := data.([]float32)
	if ok {
		d = js.TypedArrayOf(buff32)
	}
	c.Call("bufferData", target, d, usage)
	d.Release()
}

// Used to modify or update some or all of a data store for a bound buffer object.
func (c *Context) BufferSubData(target int, offset int, data interface{}) {
	buf, ok := data.([]uint8)
	d := js.TypedArrayOf([]uint8{})
	if ok {
		d = js.TypedArrayOf(buf)
	}
	buf2, ok := data.([]uint16)
	if ok {
		d = js.TypedArrayOf(buf2)
	}
	buf3, ok := data.([]uint32)
	if ok {
		d = js.TypedArrayOf(buf3)
	}
	buff32, ok := data.([]float32)
	if ok {
		d = js.TypedArrayOf(buff32)
	}
	c.Call("bufferSubData", target, offset, d)
	d.Release()
}

// Returns whether the currently bound WebGLFramebuffer is complete.
// If not complete, returns the reason why.
func (c *Context) CheckFramebufferStatus(target int) int {
	return c.Call("checkFramebufferStatus", target).Int()
}

// Sets all pixels in a specific buffer to the same value.
func (c *Context) Clear(flags int) {
	c.Call("clear", flags)
}

// Specifies color values to use by the clear method to clear the color buffer.
func (c *Context) ClearColor(r, g, b, a float32) {
	c.Call("clearColor", r, g, b, a)
}

// Clears the depth buffer to a specific value.
func (c *Context) ClearDepth(depth float64) {
	c.Call("clearDepth", depth)
}

func (c *Context) ClearStencil(s int) {
	c.Call("clearStencil", s)
}

// Lets you set whether individual colors can be written when
// drawing or rendering to a framebuffer.
func (c *Context) ColorMask(r, g, b, a bool) {
	c.Call("colorMask", r, g, b, a)
}

// Compiles the GLSL shader source into binary data used by the WebGLProgram object.
func (c *Context) CompileShader(shader *Shader) {
	c.Call("compileShader", shader.Value)

	status := c.Call("getShaderParameter", shader.Value, c.COMPILE_STATUS)
	if !status.Bool() {
		info := c.Call("getShaderInfoLog", shader.Value)
		log.Print("Error: could not compile shader:", info.String())
	}
}

// Copies a rectangle of pixels from the current WebGLFramebuffer into a texture image.
func (c *Context) CopyTexImage2D(target, level, internal, x, y, w, h, border int) {
	c.Call("copyTexImage2D", target, level, internal, x, y, w, h, border)
}

// Replaces a portion of an existing 2D texture image with data from the current framebuffer.
func (c *Context) CopyTexSubImage2D(target, level, xoffset, yoffset, x, y, w, h int) {
	c.Call("copyTexSubImage2D", target, level, xoffset, yoffset, x, y, w, h)
}

// Creates and initializes a WebGLBuffer.
func (c *Context) CreateBuffer() *Buffer {
	b := &Buffer{c.Call("createBuffer")}
	return b
}

// Returns a WebGLFramebuffer object.
func (c *Context) CreateFramebuffer() *FrameBuffer {
	return &FrameBuffer{c.Call("createFramebuffer")}
}

// Creates an empty WebGLProgram object to which vector and fragment
// WebGLShader objects can be bound.
func (c *Context) CreateProgram() *Program {
	return &Program{c.Call("createProgram")}
}

// Creates and returns a WebGLRenderbuffer object.
func (c *Context) CreateRenderbuffer() *RenderBuffer {
	return &RenderBuffer{c.Call("createRenderbuffer")}
}

// Returns an empty vertex or fragment shader object based on the type specified.
func (c *Context) CreateShader(typ int) *Shader {
	return &Shader{c.Call("createShader", typ)}
}

// Used to generate a WebGLTexture object to which images can be bound.
func (c *Context) CreateTexture() *Texture {
	return &Texture{c.Call("createTexture")}
}

// Sets whether or not front, back, or both facing facets are able to be culled.
func (c *Context) CullFace(mode int) {
	c.Call("cullFace", mode)
}

// Delete a specific buffer.
func (c *Context) DeleteBuffer(buffer *Buffer) {
	c.Call("deleteBuffer", buffer)
}

// Deletes a specific WebGLFramebuffer object. If you delete the
// currently bound framebuffer, the default framebuffer will be bound.
// Deleting a framebuffer detaches all of its attachments.
func (c *Context) DeleteFramebuffer(framebuffer *FrameBuffer) {
	c.Call("deleteFramebuffer", framebuffer)
}

// Flags a specific WebGLProgram object for deletion if currently active.
// It will be deleted when it is no longer being used.
// Any shader objects associated with the program will be detached.
// They will be deleted if they were already flagged for deletion.
func (c *Context) DeleteProgram(program *Program) {
	c.Call("deleteProgram", program.Value)
}

// Deletes the specified renderbuffer object. If the renderbuffer is
// currently bound, it will become unbound. If the renderbuffer is
// attached to the currently bound framebuffer, it is detached.
func (c *Context) DeleteRenderbuffer(renderbuffer *RenderBuffer) {
	c.Call("deleteRenderbuffer", renderbuffer.Value)
}

// Deletes a specific shader object.
func (c *Context) DeleteShader(shader *Shader) {
	c.Call("deleteShader", shader.Value)
}

// Deletes a specific texture object.
func (c *Context) DeleteTexture(texture *Texture) {
	c.Call("deleteTexture", texture.Value)
}

// Sets whether or not you can write to the depth buffer.
func (c *Context) DepthMask(flag bool) {
	c.Call("depthMask", flag)
}

// Sets the depth range for normalized coordinates to canvas or viewport depth coordinates.
func (c *Context) DepthRange(zNear, zFar float64) {
	c.Call("depthRange", zNear, zFar)
}

// Detach a shader object from a program object.
func (c *Context) DetachShader(program *Program, shader *Shader) {
	c.Call("detachShader", program.Value, shader.Value)
}

// Turns off specific WebGL capabilities for this context.
func (c *Context) Disable(cap int) {
	c.Call("disable", cap)
}

// Turns off a vertex attribute array at a specific index position.
func (c *Context) DisableVertexAttribArray(index int) {
	c.Call("disableVertexAttribArray", index)
}

// Render geometric primitives from bound and enabled vertex data.
func (c *Context) DrawArrays(mode, first, count int) {
	c.Call("drawArrays", mode, first, count)
}

// Renders geometric primitives indexed by element array data.
func (c *Context) DrawElements(mode, count, typ, offset int) {
	c.Call("drawElements", mode, count, typ, offset)
}

// Turns on specific WebGL capabilities for this context.
func (c *Context) Enable(cap int) {
	if cap == 0 {
		return
	}
	c.Call("enable", cap)
}

// Turns on a vertex attribute at a specific index position in
// a vertex attribute array.
func (c *Context) EnableVertexAttribArray(index int) {
	c.Call("enableVertexAttribArray", index)
}

func (c *Context) Finish() {
	c.Call("finish")
}

func (c *Context) Flush() {
	c.Call("flush")
}

// Attaches a WebGLRenderbuffer object as a logical buffer to the
// currently bound WebGLFramebuffer object.
func (c *Context) FrameBufferRenderBuffer(target, attachment, renderbufferTarget int, renderbuffer *RenderBuffer) {
	c.Call("framebufferRenderBuffer", target, attachment, renderbufferTarget, renderbuffer)
}

// Attaches a texture to a WebGLFramebuffer object.
func (c *Context) FramebufferTexture2D(target, attachment, textarget int, texture *Texture, level int) {
	c.Call("framebufferTexture2D", target, attachment, textarget, texture, level)
}

// Sets whether or not polygons are considered front-facing based
// on their winding direction.
func (c *Context) FrontFace(mode int) {
	c.Call("frontFace", mode)
}

// Creates a set of textures for a WebGLTexture object with image
// dimensions from the original size of the image down to a 1x1 image.
func (c *Context) GenerateMipmap(target int) {
	c.Call("generateMipmap", target)
}

// Returns an WebGLActiveInfo object containing the size, type, and name
// of a vertex attribute at a specific index position in a program object.
func (c *Context) GetActiveAttrib(program *Program, index int) string {
	return c.Call("getActiveAttrib", program.Value, index).String()
}

// Returns an WebGLActiveInfo object containing the size, type, and name
// of a uniform attribute at a specific index position in a program object.
func (c *Context) GetActiveUniform(program *Program, index int) string {
	return c.Call("getActiveUniform", program.Value, index).String()
}

// Returns a slice of WebGLShaders bound to a WebGLProgram.
func (c *Context) GetAttachedShaders(program *Program) []*Shader {
	objs := c.Call("getAttachedShaders", program.Value)
	shaders := make([]*Shader, objs.Length())
	for i := 0; i < objs.Length(); i++ {
		shaders[i] = &Shader{objs.Index(i)}
	}
	return shaders
}

// Returns an index to the location in a program of a named attribute variable.
func (c *Context) GetAttribLocation(program *Program, name string) int {
	return c.Call("getAttribLocation", program.Value, name).Int()
}

// Returns the type of a parameter for a given buffer.
func (c *Context) GetBufferParameter(target, pname int) int {
	return c.Call("getBufferParameter", target, pname).Int()
}

// TODO: Create type specific variations.
// Returns the natural type value for a constant parameter.
func (c *Context) GetParameter(pname int) js.Value {
	return c.Call("getParameter", pname)
}

// Returns a value for the WebGL error flag and clears the flag.
func (c *Context) GetError() int {
	return c.Call("getError").Int()
}

// TODO: Create type specific variations.
// Enables a passed extension, otherwise returns null.
func (c *Context) GetExtension(name string) js.Value {
	return c.Call("getExtension", name)
}

// TODO: Create type specific variations.
// Gets a parameter value for a given target and attachment.
func (c *Context) GetFramebufferAttachmentParameter(target, attachment, pname int) js.Value {
	return c.Call("getFramebufferAttachmentParameter", target, attachment, pname)
}

// Returns the value of the program parameter that corresponds to a supplied pname
// which is interpreted as an int.
func (c *Context) GetProgramParameteri(program *Program, pname int) int {
	return c.Call("getProgramParameter", program.Value, pname).Int()
}

// Returns the value of the program parameter that corresponds to a supplied pname
// which is interpreted as a bool.
func (c *Context) GetProgramParameterb(program *Program, pname int) bool {
	return c.Call("getProgramParameter", program.Value, pname).Bool()
}

// Returns information about the last error that occurred during
// the failed linking or validation of a WebGL program object.
func (c *Context) GetProgramInfoLog(program *Program) string {
	return c.Call("getProgramInfoLog", program.Value).String()
}

// TODO: Create type specific variations.
// Returns a renderbuffer parameter from the currently bound WebGLRenderbuffer object.
func (c *Context) GetRenderbufferParameter(target, pname int) int {
	return c.Call("getRenderbufferParameter", target, pname).Int()
}

// TODO: Create type specific variations.
// Returns the value of the parameter associated with pname for a shader object.
func (c *Context) GetShaderParameter(shader *Shader, pname int) js.Value {
	return c.Call("getShaderParameter", shader.Value, pname)
}

// Returns the value of the parameter associated with pname for a shader object.
func (c *Context) GetShaderParameterb(shader *Shader, pname int) bool {
	return c.Call("getShaderParameter", shader.Value, pname).Bool()
}

// Returns errors which occur when compiling a shader.
func (c *Context) GetShaderInfoLog(shader *Shader) string {
	return c.Call("getShaderInfoLog", shader.Value).String()
}

// Returns source code string associated with a shader object.
func (c *Context) GetShaderSource(shader *Shader) string {
	return c.Call("getShaderSource", shader.Value).String()
}

// Returns a parameter from a shader object
func (c *Context) GetShaderiv(shader *Shader, pname uint32) bool {
	return c.Call("getShaderParameter", shader.Value, pname).Bool()
}

// Returns a slice of supported extension strings.
func (c *Context) GetSupportedExtensions() []string {
	ext := c.Call("getSupportedExtensions")
	extensions := make([]string, ext.Length())
	for i := 0; i < ext.Length(); i++ {
		extensions[i] = ext.Index(i).String()
	}
	return extensions
}

// TODO: Create type specific variations.
// Returns the value for a parameter on an active texture unit.
func (c *Context) GetTexParameter(target, pname int) js.Value {
	return c.Call("getTexParameter", target, pname)
}

// TODO: Create type specific variations.
// Gets the uniform value for a specific location in a program.
func (c *Context) GetUniform(program *Program, location *UniformLocation) js.Value {
	return c.Call("getUniform", program.Value, location.Value)
}

// Returns a WebGLUniformLocation object for the location
// of a uniform variable within a WebGLProgram object.
func (c *Context) GetUniformLocation(program *Program, name string) *UniformLocation {
	return &UniformLocation{c.Call("getUniformLocation", program.Value, name)}
}

// TODO: Create type specific variations.
// Returns data for a particular characteristic of a vertex
// attribute at an index in a vertex attribute array.
func (c *Context) GetVertexAttrib(index, pname int) js.Value {
	return c.Call("getVertexAttrib", index, pname)
}

// Returns the address of a specified vertex attribute.
func (c *Context) GetVertexAttribOffset(index, pname int) int {
	return c.Call("getVertexAttribOffset", index, pname).Int()
}

// public function hint(target:GLenum, mode:GLenum) : Void;

// Returns true if buffer is valid, false otherwise.
func (c *Context) IsBuffer(buffer js.Value) bool {
	return c.Call("isBuffer", buffer).Bool()
}

// Returns whether the WebGL context has been lost.
func (c *Context) IsContextLost() bool {
	return c.Call("isContextLost").Bool()
}

// Returns true if buffer is valid, false otherwise.
func (c *Context) IsFramebuffer(framebuffer *FrameBuffer) bool {
	return c.Call("isFramebuffer", framebuffer.Value).Bool()
}

// Returns true if program object is valid, false otherwise.
func (c *Context) IsProgram(program *Program) bool {
	return c.Call("isProgram", program.Value).Bool()
}

// Returns true if buffer is valid, false otherwise.
func (c *Context) IsRenderbuffer(renderbuffer *RenderBuffer) bool {
	return c.Call("isRenderbuffer", renderbuffer.Value).Bool()
}

// Returns true if shader is valid, false otherwise.
func (c *Context) IsShader(shader *Shader) bool {
	return c.Call("isShader", shader.Value).Bool()
}

// Returns true if texture is valid, false otherwise.
func (c *Context) IsTexture(texture *Texture) bool {
	return c.Call("isTexture", texture.Value).Bool()
}

// Returns whether or not a WebGL capability is enabled for this context.
func (c *Context) IsEnabled(capability int) bool {
	return c.Call("isEnabled", capability).Bool()
}

// Sets the width of lines in WebGL.
func (c *Context) LineWidth(width float32) {
	c.Call("lineWidth", float64(width))
}

// Links an attached vertex shader and an attached fragment shader
// to a program so it can be used by the graphics processing unit (GPU).
func (c *Context) LinkProgram(program *Program) {
	c.Call("linkProgram", program.Value)

	param := c.Call("getProgramParameter", program.Value, c.LINK_STATUS)
	if !param.Bool() {
		info := c.Call("getProgramInfoLog", program.Value)
		log.Print("Error: could not link program: ", info.String())
	}
}

// Sets pixel storage modes for readPixels and unpacking of textures
// with texImage2D and texSubImage2D.
func (c *Context) PixelStorei(pname, param int) {
	c.Call("pixelStorei", pname, param)
}

// Sets the implementation-specific units and scale factor
// used to calculate fragment depth values.
func (c *Context) PolygonOffset(factor, units float64) {
	c.Call("polygonOffset", factor, units)
}

// TODO: Figure out if pixels should be a slice.
// Reads pixel data into an ArrayBufferView object from a
// rectangular area in the color buffer of the active frame buffer.
func (c *Context) ReadPixels(x, y, width, height, format, typ int, pixels js.Value) {
	c.Call("readPixels", x, y, width, height, format, typ, pixels)
}

// Creates or replaces the data store for the currently bound WebGLRenderbuffer object.
func (c *Context) RenderbufferStorage(target, internalFormat, width, height int) {
	c.Call("renderbufferStorage", target, internalFormat, width, height)
}

// Sets the dimensions of the scissor box.
func (c *Context) Scissor(x, y, width, height int) {
	c.Call("scissor", x, y, width, height)
}

// Sets and replaces shader source code in a shader object.
func (c *Context) ShaderSource(shader *Shader, source string) {
	c.Call("shaderSource", shader.Value, source)
}

// public function stencilMask(mask:GLuint) : Void;
// public function stencilMaskSeparate(face:GLenum, mask:GLuint) : Void;

// Loads the supplied pixel data into a texture.
func (c *Context) TexImage2D(target, level, internalFormat, format, kind int, data interface{}) {
	switch img := data.(type) {
	case *image.NRGBA:
		d := js.TypedArrayOf(img.Pix)
		c.Call("texImage2D", target, level, internalFormat, img.Bounds().Dx(), img.Bounds().Dy(), 0, format, kind, d)
		d.Release()
	case *image.RGBA:
		d := js.TypedArrayOf(img.Pix)
		c.Call("texImage2D", target, level, internalFormat, img.Bounds().Dx(), img.Bounds().Dy(), 0, format, kind, d)
		d.Release()
	default:
		c.Call("texImage2D", target, level, internalFormat, format, kind, data)
	}
}

// Sets texture parameters for the current texture unit.
func (c *Context) TexParameteri(target int, pname int, param int) {
	c.Call("texParameteri", target, pname, param)
}

// Replaces a portion of an existing 2D texture image with all of another image.
func (c *Context) TexSubImage2D(target, level, xoffset, yoffset, format, typ int, image js.Value) {
	c.Call("texSubImage2D", target, level, xoffset, yoffset, format, typ, image)
}

// Assigns a floating point value to a uniform variable for the current program object.
func (c *Context) Uniform1f(location *UniformLocation, x float32) {
	c.Call("uniform1f", location.Value, x)
}

// Assigns a integer value to a uniform variable for the current program object.
func (c *Context) Uniform1i(location *UniformLocation, x int) {
	c.Call("uniform1i", location.Value, x)
}

// Assigns 2 floating point values to a uniform variable for the current program object.
func (c *Context) Uniform2f(location *UniformLocation, x, y float32) {
	c.Call("uniform2f", location.Value, x, y)
}

// Assigns 2 integer values to a uniform variable for the current program object.
func (c *Context) Uniform2i(location *UniformLocation, x, y int) {
	c.Call("uniform2i", location.Value, x, y)
}

// Assigns 3 floating point values to a uniform variable for the current program object.
func (c *Context) Uniform3f(location *UniformLocation, x, y, z float32) {
	c.Call("uniform3f", location.Value, x, y, z)
}

// Assigns 3 integer values to a uniform variable for the current program object.
func (c *Context) Uniform3i(location *UniformLocation, x, y, z int) {
	c.Call("uniform3i", location.Value, x, y, z)
}

// Assigns 4 floating point values to a uniform variable for the current program object.
func (c *Context) Uniform4f(location *UniformLocation, x, y, z, w float32) {
	c.Call("uniform4f", location.Value, x, y, z, w)
}

// Assigns 4 integer values to a uniform variable for the current program object.
func (c *Context) Uniform4i(location *UniformLocation, x, y, z, w int) {
	c.Call("uniform4i", location.Value, x, y, z, w)
}

// public function uniform1fv(location:WebGLUniformLocation, v:ArrayAccess<Float>) : Void;
// public function uniform1iv(location:WebGLUniformLocation, v:ArrayAccess<Long>) : Void;
// public function uniform2fv(location:WebGLUniformLocation, v:ArrayAccess<Float>) : Void;
// public function uniform2iv(location:WebGLUniformLocation, v:ArrayAccess<Long>) : Void;
// public function uniform3fv(location:WebGLUniformLocation, v:ArrayAccess<Float>) : Void;
// public function uniform3iv(location:WebGLUniformLocation, v:ArrayAccess<Long>) : Void;
// public function uniform4fv(location:WebGLUniformLocation, v:ArrayAccess<Float>) : Void;
// public function uniform4iv(location:WebGLUniformLocation, v:ArrayAccess<Long>) : Void;

// Sets values for a 2x2 floating point vector matrix into a
// uniform location as a matrix or a matrix array.
func (c *Context) UniformMatrix2fv(location *UniformLocation, transpose bool, value []float32) {
	d := js.TypedArrayOf(value)
	c.Call("uniformMatrix2fv", location.Value, transpose, d)
	d.Release()
}

// Sets values for a 3x3 floating point vector matrix into a
// uniform location as a matrix or a matrix array.
func (c *Context) UniformMatrix3fv(location *UniformLocation, transpose bool, value []float32) {
	d := js.TypedArrayOf(value)
	c.Call("uniformMatrix3fv", location.Value, transpose, d)
}

// Sets values for a 4x4 floating point vector matrix into a
// uniform location as a matrix or a matrix array.
func (c *Context) UniformMatrix4fv(location *UniformLocation, transpose bool, value []float32) {
	d := js.TypedArrayOf(value)
	c.Call("uniformMatrix4fv", location.Value, transpose, d)
}

// Set the program object to use for rendering.
func (c *Context) UseProgram(program *Program) {
	if program == nil {
		c.Call("useProgram", nil)
		return
	}
	c.Call("useProgram", program.Value)
}

// Returns whether a given program can run in the current WebGL state.
func (c *Context) ValidateProgram(program *Program) {
	c.Call("validateProgram", program.Value)
}

func (c *Context) VertexAttribPointer(index, size, typ int, normal bool, stride int, offset int) {
	c.Call("vertexAttribPointer", index, size, typ, normal, stride, offset)
}

// public function vertexAttrib1f(indx:GLuint, x:GLfloat) : Void;
// public function vertexAttrib2f(indx:GLuint, x:GLfloat, y:GLfloat) : Void;
// public function vertexAttrib3f(indx:GLuint, x:GLfloat, y:GLfloat, z:GLfloat) : Void;
// public function vertexAttrib4f(indx:GLuint, x:GLfloat, y:GLfloat, z:GLfloat, w:GLfloat) : Void;
// public function vertexAttrib1fv(indx:GLuint, values:ArrayAccess<Float>) : Void;
// public function vertexAttrib2fv(indx:GLuint, values:ArrayAccess<Float>) : Void;
// public function vertexAttrib3fv(indx:GLuint, values:ArrayAccess<Float>) : Void;
// public function vertexAttrib4fv(indx:GLuint, values:ArrayAccess<Float>) : Void;

// Represents a rectangular viewable area that contains
// the rendering results of the drawing buffer.
func (c *Context) Viewport(x, y, width, height int) {
	c.Call("viewport", x, y, width, height)
}
