# --- Day 20: Race Condition ---

The Historians are quite pixelated again. This time, a massive, black building looms over you - you're [/2017/day/24](right outside) the CPU!


While The Historians get to work, a nearby program sees that you're idle and challenges you to a <em><b>race</b></em>. Apparently, you've arrived just in time for the frequently-held <em><b>race condition</b></em> festival!


The race takes place on a particularly long and twisting code path; programs compete to see who can finish in the <em><b>fewest picoseconds</b></em>. The <span title="If we give away enough mutexes, maybe someone will use one of them to fix the race condition!">winner</span> even gets their very own [https://en.wikipedia.org/wiki/Lock_(computer_science)](mutex)!


They hand you a <em><b>map of the racetrack</b></em> (your puzzle input). For example:


<pre><code>###############
#...#...#.....#
#.#.#.#.#.###.#
#<em><b>S</b></em>#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..<em><b>E</b></em>#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############
</code></pre>
The map consists of track (<code>.</code>) - including the <em><b>start</b></em> (<code>S</code>) and <em><b>end</b></em> (<code>E</code>) positions (both of which also count as track) - and <em><b>walls</b></em> (<code>#</code>).


When a program runs through the racetrack, it starts at the start position. Then, it is allowed to move up, down, left, or right; each such move takes <em><b>1 picosecond</b></em>. The goal is to reach the end position as quickly as possible. In this example racetrack, the fastest time is <code>84</code> picoseconds.


Because there is only a single path from the start to the end and the programs all go the same speed, the races used to be pretty boring. To make things more interesting, they introduced a new rule to the races: programs are allowed to <em><b>cheat</b></em>.


The rules for cheating are very strict. <em><b>Exactly once</b></em> during a race, a program may <em><b>disable collision</b></em> for up to <em><b>2 picoseconds</b></em>. This allows the program to <em><b>pass through walls</b></em> as if they were regular track. At the end of the cheat, the program must be back on normal track again; otherwise, it will receive a [https://en.wikipedia.org/wiki/Segmentation_fault](segmentation fault) and get disqualified.


So, a program could complete the course in 72 picoseconds (saving <em><b>12 picoseconds</b></em>) by cheating for the two moves marked <code>1</code> and <code>2</code>:


<pre><code>###############
#...#...12....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############
</code></pre>
Or, a program could complete the course in 64 picoseconds (saving <em><b>20 picoseconds</b></em>) by cheating for the two moves marked <code>1</code> and <code>2</code>:


<pre><code>###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...12..#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############
</code></pre>
This cheat saves <em><b>38 picoseconds</b></em>:


<pre><code>###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.####1##.###
#...###.2.#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############
</code></pre>
This cheat saves <em><b>64 picoseconds</b></em> and takes the program directly to the end:


<pre><code>###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..21...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############
</code></pre>
Each cheat has a distinct <em><b>start position</b></em> (the position where the cheat is activated, just before the first move that is allowed to go through walls) and <em><b>end position</b></em>; cheats are uniquely identified by their start position and end position.


In this example, the total number of cheats (grouped by the amount of time they save) are as follows:


<ul>
<li>There are 14 cheats that save 2 picoseconds.</li>
<li>There are 14 cheats that save 4 picoseconds.</li>
<li>There are 2 cheats that save 6 picoseconds.</li>
<li>There are 4 cheats that save 8 picoseconds.</li>
<li>There are 2 cheats that save 10 picoseconds.</li>
<li>There are 3 cheats that save 12 picoseconds.</li>
<li>There is one cheat that saves 20 picoseconds.</li>
<li>There is one cheat that saves 36 picoseconds.</li>
<li>There is one cheat that saves 38 picoseconds.</li>
<li>There is one cheat that saves 40 picoseconds.</li>
<li>There is one cheat that saves 64 picoseconds.</li>
</ul>
You aren't sure what the conditions of the racetrack will be like, so to give yourself as many options as possible, you'll need a list of the best cheats. <em><b>How many cheats would save you at least 100 picoseconds?</b></em>


