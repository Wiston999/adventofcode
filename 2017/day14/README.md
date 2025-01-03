# --- Day 14: Disk Defragmentation ---

Suddenly, a scheduled job activates the system's [https://en.wikipedia.org/wiki/Defragmentation](disk defragmenter). Were the situation different, you might [https://www.youtube.com/watch?v=kPv1gQ5Rs8A&t=37](sit and watch it for a while), but today, you just don't have that kind of time. It's soaking up valuable system resources that are needed elsewhere, and so the only option is to help it finish its task as soon as possible.


The disk in question consists of a 128x128 grid; each square of the grid is either <em><b>free</b></em> or <em><b>used</b></em>. On this disk, the state of the grid is tracked by the bits in a sequence of [10](knot hashes).


A total of 128 knot hashes are calculated, each corresponding to a single row in the grid; each hash contains 128 bits which correspond to individual grid squares. Each bit of a hash indicates whether that square is <em><b>free</b></em> (<code>0</code>) or <em><b>used</b></em> (<code>1</code>).


The hash inputs are a key string (your puzzle input), a dash, and a number from <code>0</code> to <code>127</code> corresponding to the row.  For example, if your key string were <code>flqrgnkx</code>, then the first row would be given by the bits of the knot hash of <code>flqrgnkx-0</code>, the second row from the bits of the knot hash of <code>flqrgnkx-1</code>, and so on until the last row, <code>flqrgnkx-127</code>.


The output of a knot hash is traditionally represented by 32 hexadecimal digits; each of these digits correspond to 4 bits, for a total of <code>4 * 32 = 128</code> bits. To convert to bits, turn each hexadecimal digit to its equivalent binary value, high-bit first: <code>0</code> becomes <code>0000</code>, <code>1</code> becomes <code>0001</code>, <code>e</code> becomes <code>1110</code>, <code>f</code> becomes <code>1111</code>, and so on; a hash that begins with <code>a0c2017...</code> in hexadecimal would begin with <code>10100000110000100000000101110000...</code> in binary.


Continuing this process, the <em><b>first 8 rows and columns</b></em> for key <code>flqrgnkx</code> appear as follows, using <code>#</code> to denote used squares, and <code>.</code> to denote free ones:


<pre><code>##.#.#..-->
.#.#.#.#   
....#.#.   
#.#.##.#   
.##.#...   
##..#..#   
.#...#..   
##.#.##.-->
|      |   
V      V   
</code></pre>
In this example, <code>8108</code> squares are used across the entire 128x128 grid.


Given your actual key string, <em><b>how many squares are used</b></em>?


