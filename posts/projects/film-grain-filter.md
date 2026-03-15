---
title: "WegGL-01: film grain filter"
slug: "film-grain-filter"
parent: "projects"
description: "Exploring film grain and noise generation using WebGL shaders."
order: 1
headers: "Introduction, GPU vs CPU, WebGL Pipeline, Types of Noise"
seo_title: "Film Grain Filter"
seo_meta_description: "Exploring film grain and procedural noise using WebGL shaders."
seo_meta_property_title: "Noise Filter"
seo_meta_property_description: "A study of GPU-generated film grain and noise in WebGL."
seo_meta_og_url: "https://grafiquer.com/projects/film-grain-filter"
---

# WebGL-0: Film Grain Filter

1. 20 year old me had been listening to Them&I for a while (if you like downtempo/chill electronic vibes you should give it a try) and realised how aesthetic and vintage the front pages of his songs were. Soon, driven by my obession with old cameras and willing to replicate that effect in my pictures, I left vanilla java script for the first time and dove head first into: **WebGL 1.0**.

<div class="w-full h-[800px] py-5">
    <iframe class="block w-full h-full border-0"
        src="/projects/film-grain-filter/canvas.html"
    ></iframe>
</div>

## Brief intro to WebGL

2. WebGL stands for Web Graphics Library, it is simply an API that allows web pages to access your GPU to render graphics using rasterization. That is, in computer graphics there are two main rendering approaches:

- Ray tracing: simulates rays that hit objects in the scene, making them visible.
- Rasterization: draw vertices -> build triangles -> color them -> image

3. In other words, WebGL just exposes the rasterization pipeline of your GPU to the browser.
   Both WebGl 1.0 and 2.0 call OpenGL ES functions through JS. This means we can write programms for the GPU (shaders), and this API (OpenGL) will send it. The language in which these programms are written OpenGL Shading Language (GLSL).

4. Notice to Mariners: I started with WebGL 1.0 because I thought it would help me appreciate and understand better 2.0, no other particular reason
   The 2.0 version provides built-ins for matters that require extensions in 1.0.

JavaScript
↓
WebGL API
↓
OpenGL ES API
↓
GPU driver
↓
GPU hardware

JavaScript
↓
WebGL API
↓
OpenGL ES
↓
GLSL shaders
↓
GPU execution

If rather than believing in your GPU you prefer to understand it (like me), I encourage you to check the official manual: <https://webglfundamentals.org/>, since this is not a tutorial per se. For the TLRD:

4. How my shader works

The API is already built on the browser, so no need to download or embed a library. Actually, you can even work online with Greggman´s <https://jsfiddle.net/greggman/8djzyjL3/>

## How the shader works

How it works in a nutshell:

```javascript

<script id="vertex_shader" type="x-shader/x-vertex">
  // Which part of the image goes on which part of the triangle?
  // That is why we define texCoord attribute
  attribute vec2 a_position;
  attribute vec2 a_tex_coord;
  uniform vec2 u_resolution;
  varying vec2 v_tex_coord;

  void main() {
     // convert coordinates to clip-space
     vec2 zero_to_one = a_position / u_resolution;
     vec2 zero_to_two = zero_to_one * 2.0;
     vec2 clip_space = zero_to_two - 1.0;
     gl_Position = vec4(clip_space*vec2(1,-1), 0, 1); // flip y-axis

     // pass the texCoord to the fragment shader
     // The GPU will interpolate this value between points.
     v_tex_coord = a_tex_coord;
  }
</script>
```

```javascript

<script id="fragment_shader" type="x-shader/x-fragment">
  precision highp float;

  uniform sampler2D u_image;
  uniform vec2 u_resolution;
  varying vec2 v_tex_coord;
  uniform float blend_val;
  uniform float grainsize;

  float rand(vec2 co){
    return fract(sin(dot(co, vec2(12.9898, 78.233))) * 43758.5453);
  }

  // Brightness
  vec4 grain(vec4 fragColor){
    vec4 color = fragColor;

    vec2 cell = floor(gl_FragCoord.xy / grainsize);
    float diff = rand(cell) - 0.5;

    color.rgb += diff;
    return color;
  }

  void main() {
      vec2 uv = gl_FragCoord.xy / u_resolution;
      vec4 color = texture2D(u_image, v_tex_coord);
      vec4 grain_val = grain(color);
      gl_FragColor = mix(color, grain_val, blend_val);

  }
</script>
```

- Vertex shader
- Fragment shader (what is fragment)

- GPU as a state machine (buffers, etc.)
  example with buffer

- `gl.*` constants

5. There is a library not to build your own functions and methods

### Different kinds of noise

4. I just discovered — GPT told me — that there are programs that talk directly to the GPU, whereas JavaScript normally talks to the CPU. These programs accelerate the rendering of pixels on screen. One of them is **WebGL**, which I intend to learn over time.

   Check this resource: <https://webglfundamentals.org/>

   Things to review:
   - GPU as a state machine (buffers, etc.)
   - `gl.*` constants
   - The rendering pipeline

5. Film grain reference:
   <https://maximmcnair.com/p/webgl-film-grain>

## Applications

I might write a tutorial further on.

Resources

Learn more:
https://bruno-simon.com/

https://glslsandbox.com/
https://observablehq.com/@observable81?tab=recents

<https://maximmcnair.com/p/webgl-film-grain>

```

```
