# lantern-go

## Dependencies

- `brew install modd`
- `go get`

We use `go-migrate` to create and run migrations.
## Dev Flow

`modd` to run the main.go and run gin server

## Migrations

To create a new migration, run:

```sh
migrate create -ext sql -dir db/migrations -seq create_users_table
```

```sh
export CONNECTION_STRING="-source file://db/migrations -database 'postgres://localhost:5432/lantern-go?sslmode=disable'"
alias mu="migrate $CONNECTION_STRING up"
alias md="migrate $CONNECTION_STRING down"
alias mdd="migrate $CONNECTION_STRING drop -f"
```

## Seeding
- Queries must be run in order, since most tables "hang off" the `People` table
- To alter the amount of records, go into the corresponding `_test.go` file in `seed/` and alter the `ModelCount` variable before running below.

```sh
go test -run TestSeedPeople ./seed
go test -run TestSeedNotes ./seed
go test -run TestSeedPressurePoints ./seed
```

## Environment Variable Configuration
- Duplicate the .env.example file and create a database with the corresponding values in postgres
- Ensure that environment variable `LANTERN_ENV` is configured as `development` for dev and `production` for production. 
  - For local dev, you can do this via a `.zshrc` export: `export LANTERN_ENV=production`
  - For DigitalOcean Apps, through the Web UI
- `lantern-go` will load `LANTERN_ENV` and:
  - if `development`, will load using `.env`
  - if `production`, will load using standard os.Getenv without an `.env`
## RUNNING MIGRATIONS
This is a big to-do, but for now: we will not need to run multiple migrations - just go onto the DigitalOcean Web UI to get the PostgreSQL's component's login information - connect manually or using TablePlus/similar, and run the SQL.up migration SQL file that we have in `db/migrations`. 
