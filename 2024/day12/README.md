# --- Day 12: Garden Groups ---

Why not search for the Chief Historian near the [/2023/day/5](gardener) and his [/2023/day/21](massive farm)? There's plenty of food, so The Historians grab something to eat while they search.


You're about to settle near a complex arrangement of garden plots when some Elves ask if you can lend a hand. They'd like to set up <span title="I originally wanted to title this puzzle &quot;Fencepost Problem&quot;, but I was afraid someone would then try to count fenceposts by mistake and experience a fencepost problem.">fences</span> around each region of garden plots, but they can't figure out how much fence they need to order or how much it will cost. They hand you a map (your puzzle input) of the garden plots.


Each garden plot grows only a single type of plant and is indicated by a single letter on your map. When multiple garden plots are growing the same type of plant and are touching (horizontally or vertically), they form a <em><b>region</b></em>. For example:


<pre><code>AAAA
BBCD
BBCC
EEEC
</code></pre>
This 4x4 arrangement includes garden plots growing five different types of plants (labeled <code>A</code>, <code>B</code>, <code>C</code>, <code>D</code>, and <code>E</code>), each grouped into their own region.


In order to accurately calculate the cost of the fence around a single region, you need to know that region's <em><b>area</b></em> and <em><b>perimeter</b></em>.


The <em><b>area</b></em> of a region is simply the number of garden plots the region contains. The above map's type <code>A</code>, <code>B</code>, and <code>C</code> plants are each in a region of area <code>4</code>. The type <code>E</code> plants are in a region of area <code>3</code>; the type <code>D</code> plants are in a region of area <code>1</code>.


Each garden plot is a square and so has <em><b>four sides</b></em>. The <em><b>perimeter</b></em> of a region is the number of sides of garden plots in the region that do not touch another garden plot in the same region. The type <code>A</code> and <code>C</code> plants are each in a region with perimeter <code>10</code>. The type <code>B</code> and <code>E</code> plants are each in a region with perimeter <code>8</code>. The lone <code>D</code> plot forms its own region with perimeter <code>4</code>.


Visually indicating the sides of plots in each region that contribute to the perimeter using <code>-</code> and <code>|</code>, the above map's regions' perimeters are measured as follows:


<pre><code>+-+-+-+-+
|A A A A|
+-+-+-+-+     +-+
              |D|
+-+-+   +-+   +-+
|B B|   |C|
+   +   + +-+
|B B|   |C C|
+-+-+   +-+ +
          |C|
+-+-+-+   +-+
|E E E|
+-+-+-+
</code></pre>
Plants of the same type can appear in multiple separate regions, and regions can even appear within other regions. For example:


<pre><code>OOOOO
OXOXO
OOOOO
OXOXO
OOOOO
</code></pre>
The above map contains <em><b>five</b></em> regions, one containing all of the <code>O</code> garden plots, and the other four each containing a single <code>X</code> plot.


The four <code>X</code> regions each have area <code>1</code> and perimeter <code>4</code>. The region containing <code>21</code> type <code>O</code> plants is more complicated; in addition to its outer edge contributing a perimeter of <code>20</code>, its boundary with each <code>X</code> region contributes an additional <code>4</code> to its perimeter, for a total perimeter of <code>36</code>.


Due to "modern" business practices, the <em><b>price</b></em> of fence required for a region is found by <em><b>multiplying</b></em> that region's area by its perimeter. The <em><b>total price</b></em> of fencing all regions on a map is found by adding together the price of fence for every region on the map.


In the first example, region <code>A</code> has price <code>4 * 10 = 40</code>, region <code>B</code> has price <code>4 * 8 = 32</code>, region <code>C</code> has price <code>4 * 10 = 40</code>, region <code>D</code> has price <code>1 * 4 = 4</code>, and region <code>E</code> has price <code>3 * 8 = 24</code>. So, the total price for the first example is <code><em><b>140</b></em></code>.


In the second example, the region with all of the <code>O</code> plants has price <code>21 * 36 = 756</code>, and each of the four smaller <code>X</code> regions has price <code>1 * 4 = 4</code>, for a total price of <code><em><b>772</b></em></code> (<code>756 + 4 + 4 + 4 + 4</code>).


Here's a larger example:


<pre><code>RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
</code></pre>
It contains:


<ul>
<li>A region of <code>R</code> plants with price <code>12 * 18 = 216</code>.</li>
<li>A region of <code>I</code> plants with price <code>4 * 8 = 32</code>.</li>
<li>A region of <code>C</code> plants with price <code>14 * 28 = 392</code>.</li>
<li>A region of <code>F</code> plants with price <code>10 * 18 = 180</code>.</li>
<li>A region of <code>V</code> plants with price <code>13 * 20 = 260</code>.</li>
<li>A region of <code>J</code> plants with price <code>11 * 20 = 220</code>.</li>
<li>A region of <code>C</code> plants with price <code>1 * 4 = 4</code>.</li>
<li>A region of <code>E</code> plants with price <code>13 * 18 = 234</code>.</li>
<li>A region of <code>I</code> plants with price <code>14 * 22 = 308</code>.</li>
<li>A region of <code>M</code> plants with price <code>5 * 12 = 60</code>.</li>
<li>A region of <code>S</code> plants with price <code>3 * 8 = 24</code>.</li>
</ul>
So, it has a total price of <code><em><b>1930</b></em></code>.


<em><b>What is the total price of fencing all regions on your map?</b></em>


