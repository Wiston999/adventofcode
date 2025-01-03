# --- Day 14: Restroom Redoubt ---

One of The Historians needs to use the bathroom; fortunately, you know there's a bathroom near an unvisited location on their list, and so you're all quickly teleported directly to the lobby of Easter Bunny Headquarters.


Unfortunately, EBHQ seems to have "improved" bathroom security <em><b>again</b></em> after your last [/2016/day/2](visit). The area outside the bathroom is swarming with robots!


To get The Historian safely to the bathroom, you'll need a way to predict where the robots will be in the future. Fortunately, they all seem to be moving on the tile floor in predictable <em><b>straight lines</b></em>.


You make a list (your puzzle input) of all of the robots' current <em><b>positions</b></em> (<code>p</code>) and <em><b>velocities</b></em> (<code>v</code>), one robot per line. For example:


<pre><code>p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3
</code></pre>
Each robot's position is given as <code>p=x,y</code> where <code>x</code> represents the number of tiles the robot is from the left wall and <code>y</code> represents the number of tiles from the top wall (when viewed from above). So, a position of <code>p=0,0</code> means the robot is all the way in the top-left corner.


Each robot's velocity is given as <code>v=x,y</code> where <code>x</code> and <code>y</code> are given in <em><b>tiles per second</b></em>. Positive <code>x</code> means the robot is moving to the <em><b>right</b></em>, and positive <code>y</code> means the robot is moving <em><b>down</b></em>. So, a velocity of <code>v=1,-2</code> means that each second, the robot moves <code>1</code> tile to the right and <code>2</code> tiles up.


The robots outside the actual bathroom are in a space which is <code>101</code> tiles wide and <code>103</code> tiles tall (when viewed from above). However, in this example, the robots are in a space which is only <code>11</code> tiles wide and <code>7</code> tiles tall.


The robots are good at navigating over/under each other (due to a combination of springs, extendable legs, and quadcopters), so they can share the same tile and don't interact with each other. Visually, the number of robots on each tile in this example looks like this:


<pre><code>1.12.......
...........
...........
......11.11
1.1........
.........1.
.......1...
</code></pre>
These robots have a unique feature for maximum bathroom security: they can <em><b>teleport</b></em>. When a robot would run into an edge of the space they're in, they instead <em><b>teleport to the other side</b></em>, effectively wrapping around the edges. Here is what robot <code>p=2,4 v=2,-3</code> does for the first few seconds:


<pre><code>Initial state:
...........
...........
...........
...........
..1........
...........
...........

After 1 second:
...........
....1......
...........
...........
...........
...........
...........

After 2 seconds:
...........
...........
...........
...........
...........
......1....
...........

After 3 seconds:
...........
...........
........1..
...........
...........
...........
...........

After 4 seconds:
...........
...........
...........
...........
...........
...........
..........1

After 5 seconds:
...........
...........
...........
.1.........
...........
...........
...........
</code></pre>
The Historian can't wait much longer, so you don't have to simulate the robots for very long. Where will the robots be after <code>100</code> seconds?


In the above example, the number of robots on each tile after 100 seconds has elapsed looks like this:


<pre><code>......2..1.
...........
1..........
.11........
.....1.....
...12......
.1....1....
</code></pre>
To determine the safest area, count the <em><b>number of robots in each quadrant</b></em> after 100 seconds. Robots that are exactly in the middle (horizontally or vertically) don't count as being in any quadrant, so the only relevant robots are:


<pre><code>..... 2..1.
..... .....
1.... .....
           
..... .....
...12 .....
.1... 1....
</code></pre>
In this example, the quadrants contain <code>1</code>, <code>3</code>, <code>4</code>, and <code>1</code> robot. Multiplying these together gives a total <em><b>safety factor</b></em> of <code><em><b>12</b></em></code>.


Predict the motion of the robots in your list within a space which is <code>101</code> tiles wide and <code>103</code> tiles tall. <em><b>What will the safety factor be after exactly 100 seconds have elapsed?</b></em>


