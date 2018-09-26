# API

## Installation

### With Docker

```bash
docker pull astrocorp/signal_api
```

### With Go

```
git clone git@github.com:astrocorp42/signal.git
cd api
make
```

## Usage

The following environment variables are required:

| Var | Description |
| --- | ----------- |
| **DATABASE_URL** | The URL to your postges databases (e.g. "postgresql://user:pass@host/mydb...") |
| **JWT_SECRET** | An unguessable (50 chars min) secret to sign the JWTs |
