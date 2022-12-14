# --- Day 14: Regolith Reservoir ---

The distress signal leads you to a giant waterfall! Actually, hang on - the signal seems like it's coming from the waterfall itself, and that doesn't make any sense. However, you do notice a little path that leads <em><b>behind</b></em> the waterfall.


Correction: the distress signal leads you behind a giant waterfall! There seems to be a large cave system here, and the signal definitely leads further inside.


As you begin to make your way deeper underground, you feel the ground rumble for a moment. Sand begins pouring into the cave! If you don't quickly figure out where the sand is going, you could quickly become trapped!


Fortunately, your [/2018/day/17](familiarity) with analyzing the path of falling material will come in handy here. You scan a two-dimensional vertical slice of the cave above you (your puzzle input) and discover that it is mostly <em><b>air</b></em> with structures made of <em><b>rock</b></em>.


Your scan traces the path of each solid rock structure and reports the <code>x,y</code> coordinates that form the shape of the path, where <code>x</code> represents distance to the right and <code>y</code> represents distance down. Each path appears as a single line of text in your scan. After the first point of each path, each point indicates the end of a straight horizontal or vertical line to be drawn from the previous point. For example:


<pre><code>498,4 -&gt; 498,6 -&gt; 496,6
503,4 -&gt; 502,4 -&gt; 502,9 -&gt; 494,9
</code></pre>
This scan means that there are two paths of rock; the first path consists of two straight lines, and the second path consists of three straight lines. (Specifically, the first path consists of a line of rock from <code>498,4</code> through <code>498,6</code> and another line of rock from <code>498,6</code> through <code>496,6</code>.)


The sand is pouring into the cave from point <code>500,0</code>.


Drawing rock as <code>#</code>, air as <code>.</code>, and the source of the sand as <code>+</code>, this becomes:


<pre><code>
  4     5  5
  9     0  0
  4     0  3
0 ......+...
1 ..........
2 ..........
3 ..........
4 ....#...##
5 ....#...#.
6 ..###...#.
7 ........#.
8 ........#.
9 #########.
</code></pre>
Sand is produced <em><b>one unit at a time</b></em>, and the next unit of sand is not produced until the previous unit of sand <em><b>comes to rest</b></em>. A unit of sand is large enough to fill one tile of air in your scan.


A unit of sand always falls <em><b>down one step</b></em> if possible. If the tile immediately below is blocked (by rock or sand), the unit of sand attempts to instead move diagonally <em><b>one step down and to the left</b></em>. If that tile is blocked, the unit of sand attempts to instead move diagonally <em><b>one step down and to the right</b></em>. Sand keeps moving as long as it is able to do so, at each step trying to move down, then down-left, then down-right. If all three possible destinations are blocked, the unit of sand <em><b>comes to rest</b></em> and no longer moves, at which point the next unit of sand is created back at the source.


So, drawing sand that has come to rest as <code>o</code>, the first unit of sand simply falls straight down and then stops:


<pre><code>......+...
..........
..........
..........
....#...##
....#...#.
..###...#.
........#.
......<em><b>o</b></em>.#.
#########.
</code></pre>
The second unit of sand then falls straight down, lands on the first one, and then comes to rest to its left:


<pre><code>......+...
..........
..........
..........
....#...##
....#...#.
..###...#.
........#.
.....oo.#.
#########.
</code></pre>
After a total of five units of sand have come to rest, they form this pattern:


<pre><code>......+...
..........
..........
..........
....#...##
....#...#.
..###...#.
......o.#.
....oooo#.
#########.
</code></pre>
After a total of 22 units of sand:


<pre><code>......+...
..........
......o...
.....ooo..
....#ooo##
....#ooo#.
..###ooo#.
....oooo#.
...ooooo#.
#########.
</code></pre>
Finally, only two more units of sand can possibly come to rest:


<pre><code>......+...
..........
......o...
.....ooo..
....#ooo##
...<em><b>o</b></em>#ooo#.
..###ooo#.
....oooo#.
.<em><b>o</b></em>.ooooo#.
#########.
</code></pre>
Once all <code><em><b>24</b></em></code> units of sand shown above have come to rest, all further sand flows out the bottom, falling into the endless void. Just for fun, the path any new sand takes before falling forever is shown here with <code>~</code>:


<pre><code>.......+...
.......~...
......~o...
.....~ooo..
....~#ooo##
...~o#ooo#.
..~###ooo#.
..~..oooo#.
.~o.ooooo#.
~#########.
~..........
~..........
~..........
</code></pre>
Using your scan, simulate the falling sand. <em><b>How many units of sand come to rest before sand starts flowing into the abyss below?</b></em>


