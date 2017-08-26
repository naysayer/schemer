# Schemer
## Overview
Schemer is a command line tool that takes in a schema file from a database and outputs native golang structs, that match the table names, columns, and their corresponding types. 

*Please note: This is considered experimental, and should not be fit for production of any type. It should be viewed largely as a proof of concept. 

## Getting started
Please note: The only database this app currently supports is postgres.

Place your schema file within the schema.txt file that is at the woot directory of this program, and run the following:
```
$ ./schemer
```
## License
MIT