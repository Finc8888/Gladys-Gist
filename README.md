# Gladys Gist Project

## Description

Project by tutorial of the book Let's Go by Alex Edwards(2 edition)

## Go's documentation

```bash
go doc fmt
```

or

```bash
go install golang.org/x/pkgsite/cmd/pkgsite@latest pkgsite -open
```

## Creating module

```bash
go mod init gist.gladys.net
```

## Run project

```bash
go run ./cmd/web
go run gist.gladys.net/cmd/web
```

### Run with options
```bash
go run ./cmd/web -help
go run ./cmd/web -addr=":5000"
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
### GET request
```bash
curl http://localhost:4000/gist/view?id=1st/view?id=1
```
### Range request
```bash
curl -i -H "Range: bytes=100-199" --output - http://localhost:4000/static/img/logo.png
```
