# --- Day 4: Camp Cleanup ---

Space needs to be cleared before the last supplies can be unloaded from the ships, and so several Elves have been assigned the job of cleaning up sections of the camp. Every section has a unique <em><b>ID number</b></em>, and each Elf is assigned a range of section IDs.


However, as some of the Elves compare their section assignments with each other, they've noticed that many of the assignments <em><b>overlap</b></em>. To try to quickly find overlaps and reduce duplicated effort, the Elves pair up and make a <em><b>big list of the section assignments for each pair</b></em> (your puzzle input).


For example, consider the following list of section assignment pairs:


<pre><code>2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
</code></pre>
For the first few pairs, this list means:


<ul>
<li>Within the first pair of Elves, the first Elf was assigned sections <code>2-4</code> (sections <code>2</code>, <code>3</code>, and <code>4</code>), while the second Elf was assigned sections <code>6-8</code> (sections <code>6</code>, <code>7</code>, <code>8</code>).</li>
<li>The Elves in the second pair were each assigned two sections.</li>
<li>The Elves in the third pair were each assigned three sections: one got sections <code>5</code>, <code>6</code>, and <code>7</code>, while the other also got <code>7</code>, plus <code>8</code> and <code>9</code>.</li>
</ul>
This example list uses single-digit section IDs to make it easier to draw; your actual list might contain larger numbers. Visually, these pairs of section assignments look like this:


<pre><code>.234.....  2-4
.....678.  6-8

.23......  2-3
...45....  4-5

....567..  5-7
......789  7-9

.2345678.  2-8
..34567..  3-7

.....6...  6-6
...456...  4-6

.23456...  2-6
...45678.  4-8
</code></pre>
Some of the pairs have noticed that one of their assignments <em><b>fully contains</b></em> the other. For example, <code>2-8</code> fully contains <code>3-7</code>, and <code>6-6</code> is fully contained by <code>4-6</code>. In pairs where one assignment fully contains the other, one Elf in the pair would be exclusively cleaning sections their partner will already be cleaning, so these seem like the most in need of reconsideration. In this example, there are <code><em><b>2</b></em></code> such pairs.


<em><b>In how many assignment pairs does one range fully contain the other?</b></em>


