# --- Day 14: Extended Polymerization ---

The incredible pressures at this depth are starting to put a strain on your submarine. The submarine has [https://en.wikipedia.org/wiki/Polymerization](polymerization) equipment that would produce suitable materials to reinforce the submarine, and the nearby volcanically-active caves should even have the necessary input elements in sufficient quantities.


The submarine manual contains <span title="HO&#xa;&#xa;HO -&gt; OH">instructions</span> for finding the optimal polymer formula; specifically, it offers a <em><b>polymer template</b></em> and a list of <em><b>pair insertion</b></em> rules (your puzzle input). You just need to work out what polymer would result after repeating the pair insertion process a few times.


For example:


<pre><code>NNCB

CH -&gt; B
HH -&gt; N
CB -&gt; H
NH -&gt; C
HB -&gt; C
HC -&gt; B
HN -&gt; C
NN -&gt; C
BH -&gt; H
NC -&gt; B
NB -&gt; B
BN -&gt; B
BB -&gt; N
BC -&gt; B
CC -&gt; N
CN -&gt; C
</code></pre>
The first line is the <em><b>polymer template</b></em> - this is the starting point of the process.


The following section defines the <em><b>pair insertion</b></em> rules. A rule like <code>AB -&gt; C</code> means that when elements <code>A</code> and <code>B</code> are immediately adjacent, element <code>C</code> should be inserted between them. These insertions all happen simultaneously.


So, starting with the polymer template <code>NNCB</code>, the first step simultaneously considers all three pairs:


<ul>
<li>The first pair (<code>NN</code>) matches the rule <code>NN -&gt; C</code>, so element <code><em><b>C</b></em></code> is inserted between the first <code>N</code> and the second <code>N</code>.</li>
<li>The second pair (<code>NC</code>) matches the rule <code>NC -&gt; B</code>, so element <code><em><b>B</b></em></code> is inserted between the <code>N</code> and the <code>C</code>.</li>
<li>The third pair (<code>CB</code>) matches the rule <code>CB -&gt; H</code>, so element <code><em><b>H</b></em></code> is inserted between the <code>C</code> and the <code>B</code>.</li>
</ul>
Note that these pairs overlap: the second element of one pair is the first element of the next pair. Also, because all pairs are considered simultaneously, inserted elements are not considered to be part of a pair until the next step.


After the first step of this process, the polymer becomes <code>N<em><b>C</b></em>N<em><b>B</b></em>C<em><b>H</b></em>B</code>.


Here are the results of a few steps using the above rules:


<pre><code>Template:     NNCB
After step 1: NCNBCHB
After step 2: NBCCNBBBCBHCB
After step 3: NBBBCNCCNBBNBNBBCHBHHBCHB
After step 4: NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB
</code></pre>
This polymer grows quickly. After step 5, it has length 97; After step 10, it has length 3073. After step 10, <code>B</code> occurs 1749 times, <code>C</code> occurs 298 times, <code>H</code> occurs 161 times, and <code>N</code> occurs 865 times; taking the quantity of the most common element (<code>B</code>, 1749) and subtracting the quantity of the least common element (<code>H</code>, 161) produces <code>1749 - 161 = <em><b>1588</b></em></code>.


Apply 10 steps of pair insertion to the polymer template and find the most and least common elements in the result. <em><b>What do you get if you take the quantity of the most common element and subtract the quantity of the least common element?</b></em>


