# --- Day 15: Chiton ---

You've almost reached the exit of the cave, but the walls are getting closer together. Your submarine can barely still fit, though; the main problem is that the walls of the cave are covered in [https://en.wikipedia.org/wiki/Chiton](chitons), and it would be best not to bump any of them.


The cavern is large, but has a very low ceiling, restricting your motion to two dimensions. The shape of the cavern resembles a square; a quick scan of chiton density produces a map of <em><b>risk level</b></em> throughout the cave (your puzzle input). For example:


<pre><code>1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581
</code></pre>
You start in the top left position, your destination is the bottom right position, and you <span title="Can't go diagonal until we can repair the caterpillar unit. Could be the liquid helium or the superconductors.">cannot move diagonally</span>. The number at each position is its <em><b>risk level</b></em>; to determine the total risk of an entire path, add up the risk levels of each position you <em><b>enter</b></em> (that is, don't count the risk level of your starting position unless you enter it; leaving it adds no risk to your total).


Your goal is to find a path with the <em><b>lowest total risk</b></em>. In this example, a path with the lowest total risk is highlighted here:


<pre><code><em><b>1</b></em>163751742
<em><b>1</b></em>381373672
<em><b>2136511</b></em>328
369493<em><b>15</b></em>69
7463417<em><b>1</b></em>11
1319128<em><b>13</b></em>7
13599124<em><b>2</b></em>1
31254216<em><b>3</b></em>9
12931385<em><b>21</b></em>
231194458<em><b>1</b></em>
</code></pre>
The total risk of this path is <code><em><b>40</b></em></code> (the starting position is never entered, so its risk is not counted).


<em><b>What is the lowest total risk of any path from the top left to the bottom right?</b></em>


