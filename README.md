# Advent of Code 2017

These are my solutions for the Advent of Code 2017 challenges. The whole point of this is to try and get more comfortable with Go, so this code will probably suffer some rewrites (even for past days as I learn better ways to do things.)

## Running it

The source tree already contains my inputs. To run day 1, for instance, follow these steps:

```shell
$ cd 01
$ go test ./...
$ cat input | go run $(ls -1 *.go | grep -v _test.go)
```

## Feedback

If you're ever so inclined, I'd love feedback in the form of Issues, Pull Requests or angry tweets.
