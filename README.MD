## Exemplo de uso de biblioteca de logs para golang

```shell
  go mod init github.com/igorferreira/go-logs
  go build main.go
  go run main.go
```

Saida:

```shell

{"level":"info","message":"Teste de logs","name":"meulog","time":"2020-12-28T23:32:08-03:00"}
{"fields":{"id":1},"level":"info","message":"adicionado um campo id no log","name":"meulog","time":"2020-12-28T23:32:08-03:00"}

```