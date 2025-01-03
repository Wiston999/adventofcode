# --- Day 23: Unstable Diffusion ---

You enter a large crater of gray dirt where the grove is supposed to be. All around you, plants you imagine were expected to be full of fruit are instead withered and broken. A large group of Elves has formed in the middle of the grove.


"...but this volcano has been dormant for months. Without ash, the fruit can't grow!"


You look up to see a massive, snow-capped mountain towering above you.


"It's not like there are other active volcanoes here; we've looked everywhere."


"But our scanners show active magma flows; clearly it's going <em><b>somewhere</b></em>."


They finally notice you at the edge of the grove, your pack almost overflowing from the random <em class="star">star</em> fruit you've been collecting. Behind you, elephants and monkeys explore the grove, looking concerned. Then, the Elves recognize the ash cloud slowly spreading above your recent detour.


"Why do you--" "How is--" "Did you just--"


Before any of them can form a complete question, another Elf speaks up: "Okay, new plan. We have almost enough fruit already, and ash from the plume should spread here eventually. If we quickly plant new seedlings now, we can still make it to the extraction point. Spread out!"


The Elves each reach into their pack and pull out a tiny plant. The plants rely on important nutrients from the ash, so they can't be planted too close together.


There isn't enough time to let the Elves figure out where to plant the seedlings themselves; you quickly scan the grove (your puzzle input) and note their positions.


For example:


<pre><code>....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..
</code></pre>
The scan shows Elves <code>#</code> and empty ground <code>.</code>; outside your scan, more empty ground extends a long way in every direction. The scan is oriented so that <em><b>north is up</b></em>; orthogonal directions are written N (north), S (south), W (west), and E (east), while diagonal directions are written NE, NW, SE, SW.


The Elves follow a time-consuming process to figure out where they should each go; you can speed up this process considerably. The process consists of some number of <em><b>rounds</b></em> during which Elves alternate between considering where to move and actually moving.


During the <em><b>first half</b></em> of each round, each Elf considers the eight positions adjacent to themself. If no other Elves are in one of those eight positions, the Elf <em><b>does not do anything</b></em> during this round. Otherwise, the Elf looks in each of four directions in the following order and <em><b>proposes</b></em> moving one step in the <em><b>first valid direction</b></em>:


<ul>
<li>If there is no Elf in the N, NE, or NW adjacent positions, the Elf proposes moving <em><b>north</b></em> one step.</li>
<li>If there is no Elf in the S, SE, or SW adjacent positions, the Elf proposes moving <em><b>south</b></em> one step.</li>
<li>If there is no Elf in the W, NW, or SW adjacent positions, the Elf proposes moving <em><b>west</b></em> one step.</li>
<li>If there is no Elf in the E, NE, or SE adjacent positions, the Elf proposes moving <em><b>east</b></em> one step.</li>
</ul>
After each Elf has had a chance to propose a move, the <em><b>second half</b></em> of the round can begin. Simultaneously, each Elf moves to their proposed destination tile if they were the <em><b>only</b></em> Elf to propose moving to that position. If two or more Elves propose moving to the same position, <em><b>none</b></em> of those Elves move.


Finally, at the end of the round, the <em><b>first direction</b></em> the Elves considered is moved to the end of the list of directions. For example, during the second round, the Elves would try proposing a move to the south first, then west, then east, then north. On the third round, the Elves would first consider west, then east, then north, then south.


As a smaller example, consider just these five Elves:


<pre><code>.....
..##.
..#..
.....
..##.
.....
</code></pre>
The northernmost two Elves and southernmost two Elves all propose moving north, while the middle Elf cannot move north and proposes moving south. The middle Elf proposes the same destination as the southwest Elf, so neither of them move, but the other three do:


<pre><code>..##.
.....
..#..
...#.
..#..
.....
</code></pre>
Next, the northernmost two Elves and the southernmost Elf all propose moving south. Of the remaining middle two Elves, the west one cannot move south and proposes moving west, while the east one cannot move south <em><b>or</b></em> west and proposes moving east. All five Elves succeed in moving to their proposed positions:


<pre><code>.....
..##.
.#...
....#
.....
..#..
</code></pre>
Finally, the southernmost two Elves choose not to move at all. Of the remaining three Elves, the west one proposes moving west, the east one proposes moving east, and the middle one proposes moving north; all three succeed in moving:


<pre><code>..#..
....#
#....
....#
.....
..#..
</code></pre>
At this point, no Elves need to move, and so the process ends.


The larger example above proceeds as follows:


<pre><code>== Initial State ==
..............
..............
.......#......
.....###.#....
...#...#.#....
....#...##....
...#.###......
...##.#.##....
....#..#......
..............
..............
..............

== End of Round 1 ==
..............
.......#......
.....#...#....
...#..#.#.....
.......#..#...
....#.#.##....
..#..#.#......
..#.#.#.##....
..............
....#..#......
..............
..............

== End of Round 2 ==
..............
.......#......
....#.....#...
...#..#.#.....
.......#...#..
...#..#.#.....
.#...#.#.#....
..............
..#.#.#.##....
....#..#......
..............
..............

== End of Round 3 ==
..............
.......#......
.....#....#...
..#..#...#....
.......#...#..
...#..#.#.....
.#..#.....#...
.......##.....
..##.#....#...
...#..........
.......#......
..............

== End of Round 4 ==
..............
.......#......
......#....#..
..#...##......
...#.....#.#..
.........#....
.#...###..#...
..#......#....
....##....#...
....#.........
.......#......
..............

== End of Round 5 ==
.......#......
..............
..#..#.....#..
.........#....
......##...#..
.#.#.####.....
...........#..
....##..#.....
..#...........
..........#...
....#..#......
..............
</code></pre>
After a few more rounds...


<pre><code>== End of Round 10 ==
.......#......
...........#..
..#.#..#......
......#.......
...#.....#..#.
.#......##....
.....##.......
..#........#..
....#.#..#....
..............
....#..#..#...
..............
</code></pre>
To make sure they're on the right track, the Elves like to check after round 10 that they're making good progress toward covering enough ground. To do this, count the number of empty ground tiles contained by the smallest rectangle that contains every Elf. (The edges of the rectangle should be aligned to the N/S/E/W directions; the Elves do not have the patience to calculate <span title="Arbitrary Rectangles is my Piet Mondrian cover band.">arbitrary rectangles</span>.) In the above example, that rectangle is:


<pre><code>......#.....
..........#.
.#.#..#.....
.....#......
..#.....#..#
#......##...
....##......
.#........#.
...#.#..#...
............
...#..#..#..
</code></pre>
In this region, the number of empty ground tiles is <code><em><b>110</b></em></code>.


Simulate the Elves' process and find the smallest rectangle that contains the Elves after 10 rounds. <em><b>How many empty ground tiles does that rectangle contain?</b></em>


