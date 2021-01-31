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

## ENV

Plan:
- Detect env value for environment
- conditionally set db strign to localhost or actually using ENV file's values like `DB_HOST` or whatever to build up the connection string
- When moved to Digitalocean we can add the same vars through the GUI