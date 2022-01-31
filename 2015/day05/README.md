# --- Day 5: Doesn&apos;t He Have Intern-Elves For This? ---

Santa needs help figuring out which strings in his text file are naughty or nice.


A <em><b>nice string</b></em> is one with all of the following properties:


<ul>
<li>It contains at least three vowels (<code>aeiou</code> only), like <code>aei</code>, <code>xazegov</code>, or <code title="John Madden John Madden John Madden">aeiouaeiouaeiou</code>.</li>
<li>It contains at least one letter that appears twice in a row, like <code>xx</code>, <code>abcdde</code> (<code>dd</code>), or <code>aabbccdd</code> (<code>aa</code>, <code>bb</code>, <code>cc</code>, or <code>dd</code>).</li>
<li>It does <em><b>not</b></em> contain the strings <code>ab</code>, <code>cd</code>, <code>pq</code>, or <code>xy</code>, even if they are part of one of the other requirements.</li>
</ul>
For example:


<ul>
<li><code>ugknbfddgicrmopn</code> is nice because it has at least three vowels (<code>u...i...o...</code>), a double letter (<code>...dd...</code>), and none of the disallowed substrings.</li>
<li><code>aaa</code> is nice because it has at least three vowels and a double letter, even though the letters used by different rules overlap.</li>
<li><code>jchzalrnumimnmhp</code> is naughty because it has no double letter.</li>
<li><code>haegwjzuvuyypxyu</code> is naughty because it contains the string <code>xy</code>.</li>
<li><code>dvszwmarrgswjxmb</code> is naughty because it contains only one vowel.</li>
</ul>
How many strings are nice?


