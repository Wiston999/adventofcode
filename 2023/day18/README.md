# --- Day 18: Lavaduct Lagoon ---

Thanks to your efforts, the machine parts factory is one of the first factories up and running since the lavafall came back. However, to catch up with the large backlog of parts requests, the factory will also need a <em><b>large supply of lava</b></em> for a while; the Elves have already started creating a large lagoon nearby for this purpose.


However, they aren't sure the lagoon will be big enough; they've asked you to take a look at the <em><b>dig plan</b></em> (your puzzle input). For example:


<pre><code>R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)
</code></pre>
The digger starts in a 1 meter cube hole in the ground. They then dig the specified number of meters <em><b>up</b></em> (<code>U</code>), <em><b>down</b></em> (<code>D</code>), <em><b>left</b></em> (<code>L</code>), or <em><b>right</b></em> (<code>R</code>), clearing full 1 meter cubes as they go. The directions are given as seen from above, so if "up" were north, then "right" would be east, and so on. Each trench is also listed with <em><b>the color that the edge of the trench should be painted</b></em> as an [https://en.wikipedia.org/wiki/RGB_color_model#Numeric_representations](RGB hexadecimal color code).


When viewed from above, the above example dig plan would result in the following loop of <em><b>trench</b></em> (<code>#</code>) having been dug out from otherwise <em><b>ground-level terrain</b></em> (<code>.</code>):


<pre><code>#######
#.....#
###...#
..#...#
..#...#
###.###
#...#..
##..###
.#....#
.######
</code></pre>
At this point, the trench could contain 38 cubic meters of lava. However, this is just the edge of the lagoon; the next step is to <em><b>dig out the interior</b></em> so that it is one meter deep as well:


<pre><code>#######
#######
#######
..#####
..#####
#######
#####..
#######
.######
.######
</code></pre>
Now, the lagoon can contain a much more respectable <code><em><b>62</b></em></code> cubic meters of lava. While the interior is dug out, the edges are also painted according to the color codes in the dig plan.


The Elves are concerned the lagoon won't be large enough; if they follow their dig plan, <em><b>how many cubic meters of lava could it hold?</b></em>


