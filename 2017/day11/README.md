# --- Day 11: Hex Ed ---

Crossing the bridge, you've barely reached the other side of the stream when a program comes up to you, clearly in distress.  "It's my child process," she says, "he's gotten lost in an infinite grid!"


Fortunately for her, you have plenty of experience with infinite grids.


Unfortunately for you, it's a [https://en.wikipedia.org/wiki/Hexagonal_tiling](hex grid).


The hexagons ("hexes") in <span title="Raindrops on roses and whiskers on kittens.">this grid</span> are aligned such that adjacent hexes can be found to the north, northeast, southeast, south, southwest, and northwest:


<pre><code>  \ n  /
nw +--+ ne
  /    \
-+      +-
  \    /
sw +--+ se
  / s  \
</code></pre>
You have the path the child process took. Starting where he started, you need to determine the fewest number of steps required to reach him. (A "step" means to move from the hex you are in to any adjacent hex.)


For example:


<ul>
<li><code>ne,ne,ne</code> is <code>3</code> steps away.</li>
<li><code>ne,ne,sw,sw</code> is <code>0</code> steps away (back where you started).</li>
<li><code>ne,ne,s,s</code> is <code>2</code> steps away (<code>se,se</code>).</li>
<li><code>se,sw,se,sw,sw</code> is <code>3</code> steps away (<code>s,s,sw</code>).</li>
</ul>
