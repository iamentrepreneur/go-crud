# Helpers

```docker run --name=gocrud-db -e POSTGRES_PASSWORD='password' -p 5432:5432 -d --rm postgres```

```migrate create -ext sql -dir ./schema -seq init```

```migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up```

```migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up```

## Прокидываем migrate и обновляем PATH
```export PATH=$PATH:$HOME/go/bin```
```source ~/.zshrc```