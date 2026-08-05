# Grafiquer

My personal website and experimental digital portfolio, built with Go and native web technologies.

Grafiquer brings together software development, technical writing, interface design, mathematics, and generative graphics in a single server-rendered application.

## Features

- Personal landing and about pages.
- Project portfolio with dedicated project pages.
- Markdown-based articles and project descriptions.
- Server-rendered HTML components.
- Syntax highlighting for technical content.
- SEO and social-sharing metadata.
- Interactive graphics and WebGL experiments.
- Custom request logging and error handling.
- Responsive interface built without a large frontend framework.

## Technologies

- Go
- templ
- HTMX
- HTML and CSS
- JavaScript
- Markdown
- Goldmark
- WebGL and GLSL

## Project structure

- `web/landing/` — homepage handlers and components.
- `web/about/` — personal profile page.
- `web/articles/` — article rendering.
- `web/projects/` — project catalogue and detail pages.
- `web/shared/` — reusable interface components.
- `web/static/` — stylesheets, scripts, images, and browser assets.
- `posts/` — Markdown content and metadata.
- `projects/` — interactive project implementations.
- `logger/` — structured HTTP logging.
- `types/` — shared content and SEO models.

## Getting started

Requirements:

- Go 1.24 or later

Install dependencies and run the server:

```sh
go mod download
go run .
```

Then open:

```text
http://localhost:8080
```

If template files are modified, regenerate the templ output before running the application:

```sh
templ generate
```

## Included experiments

The portfolio includes visual and mathematical projects such as a film-grain filter and Bézier-curve-based animation experiments.

## Motivation

I built Grafiquer as a space where technical work and artistic exploration can coexist. It is both a portfolio and a platform for experimenting with software architecture, writing, interaction design, and computational graphics.

## Status

Grafiquer is an evolving personal project. New articles, software experiments, and visual projects are added over time.
