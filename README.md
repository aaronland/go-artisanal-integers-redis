# go-artisanal-integers-redis

No, really.

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
