# Language-Helper

Simple language learning helper backend using `go` and the `gin framework`.

### `dewit.sh`

Runs the application without building it.
### **Dockerization**

Build Multistage Docker image for go backend:
```sh
docker build -t langhelp/backend . -f docker/Dockerfile
```

Run docker image:
```sh
docker run -d -p 8080:8080 --name go_backend langhelp/backend
```
