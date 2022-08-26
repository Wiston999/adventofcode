# --- Day 22: Sporifica Virus ---

Diagnostics indicate that the local <em><b>grid computing cluster</b></em> has been contaminated with the <em><b>Sporifica Virus</b></em>. The grid computing cluster is a seemingly-<span title="The infinite is possible at AdventOfCodeCom.">infinite</span> two-dimensional grid of compute nodes.  Each node is either <em><b>clean</b></em> or <em><b>infected</b></em> by the virus.<p>
<p>To [https://en.wikipedia.org/wiki/Morris_worm#The_mistake](prevent overloading) the nodes (which would render them useless to the virus) or detection by system administrators, exactly one <em><b>virus carrier</b></em> moves through the network, infecting or cleaning nodes as it moves. The virus carrier is always located on a single node in the network (the <em><b>current node</b></em>) and keeps track of the <em><b>direction</b></em> it is facing.


To avoid detection, the virus carrier works in bursts; in each burst, it <em><b>wakes up</b></em>, does some <em><b>work</b></em>, and goes back to <em><b>sleep</b></em>. The following steps are all executed <em><b>in order</b></em> one time each burst:


<ul>
<li>If the <em><b>current node</b></em> is <em><b>infected</b></em>, it turns to its <em><b>right</b></em>.  Otherwise, it turns to its <em><b>left</b></em>. (Turning is done in-place; the <em><b>current node</b></em> does not change.)</li>
<li>If the <em><b>current node</b></em> is <em><b>clean</b></em>, it becomes <em><b>infected</b></em>.  Otherwise, it becomes <em><b>cleaned</b></em>. (This is done <em><b>after</b></em> the node is considered for the purposes of changing direction.)</li>
<li>The virus carrier [https://www.youtube.com/watch?v=2vj37yeQQHg](moves) <em><b>forward</b></em> one node in the direction it is facing.</li>
</ul>
Diagnostics have also provided a <em><b>map of the node infection status</b></em> (your puzzle input).  <em><b>Clean</b></em> nodes are shown as <code>.</code>; <em><b>infected</b></em> nodes are shown as <code>#</code>.  This map only shows the center of the grid; there are many more nodes beyond those shown, but none of them are currently infected.


The virus carrier begins in the middle of the map facing <em><b>up</b></em>.


For example, suppose you are given a map like this:


<pre><code>..#
#..
...
</code></pre>
Then, the middle of the infinite grid looks like this, with the virus carrier's position marked with <code>[ ]</code>:


<pre><code>. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . # . . .
. . . #[.]. . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
</code></pre>
The virus carrier is on a <em><b>clean</b></em> node, so it turns <em><b>left</b></em>, <em><b>infects</b></em> the node, and moves left:


<pre><code>. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . # . . .
. . .[#]# . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
</code></pre>
The virus carrier is on an <em><b>infected</b></em> node, so it turns <em><b>right</b></em>, <em><b>cleans</b></em> the node, and moves up:


<pre><code>. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . .[.]. # . . .
. . . . # . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
</code></pre>
Four times in a row, the virus carrier finds a <em><b>clean</b></em>, <em><b>infects</b></em> it, turns <em><b>left</b></em>, and moves forward, ending in the same place and still facing up:


<pre><code>. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . #[#]. # . . .
. . # # # . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
</code></pre>
Now on the same node as before, it sees an infection, which causes it to turn <em><b>right</b></em>, <em><b>clean</b></em> the node, and move forward:


<pre><code>. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . # .[.]# . . .
. . # # # . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
</code></pre>
After the above actions, a total of <code>7</code> bursts of activity had taken place. Of them, <code>5</code> bursts of activity caused an infection.


After a total of <code>70</code>, the grid looks like this, with the virus carrier facing up:


<pre><code>. . . . . # # . .
. . . . # . . # .
. . . # . . . . #
. . # . #[.]. . #
. . # . # . . # .
. . . . . # # . .
. . . . . . . . .
. . . . . . . . .
</code></pre>
By this time, <code>41</code> bursts of activity caused an infection (though most of those nodes have since been cleaned).


After a total of <code>10000</code> bursts of activity, <code>5587</code> bursts will have caused an infection.


Given your actual map, after <code>10000</code> bursts of activity, <em><b>how many bursts cause a node to become infected</b></em>? (Do not count nodes that begin infected.)


