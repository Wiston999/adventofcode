# --- Day 9: Disk Fragmenter ---

Another push of the button leaves you in the familiar hallways of some friendly [/2021/day/23](amphipods)! Good thing you each somehow got your own personal mini submarine. The Historians jet away in search of the Chief, mostly by driving directly into walls.


While The Historians quickly figure out how to pilot these things, you notice an amphipod in the corner struggling with his computer. He's trying to make more contiguous free space by compacting all of the files, but his program isn't working; you offer to help.


He shows you the <em><b>disk map</b></em> (your puzzle input) he's already generated. For example:


<pre><code>2333133121414131402</code></pre>
The disk map uses a dense format to represent the layout of <em><b>files</b></em> and <em><b>free space</b></em> on the disk. The digits alternate between indicating the length of a file and the length of free space.


So, a disk map like <code>12345</code> would represent a one-block file, two blocks of free space, a three-block file, four blocks of free space, and then a five-block file. A disk map like <code>90909</code> would represent three nine-block files in a row (with no free space between them).


Each file on disk also has an <em><b>ID number</b></em> based on the order of the files as they appear <em><b>before</b></em> they are rearranged, starting with ID <code>0</code>. So, the disk map <code>12345</code> has three files: a one-block file with ID <code>0</code>, a three-block file with ID <code>1</code>, and a five-block file with ID <code>2</code>. Using one character for each block where digits are the file ID and <code>.</code> is free space, the disk map <code>12345</code> represents these individual blocks:


<pre><code>0..111....22222</code></pre>
The first example above, <code>2333133121414131402</code>, represents these individual blocks:


<pre><code>00...111...2...333.44.5555.6666.777.888899</code></pre>
The amphipod would like to <em><b>move file blocks one at a time</b></em> from the end of the disk to the leftmost free space block (until there are no gaps remaining between file blocks). For the disk map <code>12345</code>, the process looks like this:


<pre><code>0..111....22222
02.111....2222.
022111....222..
0221112...22...
02211122..2....
022111222......
</code></pre>
The first example requires a few more steps:


<pre><code>00...111...2...333.44.5555.6666.777.888899
009..111...2...333.44.5555.6666.777.88889.
0099.111...2...333.44.5555.6666.777.8888..
00998111...2...333.44.5555.6666.777.888...
009981118..2...333.44.5555.6666.777.88....
0099811188.2...333.44.5555.6666.777.8.....
009981118882...333.44.5555.6666.777.......
0099811188827..333.44.5555.6666.77........
00998111888277.333.44.5555.6666.7.........
009981118882777333.44.5555.6666...........
009981118882777333644.5555.666............
00998111888277733364465555.66.............
0099811188827773336446555566..............
</code></pre>
The final step of this file-compacting process is to update the <em><b>filesystem checksum</b></em>. To calculate the checksum, add up the result of multiplying each of these blocks' position with the file ID number it contains. The leftmost block is in position <code>0</code>. If a block contains free space, skip it instead.


Continuing the first example, the first few blocks' position multiplied by its file ID number are <code>0 * 0 = 0</code>, <code>1 * 0 = 0</code>, <code>2 * 9 = 18</code>, <code>3 * 9 = 27</code>, <code>4 * 8 = 32</code>, and so on. In this example, the checksum is the sum of these, <code><em><b>1928</b></em></code>.


<span title="Bonus points if you make a cool animation of this process.">Compact the amphipod's hard drive</span> using the process he requested. <em><b>What is the resulting filesystem checksum?</b></em> <span class="quiet">(Be careful copy/pasting the input for this puzzle; it is a single, very long line.)</span>

