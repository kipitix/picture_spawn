# picture_spawn

Desktop application to grab pictures for wallpaper from site wallpaperswide.com and set them to desktop.

## Parts

### [sourceparser](./cmd/sourceparser/main.go)

### [dataserver](./cmd/dataserver/main.go)

## Used go modules

- Logger is [zerolog](https://github.com/rs/zerolog)
- Web page parser is [geziyor](https://github.com/geziyor/geziyor)
- Web server is [net/http](https://pkg.go.dev/net/http)

## Useful links

[Proper HTTP server shutdown](https://dev.to/mokiat/proper-http-shutdown-in-go-3fji)
[Go project structure](https://github.com/golang-standards/project-layout)
[Go and DDD](https://github.com/sklinkert/go-ddd)
[Useful Go idiom](https://duncanleung.com/go-idiom-accept-interfaces-return-types/)
