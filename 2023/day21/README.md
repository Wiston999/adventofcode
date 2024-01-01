# --- Day 21: Step Counter ---

You manage to catch the [7](airship) right as it's dropping someone else off on their all-expenses-paid trip to Desert Island! It even helpfully drops you off near the [5](gardener) and his massive farm.


"You got the sand flowing again! Great work! Now we just need to wait until we have enough sand to filter the water for Snow Island and we'll have snow again in no time."


While you wait, one of the Elves that works with the gardener heard how good you are at solving problems and would like your help. He needs to get his [https://en.wikipedia.org/wiki/Pedometer](steps) in for the day, and so he'd like to know <em><b>which garden plots he can reach with exactly his remaining <code>64</code> steps</b></em>.


He gives you an up-to-date map (your puzzle input) of his starting position (<code>S</code>), garden plots (<code>.</code>), and rocks (<code>#</code>). For example:


<pre><code>...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
</code></pre>
The Elf starts at the starting position (<code>S</code>) which also counts as a garden plot. Then, he can take one step north, south, east, or west, but only onto tiles that are garden plots. This would allow him to reach any of the tiles marked <code>O</code>:


<pre><code>...........
.....###.#.
.###.##..#.
..#.#...#..
....#O#....
.##.OS####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
</code></pre>
Then, he takes a second step. Since at this point he could be at <em><b>either</b></em> tile marked <code>O</code>, his second step would allow him to reach any garden plot that is one step north, south, east, or west of <em><b>any</b></em> tile that he could have reached after the first step:


<pre><code>...........
.....###.#.
.###.##..#.
..#.#O..#..
....#.#....
.##O.O####.
.##.O#...#.
.......##..
.##.#.####.
.##..##.##.
...........
</code></pre>
After two steps, he could be at any of the tiles marked <code>O</code> above, including the starting position (either by going north-then-south or by going west-then-east).


A single third step leads to even more possibilities:


<pre><code>...........
.....###.#.
.###.##..#.
..#.#.O.#..
...O#O#....
.##.OS####.
.##O.#...#.
....O..##..
.##.#.####.
.##..##.##.
...........
</code></pre>
He will continue like this until his steps for the day have been exhausted. After a total of <code>6</code> steps, he could reach any of the garden plots marked <code>O</code>:


<pre><code>...........
.....###.#.
.###.##.O#.
.O#O#O.O#..
O.O.#.#.O..
.##O.O####.
.##.O#O..#.
.O.O.O.##..
.##.#.####.
.##O.##.##.
...........
</code></pre>
In this example, if the Elf's goal was to get exactly <code>6</code> more steps today, he could use them to reach any of <code><em><b>16</b></em></code> garden plots.


However, the Elf <em><b>actually needs to get <code>64</code> steps today</b></em>, and the map he's handed you is much larger than the example map.


Starting from the garden plot marked <code>S</code> on your map, <em><b>how many garden plots could the Elf reach in exactly <code>64</code> steps?</b></em>


