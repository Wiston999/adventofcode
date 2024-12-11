# --- Day 11: Plutonian Pebbles ---

The ancient civilization on [/2019/day/20](Pluto) was known for its ability to manipulate spacetime, and while The Historians explore their infinite corridors, you've noticed a strange set of physics-defying stones.


At first glance, they seem like normal stones: they're arranged in a perfectly <em><b>straight line</b></em>, and each stone has a <em><b>number</b></em> engraved on it.


The strange part is that every time you <span title="No, they're not statues. Why do you ask?">blink</span>, the stones <em><b>change</b></em>.


Sometimes, the number engraved on a stone changes. Other times, a stone might <em><b>split in two</b></em>, causing all the other stones to shift over a bit to make room in their perfectly straight line.


As you observe them for a while, you find that the stones have a consistent behavior. Every time you blink, the stones each <em><b>simultaneously</b></em> change according to the <em><b>first applicable rule</b></em> in this list:


<ul>
<li>If the stone is engraved with the number <code>0</code>, it is replaced by a stone engraved with the number <code>1</code>.</li>
<li>If the stone is engraved with a number that has an <em><b>even</b></em> number of digits, it is replaced by <em><b>two stones</b></em>. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: <code>1000</code> would become stones <code>10</code> and <code>0</code>.)</li>
<li>If none of the other rules apply, the stone is replaced by a new stone; the old stone's number <em><b>multiplied by 2024</b></em> is engraved on the new stone.</li>
</ul>
No matter how the stones change, their <em><b>order is preserved</b></em>, and they stay on their perfectly straight line.


How will the stones evolve if you keep blinking at them? You take a note of the number engraved on each stone in the line (your puzzle input).


If you have an arrangement of five stones engraved with the numbers <code>0 1 10 99 999</code> and you blink once, the stones transform as follows:


<ul>
<li>The first stone, <code>0</code>, becomes a stone marked <code>1</code>.</li>
<li>The second stone, <code>1</code>, is multiplied by 2024 to become <code>2024</code>.</li>
<li>The third stone, <code>10</code>, is split into a stone marked <code>1</code> followed by a stone marked <code>0</code>.</li>
<li>The fourth stone, <code>99</code>, is split into two stones marked <code>9</code>.</li>
<li>The fifth stone, <code>999</code>, is replaced by a stone marked <code>2021976</code>.</li>
</ul>
So, after blinking once, your five stones would become an arrangement of seven stones engraved with the numbers <code>1 2024 1 0 9 9 2021976</code>.


Here is a longer example:


<pre><code>Initial arrangement:
125 17

After 1 blink:
253000 1 7

After 2 blinks:
253 0 2024 14168

After 3 blinks:
512072 1 20 24 28676032

After 4 blinks:
512 72 2024 2 0 2 4 2867 6032

After 5 blinks:
1036288 7 2 20 24 4048 1 4048 8096 28 67 60 32

After 6 blinks:
2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2
</code></pre>
In this example, after blinking six times, you would have <code>22</code> stones. After blinking 25 times, you would have <code><em><b>55312</b></em></code> stones!


Consider the arrangement of stones in front of you. <em><b>How many stones will you have after blinking 25 times?</b></em>


