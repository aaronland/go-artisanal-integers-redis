# go-artisanal-integers-redis

No, really.

## Running as an AWS Lambda function

First run the `lambda` target in the included Makefile. For example:

```
$> make lambda
if test -d pkg; then rm -rf pkg; fi
if test -d src/github.com/aaronland/go-artisanal-integers-redis; then rm -rf src/github.com/aaronland/go-artisanal-integers-redis; fi
mkdir -p src/github.com/aaronland/go-artisanal-integers-redis/
cp *.go src/github.com/aaronland/go-artisanal-integers-redis/
cp -r engine src/github.com/aaronland/go-artisanal-integers-redis/
cp -r vendor/* src/
if test -f main; then rm -f main; fi
if test -f deployment.zip; then rm -f deployment.zip; fi
zip deployment.zip main
  adding: main (deflated 66%)
rm -f main
```

_Something something something create your AWS Elasticache cluster here..._

_Something something something create your AWS Lambda function here..._

Make sure the set the following environment variables:

| Environment variable | Value |
| --- | --- |
| `ARTISANAL_DSN` | _Something like `redis://{HOST}.cache.amazonaws.com:6379`_ |
| `ARTISANAL_PROTOCOL` | `lambda` |

## Performance

#### Redis

Running `intd` backed by Redis on a vanilla Vagrant machine (running Ubuntu 14.04) on a laptop against 1000 concurrent users, using siege:

```
siege -c 1000 http://localhost:8080
** SIEGE 3.0.5
** Preparing 1000 concurrent users for battle.
The server is now under siege...^C
Lifting the server siege...      done.

Transactions:			110761 hits
Availability:			100.00 %
Elapsed time:			63.92 secs
Data transferred:		0.59 MB
Response time:			0.06 secs
Transaction rate:		1732.81 trans/sec
Throughput:			0.01 MB/sec
Concurrency:			98.32
Successful transactions:	110761
Failed transactions:		0
Longest transaction:		6.24
Shortest transaction:		0.00
```

## See also

* https://github.com/aaronland/go-artisanal-integers
