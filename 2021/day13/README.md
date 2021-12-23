# --- Day 13: Transparent Origami ---

You reach another volcanically active part of the cave. It would be nice if you could do some kind of thermal imaging so you could tell ahead of time which caves are too hot to safely enter.


Fortunately, the submarine seems to be equipped with a thermal camera! When you activate it, you are greeted with:


<pre><code>Congratulations on your purchase! To activate this infrared thermal imaging
camera system, please enter the code found on page 1 of the manual.
</code></pre>
Apparently, the Elves have never used this feature. To your surprise, you manage to find the manual; as you go to open it, page 1 falls out. It's a large sheet of [https://en.wikipedia.org/wiki/Transparency_(projection)](transparent paper)! The transparent paper is marked with random dots and includes instructions on how to fold it up (your puzzle input). For example:


<pre><code>6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5
</code></pre>
The first section is a list of dots on the transparent paper. <code>0,0</code> represents the top-left coordinate.  The first value, <code>x</code>, increases to the right.  The second value, <code>y</code>, increases downward.  So, the coordinate <code>3,0</code> is to the right of <code>0,0</code>, and the coordinate <code>0,7</code> is below <code>0,0</code>. The coordinates in this example form the following pattern, where <code>#</code> is a dot on the paper and <code>.</code> is an empty, unmarked position:


<pre><code>...#..#..#.
....#......
...........
#..........
...#....#.#
...........
...........
...........
...........
...........
.#....#.##.
....#......
......#...#
#..........
#.#........
</code></pre>
Then, there is a list of <em><b>fold instructions</b></em>. Each instruction indicates a line on the transparent paper and wants you to fold the paper <em><b>up</b></em> (for horizontal <code>y=...</code> lines) or <em><b>left</b></em> (for vertical <code>x=...</code> lines). In this example, the first fold instruction is <code>fold along y=7</code>, which designates the line formed by all of the positions where <code>y</code> is <code>7</code> (marked here with <code>-</code>):


<pre><code>...#..#..#.
....#......
...........
#..........
...#....#.#
...........
...........
-----------
...........
...........
.#....#.##.
....#......
......#...#
#..........
#.#........
</code></pre>
Because this is a horizontal line, fold the bottom half <em><b>up</b></em>. Some of the dots might end up overlapping after the fold is complete, but dots will never appear exactly on a fold line. The result of doing this fold looks like this:


<pre><code>#.##..#..#.
#...#......
......#...#
#...#......
.#.#..#.###
...........
...........
</code></pre>
Now, only <code>17</code> dots are visible.


Notice, for example, the two dots in the bottom left corner before the transparent paper is folded; after the fold is complete, those dots appear in the top left corner (at <code>0,0</code> and <code>0,1</code>). Because the paper is transparent, the dot just below them in the result (at <code>0,3</code>) remains visible, as it can be seen through the transparent paper.


Also notice that some dots can end up <em><b>overlapping</b></em>; in this case, the dots merge together and become a single dot.


The second fold instruction is <code>fold along x=5</code>, which indicates this line:


<pre><code>#.##.|#..#.
#...#|.....
.....|#...#
#...#|.....
.#.#.|#.###
.....|.....
.....|.....
</code></pre>
Because this is a vertical line, fold <em><b>left</b></em>:


<pre><code>#####
#...#
#...#
#...#
#####
.....
.....
</code></pre>
The instructions made a square!


The transparent paper is pretty big, so for now, focus on just completing the first fold. After the first fold in the example above, <code><em><b>17</b></em></code> dots are visible - dots that end up overlapping after the fold is completed count as a single dot.


<em><b>How many dots are visible after completing just the first fold instruction on your transparent paper?</b></em>


