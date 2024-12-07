# Day 6 Solutions in Go / Golang

The code for part A of this challenge loads the text file into an array of strings.  A map[int]bool is created as an existential dictionary to see if an X already occupies a spot.  An array of int pairs with the x/y increments for the four 90 degree directions is used to move the logical marker around the 2D grid.

---

The code for part B uses the core of the processing loop for part A to run multiple simulations.  It will attempt to simulate the 'O' symbol in every position on the grid as long as that position currently only holds a '.'  If the simulation runs too long, the sum counter is incremented and the simulation ends.  I multiplied the dimensions of the array and added 1 as a test limit to see if we've been looping too long.

---

In a recent episode of my podcast ( Stray Pointers ) two friends and I discuss the AoC competition from prior years and our intended approaches to this year's competition.  You might want to give it a listen.

https://straypointers.com/e/s2e16.htm

We'll be meeting up again soon for a new episode to discuss how we're faring in this year's competition.
