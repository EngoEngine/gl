// Copyright 2014 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !netgo,!android

package webgl

import (
	"fmt"
	"image"
	"log"
	"reflect"
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
)

type Texture struct{ uint32 }
type Buffer struct{ uint32 }
type FrameBuffer struct{ uint32 }
type RenderBuffer struct{ uint32 }
type Program struct{ uint32 }
type UniformLocation struct{ int32 }
type Shader struct{ uint32 }

type Context struct {
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
	INFO_LOG_LENGTH                              uint32
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
	NEAREST                                      int
	NEAREST_MIPMAP_LINEAR                        int
	NEAREST_MIPMAP_NEAREST                       int
	NEVER                                        int
	NICEST                                       int
	NONE                                         int
	NOTEQUAL                                     int
	NO_ERROR                                     int
	NUM_COMPRESSED_TEXTURE_FORMATS               int
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
	SHADER_COMPILER                              int
	SHADER_SOURCE_LENGTH                         int
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
	VENDOR                                       int
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

func NewContext() *Context {
	if err := gl.Init(); err != nil {
		log.Fatal(err)
	}
	return &Context{
		ARRAY_BUFFER:               gl.ARRAY_BUFFER,
		ARRAY_BUFFER_BINDING:       gl.ARRAY_BUFFER_BINDING,
		ATTACHED_SHADERS:           gl.ATTACHED_SHADERS,
		BACK:                       gl.BACK,
		BLEND:                      gl.BLEND,
		BLEND_COLOR:                gl.BLEND_COLOR,
		BLEND_DST_ALPHA:            gl.BLEND_DST_ALPHA,
		BLEND_DST_RGB:              gl.BLEND_DST_RGB,
		BLEND_EQUATION:             gl.BLEND_EQUATION,
		BLEND_EQUATION_ALPHA:       gl.BLEND_EQUATION_ALPHA,
		BLEND_EQUATION_RGB:         gl.BLEND_EQUATION_RGB,
		BLEND_SRC_ALPHA:            gl.BLEND_SRC_ALPHA,
		BLEND_SRC_RGB:              gl.BLEND_SRC_RGB,
		BLUE_BITS:                  gl.BLUE_BITS,
		BOOL:                       gl.BOOL,
		BOOL_VEC2:                  gl.BOOL_VEC2,
		BOOL_VEC3:                  gl.BOOL_VEC3,
		BOOL_VEC4:                  gl.BOOL_VEC4,
		BUFFER_SIZE:                gl.BUFFER_SIZE,
		BUFFER_USAGE:               gl.BUFFER_USAGE,
		BYTE:                       gl.BYTE,
		CCW:                        gl.CCW,
		CLAMP_TO_EDGE:              gl.CLAMP_TO_EDGE,
		COLOR_ATTACHMENT0:          gl.COLOR_ATTACHMENT0,
		COLOR_BUFFER_BIT:           gl.COLOR_BUFFER_BIT,
		COLOR_CLEAR_VALUE:          gl.COLOR_CLEAR_VALUE,
		COLOR_WRITEMASK:            gl.COLOR_WRITEMASK,
		COMPILE_STATUS:             gl.COMPILE_STATUS,
		COMPRESSED_TEXTURE_FORMATS: gl.COMPRESSED_TEXTURE_FORMATS,
		CONSTANT_ALPHA:             gl.CONSTANT_ALPHA,
		CONSTANT_COLOR:             gl.CONSTANT_COLOR,
		CULL_FACE:                  gl.CULL_FACE,
		CULL_FACE_MODE:             gl.CULL_FACE_MODE,
		CURRENT_PROGRAM:            gl.CURRENT_PROGRAM,
		CURRENT_VERTEX_ATTRIB:      gl.CURRENT_VERTEX_ATTRIB,
		CW:                           gl.CW,
		DECR:                         gl.DECR,
		DECR_WRAP:                    gl.DECR_WRAP,
		DELETE_STATUS:                gl.DELETE_STATUS,
		DEPTH_ATTACHMENT:             gl.DEPTH_ATTACHMENT,
		DEPTH_BITS:                   gl.DEPTH_BITS,
		DEPTH_BUFFER_BIT:             gl.DEPTH_BUFFER_BIT,
		DEPTH_CLEAR_VALUE:            gl.DEPTH_CLEAR_VALUE,
		DEPTH_COMPONENT:              gl.DEPTH_COMPONENT,
		DEPTH_COMPONENT16:            gl.DEPTH_COMPONENT16,
		DEPTH_FUNC:                   gl.DEPTH_FUNC,
		DEPTH_RANGE:                  gl.DEPTH_RANGE,
		DEPTH_STENCIL:                gl.DEPTH_STENCIL,
		DEPTH_STENCIL_ATTACHMENT:     gl.DEPTH_STENCIL_ATTACHMENT,
		DEPTH_TEST:                   gl.DEPTH_TEST,
		DEPTH_WRITEMASK:              gl.DEPTH_WRITEMASK,
		DITHER:                       gl.DITHER,
		DONT_CARE:                    gl.DONT_CARE,
		DST_ALPHA:                    gl.DST_ALPHA,
		DST_COLOR:                    gl.DST_COLOR,
		DYNAMIC_DRAW:                 gl.DYNAMIC_DRAW,
		ELEMENT_ARRAY_BUFFER:         gl.ELEMENT_ARRAY_BUFFER,
		ELEMENT_ARRAY_BUFFER_BINDING: gl.ELEMENT_ARRAY_BUFFER_BINDING,
		EQUAL:                                        gl.EQUAL,
		FASTEST:                                      gl.FASTEST,
		FLOAT:                                        gl.FLOAT,
		FLOAT_MAT2:                                   gl.FLOAT_MAT2,
		FLOAT_MAT3:                                   gl.FLOAT_MAT3,
		FLOAT_MAT4:                                   gl.FLOAT_MAT4,
		FLOAT_VEC2:                                   gl.FLOAT_VEC2,
		FLOAT_VEC3:                                   gl.FLOAT_VEC3,
		FLOAT_VEC4:                                   gl.FLOAT_VEC4,
		FRAGMENT_SHADER:                              gl.FRAGMENT_SHADER,
		FRAMEBUFFER:                                  gl.FRAMEBUFFER,
		FRAMEBUFFER_ATTACHMENT_OBJECT_NAME:           gl.FRAMEBUFFER_ATTACHMENT_OBJECT_NAME,
		FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE:           gl.FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE,
		FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE: gl.FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE,
		FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL:         gl.FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL,
		FRAMEBUFFER_BINDING:                          gl.FRAMEBUFFER_BINDING,
		FRAMEBUFFER_COMPLETE:                         gl.FRAMEBUFFER_COMPLETE,
		FRAMEBUFFER_INCOMPLETE_ATTACHMENT:            gl.FRAMEBUFFER_INCOMPLETE_ATTACHMENT,
		FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT:    gl.FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT,
		FRAMEBUFFER_UNSUPPORTED:                      gl.FRAMEBUFFER_UNSUPPORTED,
		FRONT:                         gl.FRONT,
		FRONT_AND_BACK:                gl.FRONT_AND_BACK,
		FRONT_FACE:                    gl.FRONT_FACE,
		FUNC_ADD:                      gl.FUNC_ADD,
		FUNC_REVERSE_SUBTRACT:         gl.FUNC_REVERSE_SUBTRACT,
		FUNC_SUBTRACT:                 gl.FUNC_SUBTRACT,
		GENERATE_MIPMAP_HINT:          gl.GENERATE_MIPMAP_HINT,
		GEQUAL:                        gl.GEQUAL,
		GREATER:                       gl.GREATER,
		GREEN_BITS:                    gl.GREEN_BITS,
		HIGH_FLOAT:                    gl.HIGH_FLOAT,
		HIGH_INT:                      gl.HIGH_INT,
		INCR:                          gl.INCR,
		INCR_WRAP:                     gl.INCR_WRAP,
		INFO_LOG_LENGTH:               gl.INFO_LOG_LENGTH,
		INT:                           gl.INT,
		INT_VEC2:                      gl.INT_VEC2,
		INT_VEC3:                      gl.INT_VEC3,
		INT_VEC4:                      gl.INT_VEC4,
		INVALID_ENUM:                  gl.INVALID_ENUM,
		INVALID_FRAMEBUFFER_OPERATION: gl.INVALID_FRAMEBUFFER_OPERATION,
		INVALID_OPERATION:             gl.INVALID_OPERATION,
		INVALID_VALUE:                 gl.INVALID_VALUE,
		INVERT:                        gl.INVERT,
		KEEP:                          gl.KEEP,
		LEQUAL:                        gl.LEQUAL,
		LESS:                          gl.LESS,
		LINEAR:                        gl.LINEAR,
		LINEAR_MIPMAP_LINEAR:          gl.LINEAR_MIPMAP_LINEAR,
		LINEAR_MIPMAP_NEAREST:         gl.LINEAR_MIPMAP_NEAREST,
		LINES:                            gl.LINES,
		LINE_LOOP:                        gl.LINE_LOOP,
		LINE_STRIP:                       gl.LINE_STRIP,
		LINE_WIDTH:                       gl.LINE_WIDTH,
		LINK_STATUS:                      gl.LINK_STATUS,
		LOW_FLOAT:                        gl.LOW_FLOAT,
		LOW_INT:                          gl.LOW_INT,
		LUMINANCE:                        gl.LUMINANCE,
		LUMINANCE_ALPHA:                  gl.LUMINANCE_ALPHA,
		MAX_COMBINED_TEXTURE_IMAGE_UNITS: gl.MAX_COMBINED_TEXTURE_IMAGE_UNITS,
		MAX_CUBE_MAP_TEXTURE_SIZE:        gl.MAX_CUBE_MAP_TEXTURE_SIZE,
		MAX_FRAGMENT_UNIFORM_VECTORS:     gl.MAX_FRAGMENT_UNIFORM_VECTORS,
		MAX_RENDERBUFFER_SIZE:            gl.MAX_RENDERBUFFER_SIZE,
		MAX_TEXTURE_IMAGE_UNITS:          gl.MAX_TEXTURE_IMAGE_UNITS,
		MAX_TEXTURE_SIZE:                 gl.MAX_TEXTURE_SIZE,
		MAX_VARYING_VECTORS:              gl.MAX_VARYING_VECTORS,
		MAX_VERTEX_ATTRIBS:               gl.MAX_VERTEX_ATTRIBS,
		MAX_VERTEX_TEXTURE_IMAGE_UNITS:   gl.MAX_VERTEX_TEXTURE_IMAGE_UNITS,
		MAX_VERTEX_UNIFORM_VECTORS:       gl.MAX_VERTEX_UNIFORM_VECTORS,
		MAX_VIEWPORT_DIMS:                gl.MAX_VIEWPORT_DIMS,
		MEDIUM_FLOAT:                     gl.MEDIUM_FLOAT,
		MEDIUM_INT:                       gl.MEDIUM_INT,
		MIRRORED_REPEAT:                  gl.MIRRORED_REPEAT,
		NEAREST:                          gl.NEAREST,
		NEAREST_MIPMAP_LINEAR:            gl.NEAREST_MIPMAP_LINEAR,
		NEAREST_MIPMAP_NEAREST:           gl.NEAREST_MIPMAP_NEAREST,
		NEVER:                          gl.NEVER,
		NICEST:                         gl.NICEST,
		NONE:                           gl.NONE,
		NOTEQUAL:                       gl.NOTEQUAL,
		NO_ERROR:                       gl.NO_ERROR,
		NUM_COMPRESSED_TEXTURE_FORMATS: gl.NUM_COMPRESSED_TEXTURE_FORMATS,
		ONE: gl.ONE,
		ONE_MINUS_CONSTANT_ALPHA:     gl.ONE_MINUS_CONSTANT_ALPHA,
		ONE_MINUS_CONSTANT_COLOR:     gl.ONE_MINUS_CONSTANT_COLOR,
		ONE_MINUS_DST_ALPHA:          gl.ONE_MINUS_DST_ALPHA,
		ONE_MINUS_DST_COLOR:          gl.ONE_MINUS_DST_COLOR,
		ONE_MINUS_SRC_ALPHA:          gl.ONE_MINUS_SRC_ALPHA,
		ONE_MINUS_SRC_COLOR:          gl.ONE_MINUS_SRC_COLOR,
		OUT_OF_MEMORY:                gl.OUT_OF_MEMORY,
		PACK_ALIGNMENT:               gl.PACK_ALIGNMENT,
		POINTS:                       gl.POINTS,
		POLYGON_OFFSET_FACTOR:        gl.POLYGON_OFFSET_FACTOR,
		POLYGON_OFFSET_FILL:          gl.POLYGON_OFFSET_FILL,
		POLYGON_OFFSET_UNITS:         gl.POLYGON_OFFSET_UNITS,
		RED_BITS:                     gl.RED_BITS,
		RENDERBUFFER:                 gl.RENDERBUFFER,
		RENDERBUFFER_ALPHA_SIZE:      gl.RENDERBUFFER_ALPHA_SIZE,
		RENDERBUFFER_BINDING:         gl.RENDERBUFFER_BINDING,
		RENDERBUFFER_BLUE_SIZE:       gl.RENDERBUFFER_BLUE_SIZE,
		RENDERBUFFER_DEPTH_SIZE:      gl.RENDERBUFFER_DEPTH_SIZE,
		RENDERBUFFER_GREEN_SIZE:      gl.RENDERBUFFER_GREEN_SIZE,
		RENDERBUFFER_HEIGHT:          gl.RENDERBUFFER_HEIGHT,
		RENDERBUFFER_INTERNAL_FORMAT: gl.RENDERBUFFER_INTERNAL_FORMAT,
		RENDERBUFFER_RED_SIZE:        gl.RENDERBUFFER_RED_SIZE,
		RENDERBUFFER_STENCIL_SIZE:    gl.RENDERBUFFER_STENCIL_SIZE,
		RENDERBUFFER_WIDTH:           gl.RENDERBUFFER_WIDTH,
		RENDERER:                     gl.RENDERER,
		REPEAT:                       gl.REPEAT,
		REPLACE:                      gl.REPLACE,
		RGB:                          gl.RGB,
		RGB5_A1:                      gl.RGB5_A1,
		RGB565:                       gl.RGB565,
		RGBA:                         gl.RGBA,
		RGBA4:                        gl.RGBA4,
		SAMPLER_2D:                   gl.SAMPLER_2D,
		SAMPLER_CUBE:                 gl.SAMPLER_CUBE,
		SAMPLES:                      gl.SAMPLES,
		SAMPLE_ALPHA_TO_COVERAGE:     gl.SAMPLE_ALPHA_TO_COVERAGE,
		SAMPLE_BUFFERS:               gl.SAMPLE_BUFFERS,
		SAMPLE_COVERAGE:              gl.SAMPLE_COVERAGE,
		SAMPLE_COVERAGE_INVERT:       gl.SAMPLE_COVERAGE_INVERT,
		SAMPLE_COVERAGE_VALUE:        gl.SAMPLE_COVERAGE_VALUE,
		SCISSOR_BOX:                  gl.SCISSOR_BOX,
		SCISSOR_TEST:                 gl.SCISSOR_TEST,
		SHADER_COMPILER:              gl.SHADER_COMPILER,
		SHADER_SOURCE_LENGTH:         gl.SHADER_SOURCE_LENGTH,
		SHADER_TYPE:                  gl.SHADER_TYPE,
		SHADING_LANGUAGE_VERSION:     gl.SHADING_LANGUAGE_VERSION,
		SHORT:                        gl.SHORT,
		SRC_ALPHA:                    gl.SRC_ALPHA,
		SRC_ALPHA_SATURATE:           gl.SRC_ALPHA_SATURATE,
		SRC_COLOR:                    gl.SRC_COLOR,
		STATIC_DRAW:                  gl.STATIC_DRAW,
		STENCIL_ATTACHMENT:           gl.STENCIL_ATTACHMENT,
		STENCIL_BACK_FAIL:            gl.STENCIL_BACK_FAIL,
		STENCIL_BACK_FUNC:            gl.STENCIL_BACK_FUNC,
		STENCIL_BACK_PASS_DEPTH_FAIL: gl.STENCIL_BACK_PASS_DEPTH_FAIL,
		STENCIL_BACK_PASS_DEPTH_PASS: gl.STENCIL_BACK_PASS_DEPTH_PASS,
		STENCIL_BACK_REF:             gl.STENCIL_BACK_REF,
		STENCIL_BACK_VALUE_MASK:      gl.STENCIL_BACK_VALUE_MASK,
		STENCIL_BACK_WRITEMASK:       gl.STENCIL_BACK_WRITEMASK,
		STENCIL_BITS:                 gl.STENCIL_BITS,
		STENCIL_BUFFER_BIT:           gl.STENCIL_BUFFER_BIT,
		STENCIL_CLEAR_VALUE:          gl.STENCIL_CLEAR_VALUE,
		STENCIL_FAIL:                 gl.STENCIL_FAIL,
		STENCIL_FUNC:                 gl.STENCIL_FUNC,
		STENCIL_INDEX:                gl.STENCIL_INDEX,
		STENCIL_INDEX8:               gl.STENCIL_INDEX8,
		STENCIL_PASS_DEPTH_FAIL:      gl.STENCIL_PASS_DEPTH_FAIL,
		STENCIL_PASS_DEPTH_PASS:      gl.STENCIL_PASS_DEPTH_PASS,
		STENCIL_REF:                  gl.STENCIL_REF,
		STENCIL_TEST:                 gl.STENCIL_TEST,
		STENCIL_VALUE_MASK:           gl.STENCIL_VALUE_MASK,
		STENCIL_WRITEMASK:            gl.STENCIL_WRITEMASK,
		STREAM_DRAW:                  gl.STREAM_DRAW,
		SUBPIXEL_BITS:                gl.SUBPIXEL_BITS,
		TEXTURE:                      gl.TEXTURE,
		TEXTURE0:                     gl.TEXTURE0,
		TEXTURE1:                     gl.TEXTURE1,
		TEXTURE2:                     gl.TEXTURE2,
		TEXTURE3:                     gl.TEXTURE3,
		TEXTURE4:                     gl.TEXTURE4,
		TEXTURE5:                     gl.TEXTURE5,
		TEXTURE6:                     gl.TEXTURE6,
		TEXTURE7:                     gl.TEXTURE7,
		TEXTURE8:                     gl.TEXTURE8,
		TEXTURE9:                     gl.TEXTURE9,
		TEXTURE10:                    gl.TEXTURE10,
		TEXTURE11:                    gl.TEXTURE11,
		TEXTURE12:                    gl.TEXTURE12,
		TEXTURE13:                    gl.TEXTURE13,
		TEXTURE14:                    gl.TEXTURE14,
		TEXTURE15:                    gl.TEXTURE15,
		TEXTURE16:                    gl.TEXTURE16,
		TEXTURE17:                    gl.TEXTURE17,
		TEXTURE18:                    gl.TEXTURE18,
		TEXTURE19:                    gl.TEXTURE19,
		TEXTURE20:                    gl.TEXTURE20,
		TEXTURE21:                    gl.TEXTURE21,
		TEXTURE22:                    gl.TEXTURE22,
		TEXTURE23:                    gl.TEXTURE23,
		TEXTURE24:                    gl.TEXTURE24,
		TEXTURE25:                    gl.TEXTURE25,
		TEXTURE26:                    gl.TEXTURE26,
		TEXTURE27:                    gl.TEXTURE27,
		TEXTURE28:                    gl.TEXTURE28,
		TEXTURE29:                    gl.TEXTURE29,
		TEXTURE30:                    gl.TEXTURE30,
		TEXTURE31:                    gl.TEXTURE31,
		TEXTURE_2D:                   gl.TEXTURE_2D,
		TEXTURE_BINDING_2D:           gl.TEXTURE_BINDING_2D,
		TEXTURE_BINDING_CUBE_MAP:     gl.TEXTURE_BINDING_CUBE_MAP,
		TEXTURE_CUBE_MAP:             gl.TEXTURE_CUBE_MAP,
		TEXTURE_CUBE_MAP_NEGATIVE_X:  gl.TEXTURE_CUBE_MAP_NEGATIVE_X,
		TEXTURE_CUBE_MAP_NEGATIVE_Y:  gl.TEXTURE_CUBE_MAP_NEGATIVE_Y,
		TEXTURE_CUBE_MAP_NEGATIVE_Z:  gl.TEXTURE_CUBE_MAP_NEGATIVE_Z,
		TEXTURE_CUBE_MAP_POSITIVE_X:  gl.TEXTURE_CUBE_MAP_POSITIVE_X,
		TEXTURE_CUBE_MAP_POSITIVE_Y:  gl.TEXTURE_CUBE_MAP_POSITIVE_Y,
		TEXTURE_CUBE_MAP_POSITIVE_Z:  gl.TEXTURE_CUBE_MAP_POSITIVE_Z,
		TEXTURE_MAG_FILTER:           gl.TEXTURE_MAG_FILTER,
		TEXTURE_MIN_FILTER:           gl.TEXTURE_MIN_FILTER,
		TEXTURE_WRAP_S:               gl.TEXTURE_WRAP_S,
		TEXTURE_WRAP_T:               gl.TEXTURE_WRAP_T,
		TRIANGLES:                    gl.TRIANGLES,
		TRIANGLE_FAN:                 gl.TRIANGLE_FAN,
		TRIANGLE_STRIP:               gl.TRIANGLE_STRIP,
		UNPACK_ALIGNMENT:             gl.UNPACK_ALIGNMENT,
		UNSIGNED_BYTE:                gl.UNSIGNED_BYTE,
		UNSIGNED_INT:                 gl.UNSIGNED_INT,
		UNSIGNED_SHORT:               gl.UNSIGNED_SHORT,
		UNSIGNED_SHORT_4_4_4_4:       gl.UNSIGNED_SHORT_4_4_4_4,
		UNSIGNED_SHORT_5_5_5_1:       gl.UNSIGNED_SHORT_5_5_5_1,
		UNSIGNED_SHORT_5_6_5:         gl.UNSIGNED_SHORT_5_6_5,
		VALIDATE_STATUS:              gl.VALIDATE_STATUS,
		VENDOR:                       gl.VENDOR,
		VERSION:                      gl.VERSION,
		VERTEX_ATTRIB_ARRAY_BUFFER_BINDING: gl.VERTEX_ATTRIB_ARRAY_BUFFER_BINDING,
		VERTEX_ATTRIB_ARRAY_ENABLED:        gl.VERTEX_ATTRIB_ARRAY_ENABLED,
		VERTEX_ATTRIB_ARRAY_NORMALIZED:     gl.VERTEX_ATTRIB_ARRAY_NORMALIZED,
		VERTEX_ATTRIB_ARRAY_POINTER:        gl.VERTEX_ATTRIB_ARRAY_POINTER,
		VERTEX_ATTRIB_ARRAY_SIZE:           gl.VERTEX_ATTRIB_ARRAY_SIZE,
		VERTEX_ATTRIB_ARRAY_STRIDE:         gl.VERTEX_ATTRIB_ARRAY_STRIDE,
		VERTEX_ATTRIB_ARRAY_TYPE:           gl.VERTEX_ATTRIB_ARRAY_TYPE,
		VERTEX_SHADER:                      gl.VERTEX_SHADER,
		VIEWPORT:                           gl.VIEWPORT,
		ZERO:                               gl.ZERO,
		TRUE:                               gl.TRUE,
	}
}

func (c *Context) CreateShader(typ int) *Shader {
	shader := &Shader{gl.CreateShader(uint32(typ))}
	return shader
}

func (c *Context) ShaderSource(shader *Shader, source string) {
	glsource, free := gl.Strs(source + "\x00")
	gl.ShaderSource(shader.uint32, 1, glsource, nil)
	free()
}

func (c *Context) CompileShader(shader *Shader) {
	gl.CompileShader(shader.uint32)
}

// Ptr takes a slice or pointer (to a singular scalar value or the first
// element of an array or slice) and returns its GL-compatible address.
func (c *Context) Ptr(data interface{}) unsafe.Pointer {
	return gl.Ptr(data)
}

// Str takes a null-terminated Go string and returns its GL-compatible address.
// This function reaches into Go string storage in an unsafe way so the caller
// must ensure the string is not garbage collected.
func (c *Context) Str(str string) *uint8 {
	return gl.Str(str)
}

// GoStr takes a null-terminated string returned by OpenGL and constructs a
// corresponding Go string.
func (c *Context) GoStr(cstr *uint8) string {
	return gl.GoStr(cstr)
}

// DeleteShader will free the shader memory. You should call this in case of
// a compilation error to avoid leaking memory
func (c *Context) DeleteShader(shader *Shader) {
	gl.DeleteShader(shader.uint32)
}

// Returns a parameter from a shader object
func (c *Context) GetShaderiv(shader *Shader, pname uint32, params *int32) {
	gl.GetShaderiv(shader.uint32, pname, params)
}

// GetShaderInfoLog is a method you can call to get the compilation logs of a shader
// maxLength​ is the size of infoLog​;
// this tells OpenGL how many bytes at maximum it will write into infoLog​.
// length​ is a return value, specifying how many bytes it actually wrote into infoLog​;
// you may pass NULL if you don't care.
func (c *Context) GetShaderInfoLog(shader *Shader, bufSize int32, length *int32, infoLog *uint8) {
	gl.GetShaderInfoLog(shader.uint32, bufSize, length, infoLog)
}

func (c *Context) CreateProgram() *Program {
	return &Program{gl.CreateProgram()}
}

func (c *Context) AttachShader(program *Program, shader *Shader) {
	gl.AttachShader(program.uint32, shader.uint32)
}

func (c *Context) LinkProgram(program *Program) {
	gl.LinkProgram(program.uint32)
}

func (c *Context) CreateTexture() *Texture {
	var loc uint32
	gl.GenTextures(1, &loc)
	return &Texture{loc}
}

func (c *Context) BindTexture(target int, texture *Texture) {
	if texture == nil {
		gl.BindTexture(uint32(target), 0)
		return
	}
	gl.BindTexture(uint32(target), texture.uint32)
}

func (c *Context) TexParameteri(target int, pname int, param int) {
	gl.TexParameteri(uint32(target), uint32(pname), int32(param))
}

func (c *Context) TexImage2D(target, level, internalFormat, format, kind int, data interface{}) {
	var pix []uint8
	width := 0
	height := 0
	if data == nil {
		pix = nil
	} else {

		switch img := data.(type) {
		case *image.NRGBA:
			width = img.Bounds().Dx()
			height = img.Bounds().Dy()
			pix = img.Pix
		case *image.RGBA:
			width = img.Bounds().Dx()
			height = img.Bounds().Dy()
			pix = img.Pix
		default:
			panic(fmt.Errorf("Image type unsupported: %T", img))
		}
	}
	gl.TexImage2D(uint32(target), int32(level), int32(internalFormat), int32(width), int32(height), int32(0), uint32(format), uint32(kind), gl.Ptr(pix))
}

func (c *Context) GetAttribLocation(program *Program, name string) int {
	return int(gl.GetAttribLocation(program.uint32, gl.Str(name+"\x00")))
}

func (c *Context) GetUniformLocation(program *Program, name string) *UniformLocation {
	return &UniformLocation{gl.GetUniformLocation(program.uint32, gl.Str(name+"\x00"))}
}

func (c *Context) GetError() int {
	return int(gl.GetError())
}

func (c *Context) CreateBuffer() *Buffer {
	var loc uint32
	gl.GenBuffers(1, &loc)
	return &Buffer{loc}
}

func (c *Context) BindBuffer(target int, buffer *Buffer) {
	if buffer == nil {
		gl.BindBuffer(uint32(target), 0)
		return
	}
	gl.BindBuffer(uint32(target), buffer.uint32)
}

func (c *Context) BufferData(target int, data interface{}, usage int) {
	s := uintptr(reflect.ValueOf(data).Len()) * reflect.TypeOf(data).Elem().Size()
	gl.BufferData(uint32(target), int(s), gl.Ptr(data), uint32(usage))
}

func (c *Context) EnableVertexAttribArray(index int) {
	gl.EnableVertexAttribArray(uint32(index))
}

func (c *Context) VertexAttribPointer(index, size, typ int, normal bool, stride int, offset int) {
	gl.VertexAttribPointer(uint32(index), int32(size), uint32(typ), normal, int32(stride), gl.PtrOffset(offset))
}

func (c *Context) Enable(flag int) {
	gl.Enable(uint32(flag))
}

func (c *Context) BlendFunc(src, dst int) {
	gl.BlendFunc(uint32(src), uint32(dst))
}

func (c *Context) UniformMatrix2fv(location *UniformLocation, transpose bool, value []float32) {
	// TODO: count value of 1 is currently hardcoded.
	//       Perhaps it should be len(value) / 16 or something else?
	//       In OpenGL 2.1 it is a manually supplied parameter, but WebGL does not have it.
	//       Not sure if WebGL automatically deduces it and supports count values greater than 1, or if 1 is always assumed.
	gl.UniformMatrix2fv(location.int32, 1, transpose, &value[0])
}

func (c *Context) UniformMatrix4fv(location *UniformLocation, transpose bool, value []float32) {
	// TODO: count value of 1 is currently hardcoded.
	//       Perhaps it should be len(value) / 16 or something else?
	//       In OpenGL 2.1 it is a manually supplied parameter, but WebGL does not have it.
	//       Not sure if WebGL automatically deduces it and supports count values greater than 1, or if 1 is always assumed.
	gl.UniformMatrix4fv(location.int32, 1, transpose, &value[0])
}

func (c *Context) UseProgram(program *Program) {
	if program == nil {
		gl.UseProgram(0)
		return
	}
	gl.UseProgram(program.uint32)
}

func (c *Context) ValidateProgram(program *Program) {
	if program == nil {
		gl.ValidateProgram(0)
		return
	}
	gl.ValidateProgram(program.uint32)
}

// Specify the value of a uniform variable for the current program object
func (c *Context) Uniform1f(location *UniformLocation, x float32) {
	gl.Uniform1f(location.int32, x)
}

func (c *Context) Uniform2f(location *UniformLocation, x, y float32) {
	gl.Uniform2f(location.int32, x, y)
}

func (c *Context) Uniform3f(location *UniformLocation, x, y, z float32) {
	gl.Uniform3f(location.int32, x, y, z)
}

func (c *Context) BufferSubData(target int, offset int, data interface{}) {
	size := uintptr(reflect.ValueOf(data).Len()) * reflect.TypeOf(data).Elem().Size()
	gl.BufferSubData(uint32(target), offset, int(size), gl.Ptr(data))
}

func (c *Context) DrawArrays(mode, first, count int) {
	gl.DrawArrays(uint32(mode), int32(first), int32(count))
}

func (c *Context) DrawElements(mode, count, typ, offset int) {

	gl.DrawElements(uint32(mode), int32(count), uint32(typ), gl.PtrOffset(offset))
}

func (c *Context) ClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
}

func (c *Context) Viewport(x, y, width, height int) {
	gl.Viewport(int32(x), int32(y), int32(width), int32(height))
}

func (c *Context) Clear(flags int) {
	gl.Clear(uint32(flags))
}

func (c *Context) Translatef(x, y, z float32) {
	gl.Translatef(x, y, z)
}

func (c *Context) Rotatef(angle, x, y, z float32) {
	gl.Rotatef(angle, x, y, z)
}

func (c *Context) MatrixMode(mode uint32) {
	gl.MatrixMode(mode)
}

func (c *Context) LoadIdentity() {
	gl.LoadIdentity()
}

func (c *Context) PushMatrix() {
	gl.PushMatrix()
}

func (c *Context) PopMatrix() {
	gl.PopMatrix()
}

func (c *Context) Frustum(left, right, bottom, top, zNear, zFar float64) {
	gl.Frustum(left, right, bottom, top, zNear, zFar)
}
