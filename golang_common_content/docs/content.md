# Annotations

## comands
You can create a module in go using:
```
go mod init module_name
```
After to build it, just execute inside your folder when you have your `go.md`:
```
go build
```
Installing external dependencie for example:
```
go get github.com/badoux/checkmail
```
Cleaning your go.mod case you had removed some dependecie
```
go mod tidy
```

## Inf
- External function modules are writing using it first letter Upercase;
- Every variable is already initiated, strings with "", and numbers with 0
