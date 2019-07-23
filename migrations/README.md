### Quickstart

Setup the user that will be talking to the db:
`createuser --interactive --pwprompt`

### How to run Migrations:
`brew install golang-migrate`
`migrate -path migrations/ -database postgres://localhost:5432/lizzie?sslmode=disable up n_migrations`

### What happens when you get a dirty migration?
run `migrate -path migrations/ -database postgres://localhost:5432/lizzie?sslmode=disable force n_latest_migration`

### How to handle time
- For tables that are frequently used:
  - Store unix time as `unixt`
- For other tables:
  - Store unix time AND local ts as `unixt` and `ts`
  - local ts helps me do analysis
  - unix time helps me convert to whatever I want
  - Why not UTC?
    - Doesn't help with analysis too much
    - Isn't as universal as unix time
-  Store `unixt` up to millisecond precision as a `bigInt` data type

### Table naming convention:
  - Use singular table names. Many reasons on SO
  - Use singular column names.
  - Don't use camelCase use snake_case bc we don't want accidents with double quotes

