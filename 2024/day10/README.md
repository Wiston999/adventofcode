# --- Day 10: Hoof It ---

You all arrive at a [/2023/day/15](Lava Production Facility) on a floating island in the sky. As the others begin to search the massive industrial complex, you feel a small nose boop your leg and look down to discover a <span title="i knew you would come back">reindeer</span> wearing a hard hat.


The reindeer is holding a book titled "Lava Island Hiking Guide". However, when you open the book, you discover that most of it seems to have been scorched by lava! As you're about to ask how you can help, the reindeer brings you a blank [https://en.wikipedia.org/wiki/Topographic_map](topographic map) of the surrounding area (your puzzle input) and looks up at you excitedly.


Perhaps you can help fill in the missing hiking trails?


The topographic map indicates the <em><b>height</b></em> at each position using a scale from <code>0</code> (lowest) to <code>9</code> (highest). For example:


<pre><code>0123
1234
8765
9876
</code></pre>
Based on un-scorched scraps of the book, you determine that a good hiking trail is <em><b>as long as possible</b></em> and has an <em><b>even, gradual, uphill slope</b></em>. For all practical purposes, this means that a <em><b>hiking trail</b></em> is any path that starts at height <code>0</code>, ends at height <code>9</code>, and always increases by a height of exactly 1 at each step. Hiking trails never include diagonal steps - only up, down, left, or right (from the perspective of the map).


You look up from the map and notice that the reindeer has helpfully begun to construct a small pile of pencils, markers, rulers, compasses, stickers, and other equipment you might need to update the map with hiking trails.


A <em><b>trailhead</b></em> is any position that starts one or more hiking trails - here, these positions will always have height <code>0</code>. Assembling more fragments of pages, you establish that a trailhead's <em><b>score</b></em> is the number of <code>9</code>-height positions reachable from that trailhead via a hiking trail. In the above example, the single trailhead in the top left corner has a score of <code>1</code> because it can reach a single <code>9</code> (the one in the bottom left).


This trailhead has a score of <code>2</code>:


<pre><code>...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9
</code></pre>
(The positions marked <code>.</code> are impassable tiles to simplify these examples; they do not appear on your actual topographic map.)


This trailhead has a score of <code>4</code> because every <code>9</code> is reachable via a hiking trail except the one immediately to the left of the trailhead:


<pre><code>..90..9
...1.98
...2..7
6543456
765.987
876....
987....
</code></pre>
This topographic map contains <em><b>two</b></em> trailheads; the trailhead at the top has a score of <code>1</code>, while the trailhead at the bottom has a score of <code>2</code>:


<pre><code>10..9..
2...8..
3...7..
4567654
...8..3
...9..2
.....01
</code></pre>
Here's a larger example:


<pre><code>89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
</code></pre>
This larger example has 9 trailheads. Considering the trailheads in reading order, they have scores of <code>5</code>, <code>6</code>, <code>5</code>, <code>3</code>, <code>1</code>, <code>3</code>, <code>5</code>, <code>3</code>, and <code>5</code>. Adding these scores together, the sum of the scores of all trailheads is <code><em><b>36</b></em></code>.


The reindeer gleefully carries over a protractor and adds it to the pile. <em><b>What is the sum of the scores of all trailheads on your topographic map?</b></em>


