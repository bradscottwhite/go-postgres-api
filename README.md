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


To Test
To fetch data:
```
curl https://go-postgres-api-production.up.railway.app/allbooks
```
To write data:
```
curl https://go-postgres-api-production.up.railway.app/addbook \
	--include \
	--header "Content-Type: application/json" \
	--request "POST" \
	--data '{"Title": "First", "Author": "Brad"}'
```
