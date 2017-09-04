# Schemer
## Please note:
This is a toy project so use with caution. 
## Overview
Schemer is a command line tool that takes in a schema file from a database and outputs native golang structs, that match the table names, columns, and their corresponding types. 

## Getting started
Please note: The only database this program currently supports is postgres.

### Running the program / binary
Start the binary with the -schema option set to the filepath of the file you wish to have parsed.
```
$ ./schemer -schema=myschemafile.txt
```

## Defining new schema type databases to parse
In the event that one would want to expand this program to be able to parse say mysql schema files or the like; one would need to first conform to the Structure intface within the api/structure package. Check the app/postgres package for an example of how to do this.
## License
MIT