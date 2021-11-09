# Run the docker container with api

```bash
docker run --rm -it -p 3000:3000  docker.io/zion/ethidx:latest
```

The swagger doc will be accessible by the url: [http://localhost:3000/swagger/](http://localhost:3000/swagger/)


# Build docker image

```bash
make docker
```