# Go Actions

![lint](https://github.com/arifseft/go-actions/workflows/lint/badge.svg?branch=master)
![lint](https://github.com/arifseft/go-actions/workflows/test/badge.svg?branch=master)

## How to run

```bash
# running with docker

# running in development mode, you can use live-reload when safe file
docker-compose up development

# running in production image
docker-compose up -d production
docker-compose logs --tail=100 -f production # monitoring production container
docker-compose exec production sh # access bash on production container
```

```bash
# running in local/without docker
make install
make run
```

## Run tests

```bash
make test
```
