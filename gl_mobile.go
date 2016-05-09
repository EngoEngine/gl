//+build android

package gl

import (
	"log"

	"golang.org/x/mobile/gl"
	"image"
	"reflect"
	"unsafe"
)

type Texture struct{ gl.Texture }
type Buffer struct{ gl.Buffer }
type FrameBuffer struct{ gl.Framebuffer }
type RenderBuffer struct{ gl.Renderbuffer }
type Program struct{ gl.Program }
type UniformLocation struct{ gl.Uniform }
type Shader struct{ gl.Shader }

type Context struct {
	ctx    gl.Context
	worker gl.Worker

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
	CLAMP_TO_BORDER                              int
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
	MULTISAMPLE                                  int
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

func NewContext(DrawContext interface{}) *Context {
	c := &Context{
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
	//c.Ctx, c.Worker = gl.NewContext()
	c.ctx = DrawContext.(gl.Context)

	return c
}

// The GL_BLEND_COLOR may be used to calculate the source and destination blending factors.
func (c *Context) BlendColor(r, g, b, a float32) {
	c.ctx.BlendColor(float32(r), float32(g), float32(b), float32(a))
}

// Sets the equation used to blend RGB and Alpha values of an incoming source
// fragment with a destination values as stored in the fragment's frame buffer.
func (c *Context) BlendEquation(mode int) {
	c.ctx.BlendEquation(gl.Enum(mode))
}

// Controls the blending of an incoming source fragment's R, G, B, and A values
// with a destination R, G, B, and A values as stored in the fragment's WebGLFramebuffer.
func (c *Context) BlendEquationSeparate(modeRGB, modeAlpha int) {
	c.ctx.BlendEquationSeparate(gl.Enum(modeRGB), gl.Enum(modeAlpha))
}

// Sets the blending factors used to combine source and destination pixels.
func (c *Context) BlendFunc(sfactor, dfactor int) {
	c.ctx.BlendFunc(gl.Enum(sfactor), gl.Enum(dfactor))
}

// Sets the weighting factors that are used by blendEquationSeparate.
func (c *Context) BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha int) {
	c.ctx.BlendFuncSeparate(gl.Enum(srcRGB), gl.Enum(dstRGB), gl.Enum(srcAlpha), gl.Enum(dstAlpha))
}

// Sets a function to use to compare incoming pixel depth to the
// current depth buffer value.
func (c *Context) DepthFunc(fun int) {
	c.ctx.DepthFunc(gl.Enum(fun))
}

func (c *Context) SampleCoverage(value float32, invert bool) {
	c.ctx.SampleCoverage(float32(value), invert)
}

func (c *Context) StencilFunc(function, ref, mask int) {
	c.ctx.StencilFunc(gl.Enum(function), ref, uint32(mask))
}

func (c *Context) StencilFuncSeparate(face, function, ref, mask int) {
	c.ctx.StencilFuncSeparate(gl.Enum(face), gl.Enum(function), ref, uint32(mask))
}

// public function stencilOp(fail:GLenum, zfail:GLenum, zpass:GLenum) : Void;
// public function stencilOpSeparate(face:GLenum, fail:GLenum, zfail:GLenum, zpass:GLenum) : Void;

// FrameBuffer

// ---------------------------------------------------------------------------

// Specifies the active texture unit.
func (c *Context) ActiveTexture(texture int) {
	c.ctx.ActiveTexture(gl.Enum(texture))
}

// Attaches a WebGLShader object to a WebGLProgram object.
func (c *Context) AttachShader(program *Program, shader *Shader) {
	c.ctx.AttachShader(program.Program, shader.Shader)
}

// Binds a generic vertex index to a user-defined attribute variable.
func (c *Context) BindAttribLocation(program *Program, index int, name string) {
	c.ctx.BindAttribLocation(program.Program, gl.Attrib{uint(index)}, name)
}

// Associates a buffer with a buffer target.
func (c *Context) BindBuffer(target int, buffer *Buffer) {
	if buffer == nil {
		c.ctx.BindBuffer(gl.Enum(target), gl.Buffer{0})
		return
	}
	c.ctx.BindBuffer(gl.Enum(target), buffer.Buffer)
}

// Associates a WebGLFramebuffer object with the FRAMEBUFFER bind target.
func (c *Context) BindFramebuffer(target int, framebuffer *FrameBuffer) {
	if framebuffer == nil {
		c.ctx.BindFramebuffer(gl.Enum(target), gl.Framebuffer{0})
		return
	}
	c.ctx.BindFramebuffer(gl.Enum(target), framebuffer.Framebuffer)
}

// Binds a WebGLRenderbuffer object to be used for rendering.
func (c *Context) BindRenderbuffer(target int, renderbuffer *RenderBuffer) {
	if renderbuffer == nil {
		c.ctx.BindRenderbuffer(gl.Enum(target), gl.Renderbuffer{0})
		return
	}
	c.ctx.BindRenderbuffer(gl.Enum(target), renderbuffer.Renderbuffer)
}

// Binds a named texture object to a target.
func (c *Context) BindTexture(target int, texture *Texture) {
	if texture == nil {
		c.ctx.BindTexture(gl.Enum(target), gl.Texture{0})
		return
	}
	c.ctx.BindTexture(gl.Enum(target), texture.Texture)
}

// Creates a buffer in memory and initializes it with array data.
// If no array is provided, the contents of the buffer is initialized to 0.
func (c *Context) BufferData(target int, data interface{}, usage int) {
	var b []byte
	switch arr := data.(type) {
	case []uint16:
		fheader := (*reflect.SliceHeader)(unsafe.Pointer(&arr))

		header := (*reflect.SliceHeader)(unsafe.Pointer(&b))
		header.Cap = 2 * fheader.Cap
		header.Len = 2 * fheader.Len
		header.Data = fheader.Data

	case []float32:
		fheader := (*reflect.SliceHeader)(unsafe.Pointer(&arr))

		header := (*reflect.SliceHeader)(unsafe.Pointer(&b))
		header.Cap = 4 * fheader.Cap
		header.Len = 4 * fheader.Len
		header.Data = fheader.Data
	}

	c.ctx.BufferData(gl.Enum(target), b, gl.Enum(usage))
}

// Used to modify or update some or all of a data store for a bound buffer object.
func (c *Context) BufferSubData(target int, offset int, data interface{}) {
	var b []byte
	switch arr := data.(type) {
	case []uint16:
		fheader := (*reflect.SliceHeader)(unsafe.Pointer(&arr))

		header := (*reflect.SliceHeader)(unsafe.Pointer(&b))
		header.Cap = 2 * fheader.Cap
		header.Len = 2 * fheader.Len
		header.Data = fheader.Data

	case []float32:
		fheader := (*reflect.SliceHeader)(unsafe.Pointer(&arr))

		header := (*reflect.SliceHeader)(unsafe.Pointer(&b))
		header.Cap = 4 * fheader.Cap
		header.Len = 4 * fheader.Len
		header.Data = fheader.Data
	}

	c.ctx.BufferSubData(gl.Enum(target), offset, b)
}

// Returns whether the currently bound WebGLFramebuffer is complete.
// If not complete, returns the reason why.
func (c *Context) CheckFramebufferStatus(target int) int {
	return int(c.ctx.CheckFramebufferStatus(gl.Enum(target)))
}

// Sets all pixels in a specific buffer to the same value.
func (c *Context) Clear(flags int) {
	c.ctx.Clear(gl.Enum(flags))
}

// Specifies color values to use by the clear method to clear the color buffer.
func (c *Context) ClearColor(r, g, b, a float32) {
	c.ctx.ClearColor(r, g, b, a)
}

// Clears the depth buffer to a specific value.
func (c *Context) ClearDepth(depth float32) {
	c.ctx.ClearDepthf(depth)
}

func (c *Context) ClearStencil(s int) {
	c.ctx.ClearStencil(s)
}

// Lets you set whether individual colors can be written when
// drawing or rendering to a framebuffer.
func (c *Context) ColorMask(r, g, b, a bool) {
	c.ctx.ColorMask(r, g, b, a)
}

// CompileShader compiles the GLSL shader source into binary data used by the WebGLProgram object.
func (c *Context) CompileShader(shader *Shader) {
	c.ctx.CompileShader(shader.Shader)
}

// Copies a rectangle of pixels from the current WebGLFramebuffer into a texture image.
func (c *Context) CopyTexImage2D(target, level, internal, x, y, w, h, border int) {
	c.ctx.CopyTexImage2D(gl.Enum(target), level, gl.Enum(internal), x, y, w, h, border)
}

// Replaces a portion of an existing 2D texture image with data from the current framebuffer.
func (c *Context) CopyTexSubImage2D(target, level, xoffset, yoffset, x, y, w, h int) {
	c.ctx.CopyTexSubImage2D(gl.Enum(target), level, xoffset, yoffset, x, y, w, h)
}

// Creates and initializes a WebGLBuffer.
func (c *Context) CreateBuffer() *Buffer {
	return &Buffer{c.ctx.CreateBuffer()}
}

// Returns a WebGLFramebuffer object.
func (c *Context) CreateFramebuffer() *FrameBuffer {
	return &FrameBuffer{c.ctx.CreateFramebuffer()}
}

// Creates an empty WebGLProgram object to which vector and fragment
// WebGLShader objects can be bound.
func (c *Context) CreateProgram() *Program {
	return &Program{c.ctx.CreateProgram()}
}

// Creates and returns a WebGLRenderbuffer object.
func (c *Context) CreateRenderbuffer() *RenderBuffer {
	return &RenderBuffer{c.ctx.CreateRenderbuffer()}
}

// CreateShader returns an empty vertex or fragment shader object based on the type specified.
func (c *Context) CreateShader(typ int) *Shader {
	shader := &Shader{c.ctx.CreateShader(gl.Enum(typ))}
	return shader
}

// Used to generate a WebGLTexture object to which images can be bound.
func (c *Context) CreateTexture() *Texture {
	return &Texture{c.ctx.CreateTexture()}
}

// Sets whether or not front, back, or both facing facets are able to be culled.
func (c *Context) CullFace(mode int) {
	c.ctx.CullFace(gl.Enum(mode))
}

// Delete a specific buffer.
func (c *Context) DeleteBuffer(buffer *Buffer) {
	c.ctx.DeleteBuffer(buffer.Buffer)
}

// Deletes a specific WebGLFramebuffer object. If you delete the
// currently bound framebuffer, the default framebuffer will be bound.
// Deleting a framebuffer detaches all of its attachments.
func (c *Context) DeleteFramebuffer(framebuffer *FrameBuffer) {
	c.ctx.DeleteFramebuffer(framebuffer.Framebuffer)
}

// Flags a specific WebGLProgram object for deletion if currently active.
// It will be deleted when it is no longer being used.
// Any shader objects associated with the program will be detached.
// They will be deleted if they were already flagged for deletion.
func (c *Context) DeleteProgram(program *Program) {
	c.ctx.DeleteProgram(program.Program)
}

// Deletes the specified renderbuffer object. If the renderbuffer is
// currently bound, it will become unbound. If the renderbuffer is
// attached to the currently bound framebuffer, it is detached.
func (c *Context) DeleteRenderbuffer(renderbuffer *RenderBuffer) {
	c.ctx.DeleteRenderbuffer(renderbuffer.Renderbuffer)
}

// Deletes a specific shader object.
func (c *Context) DeleteShader(shader *Shader) {
	c.ctx.DeleteShader(shader.Shader)
}

// Deletes a specific texture object.
func (c *Context) DeleteTexture(texture *Texture) {
	c.ctx.DeleteTexture(texture.Texture)
}

// Sets whether or not you can write to the depth buffer.
func (c *Context) DepthMask(flag bool) {
	c.ctx.DepthMask(flag)
}

// Sets the depth range for normalized coordinates to canvas or viewport depth coordinates.
func (c *Context) DepthRange(zNear, zFar float32) {
	c.ctx.DepthRangef(float32(zNear), float32(zFar))
}

// Detach a shader object from a program object.
func (c *Context) DetachShader(program *Program, shader *Shader) {
	c.ctx.DetachShader(program.Program, shader.Shader)
}

// Turns off specific WebGL capabilities for this context.
func (c *Context) Disable(cap int) {
	c.ctx.Disable(gl.Enum(cap))
}

// Turns off a vertex attribute array at a specific index position.
func (c *Context) DisableVertexAttribArray(index int) {
	c.ctx.DisableVertexAttribArray(gl.Attrib{uint(index)})
}

// Render geometric primitives from bound and enabled vertex data.
func (c *Context) DrawArrays(mode, first, count int) {
	c.ctx.DrawArrays(gl.Enum(mode), first, count)
}

// Renders geometric primitives indexed by element array data.
func (c *Context) DrawElements(mode, count, typ, offset int) {
	c.ctx.DrawElements(gl.Enum(mode), count, gl.Enum(typ), offset)
}

// Turns on specific WebGL capabilities for this context.
func (c *Context) Enable(cap int) {
	c.ctx.Enable(gl.Enum(cap))
}

// Turns on a vertex attribute at a specific index position in
// a vertex attribute array.
func (c *Context) EnableVertexAttribArray(index int) {
	c.ctx.EnableVertexAttribArray(gl.Attrib{uint(index)})
}

func (c *Context) Finish() {
	c.ctx.Finish()
}

func (c *Context) Flush() {
	c.ctx.Flush()
}

// Attaches a WebGLRenderbuffer object as a logical buffer to the
// currently bound WebGLFramebuffer object.
func (c *Context) FrameBufferRenderBuffer(target, attachment, renderbufferTarget int, renderbuffer *RenderBuffer) {
	c.ctx.FramebufferRenderbuffer(gl.Enum(target), gl.Enum(attachment), gl.Enum(renderbufferTarget), renderbuffer.Renderbuffer)
}

// Attaches a texture to a WebGLFramebuffer object.
func (c *Context) FramebufferTexture2D(target, attachment, textarget int, texture *Texture, level int) {
	c.ctx.FramebufferTexture2D(gl.Enum(target), gl.Enum(attachment), gl.Enum(textarget), texture.Texture, level)
}

// Sets whether or not polygons are considered front-facing based
// on their winding direction.
func (c *Context) FrontFace(mode int) {
	c.ctx.FrontFace(gl.Enum(mode))
}

// Creates a set of textures for a WebGLTexture object with image
// dimensions from the original size of the image down to a 1x1 image.
func (c *Context) GenerateMipmap(target int) {
	c.ctx.GenerateMipmap(gl.Enum(target))
}

// Returns an WebGLActiveInfo object containing the size, type, and name
// of a vertex attribute at a specific index position in a program object.
func (c *Context) GetActiveAttrib(program *Program, index int) (name string, size int, ty int) {
	n, s, typ := c.ctx.GetActiveAttrib(program.Program, uint32(index))
	return n, s, int(typ)
}

// Returns an WebGLActiveInfo object containing the size, type, and name
// of a uniform attribute at a specific index position in a program object.
func (c *Context) GetActiveUniform(program *Program, index int) (name string, size int, ty int) {
	n, s, typ := c.ctx.GetActiveUniform(program.Program, uint32(index))
	return n, s, int(typ)
}

// Returns a slice of WebGLShaders bound to a WebGLProgram.
func (c *Context) GetAttachedShaders(program *Program) []*Shader {
	objs := c.ctx.GetAttachedShaders(program.Program)
	shaders := make([]*Shader, len(objs))
	for i := 0; i < len(shaders); i++ {
		shaders[i] = &Shader{objs[i]}
	}
	return shaders
}

// Returns an index to the location in a program of a named attribute variable.
func (c *Context) GetAttribLocation(program *Program, name string) int {
	return int(c.ctx.GetAttribLocation(program.Program, name).Value)
}

// Returns the type of a parameter for a given buffer.
func (c *Context) GetBufferParameter(target, pname int) int {
	return c.ctx.GetBufferParameteri(gl.Enum(target), gl.Enum(pname))
}

// Returns a value for the WebGL error flag and clears the flag.
func (c *Context) GetError() int {
	return int(c.ctx.GetError())
}

// Enables a passed extension, otherwise returns null.
func (c *Context) GetExtension(name string) {
	// TODO: doesn't seem to exist on mobile?
}

// TODO: Create type specific variations.
// Gets a parameter value for a given target and attachment.
func (c *Context) GetFramebufferAttachmentParameter(target, attachment, pname int) int {
	return c.ctx.GetFramebufferAttachmentParameteri(gl.Enum(target), gl.Enum(attachment), gl.Enum(pname))
}

// Returns the value of the program parameter that corresponds to a supplied pname
// which is interpreted as an int.
func (c *Context) GetProgramParameteri(program *Program, pname int) int {
	return c.ctx.GetProgrami(program.Program, gl.Enum(pname))
}

// Returns the value of the program parameter that corresponds to a supplied pname
// which is interpreted as a bool.
func (c *Context) GetProgramParameterb(program *Program, pname int) bool {
	log.Println("GetProgramParameterb not found on mobile system")
	return false
}

// Returns information about the last error that occurred during
// the failed linking or validation of a WebGL program object.
func (c *Context) GetProgramInfoLog(program *Program) string {
	return c.ctx.GetProgramInfoLog(program.Program)
}

// Returns a renderbuffer parameter from the currently bound WebGLRenderbuffer object.
func (c *Context) GetRenderbufferParameter(target, pname int) int {
	return c.ctx.GetRenderbufferParameteri(gl.Enum(target), gl.Enum(pname))
}

// Returns the value of the parameter associated with pname for a shader object.
func (c *Context) GetShaderParameter(shader *Shader, pname int) int {
	return c.ctx.GetShaderi(shader.Shader, gl.Enum(pname))
}

// Returns the value of the parameter associated with pname for a shader object.
func (c *Context) GetShaderParameterb(shader *Shader, pname int) bool {
	log.Println("GetShaderParameterb not found on mobile system")
	return false
}

// Returns errors which occur when compiling a shader.
func (c *Context) GetShaderInfoLog(shader *Shader) string {
	return c.ctx.GetShaderInfoLog(shader.Shader)
}

func (c *Context) GetShaderiv(shader *Shader, param uint32) bool {
	return c.ctx.GetShaderi(shader.Shader, gl.Enum(param)) == gl.TRUE
}

// Returns source code string associated with a shader object.
func (c *Context) GetShaderSource(shader *Shader) string {
	return c.ctx.GetShaderSource(shader.Shader)
}

// Returns a slice of supported extension strings.
func (c *Context) GetSupportedExtensions() []string {
	return nil // TODO: extensions dont exist on mobile?
}

// Returns the value for a parameter on an active texture unit.
func (c *Context) GetTexParameterfv(target, pname int) {
	log.Println("Warning: GetTexParameterfv is not yet implemented")
}

// Gets the uniform value for a specific location in a program.
func (c *Context) GetUniformfv(program *Program, location *UniformLocation) {
	log.Println("Warning: GetUniformfv is not yet implemented")
}

// Returns a WebGLUniformLocation object for the location
// of a uniform variable within a WebGLProgram object.
func (c *Context) GetUniformLocation(program *Program, name string) *UniformLocation {
	return &UniformLocation{c.ctx.GetUniformLocation(program.Program, name)}
}

// TODO: Create type specific variations.
// Returns data for a particular characteristic of a vertex
// attribute at an index in a vertex attribute array.
func (c *Context) GetVertexAttribfv(index, pname int) {
	log.Println("Warning: GetVertexAttribfv is not yet implemented")
}

// Returns the address of a specified vertex attribute.
func (c *Context) GetVertexAttribOffset(index, pname int) int {
	log.Println("GetVertexAttribOffset not found on mobile system")
	return -1
}

// public function hint(target:GLenum, mode:GLenum) : Void;

// Returns true if buffer is valid, false otherwise.
func (c *Context) IsBuffer(buffer *Buffer) bool {
	return c.ctx.IsBuffer(buffer.Buffer)
}

// Returns whether the WebGL context has been lost.
func (c *Context) IsContextLost() bool {
	return false // because it can't be lost in android
}

// Returns true if buffer is valid, false otherwise.
func (c *Context) IsFramebuffer(framebuffer *FrameBuffer) bool {
	return c.ctx.IsFramebuffer(framebuffer.Framebuffer)
}

// Returns true if program object is valid, false otherwise.
func (c *Context) IsProgram(program *Program) bool {
	return c.ctx.IsProgram(program.Program)
}

// Returns true if buffer is valid, false otherwise.
func (c *Context) IsRenderbuffer(renderbuffer *RenderBuffer) bool {
	return c.ctx.IsRenderbuffer(renderbuffer.Renderbuffer)
}

// Returns true if shader is valid, false otherwise.
func (c *Context) IsShader(shader *Shader) bool {
	return c.ctx.IsShader(shader.Shader)
}

// Returns true if texture is valid, false otherwise.
func (c *Context) IsTexture(texture *Texture) bool {
	return c.ctx.IsTexture(texture.Texture)
}

// Returns whether or not a WebGL capability is enabled for this context.
func (c *Context) IsEnabled(capability int) bool {
	return c.ctx.IsEnabled(gl.Enum(capability))
}

// Sets the width of lines in WebGL.
func (c *Context) LineWidth(width float32) {
	c.ctx.LineWidth(width)
}

// Links an attached vertex shader and an attached fragment shader
// to a program so it can be used by the graphics processing unit (GPU).
func (c *Context) LinkProgram(program *Program) {
	c.ctx.LinkProgram(program.Program)

	result := c.ctx.GetProgrami(program.Program, gl.LINK_STATUS)

	if result != gl.TRUE {
		info := c.ctx.GetProgramInfoLog(program.Program)
		log.Print("Error: could not link program: ", info)
	}
}

// Sets pixel storage modes for readPixels and unpacking of textures
// with texImage2D and texSubImage2D.
func (c *Context) PixelStorei(pname, param int) {
	c.ctx.PixelStorei(gl.Enum(pname), int32(param))
}

// Sets the implementation-specific units and scale factor
// used to calculate fragment depth values.
func (c *Context) PolygonOffset(factor, units float32) {
	c.ctx.PolygonOffset(factor, units)
}

/*
// TODO: Figure out if pixels should be a slice.
// Reads pixel data into an ArrayBufferView object from a
// rectangular area in the color buffer of the active frame buffer.
func (c *Context) ReadPixels(x, y, width, height, format, typ int, pixels js.Object) {
	c.ctx.ReadPixels(x, y, width, height, format, typ, pixels)
}
*/

// Creates or replaces the data store for the currently bound WebGLRenderbuffer object.
func (c *Context) RenderbufferStorage(target, internalFormat, width, height int) {
	c.ctx.RenderbufferStorage(gl.Enum(target), gl.Enum(internalFormat), width, height)
}

// Sets the dimensions of the scissor box.
func (c *Context) Scissor(x, y, width, height int) {
	c.ctx.Scissor(int32(x), int32(y), int32(width), int32(height))
}

// ShaderSource sets and replaces shader source code in a shader object.
func (c *Context) ShaderSource(shader *Shader, source string) {
	c.ctx.ShaderSource(shader.Shader, source)
}

// public function stencilMask(mask:GLuint) : Void;
// public function stencilMaskSeparate(face:GLenum, mask:GLuint) : Void;

// Loads the supplied pixel data into a texture.
func (c *Context) TexImage2D(target, level, internalFormat, format, kind int, data interface{}) {
	if format != internalFormat {
		log.Println("Warning: format and internalFormat should be the same for TexImage2D on mobile system")
	}

	switch img := data.(type) {
	case *image.NRGBA:
		c.ctx.TexImage2D(gl.Enum(target), level, img.Bounds().Dx(), img.Bounds().Dy(), gl.Enum(format), gl.Enum(kind), *(*[]byte)(unsafe.Pointer(&img.Pix)))
	case *image.RGBA:
		c.ctx.TexImage2D(gl.Enum(target), level, img.Bounds().Dx(), img.Bounds().Dy(), gl.Enum(format), gl.Enum(kind), *(*[]byte)(unsafe.Pointer(&img.Pix)))
	default:
		log.Println("Warning: TexImage2D does not support your requested type (yet)")
	}
}

// Sets texture parameters for the current texture unit.
func (c *Context) TexParameteri(target int, pname int, param int) {
	c.ctx.TexParameteri(gl.Enum(target), gl.Enum(pname), param)
}

/*
// Replaces a portion of an existing 2D texture image with all of another image.
func (c *Context) TexSubImage2D(gl.Enum(target), level, xoffset, yoffset, format, typ int, image js.Object) {
	c.ctx.TexSubImage2D(gl.Enum(target), level, xoffset, yoffset, format, typ, image)
}
*/

// Assigns a floating point value to a uniform variable for the current program object.
func (c *Context) Uniform1f(location *UniformLocation, x float32) {
	c.ctx.Uniform1f(location.Uniform, x)
}

// Assigns a integer value to a uniform variable for the current program object.
func (c *Context) Uniform1i(location *UniformLocation, x int) {
	c.ctx.Uniform1i(location.Uniform, x)
}

// Assigns 2 floating point values to a uniform variable for the current program object.
func (c *Context) Uniform2f(location *UniformLocation, x, y float32) {
	c.ctx.Uniform2f(location.Uniform, x, y)
}

// Assigns 2 integer values to a uniform variable for the current program object.
func (c *Context) Uniform2i(location *UniformLocation, x, y int) {
	c.ctx.Uniform2i(location.Uniform, x, y)
}

// Assigns 3 floating point values to a uniform variable for the current program object.
func (c *Context) Uniform3f(location *UniformLocation, x, y, z float32) {
	c.ctx.Uniform3f(location.Uniform, x, y, z)
}

// Assigns 3 integer values to a uniform variable for the current program object.
func (c *Context) Uniform3i(location *UniformLocation, x, y, z int) {
	c.ctx.Uniform3i(location.Uniform, int32(x), int32(y), int32(z))
}

// Assigns 4 floating point values to a uniform variable for the current program object.
func (c *Context) Uniform4f(location *UniformLocation, x, y, z, w float32) {
	c.ctx.Uniform4f(location.Uniform, x, y, z, w)
}

// Assigns 4 integer values to a uniform variable for the current program object.
func (c *Context) Uniform4i(location *UniformLocation, x, y, z, w int) {
	c.ctx.Uniform4i(location.Uniform, int32(x), int32(y), int32(z), int32(w))
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
// TODO: document that 'transpose' is not supported
func (c *Context) UniformMatrix2fv(location *UniformLocation, transpose bool, value []float32) {
	c.ctx.UniformMatrix2fv(location.Uniform, value)
}

// Sets values for a 3x3 floating point vector matrix into a
// uniform location as a matrix or a matrix array.
func (c *Context) UniformMatrix3fv(location *UniformLocation, transpose bool, value []float32) {
	c.ctx.UniformMatrix3fv(location.Uniform, value)
}

// Sets values for a 4x4 floating point vector matrix into a
// uniform location as a matrix or a matrix array.
func (c *Context) UniformMatrix4fv(location *UniformLocation, transpose bool, value []float32) {
	c.ctx.UniformMatrix4fv(location.Uniform, value)
}

// Set the program object to use for rendering.
func (c *Context) UseProgram(program *Program) {
	c.ctx.UseProgram(program.Program)
}

// Returns whether a given program can run in the current WebGL state.
func (c *Context) ValidateProgram(program *Program) {
	c.ctx.ValidateProgram(program.Program)
}

func (c *Context) VertexAttribPointer(index, size, typ int, normal bool, stride int, offset int) {
	c.ctx.VertexAttribPointer(gl.Attrib{uint(index)}, size, gl.Enum(typ), normal, stride, offset)
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
	c.ctx.Viewport(x, y, width, height)
}
