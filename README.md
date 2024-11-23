# quake_report

> Parser de logs do jogo Quake

## Sobre

Faz o parse do arquivo de log `qgames.log` e produz um relatório no formato JSON com as informações:

- Dados agrupados para cada partida
- Ranking de jogadores baseado no número de kills
- Relatório de mortes agrupadas por causa da morte

Um exemplo do relatório gerado pode ser visto no arquivo [report.json](report.json).

**Atenção:** Nenhuma validação está sendo feita. Confiando totalmente no formato do arquivo.

## Funcionalidades

- Leitura do arquivo de Log
- Agrupamento de dados por partida
- Coleta de dados de morte
- Geração de relatório no formato JSON

## Estrutura de Arquivos

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

## Pré-requisitos

É necessário o Go 1.23 ou Docker instalado na máquina.

## Instalação

```bash
git clone https://github.com/gustavohenrique/quake_report
cd quake_report
```

## Rodando

### Usando apenas Go

```bash
make install
make lint   
make tests  
make run

# É possível passar o caminho do arquivo de log via variável de ambiente
LOG_FILE=/tmp/file.log make run
```

### Usando Docker

```bash
docker build . -t=quake_report
docker run --rm -v $PWD:/app quake_report
```

## Contribuindo

Contribuições são sempre bem-vindas!

1. Crie um fork do projeto
2. Crie uma nova branch
3. Commit suas alterações
4. Faça um push da sua branch
5. Abra um Pull Request

## Licença

Distribuído sobre a licença MIT. Veja o arquivo `LICENSE` para mais informações.

