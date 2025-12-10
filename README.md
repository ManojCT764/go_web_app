# go_web_app

DevOps implementation on a web application written in Go.

# TaskBoard


A minimal Go web application (TaskBoard) built with the standard library. Intended for practicing DevOps: containerization, CI, monitoring, and deployment.


## Features


- In-memory task store (no DB)
- HTML UI + small JSON API
- Health endpoint (`/healthz`) for probes
- Graceful shutdown and basic logging


## Run locally


```bash
go run main.go
# open http://localhost:8080