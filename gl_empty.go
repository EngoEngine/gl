// This doesn't actually send anything to the graphics card. This is so that servers
// can use engo and other parts of the engine while being headless and not depending
// on the OpenGL library. Anything that does have a return, returns a zero-value
// of whatever's supposed to be returned, except returns true for booleans in case
// success checks are used.

// +build nogl

package gl

import (
	"unsafe"
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
	RGBA8                                        int
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
	return &Context{}
}

func (c *Context) CreateShader(typ int) *Shader {
	shader := &Shader{0}
	return shader
}

func (c *Context) ShaderSource(shader *Shader, source string) {}

func (c *Context) CompileShader(shader *Shader) {}

func (c *Context) Ptr(data interface{}) unsafe.Pointer {
	var ptr unsafe.Pointer
	return ptr
}

func (c *Context) Str(str string) *uint8 {
	var num uint8
	return &num
}

func (c *Context) GoStr(cstr *uint8) string {
	return ""
}

func (c *Context) DeleteShader(shader *Shader) {}

func (c *Context) DeleteTexture(texture *Texture) {}

func (c *Context) GetShaderiv(shader *Shader, pname uint32) bool {
	return true
}

func (c *Context) GetShaderInfoLog(shader *Shader) string {
	return ""
}

func (c *Context) CreateProgram() *Program {
	return &Program{0}
}

func (c *Context) DeleteProgram(program *Program) {}

func (c *Context) BindAttribLocation(program *Program, index int, name string) {}

func (c *Context) GetProgramParameteri(program *Program, pname int) int {
	return 0
}

func (c *Context) GetProgramParameterb(program *Program, pname int) bool {
	return true
}

func (c *Context) GetProgramInfoLog(program *Program) string {
	return ""
}

func (c *Context) AttachShader(program *Program, shader *Shader) {}

func (c *Context) LineWidth(width float32) {}

func (c *Context) LinkProgram(program *Program) {}

func (c *Context) CreateTexture() *Texture {
	return &Texture{0}
}

func (c *Context) BindTexture(target int, texture *Texture) {}

func (c *Context) ActiveTexture(target int) {}

func (c *Context) TexParameteri(target int, pname int, param int) {}

func (c *Context) TexImage2D(target, level, internalFormat, format, kind int, data interface{}) {}

func (c *Context) GetAttribLocation(program *Program, name string) int {
	return 0
}

func (c *Context) GetUniformLocation(program *Program, name string) *UniformLocation {
	return &UniformLocation{0}
}

func (c *Context) GetError() int {
	return 0
}

func (c *Context) CreateBuffer() *Buffer {
	return &Buffer{0}
}

func (c *Context) DeleteBuffer(buffer *Buffer) {}

func (c *Context) BindBuffer(target int, buffer *Buffer) {}

func (c *Context) BufferData(target int, data interface{}, usage int) {}

func (c *Context) EnableVertexAttribArray(index int) {}

func (c *Context) DisableVertexAttribArray(index int) {}

func (c *Context) VertexAttribPointer(index, size, typ int, normal bool, stride int, offset int) {}

func (c *Context) Enable(flag int) {}

func (c *Context) Disable(flag int) {}

func (c *Context) BlendFunc(src, dst int) {}

func (c *Context) BlendEquation(mode int) {}

func (c *Context) UniformMatrix2fv(location *UniformLocation, transpose bool, value []float32) {}

func (c *Context) UniformMatrix3fv(location *UniformLocation, transpose bool, value []float32) {}

func (c *Context) UniformMatrix4fv(location *UniformLocation, transpose bool, value []float32) {}

func (c *Context) UseProgram(program *Program) {}

func (c *Context) ValidateProgram(program *Program) {}

func (c *Context) Uniform1f(location *UniformLocation, x float32) {}

func (c *Context) Uniform1i(location *UniformLocation, x int) {}

func (c *Context) Uniform2f(location *UniformLocation, x, y float32) {}

func (c *Context) Uniform3f(location *UniformLocation, x, y, z float32) {}

func (c *Context) Uniform4f(location *UniformLocation, x, y, z, w float32) {}

func (c *Context) BufferSubData(target int, offset int, data interface{}) {}

func (c *Context) DrawArrays(mode, first, count int) {}

func (c *Context) DrawElements(mode, count, typ, offset int) {}

func (c *Context) ClearColor(r, g, b, a float32) {}

func (c *Context) Viewport(x, y, width, height int) {}

func (c *Context) GetViewport() [4]int32 {
	return [4]int32{0, 0, 0, 0}
}

func (c *Context) Scissor(x, y, width, height int) {}

func (c *Context) Clear(flags int) {}

func (c *Context) MatrixMode(mode uint32) {}

func (c *Context) LoadIdentity() {}

func (c *Context) PushMatrix() {}

func (c *Context) PopMatrix() {}

func (c *Context) CreateRenderBuffer() *RenderBuffer {
	return &RenderBuffer{0}
}

func (c *Context) DeleteRenderBuffer(rb *RenderBuffer) {
}

func (c *Context) BindRenderBuffer(rb *RenderBuffer) {
}

func (c *Context) RenderBufferStorage(internalFormat uint32, width, height int) {
}

func (c *Context) CreateFrameBuffer() *FrameBuffer {
	return &FrameBuffer{0}
}

func (c *Context) DeleteFrameBuffer(fb *FrameBuffer) {
}

func (c *Context) BindFrameBuffer(fb *FrameBuffer) {
}

func (c *Context) FrameBufferTexture2D(target, attachment, texTarget uint32, t *Texture, level int) {
}

func (c *Context) FrameBufferRenderBuffer(target, attachment uint32, rb *RenderBuffer) {
}
