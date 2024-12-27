# Day 11 Solutions in Go / Golang

A friend was telling me about problem #11, so I thought I'd give it a try.

Part A only required a recursive function to perform the calculation.

When part B required 75 "blinks" my program started consuming a lot of resources.  To remedy this, I simulated "memoizing" the function so that if the stone face value and the depth of the traversal had been seen before, the function would return the previously computed sum.  This permitted the B portion of the problem to run almost instantly by caching those common search results.

---

In a recent episode of my podcast ( Stray Pointers ) two friends and I discuss the AoC competition from prior years and our intended approaches to this year's competition.  You might want to give it a listen.

https://straypointers.com/e/s2e16.htm

# New - 27 December 2024

We just published our AoC2024 wrap-up video podcast:

https://straypointers.com/e/s2e20.htm

