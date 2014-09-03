// Copyright 2014 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webgl

/*
Objects:

Texture
Buffer
FrameBuffer
RenderBuffer
Program
UniformLocation
Shader

PerFragment:

void blendColor(float red, float green, float blue, float alpha)
void blendEquation(enum mode)
void blendEquationSeparate(enum modeRGB, enum modeAlpha)
void blendFunc(enum sfactor, enum dfactor)
void blendFuncSeparate(enum srcRGB, enum dstRGB, enum srcAlpha, enum dstAlpha)
void depthFunc(enum func)
void sampleCoverage(float value, bool invert)
void stencilFunc(int func, int ref, uint mask)
void stencilFuncSeparate(enum face, enum func, int ref, uint mask)
void stencilOp(enum fail, enum zfail, enum zpass)
void stencilOpSeparate(enum face, enum fail, enum zfail, enum zpass)

FrameBuffer:

void clear(ulong mask)
void clearColor(float red, float green, float blue, float alpha)
void clearDepth(float depth)
void clearStencil(int s)
void colorMask(bool red, bool green, bool blue, bool alpha)
void depthMask(bool flag)
void stencilMask(uint mask)
void stencilMaskSeparate(enum face, uint mask)
void bindFramebuffer(enum target, Object framebuffer)
enum checkFramebufferStatus(enum target)
Object createFramebuffer()
void deleteFramebuffer(Object buffer)
void framebufferRenderbuffer(enum target, enum attachment, enum renderbuffertarget, Object renderbuffer)
bool isFramebuffer(Object framebuffer)
void framebufferTexture2D(enum target, enum attachment, enum textarget, Object texture, int level)
any getFramebufferAttachmentParameter(enum target, enum attachment, enum pname)

Buffer:

void bindBuffer(enum target, Object buffer)
void bufferData(enum target, long size, enum usage)
void bufferData(enum target, Object data, enum usage)
void bufferSubData(enum target, long offset, Object data)
Object createBuffer()
void deleteBuffer(Object buffer)
any getBufferParameter(enum target, enum pname)
bool isBuffer(Object buffer)

View:

void depthRange(float zNear, float zFar)
void scissor(int x, int y, long width, long height)
void viewport(int x, int y, long width, long height)

Rasterization:

void cullFace(enum mode)
void frontFace(enum mode)
void lineWidth(float width)
void polygonOffset(float factor, float units)

Shaders:

void attachShader(Object program, Object shader)
void bindAttribLocation(Object program, uint index, string name)
void compileShader(Object shader)
Object createProgram()
Object createShader(enum type)
void deleteProgram(Object program)
void deleteShader(Object shader)
void detachShader(Object program, Object shader)
Object[ ] getAttachedShaders(Object program)
any getProgramParameter(Object program, enum pname)
string getProgramInfoLog(Object program)
any getShaderParameter(Object shader, enum pname)
string getShaderInfoLog(Object shader)
string getShaderSource(Object shader)
bool isProgram(Object program)
bool isShader(Object shader)
void linkProgram(Object program)
void shaderSource(Object shader, string source)
void useProgram(Object program)
void validateProgram(Object program)

Textures:

void activeTexture(enum texture)
void bindTexture(enum target, Object texture)
void copyTexImage2D(enum target, int level, enum internalformat, int x, int y, long width, long height, int border)
void copyTexSubImage2D(enum target, int level, int xoffset, int yoffset, int x, int y, long width, long height)
Object createTexture()
void deleteTexture(Object texture)
void generateMipmap(enum target)
any getTexParameter(enum target, enum pname)
bool isTexture(Object texture)
void texImage2D(enum target, int level, enum internalformat, long width, long height, int border, enum format, enum type, Object pixels)
void texImage2D(enum target, int level, enum internalformat, enum format, enum type, Object object)
void texParameterf(enum target, enum pname, float param)
void texParameteri(enum target, enum pname, int param)
void texSubImage2D(enum target, int level, int xoffset, int yoffset, long width, long height, enum format, enum type, Object pixels)
void texSubImage2D(enum target, int level, int xoffset, int yoffset, enum format, enum type, Object object)


Special:

void disable(enum cap)
void enable(enum cap)
void finish()
void flush()
enum getError()
any getParameter(enum pname)
void hint(enum target, enum mode)
bool isEnabled(enum cap)
void pixelStorei(enum pname, int param)

Uniforms and Attributes:

void disableVertexAttribArray(uint index)
void enableVertexAttribArray(uint index)
Object getActiveAttrib(Object program, uint index)
Object getActiveUniform(Object program, uint index)
ulong getAttribLocation(Object program, string name)
any getUniform(Object program, uint location)
uint getUniformLocation(Object program, string name)
any getVertexAttrib(uint index, enum pname)
long getVertexAttribOffset(uint index, enum pname)
void uniform[1234][fi](uint location, ...)
void uniform[1234][fi]v(uint location, Array value)
void uniformMatrix[234]fv(uint location, bool transpose, Array)
void vertexAttrib[1234]f(uint index, ...)
void vertexAttrib[1234]fv(uint index, Array value)
void vertexAttribPointer(uint index, int size, enum type, bool normalized, long stride, long offset)

RenderBuffer:

void bindRenderbuffer(enum target, Object renderbuffer)
Object createRenderbuffer()
void deleteRenderbuffer(Object renderbuffer)
any getRenderbufferParameter(enum target, enum pname)
bool isRenderbuffer(Object renderbuffer)
void renderbufferStorage(enum target, enum internalformat, long width, long height)

DrawBuffer:

void drawArrays(enum mode, int first, long count)
void drawElements(enum mode, long count, enum type, long offset)

ReadPixels:

void readPixels(int x, int y, long width, long height, enum format, enum type, Object pixels)

*/
