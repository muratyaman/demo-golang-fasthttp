# demo-golang-fasthttp
Demo HTTP server using Golang and Fast HTTP

## installation

```
go get -u github.com/valyala/fasthttp
```

## execution

```
go build ./main.go
./main
```


## references

### Golang

"Go is an open source programming language that makes it easy to build simple, reliable, and efficient software."

[https://golang.org/](https://golang.org/)

### FastHTTP

"Fast HTTP implementation for Go."

[https://github.com/valyala/fasthttp](https://github.com/valyala/fasthttp)

## benchmark

Install autocannon

[https://github.com/mcollina/autocannon](https://github.com/mcollina/autocannon)

```
npm i -g autocannon
```

Run autocannon using 10 connections for 60 seconds.

```
autocannon -c 10 -d 60 http://127.0.0.1:8080
```

### A. Plain text responses

Output:

````
Running 60s test @ http://127.0.0.1:8080
10 connections

┌─────────┬──────┬──────┬───────┬──────┬─────────┬─────────┬──────────┐
│ Stat    │ 2.5% │ 50%  │ 97.5% │ 99%  │ Avg     │ Stdev   │ Max      │
├─────────┼──────┼──────┼───────┼──────┼─────────┼─────────┼──────────┤
│ Latency │ 0 ms │ 0 ms │ 0 ms  │ 0 ms │ 0.01 ms │ 0.09 ms │ 13.94 ms │
└─────────┴──────┴──────┴───────┴──────┴─────────┴─────────┴──────────┘
┌───────────┬─────────┬─────────┬─────────┬─────────┬─────────┬────────┬─────────┐
│ Stat      │ 1%      │ 2.5%    │ 50%     │ 97.5%   │ Avg     │ Stdev  │ Min     │
├───────────┼─────────┼─────────┼─────────┼─────────┼─────────┼────────┼─────────┤
│ Req/Sec   │ 24095   │ 24975   │ 36287   │ 45663   │ 36040   │ 5034.9 │ 24091   │
├───────────┼─────────┼─────────┼─────────┼─────────┼─────────┼────────┼─────────┤
│ Bytes/Sec │ 3.59 MB │ 3.72 MB │ 5.41 MB │ 6.81 MB │ 5.37 MB │ 751 kB │ 3.59 MB │
└───────────┴─────────┴─────────┴─────────┴─────────┴─────────┴────────┴─────────┘

Req/Bytes counts sampled once per second.

2162k requests in 60.08s, 322 MB read
````

### B. JSON responses

Output:

````
Running 60s test @ http://127.0.0.1:8080
10 connections

┌─────────┬──────┬──────┬───────┬──────┬─────────┬─────────┬──────────┐
│ Stat    │ 2.5% │ 50%  │ 97.5% │ 99%  │ Avg     │ Stdev   │ Max      │
├─────────┼──────┼──────┼───────┼──────┼─────────┼─────────┼──────────┤
│ Latency │ 0 ms │ 0 ms │ 0 ms  │ 0 ms │ 0.01 ms │ 0.12 ms │ 21.43 ms │
└─────────┴──────┴──────┴───────┴──────┴─────────┴─────────┴──────────┘
┌───────────┬─────────┬─────────┬─────────┬─────────┬──────────┬─────────┬─────────┐
│ Stat      │ 1%      │ 2.5%    │ 50%     │ 97.5%   │ Avg      │ Stdev   │ Min     │
├───────────┼─────────┼─────────┼─────────┼─────────┼──────────┼─────────┼─────────┤
│ Req/Sec   │ 22719   │ 27631   │ 36991   │ 41599   │ 35687.74 │ 4154.44 │ 22719   │
├───────────┼─────────┼─────────┼─────────┼─────────┼──────────┼─────────┼─────────┤
│ Bytes/Sec │ 5.09 MB │ 6.18 MB │ 8.28 MB │ 9.31 MB │ 7.99 MB  │ 930 kB  │ 5.09 MB │
└───────────┴─────────┴─────────┴─────────┴─────────┴──────────┴─────────┴─────────┘

Req/Bytes counts sampled once per second.

2141k requests in 60.07s, 479 MB read
````