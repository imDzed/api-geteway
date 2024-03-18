go test -cover ./service-employee/services/ -coverprofile=cover.out && go tool cover -func=cover.out
go test -cover ./service-user/services/ -coverprofile=cover.out && go tool cover -func=cover.out