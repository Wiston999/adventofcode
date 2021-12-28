# --- Day 5: Hydrothermal Venture ---

You come across a field of [https://en.wikipedia.org/wiki/Hydrothermal_vent](hydrothermal vents) on the ocean floor! These vents constantly produce large, opaque clouds, so it would be best to avoid them if possible.


They tend to form in <em><b>lines</b></em>; the submarine helpfully produces a list of nearby <span title="Maybe they're Bresenham vents.">lines of vents</span> (your puzzle input) for you to review. For example:


<pre><code>0,9 -&gt; 5,9
8,0 -&gt; 0,8
9,4 -&gt; 3,4
2,2 -&gt; 2,1
7,0 -&gt; 7,4
6,4 -&gt; 2,0
0,9 -&gt; 2,9
3,4 -&gt; 1,4
0,0 -&gt; 8,8
5,5 -&gt; 8,2
</code></pre>
Each line of vents is given as a line segment in the format <code>x1,y1 -&gt; x2,y2</code> where <code>x1</code>,<code>y1</code> are the coordinates of one end the line segment and <code>x2</code>,<code>y2</code> are the coordinates of the other end. These line segments include the points at both ends. In other words:


<ul>
<li>An entry like <code>1,1 -&gt; 1,3</code> covers points <code>1,1</code>, <code>1,2</code>, and <code>1,3</code>.</li>
<li>An entry like <code>9,7 -&gt; 7,7</code> covers points <code>9,7</code>, <code>8,7</code>, and <code>7,7</code>.</li>
</ul>
For now, <em><b>only consider horizontal and vertical lines</b></em>: lines where either <code>x1 = x2</code> or <code>y1 = y2</code>.


So, the horizontal and vertical lines from the above list would produce the following diagram:


<pre><code>.......1..
..1....1..
..1....1..
.......1..
.112111211
..........
..........
..........
..........
222111....
</code></pre>
In this diagram, the top left corner is <code>0,0</code> and the bottom right corner is <code>9,9</code>. Each position is shown as <em><b>the number of lines which cover that point</b></em> or <code>.</code> if no line covers that point. The top-left pair of <code>1</code>s, for example, comes from <code>2,2 -&gt; 2,1</code>; the very bottom row is formed by the overlapping lines <code>0,9 -&gt; 5,9</code> and <code>0,9 -&gt; 2,9</code>.


To avoid the most dangerous areas, you need to determine <em><b>the number of points where at least two lines overlap</b></em>. In the above example, this is anywhere in the diagram with a <code>2</code> or larger - a total of <code><em><b>5</b></em></code> points.


Consider only horizontal and vertical lines. <em><b>At how many points do at least two lines overlap?</b></em>


