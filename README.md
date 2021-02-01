# lantern-go

## Dependencies

- `brew install modd`

## Dev Flow

`modd` to run the main.go and run gin server

## Migrations

TODO: Use a Makefile or something LOL

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
- Ensure that environment variable `LANTERN_ENV` is configured as `development` for dev and `production` for production. 
  - For local dev, you can do this via a `.zshrc` export: `export LANTERN_ENV=production`
  - For DigitalOcean Apps, through the Web UI
- `lantern-go` will load `LANTERN_ENV` and:
  - if `development`, will load using `.env`
  - if `production`, will load using standard os.Getenv without an `.env`
- Duplicate the .env.example file and create a database with the corresponding values in postgres