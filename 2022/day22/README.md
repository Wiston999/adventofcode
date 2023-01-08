# --- Day 22: Monkey Map ---

The monkeys take you on a surprisingly easy trail through the jungle. They're even going in roughly the right direction according to your handheld device's Grove Positioning System.


As you walk, the monkeys explain that the grove is protected by a <em><b>force field</b></em>. To pass through the force field, you have to enter a password; doing so involves tracing a specific <em><b>path</b></em> on a strangely-shaped board.


At least, you're pretty sure that's what you have to do; the elephants aren't exactly fluent in monkey.


The monkeys give you notes that they took when they last saw the password entered (your puzzle input).


For example:


<pre><code>        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5
</code></pre>
The first half of the monkeys' notes is a <em><b>map of the board</b></em>. It is comprised of a set of <em><b>open tiles</b></em> (on which you can move, drawn <code>.</code>) and <em><b>solid walls</b></em> (tiles which you cannot enter, drawn <code>#</code>).


The second half is a description of <em><b>the path you must follow</b></em>. It consists of alternating numbers and letters:


<ul>
<li>A <em><b>number</b></em> indicates the <em><b>number of tiles to move</b></em> in the direction you are facing. If you run into a wall, you stop moving forward and continue with the next instruction.</li>
<li>A <em><b>letter</b></em> indicates whether to turn 90 degrees <em><b>clockwise</b></em> (<code>R</code>) or <em><b><span title="Or &quot;anticlockwise&quot;, if you're anti-counterclockwise.">counterclockwise</span></b></em> (<code>L</code>). Turning happens in-place; it does not change your current tile.</li>
</ul>
So, a path like <code>10R5</code> means "go forward 10 tiles, then turn clockwise 90 degrees, then go forward 5 tiles".


You begin the path in the leftmost open tile of the top row of tiles. Initially, you are facing <em><b>to the right</b></em> (from the perspective of how the map is drawn).


If a movement instruction would take you off of the map, you <em><b>wrap around</b></em> to the other side of the board. In other words, if your next tile is off of the board, you should instead look in the direction opposite of your current facing as far as you can until you find the opposite edge of the board, then reappear there.


For example, if you are at <code>A</code> and facing to the right, the tile in front of you is marked <code>B</code>; if you are at <code>C</code> and facing down, the tile in front of you is marked <code>D</code>:


<pre><code>        ...#
        .#..
        #...
        ....
...#.<em><b>D</b></em>.....#
........#...
<em><b>B</b></em>.#....#...<em><b>A</b></em>
.....<em><b>C</b></em>....#.
        ...#....
        .....#..
        .#......
        ......#.
</code></pre>
It is possible for the next tile (after wrapping around) to be a <em><b>wall</b></em>; this still counts as there being a wall in front of you, and so movement stops before you actually wrap to the other side of the board.


By drawing the <em><b>last facing you had</b></em> with an arrow on each tile you visit, the full path taken by the above example looks like this:


<pre><code>        &gt;&gt;v#    
        .#v.    
        #.v.    
        ..v.    
...#...v..v#    
&gt;&gt;&gt;v...<em><b>&gt;</b></em>#.&gt;&gt;    
..#v...#....    
...&gt;&gt;&gt;&gt;v..#.    
        ...#....
        .....#..
        .#......
        ......#.
</code></pre>
To finish providing the password to this strange input device, you need to determine numbers for your final <em><b>row</b></em>, <em><b>column</b></em>, and <em><b>facing</b></em> as your final position appears from the perspective of the original map. Rows start from <code>1</code> at the top and count downward; columns start from <code>1</code> at the left and count rightward. (In the above example, row 1, column 1 refers to the empty space with no tile on it in the top-left corner.) Facing is <code>0</code> for right (<code>&gt;</code>), <code>1</code> for down (<code>v</code>), <code>2</code> for left (<code>&lt;</code>), and <code>3</code> for up (<code>^</code>). The <em><b>final password</b></em> is the sum of 1000 times the row, 4 times the column, and the facing.


In the above example, the final row is <code>6</code>, the final column is <code>8</code>, and the final facing is <code>0</code>. So, the final password is 1000 * 6 + 4 * 8 + 0: <code><em><b>6032</b></em></code>.


Follow the path given in the monkeys' notes. <em><b>What is the final password?</b></em>


