# --- Day 23: Coprocessor Conflagration ---

You decide to head directly to the CPU and fix the printer from there. As you get close, you find an <em><b>experimental coprocessor</b></em> doing so much work that the local programs are afraid it will [https://en.wikipedia.org/wiki/Halt_and_Catch_Fire](halt and catch fire). This would cause serious issues for the rest of the computer, so you head in and see what you can do.


The code it's running seems to be a variant of the kind you saw recently on that [18](tablet). The general functionality seems <em><b>very similar</b></em>, but some of the instructions are different:


<ul>
<li><code>set X Y</code> <em><b>sets</b></em> register <code>X</code> to the value of <code>Y</code>.</li>
<li><code>sub X Y</code> <em><b>decreases</b></em> register <code>X</code> by the value of <code>Y</code>.</li>
<li><code>mul X Y</code> sets register <code>X</code> to the result of <em><b>multiplying</b></em> the value contained in register <code>X</code> by the value of <code>Y</code>.</li>
<li><code>jnz X Y</code> <em><b>jumps</b></em> with an offset of the value of <code>Y</code>, but only if the value of <code>X</code> is <em><b>not zero</b></em>. (An offset of <code>2</code> skips the next instruction, an offset of <code>-1</code> jumps to the previous instruction, and so on.)</li>
Only the instructions listed above are used. The eight registers here, named <code>a</code> through <code>h</code>, all start at <code>0</code>.


</ul>
The coprocessor is currently set to some kind of <em><b>debug mode</b></em>, which allows for testing, but prevents it from doing any meaningful work.


If you run the program (your puzzle input), <em><b>how many times is the <code>mul</code> instruction invoked?</b></em>


