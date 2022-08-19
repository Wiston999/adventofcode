# --- Day 17: Spinlock ---

Suddenly, whirling in the distance, you notice what looks like a massive, <span title="You know, as opposed to all those non-pixelated hurricanes you see on TV.">pixelated hurricane</span>: a deadly [https://en.wikipedia.org/wiki/Spinlock](spinlock). This spinlock isn't just consuming computing power, but memory, too; vast, digital mountains are being ripped from the ground and consumed by the vortex.


If you don't move quickly, fixing that printer will be the least of your problems.


This spinlock's algorithm is simple but efficient, quickly consuming everything in its path. It starts with a circular buffer containing only the value <code>0</code>, which it marks as the <em><b>current position</b></em>. It then steps forward through the circular buffer some number of steps (your puzzle input) before inserting the first new value, <code>1</code>, after the value it stopped on.  The inserted value becomes the <em><b>current position</b></em>. Then, it steps forward from there the same number of steps, and wherever it stops, inserts after it the second new value, <code>2</code>, and uses that as the new <em><b>current position</b></em> again.


It repeats this process of <em><b>stepping forward</b></em>, <em><b>inserting a new value</b></em>, and <em><b>using the location of the inserted value as the new current position</b></em> a total of <code><em><b>2017</b></em></code> times, inserting <code>2017</code> as its final operation, and ending with a total of <code>2018</code> values (including <code>0</code>) in the circular buffer.


For example, if the spinlock were to step <code>3</code> times per insert, the circular buffer would begin to evolve like this (using parentheses to mark the current position after each iteration of the algorithm):


<ul>
<li><code>(0)</code>, the initial state before any insertions.</li>
<li><code>0&nbsp;(1)</code>: the spinlock steps forward three times (<code>0</code>, <code>0</code>, <code>0</code>), and then inserts the first value, <code>1</code>, after it. <code>1</code> becomes the current position.</li>
<li><code>0&nbsp;(2)&nbsp;1</code>: the spinlock steps forward three times (<code>0</code>, <code>1</code>, <code>0</code>), and then inserts the second value, <code>2</code>, after it. <code>2</code> becomes the current position.</li>
<li><code>0&nbsp;&nbsp;2&nbsp;(3)&nbsp;1</code>: the spinlock steps forward three times (<code>1</code>, <code>0</code>, <code>2</code>), and then inserts the third value, <code>3</code>, after it. <code>3</code> becomes the current position.</li>
</ul>
And so on:


<ul>
<li><code>0&nbsp;&nbsp;2&nbsp;(4)&nbsp;3&nbsp;&nbsp;1</code></li>
<li><code>0&nbsp;(5)&nbsp;2&nbsp;&nbsp;4&nbsp;&nbsp;3&nbsp;&nbsp;1</code></li>
<li><code>0&nbsp;&nbsp;5&nbsp;&nbsp;2&nbsp;&nbsp;4&nbsp;&nbsp;3&nbsp;(6)&nbsp;1</code></li>
<li><code>0&nbsp;&nbsp;5&nbsp;(7)&nbsp;2&nbsp;&nbsp;4&nbsp;&nbsp;3&nbsp;&nbsp;6&nbsp;&nbsp;1</code></li>
<li><code>0&nbsp;&nbsp;5&nbsp;&nbsp;7&nbsp;&nbsp;2&nbsp;&nbsp;4&nbsp;&nbsp;3&nbsp;(8)&nbsp;6&nbsp;&nbsp;1</code></li>
<li><code>0&nbsp;(9)&nbsp;5&nbsp;&nbsp;7&nbsp;&nbsp;2&nbsp;&nbsp;4&nbsp;&nbsp;3&nbsp;&nbsp;8&nbsp;&nbsp;6&nbsp;&nbsp;1</code></li>
</ul>
Eventually, after 2017 insertions, the section of the circular buffer near the last insertion looks like this:


<pre><code>1512  1134  151 (2017) 638  1513  851</code></pre>
Perhaps, if you can identify the value that will ultimately be <em><b>after</b></em> the last value written (<code>2017</code>), you can short-circuit the spinlock.  In this example, that would be <code>638</code>.


<em><b>What is the value after <code>2017</code></b></em> in your completed circular buffer?


