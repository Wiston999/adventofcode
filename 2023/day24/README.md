# --- Day 24: Never Tell Me The Odds ---

It seems like something is going wrong with the snow-making process. Instead of forming snow, the water that's been absorbed into the air seems to be forming [https://en.wikipedia.org/wiki/Hail](hail)!


Maybe there's something you can do to break up the hailstones?


Due to strong, probably-magical winds, the hailstones are all flying through the air in perfectly linear trajectories. You make a note of each hailstone's <em><b>position</b></em> and <em><b>velocity</b></em> (your puzzle input). For example:


<pre><code>19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3
</code></pre>
Each line of text corresponds to the position and velocity of a single hailstone. The positions indicate where the hailstones are <em><b>right now</b></em> (at time <code>0</code>). The velocities are constant and indicate exactly how far each hailstone will move in <em><b>one nanosecond</b></em>.


Each line of text uses the format <code>px py pz @ vx vy vz</code>. For instance, the hailstone specified by <code>20, 19, 15 @  1, -5, -3</code> has initial X position <code>20</code>, Y position <code>19</code>, Z position <code>15</code>, X velocity <code>1</code>, Y velocity <code>-5</code>, and Z velocity <code>-3</code>. After one nanosecond, the hailstone would be at <code>21, 14, 12</code>.


Perhaps you won't have to do anything. How likely are the hailstones to collide with each other and smash into tiny ice crystals?


To estimate this, consider only the X and Y axes; <em><b>ignore the Z axis</b></em>. Looking <em><b>forward in time</b></em>, how many of the hailstones' <em><b>paths</b></em> will intersect within a test area? (The hailstones themselves don't have to collide, just test for intersections between the paths they will trace.)


In this example, look for intersections that happen with an X and Y position each at least <code>7</code> and at most <code>27</code>; in your actual data, you'll need to check a much larger test area. Comparing all pairs of hailstones' future paths produces the following results:


<pre><code>Hailstone A: 19, 13, 30 @ -2, 1, -2
Hailstone B: 18, 19, 22 @ -1, -1, -2
Hailstones' paths will cross <em><b>inside</b></em> the test area (at x=14.333, y=15.333).

Hailstone A: 19, 13, 30 @ -2, 1, -2
Hailstone B: 20, 25, 34 @ -2, -2, -4
Hailstones' paths will cross <em><b>inside</b></em> the test area (at x=11.667, y=16.667).

Hailstone A: 19, 13, 30 @ -2, 1, -2
Hailstone B: 12, 31, 28 @ -1, -2, -1
Hailstones' paths will cross outside the test area (at x=6.2, y=19.4).

Hailstone A: 19, 13, 30 @ -2, 1, -2
Hailstone B: 20, 19, 15 @ 1, -5, -3
Hailstones' paths crossed in the past for hailstone A.

Hailstone A: 18, 19, 22 @ -1, -1, -2
Hailstone B: 20, 25, 34 @ -2, -2, -4
Hailstones' paths are parallel; they never intersect.

Hailstone A: 18, 19, 22 @ -1, -1, -2
Hailstone B: 12, 31, 28 @ -1, -2, -1
Hailstones' paths will cross outside the test area (at x=-6, y=-5).

Hailstone A: 18, 19, 22 @ -1, -1, -2
Hailstone B: 20, 19, 15 @ 1, -5, -3
Hailstones' paths crossed in the past for both hailstones.

Hailstone A: 20, 25, 34 @ -2, -2, -4
Hailstone B: 12, 31, 28 @ -1, -2, -1
Hailstones' paths will cross outside the test area (at x=-2, y=3).

Hailstone A: 20, 25, 34 @ -2, -2, -4
Hailstone B: 20, 19, 15 @ 1, -5, -3
Hailstones' paths crossed in the past for hailstone B.

Hailstone A: 12, 31, 28 @ -1, -2, -1
Hailstone B: 20, 19, 15 @ 1, -5, -3
Hailstones' paths crossed in the past for both hailstones.
</code></pre>
So, in this example, <code><em><b>2</b></em></code> hailstones' future paths cross inside the boundaries of the test area.


However, you'll need to search a much larger test area if you want to see if any hailstones might collide. Look for intersections that happen with an X and Y position each at least <code>200000000000000</code> and at most <code>400000000000000</code>. Disregard the Z axis entirely.


Considering only the X and Y axes, check all pairs of hailstones' future paths for intersections. <em><b>How many of these intersections occur within the test area?</b></em>


