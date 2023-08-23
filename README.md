# go-postgres-api

###### To build GO app
- `go mod init ***MOD_NAME***`

[NOTE: Sometimes the version of gofiber will cause trouble]

###### To deploy to railway
- `railway init`
- Add a Postgres DB service to app and copy env vars to .env file:
```
railway add
```
(Only allows MySQL, Redis, and MongoDB)
- `railway run go run app.go`
- `railway link`
- `railway up`
- `railway domain`
