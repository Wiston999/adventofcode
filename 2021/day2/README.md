# --- Day 2: Dive! ---

Now, you need to figure out how to <span title="Tank, I need a pilot program for a B212 helicopter.">pilot this thing</span>.


It seems like the submarine can take a series of commands like <code>forward 1</code>, <code>down 2</code>, or <code>up 3</code>:


<ul>
<li><code>forward X</code> increases the horizontal position by <code>X</code> units.</li>
<li><code>down X</code> <em><b>increases</b></em> the depth by <code>X</code> units.</li>
<li><code>up X</code> <em><b>decreases</b></em> the depth by <code>X</code> units.</li>
</ul>
Note that since you're on a submarine, <code>down</code> and <code>up</code> affect your <em><b>depth</b></em>, and so they have the opposite result of what you might expect.


The submarine seems to already have a planned course (your puzzle input). You should probably figure out where it's going. For example:


<pre><code>forward 5
down 5
forward 8
up 3
down 8
forward 2
</code></pre>
Your horizontal position and depth both start at <code>0</code>. The steps above would then modify them as follows:


<ul>
<li><code>forward 5</code> adds <code>5</code> to your horizontal position, a total of <code>5</code>.</li>
<li><code>down 5</code> adds <code>5</code> to your depth, resulting in a value of <code>5</code>.</li>
<li><code>forward 8</code> adds <code>8</code> to your horizontal position, a total of <code>13</code>.</li>
<li><code>up 3</code> decreases your depth by <code>3</code>, resulting in a value of <code>2</code>.</li>
<li><code>down 8</code> adds <code>8</code> to your depth, resulting in a value of <code>10</code>.</li>
<li><code>forward 2</code> adds <code>2</code> to your horizontal position, a total of <code>15</code>.</li>
</ul>
After following these instructions, you would have a horizontal position of <code>15</code> and a depth of <code>10</code>. (Multiplying these together produces <code><em><b>150</b></em></code>.)


Calculate the horizontal position and depth you would have after following the planned course. <em><b>What do you get if you multiply your final horizontal position by your final depth?</b></em>


