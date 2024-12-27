# Day 4 Solutions in Go / Golang

The approach I took in the first part of the challenge was to first look for an 'X' and then recursively look for the remainder of the string, passing in the X and Y increments for each of eight directions around the character.

For the second portion of the challenge, I removed some of the recursive checks and I looked for the letter 'M' to start with.  After finding "MAS" diagonally, I noted the position of the letter 'A' which was common to all of the x-shapes in the center.  I incremented a value in an [int]int map.  If an x-shape is formed, the value of the 'A' at a given position will be 2.

Instead of making a record structure, I multiplied the incremented x position by 1000 and added the incremented y position to fuse the two together in one int.  If a line of input exceeds 1000 characters in length, this code will break. 

During the day, I thought about shortening the code that scans the perimeter around the given character.  I thought I could shorten those sections using an array of arrays of integers.  Please refer to new_4a.go and new_4b.go to see those approaches.

---

In a recent episode of my podcast ( Stray Pointers ) two friends and I discuss the AoC competition from prior years and our intended approaches to this year's competition.  You might want to give it a listen.

https://straypointers.com/e/s2e16.htm

# New - 27 December 2024

We just published our AoC2024 wrap-up video podcast:

https://straypointers.com/e/s2e20.htm

