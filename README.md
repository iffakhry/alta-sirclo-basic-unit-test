## Basic Unit Testing in Go

Untuk menjalankan unit test (pilih salah satu) :
1. Menjalankan semua file test dengan `go test`
```
go test ./...
```

2. go test with cover
```
go test ./... -cover
```

3. go test with report coverage
```
go test ./... -coverprofile=cover.out && go tool cover -html=cover.out
```

### Untuk menjalan file test secara spesifik berdasarkan folder (sesuaikan dengan nama folder):
```
go test ./calculator/
```
