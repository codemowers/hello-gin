Just a testing clone to play with automated deployment

# Hello cloud native Gin

Cloud native example of [Gin web framework](https://gin-gonic.com/)

* Includes Prometheus metrics middleware and endpoint
* Builds statically linked binary for minimized Docker image footprint


## Usage

Building and running locally:

```
docker-compose up --build
```

Proceed to open up http://localhost:8000

For Skaffold usage substitute `foobar` in `deployment.yaml` with your username.
