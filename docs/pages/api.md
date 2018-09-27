# API

[![Docker image](https://img.shields.io/badge/docker-astrocorp/signal__api-blue.svg)](https://hub.docker.com/r/astrocorp/signal_api)


## Usage

The following environment variables are required:

| Var | Description |
| --- | ----------- |
| **DATABASE_URL** | The URL to your postges databases (e.g. "postgresql://user:pass@host/mydb...") |
| **JWT_SECRET** | An unguessable (50 chars min) secret to sign the JWTs |
| **HOST** | Your API host (e.g. "https://signal.kerkour.com" |

```bash
$ cat .env
DATABASE_URL="xxx"
JWT_SECRET="xxx"
HOST="http://localhost:8080"
```

### With Docker
```bash
# $ docker run -d --name signal_postgres -e POSTGRES_USER="signal" -e POSTGRES_PASSWORD="xxx" postgres
$ docker run -ti --rm --name signal_api -v $PWD/.env:/signal/.env:ro astrocorp/signal_api /signal/signal -migrate # to migrate the database schema
$ docker run -d --name signal_api -v $PWD/.env:/signal/.env:ro -p 8080:8080 astrocorp/signal_api
```

### With Go

```
git clone git@github.com:astrocorp42/signal.git
cd api
make # or go install
./dist/signal -migrate # to migrate the database schema
./dist/signal
```
