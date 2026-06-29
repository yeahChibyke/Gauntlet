# Gauntlet

> A lightweight OpenAI Responses API-compatible gateway for NVIDIA NIM.

Gauntlet allows OpenAI-compatible clients (such as AI coding agents, editors, and other developer tools) to communicate with NVIDIA NIM models through a single OpenAI-style API.

The long-term vision is to provide a drop-in replacement for the OpenAI Responses API that can route requests to multiple LLM providers while exposing a consistent interface to clients.

---

## Project Status

> **⚠️ Early Prototype**

Gauntlet is currently in active development.

### Currently Implemented

- ✅ OpenAI Responses API-compatible endpoint (`/v1/responses`)
- ✅ NVIDIA NIM provider
- ✅ Request translation layer
- ✅ Response translation layer
- ✅ Streaming responses (Server-Sent Events)
- ✅ Graceful shutdown
- ✅ Structured logging
- ✅ Request IDs
- ✅ Panic recovery middleware
- ✅ Basic `/v1/models` endpoint

### In Progress

- Full OpenAI Responses streaming protocol compatibility
- Codex CLI compatibility
- VS Code Codex extension compatibility

### Planned

- OpenAI provider
- Anthropic provider
- Google Gemini provider
- Groq
- OpenRouter
- Local Ollama support
- Provider routing
- Model aliases
- Metrics
- Authentication
- Rate limiting
- Configuration profiles

---

# Motivation

Many AI tools only support OpenAI's API.

Meanwhile, excellent models exist across providers like:

- NVIDIA NIM
- Anthropic
- Google
- Groq
- OpenRouter
- Local models

Gauntlet aims to bridge that gap by acting as a lightweight compatibility layer.

```
Client
      │
      ▼
  OpenAI Responses API
      │
      ▼
  ┌──────────────┐
  │  Gauntlet    │
  └──────────────┘
      │
      ▼
 NVIDIA NIM
```

Eventually:

```
                  ┌────────────┐
                  │ OpenAI     │
                  └────────────┘
                         ▲
                         │
                  ┌────────────┐
                  │ Anthropic  │
                  └────────────┘
                         ▲
                         │
                  ┌────────────┐
Client ───────▶   │ Gauntlet   │
                  └────────────┘
                         │
                  ┌────────────┐
                  │ NVIDIA NIM │
                  └────────────┘
                         │
                  ┌────────────┐
                  │ Ollama     │
                  └────────────┘
```

---

# Architecture

```
                  HTTP

          /v1/responses
                 │
                 ▼
          HTTP Server
                 │
                 ▼
        Response Service
                 │
                 ▼
         Canonical Request
                 │
        ┌────────┴────────┐
        │                 │
        ▼                 ▼
 Translation         Provider
                         │
                         ▼
                  NVIDIA NIM API
```

---

# Features

## Responses API

```
POST /v1/responses
```

Accepts an OpenAI Responses request and forwards it to NVIDIA NIM.

---

## Streaming

Supports streaming via Server-Sent Events (SSE).

```
stream=true
```

---

## Model Discovery

```
GET /v1/models
```

Returns available models exposed by the gateway.

---

## Middleware

- Request IDs
- Structured logging
- Panic recovery

---

# Configuration

Environment variables:

```bash
NVIDIA_API_KEY=your_api_key
```

---

# Running

## Clone

```bash
git clone https://github.com/yeahChibyke/Gauntlet.git

cd Gauntlet
```

---

## Install

```bash
go mod download
```

---

## Configure

```bash
export NVIDIA_API_KEY=xxxxxxxxxxxxxxxx
```

---

## Start

```bash
go run ./cmd/gauntlet
```

Default address:

```
http://localhost:8080
```

---

# Example

## Non-streaming

```bash
curl http://localhost:8080/v1/responses \
  -H "Content-Type: application/json" \
  -d '{
    "model":"meta/llama-3.3-70b-instruct",
    "input":[
      {
        "role":"user",
        "content":[
          {
            "type":"input_text",
            "text":"Hello!"
          }
        ]
      }
    ]
  }'
```

---

## Streaming

```bash
curl -N http://localhost:8080/v1/responses \
  -H "Content-Type: application/json" \
  -d '{
    "model":"meta/llama-3.3-70b-instruct",
    "stream":true,
    "input":[
      {
        "role":"user",
        "content":[
          {
            "type":"input_text",
            "text":"Count from one to five."
          }
        ]
      }
    ]
  }'
```

---

## Models

```bash
curl http://localhost:8080/v1/models
```

---

# Directory Layout

```
cmd/
    gauntlet/

internal/

    config/

    middleware/

    protocol/
        canonical/
        models/
        responses/

    provider/
        nvidia/

    server/

    service/

    translate/
```

---

# Design Principles

Gauntlet follows a few guiding principles:

- Keep the gateway lightweight.
- Expose an OpenAI-compatible interface.
- Separate provider-specific code from protocol translation.
- Use a canonical internal representation.
- Make adding new providers straightforward.
- Prefer composition over framework-heavy abstractions.

---

# Roadmap

## Phase 1

- [x] Responses endpoint
- [x] NVIDIA provider
- [x] Translation layer
- [x] Streaming
- [x] Logging
- [x] Recovery middleware

## Phase 2

- [ ] Complete OpenAI Responses protocol compatibility
- [ ] Full Codex CLI support
- [ ] VS Code Codex support

## Phase 3

- [ ] OpenAI provider
- [ ] Anthropic provider
- [ ] Gemini provider
- [ ] Groq provider
- [ ] Ollama provider

## Phase 4

- [ ] Provider routing
- [ ] Authentication
- [ ] Metrics
- [ ] Rate limiting
- [ ] Observability

---

# Contributing

Contributions are welcome.

If you'd like to add support for another provider or improve protocol compatibility, feel free to open an issue or submit a pull request.

---

# License

MIT License.

---

# Acknowledgements

Gauntlet builds upon the OpenAI Responses API design while aiming to make alternative model providers easier to integrate into existing AI tooling.

Special thanks to the open-source Go and AI communities whose tooling and documentation made this project possible.