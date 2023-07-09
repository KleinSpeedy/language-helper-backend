# Language-Helper

Simple language learning helper backend using `go` and the `gin framework`.

### `dewit.sh`

Runs the application without building it.

### **Dockerization**

Build Docker image for go backend - DEV version:
```sh
docker build -t langhelp/backend_dev . -f docker/Dockerfile.dev
```

Run tests for go app:
```sh
docker run --rm langhelp/backend_dev go test .
```

Run docker image:
```sh
docker run -d -p 8080:8080 --name go_backend langhelp/backend_dev ./lh_backend
```
