# --- Day 8: Treetop Tree House ---

The expedition comes across a peculiar patch of tall trees all planted carefully in a grid. The Elves explain that a previous expedition planted these trees as a reforestation effort. Now, they're curious if this would be a good location for a [https://en.wikipedia.org/wiki/Tree_house](tree house).


First, determine whether there is enough tree cover here to keep a tree house <em><b>hidden</b></em>. To do this, you need to count the number of trees that are <em><b>visible from outside the grid</b></em> when looking directly along a row or column.


The Elves have already launched a [https://en.wikipedia.org/wiki/Quadcopter](quadcopter) to generate a map with the height of each tree (<span title="The Elves have already launched a quadcopter (your puzzle input).">your puzzle input</span>). For example:


<pre><code>30373
25512
65332
33549
35390
</code></pre>
Each tree is represented as a single digit whose value is its height, where <code>0</code> is the shortest and <code>9</code> is the tallest.


A tree is <em><b>visible</b></em> if all of the other trees between it and an edge of the grid are <em><b>shorter</b></em> than it. Only consider trees in the same row or column; that is, only look up, down, left, or right from any given tree.


All of the trees around the edge of the grid are <em><b>visible</b></em> - since they are already on the edge, there are no trees to block the view. In this example, that only leaves the <em><b>interior nine trees</b></em> to consider:


<ul>
<li>The top-left <code>5</code> is <em><b>visible</b></em> from the left and top. (It isn't visible from the right or bottom since other trees of height <code>5</code> are in the way.)</li>
<li>The top-middle <code>5</code> is <em><b>visible</b></em> from the top and right.</li>
<li>The top-right <code>1</code> is not visible from any direction; for it to be visible, there would need to only be trees of height <em><b>0</b></em> between it and an edge.</li>
<li>The left-middle <code>5</code> is <em><b>visible</b></em>, but only from the right.</li>
<li>The center <code>3</code> is not visible from any direction; for it to be visible, there would need to be only trees of at most height <code>2</code> between it and an edge.</li>
<li>The right-middle <code>3</code> is <em><b>visible</b></em> from the right.</li>
<li>In the bottom row, the middle <code>5</code> is <em><b>visible</b></em>, but the <code>3</code> and <code>4</code> are not.</li>
</ul>
With 16 trees visible on the edge and another 5 visible in the interior, a total of <code><em><b>21</b></em></code> trees are visible in this arrangement.


Consider your map; <em><b>how many trees are visible from outside the grid?</b></em>


