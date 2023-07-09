# Language-Helper

Simple language learning helper backend using `go` and the `gin framework`.

### `dewit.sh`

Runs the application without building it.

### **Dockerization** Development

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
docker run --rm -it -p 8080:8080 --name go_backend langhelp/backend_dev ./lh_backend
```

### **Dockerization** Production

Build Docker image for go backend - PROD version:
```sh
docker build -t langhelp/backend_prod . -f docker/Dockerfile.prod
```

Run docker image:
```sh
docker run -d -p 8080:8080 --name go_backend langhelp/backend_prod
```

Stop container:
```sh
docker stop go_backend
```