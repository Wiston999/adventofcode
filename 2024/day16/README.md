# --- Day 16: Reindeer Maze ---

It's time again for the [/2015/day/14](Reindeer Olympics)! This year, the big event is the <em><b>Reindeer Maze</b></em>, where the Reindeer compete for the <em><b><span title="I would say it's like Reindeer Golf, but knowing Reindeer, it's almost certainly nothing like Reindeer Golf.">lowest score</span></b></em>.


You and The Historians arrive to search for the Chief right as the event is about to start. It wouldn't hurt to watch a little, right?


The Reindeer start on the Start Tile (marked <code>S</code>) facing <em><b>East</b></em> and need to reach the End Tile (marked <code>E</code>). They can move forward one tile at a time (increasing their score by <code>1</code> point), but never into a wall (<code>#</code>). They can also rotate clockwise or counterclockwise 90 degrees at a time (increasing their score by <code>1000</code> points).


To figure out the best place to sit, you start by grabbing a map (your puzzle input) from a nearby kiosk. For example:


<pre><code>###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############
</code></pre>
There are many paths through this maze, but taking any of the best paths would incur a score of only <code><em><b>7036</b></em></code>. This can be achieved by taking a total of <code>36</code> steps forward and turning 90 degrees a total of <code>7</code> times:


<pre><code>
###############
#.......#....<em><b>E</b></em>#
#.#.###.#.###<em><b>^</b></em>#
#.....#.#...#<em><b>^</b></em>#
#.###.#####.#<em><b>^</b></em>#
#.#.#.......#<em><b>^</b></em>#
#.#.#####.###<em><b>^</b></em>#
#..<em><b>&gt;</b></em><em><b>&gt;</b></em><em><b>&gt;</b></em><em><b>&gt;</b></em><em><b>&gt;</b></em><em><b>&gt;</b></em><em><b>&gt;</b></em><em><b>&gt;</b></em><em><b>v</b></em>#<em><b>^</b></em>#
###<em><b>^</b></em>#.#####<em><b>v</b></em>#<em><b>^</b></em>#
#<em><b>&gt;</b></em><em><b>&gt;</b></em><em><b>^</b></em>#.....#<em><b>v</b></em>#<em><b>^</b></em>#
#<em><b>^</b></em>#.#.###.#<em><b>v</b></em>#<em><b>^</b></em>#
#<em><b>^</b></em>....#...#<em><b>v</b></em>#<em><b>^</b></em>#
#<em><b>^</b></em>###.#.#.#<em><b>v</b></em>#<em><b>^</b></em>#
#S..#.....#<em><b>&gt;</b></em><em><b>&gt;</b></em><em><b>^</b></em>#
###############
</code></pre>
Here's a second example:


<pre><code>#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################
</code></pre>
In this maze, the best paths cost <code><em><b>11048</b></em></code> points; following one such path would look like this:


<pre><code>#################
#...#...#...#..<em><b>E</b></em>#
#.#.#.#.#.#.#.#<em><b>^</b></em>#
#.#.#.#...#...#<em><b>^</b></em>#
#.#.#.#.###.#.#<em><b>^</b></em>#
#<em><b>&gt;</b></em><em><b>&gt;</b></em><em><b>v</b></em>#.#.#.....#<em><b>^</b></em>#
#<em><b>^</b></em>#<em><b>v</b></em>#.#.#.#####<em><b>^</b></em>#
#<em><b>^</b></em>#<em><b>v</b></em>..#.#.#<em><b>&gt;</b></em><em><b>&gt;</b></em><em><b>&gt;</b></em><em><b>&gt;</b></em><em><b>^</b></em>#
#<em><b>^</b></em>#<em><b>v</b></em>#####.#<em><b>^</b></em>###.#
#<em><b>^</b></em>#<em><b>v</b></em>#..<em><b>&gt;</b></em><em><b>&gt;</b></em><em><b>&gt;</b></em><em><b>&gt;</b></em><em><b>^</b></em>#...#
#<em><b>^</b></em>#<em><b>v</b></em>###<em><b>^</b></em>#####.###
#<em><b>^</b></em>#<em><b>v</b></em>#<em><b>&gt;</b></em><em><b>&gt;</b></em><em><b>^</b></em>#.....#.#
#<em><b>^</b></em>#<em><b>v</b></em>#<em><b>^</b></em>#####.###.#
#<em><b>^</b></em>#<em><b>v</b></em>#<em><b>^</b></em>........#.#
#<em><b>^</b></em>#<em><b>v</b></em>#<em><b>^</b></em>#########.#
#S#<em><b>&gt;</b></em><em><b>&gt;</b></em><em><b>^</b></em>..........#
#################
</code></pre>
Note that the path shown above includes one 90 degree turn as the very first move, rotating the Reindeer from facing East to facing North.


Analyze your map carefully. <em><b>What is the lowest score a Reindeer could possibly get?</b></em>


