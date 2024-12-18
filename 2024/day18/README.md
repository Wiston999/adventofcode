# --- Day 18: RAM Run ---

You and The Historians look a lot more pixelated than you remember. You're [/2017/day/2](inside a computer) at the North Pole!


Just as you're about to check out your surroundings, a program runs up to you. "This region of memory isn't safe! The User misunderstood what a [https://en.wikipedia.org/wiki/Pushdown_automaton](pushdown automaton) is and their algorithm is pushing whole <em><b>bytes</b></em> down on top of us! <span title="Pun intended.">Run</span>!"


The algorithm is fast - it's going to cause a byte to fall into your memory space once every [https://www.youtube.com/watch?v=9eyFDBPk4Yw](nanosecond)! Fortunately, you're <em><b>faster</b></em>, and by quickly scanning the algorithm, you create a <em><b>list of which bytes will fall</b></em> (your puzzle input) in the order they'll land in your memory space.


Your memory space is a two-dimensional grid with coordinates that range from <code>0</code> to <code>70</code> both horizontally and vertically. However, for the sake of example, suppose you're on a smaller grid with coordinates that range from <code>0</code> to <code>6</code> and the following list of incoming byte positions:


<pre><code>5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0
</code></pre>
Each byte position is given as an <code>X,Y</code> coordinate, where <code>X</code> is the distance from the left edge of your memory space and <code>Y</code> is the distance from the top edge of your memory space.


You and The Historians are currently in the top left corner of the memory space (at <code>0,0</code>) and need to reach the exit in the bottom right corner (at <code>70,70</code> in your memory space, but at <code>6,6</code> in this example). You'll need to simulate the falling bytes to plan out where it will be safe to run; for now, simulate just the first few bytes falling into your memory space.


As bytes fall into your memory space, they make that coordinate <em><b>corrupted</b></em>. Corrupted memory coordinates cannot be entered by you or The Historians, so you'll need to plan your route carefully. You also cannot leave the boundaries of the memory space; your only hope is to reach the exit.


In the above example, if you were to draw the memory space after the first <code>12</code> bytes have fallen (using <code>.</code> for safe and <code>#</code> for corrupted), it would look like this:


<pre><code>...#...
..#..#.
....#..
...#..#
..#..#.
.#..#..
#.#....
</code></pre>
You can take steps up, down, left, or right. After just 12 bytes have corrupted locations in your memory space, the shortest path from the top left corner to the exit would take <code><em><b>22</b></em></code> steps. Here (marked with <code>O</code>) is one such path:


<pre><code><em><b>O</b></em><em><b>O</b></em>.#<em><b>O</b></em><em><b>O</b></em><em><b>O</b></em>
.<em><b>O</b></em>#<em><b>O</b></em><em><b>O</b></em>#<em><b>O</b></em>
.<em><b>O</b></em><em><b>O</b></em><em><b>O</b></em>#<em><b>O</b></em><em><b>O</b></em>
...#<em><b>O</b></em><em><b>O</b></em>#
..#<em><b>O</b></em><em><b>O</b></em>#.
.#.<em><b>O</b></em>#..
#.#<em><b>O</b></em><em><b>O</b></em><em><b>O</b></em><em><b>O</b></em>
</code></pre>
Simulate the first kilobyte (<code>1024</code> bytes) falling onto your memory space. Afterward, <em><b>what is the minimum number of steps needed to reach the exit?</b></em>


