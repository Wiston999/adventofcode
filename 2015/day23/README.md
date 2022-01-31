# --- Day 23: Opening the Turing Lock ---

Little Jane Marie just got her very first computer for Christmas from some <span title="Definitely not Wintermute.">unknown benefactor</span>.  It comes with instructions and an example program, but the computer itself seems to be malfunctioning.  She's curious what the program does, and would like you to help her run it.


The manual explains that the computer supports two [https://en.wikipedia.org/wiki/Processor_register](registers) and six [https://en.wikipedia.org/wiki/Instruction_set](instructions) (truly, it goes on to remind the reader, a state-of-the-art technology). The registers are named <code>a</code> and <code>b</code>, can hold any [https://en.wikipedia.org/wiki/Natural_number](non-negative integer), and begin with a value of <code>0</code>.  The instructions are as follows:


<ul>
<li><code>hlf r</code> sets register <code>r</code> to <em><b>half</b></em> its current value, then continues with the next instruction.</li>
<li><code>tpl r</code> sets register <code>r</code> to <em><b>triple</b></em> its current value, then continues with the next instruction.</li>
<li><code>inc r</code> <em><b>increments</b></em> register <code>r</code>, adding <code>1</code> to it, then continues with the next instruction.</li>
<li><code>jmp offset</code> is a <em><b>jump</b></em>; it continues with the instruction <code>offset</code> away <em><b>relative to itself</b></em>.</li>
<li><code>jie r, offset</code> is like <code>jmp</code>, but only jumps if register <code>r</code> is <em><b>even</b></em> ("jump if even").</li>
<li><code>jio r, offset</code> is like <code>jmp</code>, but only jumps if register <code>r</code> is <code>1</code> ("jump if <em><b>one</b></em>", not odd).</li>
</ul>
All three jump instructions work with an <em><b>offset</b></em> relative to that instruction.  The offset is always written with a prefix <code>+</code> or <code>-</code> to indicate the direction of the jump (forward or backward, respectively).  For example, <code>jmp +1</code> would simply continue with the next instruction, while <code>jmp +0</code> would continuously jump back to itself forever.


The program exits when it tries to run an instruction beyond the ones defined.


For example, this program sets <code>a</code> to <code>2</code>, because the <code>jio</code> instruction causes it to skip the <code>tpl</code> instruction:


<pre><code>inc a
jio a, +2
tpl a
inc a
</code></pre>
What is <em><b>the value in register <code>b</code></b></em> when the program in your puzzle input is finished executing?


