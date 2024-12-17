# --- Day 17: Chronospatial Computer ---

The Historians push the button on their strange device, but this time, you all just feel like you're [/2018/day/6](falling).


"Situation critical", the device announces in a familiar voice. "Bootstrapping process failed. Initializing debugger...."


The small handheld device suddenly unfolds into an entire computer! The Historians look around nervously before one of them tosses it to you.


This seems to be a 3-bit computer: its program is a list of 3-bit numbers (0 through 7), like <code>0,1,2,3</code>. The computer also has three <em><b>registers</b></em> named <code>A</code>, <code>B</code>, and <code>C</code>, but these registers aren't limited to 3 bits and can instead hold any integer.


The computer knows <em><b>eight instructions</b></em>, each identified by a 3-bit number (called the instruction's <em><b>opcode</b></em>). Each instruction also reads the 3-bit number after it as an input; this is called its <em><b>operand</b></em>.


A number called the <em><b>instruction pointer</b></em> identifies the position in the program from which the next opcode will be read; it starts at <code>0</code>, pointing at the first 3-bit number in the program. Except for jump instructions, the instruction pointer increases by <code>2</code> after each instruction is processed (to move past the instruction's opcode and its operand). If the computer tries to read an opcode past the end of the program, it instead <em><b>halts</b></em>.


So, the program <code>0,1,2,3</code> would run the instruction whose opcode is <code>0</code> and pass it the operand <code>1</code>, then run the instruction having opcode <code>2</code> and pass it the operand <code>3</code>, then halt.


There are two types of operands; each instruction specifies the type of its operand. The value of a <em><b>literal operand</b></em> is the operand itself. For example, the value of the literal operand <code>7</code> is the number <code>7</code>. The value of a <em><b>combo operand</b></em> can be found as follows:


<ul>
<li>Combo operands <code>0</code> through <code>3</code> represent literal values <code>0</code> through <code>3</code>.</li>
<li>Combo operand <code>4</code> represents the value of register <code>A</code>.</li>
<li>Combo operand <code>5</code> represents the value of register <code>B</code>.</li>
<li>Combo operand <code>6</code> represents the value of register <code>C</code>.</li>
<li>Combo operand <code>7</code> is reserved and will not appear in valid programs.</li>
</ul>
The eight instructions are as follows:


The <code><em><b>adv</b></em></code> instruction (opcode <code><em><b>0</b></em></code>) performs <em><b>division</b></em>. The numerator is the value in the <code>A</code> register. The denominator is found by raising 2 to the power of the instruction's <em><b>combo</b></em> operand. (So, an operand of <code>2</code> would divide <code>A</code> by <code>4</code> (<code>2^2</code>); an operand of <code>5</code> would divide <code>A</code> by <code>2^B</code>.) The result of the division operation is <em><b>truncated</b></em> to an integer and then written to the <code>A</code> register.


The <code><em><b>bxl</b></em></code> instruction (opcode <code><em><b>1</b></em></code>) calculates the [https://en.wikipedia.org/wiki/Bitwise_operation#XOR](bitwise XOR) of register <code>B</code> and the instruction's <em><b>literal</b></em> operand, then stores the result in register <code>B</code>.


The <code><em><b>bst</b></em></code> instruction (opcode <code><em><b>2</b></em></code>) calculates the value of its <em><b>combo</b></em> operand [https://en.wikipedia.org/wiki/Modulo](modulo) 8 (thereby keeping only its lowest 3 bits), then writes that value to the <code>B</code> register.


The <code><em><b>jnz</b></em></code> instruction (opcode <code><em><b>3</b></em></code>) does <em><b>nothing</b></em> if the <code>A</code> register is <code>0</code>. However, if the <code>A</code> register is <em><b>not zero</b></em>, it <span title="The instruction does this using a little trampoline."><em><b>jumps</b></em></span> by setting the instruction pointer to the value of its <em><b>literal</b></em> operand; if this instruction jumps, the instruction pointer is <em><b>not</b></em> increased by <code>2</code> after this instruction.


The <code><em><b>bxc</b></em></code> instruction (opcode <code><em><b>4</b></em></code>) calculates the <em><b>bitwise XOR</b></em> of register <code>B</code> and register <code>C</code>, then stores the result in register <code>B</code>. (For legacy reasons, this instruction reads an operand but <em><b>ignores</b></em> it.)


The <code><em><b>out</b></em></code> instruction (opcode <code><em><b>5</b></em></code>) calculates the value of its <em><b>combo</b></em> operand modulo 8, then <em><b>outputs</b></em> that value. (If a program outputs multiple values, they are separated by commas.)


The <code><em><b>bdv</b></em></code> instruction (opcode <code><em><b>6</b></em></code>) works exactly like the <code>adv</code> instruction except that the result is stored in the <em><b><code>B</code> register</b></em>. (The numerator is still read from the <code>A</code> register.)


The <code><em><b>cdv</b></em></code> instruction (opcode <code><em><b>7</b></em></code>) works exactly like the <code>adv</code> instruction except that the result is stored in the <em><b><code>C</code> register</b></em>. (The numerator is still read from the <code>A</code> register.)


Here are some examples of instruction operation:


<ul>
<li>If register <code>C</code> contains <code>9</code>, the program <code>2,6</code> would set register <code>B</code> to <code>1</code>.</li>
<li>If register <code>A</code> contains <code>10</code>, the program <code>5,0,5,1,5,4</code> would output <code>0,1,2</code>.</li>
<li>If register <code>A</code> contains <code>2024</code>, the program <code>0,1,5,4,3,0</code> would output <code>4,2,5,6,7,7,7,7,3,1,0</code> and leave <code>0</code> in register <code>A</code>.</li>
<li>If register <code>B</code> contains <code>29</code>, the program <code>1,7</code> would set register <code>B</code> to <code>26</code>.</li>
<li>If register <code>B</code> contains <code>2024</code> and register <code>C</code> contains <code>43690</code>, the program <code>4,0</code> would set register <code>B</code> to <code>44354</code>.</li>
</ul>
The Historians' strange device has finished initializing its debugger and is displaying some <em><b>information about the program it is trying to run</b></em> (your puzzle input). For example:


<pre><code>Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
</code></pre>
Your first task is to <em><b>determine what the program is trying to output</b></em>. To do this, initialize the registers to the given values, then run the given program, collecting any output produced by <code>out</code> instructions. (Always join the values produced by <code>out</code> instructions with commas.) After the above program halts, its final output will be <code><em><b>4,6,3,5,6,3,5,2,1,0</b></em></code>.


Using the information provided by the debugger, initialize the registers to the given values, then run the program. Once it halts, <em><b>what do you get if you use commas to join the values it output into a single string?</b></em>


