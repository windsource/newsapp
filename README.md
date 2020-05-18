# News app

Go application that reads RSS feeds defined in `data/data.json` and provides all on one HTML page.

## Preparation

1. Go must have been installed
2. Get the source with `go get github.com/windsource/newsapp`

## Build

Switch to the folder in which this file resides and call

```bash
go build
```

## Run

```bash
./newsapp
```

Point your browser to http://localhost:8080.

## Docker

To create a Docker image:

```bash
make build
```

To push the image to Docker Hub:

```bash
make push
```

To run the image

```bash
docker run -p 8080:8080 windsource/newsapp
```





