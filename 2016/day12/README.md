# --- Day 12: Leonardo&apos;s Monorail ---

You finally reach the top floor of this building: a garden with a slanted glass ceiling. Looks like there are no more stars to be had.


While sitting on a nearby bench amidst some [https://www.google.com/search?q=tiger+lilies&tbm=isch](tiger lilies), you manage to decrypt some of the files you extracted from the servers downstairs.


According to these documents, Easter Bunny HQ isn't just this building - it's a collection of buildings in the nearby area. They're all connected by a local monorail, and there's another building not far from here! Unfortunately, being night, the monorail is currently not operating.


You remotely connect to the monorail control systems and discover that the boot sequence expects a password. The password-checking logic (your puzzle input) is easy to extract, but the code it uses is strange: it's <span title="Strangely, this assembunny code doesn't seem to be very good at multiplying.">assembunny</span> code designed for the [11](new computer) you just assembled. You'll have to execute the code and get the password.


The assembunny code you've extracted operates on four [https://en.wikipedia.org/wiki/Processor_register](registers) (<code>a</code>, <code>b</code>, <code>c</code>, and <code>d</code>) that start at <code>0</code> and can hold any [https://en.wikipedia.org/wiki/Integer](integer). However, it seems to make use of only a few [https://en.wikipedia.org/wiki/Instruction_set](instructions):


<ul>
<li><code>cpy x y</code> <em><b>copies</b></em> <code>x</code> (either an integer or the <em><b>value</b></em> of a register) into register <code>y</code>.</li>
<li><code>inc x</code> <em><b>increases</b></em> the value of register <code>x</code> by one.</li>
<li><code>dec x</code> <em><b>decreases</b></em> the value of register <code>x</code> by one.</li>
<li><code>jnz x y</code> <em><b>jumps</b></em> to an instruction <code>y</code> away (positive means forward; negative means backward), but only if <code>x</code> is <em><b>not zero</b></em>.</li>
</ul>
The <code>jnz</code> instruction moves relative to itself: an offset of <code>-1</code> would continue at the previous instruction, while an offset of <code>2</code> would <em><b>skip over</b></em> the next instruction.


For example:


<pre><code>cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a
</code></pre>
The above code would set register <code>a</code> to <code>41</code>, increase its value by <code>2</code>, decrease its value by <code>1</code>, and then skip the last <code>dec a</code> (because <code>a</code> is not zero, so the <code>jnz a 2</code> skips it), leaving register <code>a</code> at <code>42</code>. When you move past the last instruction, the program halts.


After executing the assembunny code in your puzzle input, <em><b>what value is left in register <code>a</code>?</b></em>


