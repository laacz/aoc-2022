# Advent of Code 2022

These are solutions for the [Advent of Code 2022](https://adventofcode.com/2022) challenges.

This year it's Go time. Or to put it punly, let's give it a Go. Or just let's Go.

```bash
# Runs the tests first. If they pass, runs solutions for the given day.
run.sh day01
```

## Day 1

Easy peasy.

## Day 2

Lemon squeezy.

## Day 3

Nothing to see here.

## Day 4

Same here.

## Day 5

Straightforward, until I got to the second part. In certain cases Go modifies target slice, instead of copying. So I had
to resort to moving crates one-by one.

## Day 6

Shortest solution to date. Guessed the second part's task, so both parts use the same implementation.

## Day 7

Easy, but had to think about the right approach longer than I'd like.

## Day 8

Got stuck on failing tests, because missed that first lines do not count.

## Day 9

This took some while, because I messed up X and Y coordinates in one place. This was a relatively easy one nonetheless.

## Day 10

A few off by one one errors, but this was sweet.

## Day 11

Integers have a nasty habit of overflowing. After a half an hour or so sorted that out.

## Day 12

Graphs. Not a fan. Took some while.

## Day 13

Oh, this one would be so much easier using a loosely typed language. Fuuuuuuu.. I don't have words. Took two days, three
rewrites. On the bright side, I learned a lot about `interface{}`, `reflect` package and myself.

## Day 14

What a relief and fun after previous day. Well, day 13 was hard only because of Go typing and me not being used to it.

## Day 15

First part was easy enough.

Second part was a bit different. First - we generate a list of all lines at manhattan distance + 1. Then we check if any
of these lines intersect (well, we don't but still resulting list of points is not huge). If they do, we add the point
to the list of possible candidates for a place to have the beacon at. Then we just iterate over all these points and
find the first one which is not covered by any sensor.

## Day 16

On the fourth day I gave up. Will have to revisit this one and theory behind it. Graphs are my nemesis.

## Day 17

Part 1 was very sweet and I enjoyed that a lot. Part 2, however, was not fun at all. Although it was fairly obvious that
what we're looking for is a cycle, actual approach and with that its solution took me several attempts.
