## Echo + Casbin + Postgres

### Run the project

`$ docker-compose up -d`

### Test cases
```bash
$ curl http://admin:@localhost:8080/admin
Admin page accessed
```

```bash
$ curl http://guest:@localhost:8080/admin
{"message":"Forbidden"}
```

```bash
$ curl http://guest:@localhost:8080/login
Login accessed
```

```bash
$ curl http://user:@localhost:8080/logout
Logout accessed
```
