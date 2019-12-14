# Martian Robots

Coding challenge.

`martian-robots` is a command line tool that reads instructions from stdin.

Results are written to stdout.

# Requirements

Go 1.13+

# Usage

Assuming that `$GOROOT/bin` is in your `$PATH`:

```
go install ./cmd/martian-robots
cat samples/sample_1.txt | martian-robots
```