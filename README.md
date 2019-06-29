# Style guide

- Try to declare all constants in the main.go file
- Name new structs: object_nameObj
- Do not put the `Obj` suffix for instances of structs
- If you are using acronyms use CAPS (even if you are declaring variables)
- If the entire variable are acronyms: ex: DBID, all instances of this acronym should be lower case: dbid
- Try to pass structs by reference to be modified if the current package shouldn't have any struct logic
    - ex: the main package doesn't handle any of the database logic it only remembers a database object
    - This has the bennefit of not having to import package-specific libraries in other packages
    - ex: the influx client package into main
- All database names and columns are snake_case
- There should only be one config object passed through the stack.
   - bc variables like IsDev is used everywhere (also LogV - log verbosity) //TODO: impliment LogV
   - DO NOT copy the config variable into an object's own config.
      - you don't know which version of the config to trust
      - if you also pass in a config variable you are passing two of the same config
- Try to use this style to create objects:
```
dbObj := DBObj{
  DBConfigObj: config.DBConfig
  DBClient: nil
}```
