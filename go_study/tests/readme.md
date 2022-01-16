### Testing

To run test

```
go test
```

Run all tests in the root

```
go test ./...
```

Run all tests in the root verbose

```
go test -v ./...
```

Run tests in parallel

```
t.parallel()
```

Run with covarage

```
go test --cover
```

Test with page

```
go tool cover -html=coverage.out
```
