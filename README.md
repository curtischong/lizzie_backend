# How to run
`cd main`
`go build -o main.o && ./main.o`
important because we don't want executables in the git history. the gitignore ignores all .o files
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
- Use `flt.printf("stuff %v\n", var)` to print dev-assigned variables
- Use `log.println(err)` to print general logs
- Try to use this style to create objects:
`
dbObj := DBObj{
  DBConfigObj: config.DBConfig
  DBClient: nil
}`


# Known Issues
 - If the DB doesn't say: "Added x!" after "connected to DB after" then it probably didn't connect to the database
 - diagnostic: run `ifconfig` to see what subnet IP (`inet`) you are on (make sure you are connected to the VPN)
      - There should be 3 subnet devices: utun0, utun1, utun2
 - Proposal: Have a statistics card sent in Lizzie Peaks every night or morning that tells you how many server events you fired in the past 24 hours. This is a sanity check so I know that the systems are running

 - This local IP thing is solved for the dev ip bc localhost will never have a different IP.
 - I think this isn't a problem bc the local ip of the prod server shouldn't change unless of a restart?