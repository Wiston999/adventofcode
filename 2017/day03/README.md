# --- Day 3: Spiral Memory ---

You come across an experimental new kind of memory stored on an <span title="Good thing we have all these infinite two-dimensional grids lying around!">infinite two-dimensional grid</span>.


Each square on the grid is allocated in a spiral pattern starting at a location marked <code>1</code> and then counting up while spiraling outward. For example, the first few squares are allocated like this:


<pre><code>17  16  15  14  13
18   5   4   3  12
19   6   1   2  11
20   7   8   9  10
21  22  23---&gt; ...
</code></pre>
While this is very space-efficient (no squares are skipped), requested data must be carried back to square <code>1</code> (the location of the only access port for this memory system) by programs that can only move up, down, left, or right. They always take the shortest path: the [https://en.wikipedia.org/wiki/Taxicab_geometry](Manhattan Distance) between the location of the data and square <code>1</code>.


For example:


<ul>
<li>Data from square <code>1</code> is carried <code>0</code> steps, since it's at the access port.</li>
<li>Data from square <code>12</code> is carried <code>3</code> steps, such as: down, left, left.</li>
<li>Data from square <code>23</code> is carried only <code>2</code> steps: up twice.</li>
<li>Data from square <code>1024</code> must be carried <code>31</code> steps.</li>
</ul>
<em><b>How many steps</b></em> are required to carry the data from the square identified in your puzzle input all the way to the access port?


