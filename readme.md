go test -coverprofile=cover.out . && go tool cover -func=cover.out
go tool cover -html=cover.out