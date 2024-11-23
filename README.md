# quake_report

> Parser de logs do jogo Quake

## Estrutura

```python
.
├── main.go                            # The application entrypoint
├── src/
│   ├── shared/                        # Shared libs imported by adapters and domain layers without DI
│   ├── adapters/                      # Manage inputs and outputs
│   │   ├── dto/                       # Simple data structure to represent external data
│   │   └── converters/                # Convert DTO to models and models to DTO
│   └── domain/                        # Where the business logic and services lives
│       ├── ports/                     # Interfaces to be used outside this layer
│       ├── services/                  # Business logic
│       └── models/
├── Dockerfile
├── .testignore                        # Skip testing for some files
├── .editorconfig
└── .gitignore
```

## Como executar?

### Usando apenas Go

```sh
make install
make lint   
make tests  
make run
```

### Usando Docker

```sh
docker build . -t=quake_report
docker run --rm -v $PWD:/app quake_report
```
