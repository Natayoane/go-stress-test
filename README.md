# Ferramenta de Teste de Carga em Go

Uma ferramenta simples de linha de comando para testes de carga em serviços web, escrita em Go.

## Funcionalidades

- Manipulação de requisições concorrentes
- Número total de requisições personalizável
- Nível de concorrência personalizável
- Relatório detalhado do teste incluindo:
  - Tempo total de execução
  - Total de requisições realizadas
  - Requisições bem-sucedidas (HTTP 200)
  - Distribuição de códigos de status
  - Requisições por segundo

## Compilação e Execução

### Compilação Local

```bash
go build -o stress-test
./stress-test --url=http://exemplo.com --requests=1000 --concurrency=10
```

### Usando Docker

Compilar a imagem Docker:
```bash
docker build -t go-stress-test .
```

Executar o container:
```bash
docker run go-stress-test --url=http://exemplo.com --requests=1000 --concurrency=10
```

## Parâmetros

- `--url`: URL do serviço a ser testado (obrigatório)
- `--requests`: Número total de requisições a serem feitas (obrigatório)
- `--concurrency`: Número de requisições concorrentes (obrigatório)

## Exemplo de Saída

```
=== Relatório do Teste de Carga ===
Tempo Total: 5.234s
Total de Requisições: 1000
Requisições Bem-sucedidas (200): 985

Distribuição de Códigos de Status:
Status 200: 985
Status 404: 10
Status 500: 5

Requisições por segundo: 191.06
``` 