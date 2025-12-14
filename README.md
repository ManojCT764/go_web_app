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
# go build -o main .
# ./main
# open http://localhost:8585


## in docker
docker build -t youknowmanoj/go-web-app:v1 . 
docker run -p 8585:8585 -it youknowmanoj/go-web-app:v2