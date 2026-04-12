# Gladys Gist Project

## Description

Project by tutorial of the book Let's Go by Alex Edwards(2 edition)

## Creating module

```bash
go mod init gist.gladys.net
```

## Run project

```bash
go run .
go run main.go
go run gist.gladys.net
```

## Command to display local ports on Linux
Go check it during http.ListenAndServe()
```bash
cat /etc/services
```
## Set the default name for git first branch

```bash
git config --global init.defaultBranch main
```

### Rename branch

```bash
git branch -m main
```

## Use curl for make requests to API
### POST request
```bash
curl -i -X POST http://localhost:4000/gist/create
```
