# --- Day 16: The Floor Will Be Lava ---

With the beam of light completely focused <em><b>somewhere</b></em>, the reindeer leads you deeper still into the Lava Production Facility. At some point, you realize that the steel facility walls have been replaced with cave, and the doorways are just cave, and the floor is cave, and you're pretty sure this is actually just a giant cave.


Finally, as you approach what must be the heart of the mountain, you see a bright light in a cavern up ahead. There, you discover that the <span title="Not anymore, there's a blanket!">beam</span> of light you so carefully focused is emerging from the cavern wall closest to the facility and pouring all of its energy into a contraption on the opposite side.


Upon closer inspection, the contraption appears to be a flat, two-dimensional square grid containing <em><b>empty space</b></em> (<code>.</code>), <em><b>mirrors</b></em> (<code>/</code> and <code>\</code>), and <em><b>splitters</b></em> (<code>|</code> and <code>-</code>).


The contraption is aligned so that most of the beam bounces around the grid, but each tile on the grid converts some of the beam's light into <em><b>heat</b></em> to melt the rock in the cavern.


You note the layout of the contraption (your puzzle input). For example:


<pre><code>.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....
</code></pre>
The beam enters in the top-left corner from the left and heading to the <em><b>right</b></em>. Then, its behavior depends on what it encounters as it moves:


<ul>
<li>If the beam encounters <em><b>empty space</b></em> (<code>.</code>), it continues in the same direction.</li>
<li>If the beam encounters a <em><b>mirror</b></em> (<code>/</code> or <code>\</code>), the beam is <em><b>reflected</b></em> 90 degrees depending on the angle of the mirror. For instance, a rightward-moving beam that encounters a <code>/</code> mirror would continue <em><b>upward</b></em> in the mirror's column, while a rightward-moving beam that encounters a <code>\</code> mirror would continue <em><b>downward</b></em> from the mirror's column.</li>
<li>If the beam encounters the <em><b>pointy end of a splitter</b></em> (<code>|</code> or <code>-</code>), the beam passes through the splitter as if the splitter were <em><b>empty space</b></em>. For instance, a rightward-moving beam that encounters a <code>-</code> splitter would continue in the same direction.</li>
<li>If the beam encounters the <em><b>flat side of a splitter</b></em> (<code>|</code> or <code>-</code>), the beam is <em><b>split into two beams</b></em> going in each of the two directions the splitter's pointy ends are pointing. For instance, a rightward-moving beam that encounters a <code>|</code> splitter would split into two beams: one that continues <em><b>upward</b></em> from the splitter's column and one that continues <em><b>downward</b></em> from the splitter's column.</li>
</ul>
Beams do not interact with other beams; a tile can have many beams passing through it at the same time. A tile is <em><b>energized</b></em> if that tile has at least one beam pass through it, reflect in it, or split in it.


In the above example, here is how the beam of light bounces around the contraption:


<pre><code>&gt;|&lt;&lt;&lt;\....
|v-.\^....
.v...|-&gt;&gt;&gt;
.v...v^.|.
.v...v^...
.v...v^..\
.v../2\\..
&lt;-&gt;-/vv|..
.|&lt;&lt;&lt;2-|.\
.v//.|.v..
</code></pre>
Beams are only shown on empty tiles; arrows indicate the direction of the beams. If a tile contains beams moving in multiple directions, the number of distinct directions is shown instead. Here is the same diagram but instead only showing whether a tile is <em><b>energized</b></em> (<code>#</code>) or not (<code>.</code>):


<pre><code>######....
.#...#....
.#...#####
.#...##...
.#...##...
.#...##...
.#..####..
########..
.#######..
.#...#.#..
</code></pre>
Ultimately, in this example, <code><em><b>46</b></em></code> tiles become <em><b>energized</b></em>.


The light isn't energizing enough tiles to produce lava; to debug the contraption, you need to start by analyzing the current situation. With the beam starting in the top-left heading right, <em><b>how many tiles end up being energized?</b></em>


