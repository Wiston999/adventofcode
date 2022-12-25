# --- Day 24: Blizzard Basin ---

With everything replanted for next year (and with elephants and monkeys to tend the grove), you and the Elves leave for the extraction point.


Partway up the mountain that shields the grove is a flat, open area that serves as the extraction point. It's a bit of a climb, but nothing the expedition can't handle.


At least, that would normally be true; now that the mountain is covered in snow, things have become more difficult than the Elves are used to.


As the expedition reaches a valley that must be traversed to reach the extraction site, you find that strong, turbulent winds are pushing small <em><b>blizzards</b></em> of snow and sharp ice around the valley. It's a good thing everyone packed warm clothes! To make it across safely, you'll need to find a way to avoid them.


Fortunately, it's easy to see all of this from the entrance to the valley, so you make a map of the valley and the blizzards (your puzzle input). For example:


<pre><code>#.#####
#.....#
#&gt;....#
#.....#
#...v.#
#.....#
#####.#
</code></pre>
The walls of the valley are drawn as <code>#</code>; everything else is ground. Clear ground - where there is currently no blizzard - is drawn as <code>.</code>. Otherwise, blizzards are drawn with an arrow indicating their direction of motion: up (<code>^</code>), down (<code>v</code>), left (<code>&lt;</code>), or right (<code>&gt;</code>).


The above map includes two blizzards, one moving right (<code>&gt;</code>) and one moving down (<code>v</code>). In one minute, each blizzard moves one position in the direction it is pointing:


<pre><code>#.#####
#.....#
#.&gt;...#
#.....#
#.....#
#...v.#
#####.#
</code></pre>
Due to <span title="I think, anyway. Do I look like a theoretical blizzacist?">conservation of blizzard energy</span>, as a blizzard reaches the wall of the valley, a new blizzard forms on the opposite side of the valley moving in the same direction. After another minute, the bottom downward-moving blizzard has been replaced with a new downward-moving blizzard at the top of the valley instead:


<pre><code>#.#####
#...v.#
#..&gt;..#
#.....#
#.....#
#.....#
#####.#
</code></pre>
Because blizzards are made of tiny snowflakes, they pass right through each other. After another minute, both blizzards temporarily occupy the same position, marked <code>2</code>:


<pre><code>#.#####
#.....#
#...2.#
#.....#
#.....#
#.....#
#####.#
</code></pre>
After another minute, the situation resolves itself, giving each blizzard back its personal space:


<pre><code>#.#####
#.....#
#....&gt;#
#...v.#
#.....#
#.....#
#####.#
</code></pre>
Finally, after yet another minute, the rightward-facing blizzard on the right is replaced with a new one on the left facing the same direction:


<pre><code>#.#####
#.....#
#&gt;....#
#.....#
#...v.#
#.....#
#####.#
</code></pre>
This process repeats at least as long as you are observing it, but probably forever.


Here is a more complex example:


<pre><code>#.######
#&gt;&gt;.&lt;^&lt;#
#.&lt;..&lt;&lt;#
#&gt;v.&gt;&lt;&gt;#
#&lt;^v^^&gt;#
######.#
</code></pre>
Your expedition begins in the only non-wall position in the top row and needs to reach the only non-wall position in the bottom row. On each minute, you can <em><b>move</b></em> up, down, left, or right, or you can <em><b>wait</b></em> in place. You and the blizzards act <em><b>simultaneously</b></em>, and you cannot share a position with a blizzard.


In the above example, the fastest way to reach your goal requires <code><em><b>18</b></em></code> steps. Drawing the position of the expedition as <code>E</code>, one way to achieve this is:


<pre><code>Initial state:
#<em><b>E</b></em>######
#&gt;&gt;.&lt;^&lt;#
#.&lt;..&lt;&lt;#
#&gt;v.&gt;&lt;&gt;#
#&lt;^v^^&gt;#
######.#

Minute 1, move down:
#.######
#<em><b>E</b></em>&gt;3.&lt;.#
#&lt;..&lt;&lt;.#
#&gt;2.22.#
#&gt;v..^&lt;#
######.#

Minute 2, move down:
#.######
#.2&gt;2..#
#<em><b>E</b></em>^22^&lt;#
#.&gt;2.^&gt;#
#.&gt;..&lt;.#
######.#

Minute 3, wait:
#.######
#&lt;^&lt;22.#
#<em><b>E</b></em>2&lt;.2.#
#&gt;&lt;2&gt;..#
#..&gt;&lt;..#
######.#

Minute 4, move up:
#.######
#<em><b>E</b></em>&lt;..22#
#&lt;&lt;.&lt;..#
#&lt;2.&gt;&gt;.#
#.^22^.#
######.#

Minute 5, move right:
#.######
#2<em><b>E</b></em>v.&lt;&gt;#
#&lt;.&lt;..&lt;#
#.^&gt;^22#
#.2..2.#
######.#

Minute 6, move right:
#.######
#&gt;2<em><b>E</b></em>&lt;.&lt;#
#.2v^2&lt;#
#&gt;..&gt;2&gt;#
#&lt;....&gt;#
######.#

Minute 7, move down:
#.######
#.22^2.#
#&lt;v<em><b>E</b></em>&lt;2.#
#&gt;&gt;v&lt;&gt;.#
#&gt;....&lt;#
######.#

Minute 8, move left:
#.######
#.&lt;&gt;2^.#
#.<em><b>E</b></em>&lt;&lt;.&lt;#
#.22..&gt;#
#.2v^2.#
######.#

Minute 9, move up:
#.######
#&lt;<em><b>E</b></em>2&gt;&gt;.#
#.&lt;&lt;.&lt;.#
#&gt;2&gt;2^.#
#.v&gt;&lt;^.#
######.#

Minute 10, move right:
#.######
#.2<em><b>E</b></em>.&gt;2#
#&lt;2v2^.#
#&lt;&gt;.&gt;2.#
#..&lt;&gt;..#
######.#

Minute 11, wait:
#.######
#2^<em><b>E</b></em>^2&gt;#
#&lt;v&lt;.^&lt;#
#..2.&gt;2#
#.&lt;..&gt;.#
######.#

Minute 12, move down:
#.######
#&gt;&gt;.&lt;^&lt;#
#.&lt;<em><b>E</b></em>.&lt;&lt;#
#&gt;v.&gt;&lt;&gt;#
#&lt;^v^^&gt;#
######.#

Minute 13, move down:
#.######
#.&gt;3.&lt;.#
#&lt;..&lt;&lt;.#
#&gt;2<em><b>E</b></em>22.#
#&gt;v..^&lt;#
######.#

Minute 14, move right:
#.######
#.2&gt;2..#
#.^22^&lt;#
#.&gt;2<em><b>E</b></em>^&gt;#
#.&gt;..&lt;.#
######.#

Minute 15, move right:
#.######
#&lt;^&lt;22.#
#.2&lt;.2.#
#&gt;&lt;2&gt;<em><b>E</b></em>.#
#..&gt;&lt;..#
######.#

Minute 16, move right:
#.######
#.&lt;..22#
#&lt;&lt;.&lt;..#
#&lt;2.&gt;&gt;<em><b>E</b></em>#
#.^22^.#
######.#

Minute 17, move down:
#.######
#2.v.&lt;&gt;#
#&lt;.&lt;..&lt;#
#.^&gt;^22#
#.2..2<em><b>E</b></em>#
######.#

Minute 18, move down:
#.######
#&gt;2.&lt;.&lt;#
#.2v^2&lt;#
#&gt;..&gt;2&gt;#
#&lt;....&gt;#
######<em><b>E</b></em>#
</code></pre>
<em><b>What is the fewest number of minutes required to avoid the blizzards and reach the goal?</b></em>


