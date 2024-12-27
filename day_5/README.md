# Day 5 Solutions in Go / Golang

Thinking about these ... I don't think I implemented them correctly.  I think they work because the numbers in the updates happen to land next to each other if I needed to compare them.  If a number X needed to be before number Y but they appear in the data as Y, some-other-number, X  ... my code won't catch that.  If I would have coded for all of those permutations, I think my code would have run somewhat longer.  Maybe too long.  I'll have to tinker with that at a later time.

I approached this solution by fusing the rules together as all of them appeared to be two digits ... a pipe ... and two digits.  I just fused all of the digits into a 4-digit number and made an existential dictionary to see if the rule existed.  In the first problem, I looked to see if the current number and the number prior appeared in reverse order in the existential "rules" map.  If so, then it was invalid.

If the current number and the next number appeared in the opposite orer in the rules map, then it was invalid.

In the "b" challenge, I swapped out the offensive entries as soon as I found them and rescanned the line fresh each time.

---

In a recent episode of my podcast ( Stray Pointers ) two friends and I discuss the AoC competition from prior years and our intended approaches to this year's competition.  You might want to give it a listen.

https://straypointers.com/e/s2e16.htm

# New - 27 December 2024

We just published our AoC2024 wrap-up video podcast:

https://straypointers.com/e/s2e20.htm

