# Advent of Code solutions

I originally started doing these exercises in Go, but eventually switched to Rust. 
The implementations are separated in folders by language. 


Solutions are in the numbered folders and I made a makefile to easily run them. I decided 
to write Big-O, Big-Theta and Big-Omega notations of each algorithm to remember 
algorithm analysis. I might be wrong though so I accept corrections via PR's.

### day 1

Day one solution is nothing fancy. A simple loop over n lines, looping over lines of size m.
It's O(m*n), both parts. While part 2 being slightly worse because of indexing. 

### day 2

Day two is also pretty simple. loop over n lines, loop over each line to find s sets. Complexity is
O(n*s)




