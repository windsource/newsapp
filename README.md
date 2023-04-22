# News app

Go application that reads RSS feeds defined in `data/data.json` and provides all on one HTML page.

## Preparation

1. Go must have been installed

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
docker build -t windsource/newsapp .
```

To run the image

```bash
docker run -p 8080:8080 windsource/newsapp
```





