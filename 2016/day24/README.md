# --- Day 24: Air Duct Spelunking ---

You've finally met your match; the doors that provide access to the roof are locked tight, and all of the controls and related electronics are inaccessible. You simply can't reach them.


The robot that cleans the air ducts, however, <em><b>can</b></em>.


It's not a very fast <span title="The Brave Little Air Duct Cleaning Robot That Could">little robot</span>, but you reconfigure it to be able to interface with some of the exposed wires that have been routed through the [https://en.wikipedia.org/wiki/HVAC](HVAC) system. If you can direct it to each of those locations, you should be able to bypass the security controls.


You extract the duct layout for this area from some blueprints you acquired and create a map with the relevant locations marked (your puzzle input). <code>0</code> is your current location, from which the cleaning robot embarks; the other numbers are (in <em><b>no particular order</b></em>) the locations the robot needs to visit at least once each. Walls are marked as <code>#</code>, and open passages are marked as <code>.</code>. Numbers behave like open passages.


For example, suppose you have a map like the following:


<pre><code>###########
#0.1.....2#
#.#######.#
#4.......3#
###########
</code></pre>
To reach all of the points of interest as quickly as possible, you would have the robot take the following path:


<ul>
<li><code>0</code> to <code>4</code> (<code>2</code> steps)</li>
<li><code>4</code> to <code>1</code> (<code>4</code> steps; it can't move diagonally)</li>
<li><code>1</code> to <code>2</code> (<code>6</code> steps)</li>
<li><code>2</code> to <code>3</code> (<code>2</code> steps)</li>
</ul>
Since the robot isn't very fast, you need to find it the <em><b>shortest route</b></em>. This path is the fewest steps (in the above example, a total of <code>14</code>) required to start at <code>0</code> and then visit every other location at least once.


Given your actual map, and starting from location <code>0</code>, what is the <em><b>fewest number of steps</b></em> required to visit every non-<code>0</code> number marked on the map at least once?


