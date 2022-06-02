# Token Writer
Simple cli tool which generates random tokens, dumps them to a file, then
reads the file and inserts the data into a DB.

## Assignment

#### Tools:

Use golang and combine it with any relational database (postgres, mysql, ...) you like.

#### Summary:

Read a file containing random tokens and store them in the database as quickly and
efficiently as possible without storing any token twice and create a list of all non-unique tokens.

#### In detail:

1. First write a token generator that creates a file with 10 million random tokens, one per line,
each consisting of seven lowercase letters a-z.
 
2. Then write a token reader that reads the file and stores
the tokens in your DB.
 
3. Naturally some tokens will occur more than once, so take care that these aren't
duplicated in the DB, but do produce a list of all non-unique tokens and their frequencies.
 
4. Find a clever way to do it efficiently in terms of network I/O, memory, and time and include documentation inline with
your code or as txt file, describing your design decisions.

5. Send us the code of your generator, reader, and documentation file if you have it separate.

6. Please also send us a description of your database layout and preferably the complete DB schema.

Feel free to reach out with any questions you may have.
We look forward to seeing your results and hope you enjoy this task!

## Running locally
Clone the project repository:

```$ git clone https://github.com/radu2020/data4life.git```

Pulling dependencies:

```$ go get github.com/mattn/go-sqlite3```

Running locally:

```$ go run main.go```

## Configuration
Configuration is hardcoded inside the main file.

|  variable name |  value |  description |
|---|---|---|
|  tokenAmount |  1.000.000 |  amount of tokens to be generated |
|  tokenLength |  7 |  fixed length of generated token |
|  fileName |  "tokens.txt" |  name of file where tokens are dumped |
|  fileWriteBatchSize |  100.000 |  size of token batch used to write to file |
|  dbWriteBatchSize |  10.000 |  size of records batch used to pause writing to db |
|  sleepDuration |  2 |  seconds to pause db write |

## Database
### Database layout description (Tokens Table)

 |  Column Name  | Data Type  | Description  |
 |---|---|---|
 |  id  |  INTEGER PRIMARY KEY AUTOINCREMENT |  record id |
 |  token |  TEXT |  token string |
 |  repeated |  INTEGER |  if token is unique in table or not, sqlite boolean is an int (0 = false, 1 = true) |
 |  frequency |  INTEGER |  token occurrences |

### Database Schema
Each time when running the program, a complete database file is generated on local file system,
called `data4life.db`.

The program is inserting all tokens inside a table inside this DB and the final
results can be easily viewed using SQLite.

### SQLite Database
The advantage of SQLite is that it is easier to install and use and the resulting
database is a single file that can be written to a USB memory stick
or emailed to a colleague.
For coding challenges or small projects it is a great tool as it requires
minimal configuration compared to other DBMS.

## View final results using SQLite
#### Install on Mac
```$ brew install sqlite```
#### Install on Linux
```$ sudo apt install sqlite3```
#### Open Local DB File
```$ sqlite3 data4life.db```
#### Select records inside Tokens Table
```sql
SELECT * FROM tokens;
SELECT * FROM tokens WHERE repeated != 0;
SELECT * FROM tokens WHERE frequency != 1;
```

## Optimizations

Writing the file is done using bufio pkg which implements buffered I/O.
The file write batch size is there just to track write progress.

Reading the file whole file into memory is done using the same bufio pkg and for
each string read we should take in account that 8 bytes are required, so for 
1M tokens there should be 80MB of memory available just for the read.

The db write batch size and the sleep duration parameters were added to fine tune
how many records can be written to the db before putting the process on hold for
a few seconds to reduce load on the DB.


## Example output

### Running the program
```
radu@mypc data4life % go run main.go
------------------------------
FileName:       tokens.txt
TokenLength:    7
TokenAmount:    1000
BatchSize:      100
DBBatchSize:    100
sleepDuration:  2
------------------------------
Writing to file progress:
------------------------------
100/1000
200/1000
300/1000
400/1000
500/1000
600/1000
700/1000
800/1000
900/1000
1000/1000
------------------------------
Finished writing to file.
------------------------------
Writing to DB progress:
------------------------------
100/1000
200/1000
300/1000
400/1000
500/1000
600/1000
700/1000
800/1000
900/1000
1000/1000
------------------------------
Finished writing to database.
------------------------------
```

### SQLite DB Select
```sql
sqlite> SELECT * FROM tokens LIMIT 10;
1|wkvrsll|0|1
2|zivabjm|0|1
3|hvidkaj|0|1
4|vuqotff|0|1
5|yhemhui|0|1
6|hzjrhej|0|1
7|aiisnba|0|1
8|hrlihqj|0|1
9|vdtaggb|0|1
10|aoaysyi|0|1
```

# That was all. Have a nice day! :)