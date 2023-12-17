# --- Day 17: Clumsy Crucible ---

The lava starts flowing rapidly once the Lava Production Facility is operational. As you <span title="see you soon?">leave</span>, the reindeer offers you a parachute, allowing you to quickly reach Gear Island.


As you descend, your bird's-eye view of Gear Island reveals why you had trouble finding anyone on your way up: half of Gear Island is empty, but the half below you is a giant factory city!


You land near the gradually-filling pool of lava at the base of your new <em><b>lavafall</b></em>. Lavaducts will eventually carry the lava throughout the city, but to make use of it immediately, Elves are loading it into large [https://en.wikipedia.org/wiki/Crucible](crucibles) on wheels.


The crucibles are top-heavy and pushed by hand. Unfortunately, the crucibles become very difficult to steer at high speeds, and so it can be hard to go in a straight line for very long.


To get Desert Island the machine parts it needs as soon as possible, you'll need to find the best way to get the crucible <em><b>from the lava pool to the machine parts factory</b></em>. To do this, you need to minimize <em><b>heat loss</b></em> while choosing a route that doesn't require the crucible to go in a <em><b>straight line</b></em> for too long.


Fortunately, the Elves here have a map (your puzzle input) that uses traffic patterns, ambient temperature, and hundreds of other parameters to calculate exactly how much heat loss can be expected for a crucible entering any particular city block.


For example:


<pre><code>2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533
</code></pre>
Each city block is marked by a single digit that represents the <em><b>amount of heat loss if the crucible enters that block</b></em>. The starting point, the lava pool, is the top-left city block; the destination, the machine parts factory, is the bottom-right city block. (Because you already start in the top-left block, you don't incur that block's heat loss unless you leave that block and then return to it.)


Because it is difficult to keep the top-heavy crucible going in a straight line for very long, it can move <em><b>at most three blocks</b></em> in a single direction before it must turn 90 degrees left or right. The crucible also can't reverse direction; after entering each city block, it may only turn left, continue straight, or turn right.


One way to <em><b>minimize heat loss</b></em> is this path:


<pre><code>2<em><b>&gt;</b></em><em><b>&gt;</b></em>34<em><b>^</b></em><em><b>&gt;</b></em><em><b>&gt;</b></em><em><b>&gt;</b></em>1323
32<em><b>v</b></em><em><b>&gt;</b></em><em><b>&gt;</b></em><em><b>&gt;</b></em>35<em><b>v</b></em>5623
32552456<em><b>v</b></em><em><b>&gt;</b></em><em><b>&gt;</b></em>54
3446585845<em><b>v</b></em>52
4546657867<em><b>v</b></em><em><b>&gt;</b></em>6
14385987984<em><b>v</b></em>4
44578769877<em><b>v</b></em>6
36378779796<em><b>v</b></em><em><b>&gt;</b></em>
465496798688<em><b>v</b></em>
456467998645<em><b>v</b></em>
12246868655<em><b>&lt;</b></em><em><b>v</b></em>
25465488877<em><b>v</b></em>5
43226746555<em><b>v</b></em><em><b>&gt;</b></em>
</code></pre>
This path never moves more than three consecutive blocks in the same direction and incurs a heat loss of only <code><em><b>102</b></em></code>.


Directing the crucible from the lava pool to the machine parts factory, but not moving more than three consecutive blocks in the same direction, <em><b>what is the least heat loss it can incur?</b></em>


