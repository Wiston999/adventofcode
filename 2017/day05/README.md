# --- Day 5: A Maze of Twisty Trampolines, All Alike ---

An urgent <span title="Later, on its turn, it sends you a sorcery.">interrupt</span> arrives from the CPU: it's trapped in a maze of jump instructions, and it would like assistance from any programs with spare cycles to help find the exit.


The message includes a list of the offsets for each jump. Jumps are relative: <code>-1</code> moves to the previous instruction, and <code>2</code> skips the next one. Start at the first instruction in the list. The goal is to follow the jumps until one leads <em><b>outside</b></em> the list.


In addition, these instructions are a little strange; after each jump, the offset of that instruction increases by <code>1</code>. So, if you come across an offset of <code>3</code>, you would move three instructions forward, but change it to a <code>4</code> for the next time it is encountered.


For example, consider the following list of jump offsets:


<pre><code>0
3
0
1
-3
</code></pre>
Positive jumps ("forward") move downward; negative jumps move upward. For legibility in this example, these offset values will be written all on one line, with the current instruction marked in parentheses. The following steps would be taken before an exit is found:


<ul>
<li><code>(0)&nbsp;3&nbsp;&nbsp;0&nbsp;&nbsp;1&nbsp;&nbsp;-3&nbsp;</code> - <em><b>before</b></em> we have taken any steps.</li>
<li><code>(1)&nbsp;3&nbsp;&nbsp;0&nbsp;&nbsp;1&nbsp;&nbsp;-3&nbsp;</code> - jump with offset <code>0</code> (that is, don't jump at all). Fortunately, the instruction is then incremented to <code>1</code>.</li>
<li><code>&nbsp;2&nbsp;(3)&nbsp;0&nbsp;&nbsp;1&nbsp;&nbsp;-3&nbsp;</code> - step forward because of the instruction we just modified. The first instruction is incremented again, now to <code>2</code>.</li>
<li><code>&nbsp;2&nbsp;&nbsp;4&nbsp;&nbsp;0&nbsp;&nbsp;1&nbsp;(-3)</code> - jump all the way to the end; leave a <code>4</code> behind.</li>
<li><code>&nbsp;2&nbsp;(4)&nbsp;0&nbsp;&nbsp;1&nbsp;&nbsp;-2&nbsp;</code> - go back to where we just were; increment <code>-3</code> to <code>-2</code>.</li>
<li><code>&nbsp;2&nbsp;&nbsp;5&nbsp;&nbsp;0&nbsp;&nbsp;1&nbsp;&nbsp;-2&nbsp;</code> - jump <code>4</code> steps forward, escaping the maze.</li>
</ul>
In this example, the exit is reached in <code>5</code> steps.


<em><b>How many steps</b></em> does it take to reach the exit?


