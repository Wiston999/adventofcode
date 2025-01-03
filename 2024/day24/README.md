# --- Day 24: Crossed Wires ---

You and The Historians arrive at the edge of a [/2022/day/23](large grove) somewhere in the jungle. After the last incident, the Elves installed a small device that monitors the fruit. While The Historians search the grove, one of them asks if you can take a look at the monitoring device; apparently, it's been malfunctioning recently.


The device seems to be trying to produce a number through some boolean logic gates. Each gate has two inputs and one output. The gates all operate on values that are either <em><b>true</b></em> (<code>1</code>) or <em><b>false</b></em> (<code>0</code>).


<ul>
<li><code>AND</code> gates output <code>1</code> if <em><b>both</b></em> inputs are <code>1</code>; if either input is <code>0</code>, these gates output <code>0</code>.</li>
<li><code>OR</code> gates output <code>1</code> if <em><b>one or both</b></em> inputs is <code>1</code>; if both inputs are <code>0</code>, these gates output <code>0</code>.</li>
<li><code>XOR</code> gates output <code>1</code> if the inputs are <em><b>different</b></em>; if the inputs are the same, these gates output <code>0</code>.</li>
</ul>
Gates wait until both inputs are received before producing output; wires can carry <code>0</code>, <code>1</code> or no value at all. There are no loops; once a gate has determined its output, the output will not change until the whole system is reset. Each wire is connected to at most one gate output, but can be connected to many gate inputs.


Rather than risk getting shocked while tinkering with the live system, you write down all of the gate connections and initial wire values (your puzzle input) so you can consider them in relative safety. For example:


<pre><code>x00: 1
x01: 1
x02: 1
y00: 0
y01: 1
y02: 0

x00 AND y00 -&gt; z00
x01 XOR y01 -&gt; z01
x02 OR y02 -&gt; z02
</code></pre>
Because gates wait for input, some wires need to start with a value (as inputs to the entire system). The first section specifies these values. For example, <code>x00: 1</code> means that the wire named <code>x00</code> starts with the value <code>1</code> (as if a gate is already outputting that value onto that wire).


The second section lists all of the gates and the wires connected to them. For example, <code>x00 AND y00 -&gt; z00</code> describes an instance of an <code>AND</code> gate which has wires <code>x00</code> and <code>y00</code> connected to its inputs and which will write its output to wire <code>z00</code>.


In this example, simulating these gates eventually causes <code>0</code> to appear on wire <code>z00</code>, <code>0</code> to appear on wire <code>z01</code>, and <code>1</code> to appear on wire <code>z02</code>.


Ultimately, the system is trying to produce a <em><b>number</b></em> by combining the bits on all wires starting with <code>z</code>. <code>z00</code> is the least significant bit, then <code>z01</code>, then <code>z02</code>, and so on.


In this example, the three output bits form the binary number <code>100</code> which is equal to the decimal number <code><em><b>4</b></em></code>.


Here's a larger example:


<pre><code>x00: 1
x01: 0
x02: 1
x03: 1
x04: 0
y00: 1
y01: 1
y02: 1
y03: 1
y04: 1

ntg XOR fgs -> mjb
y02 OR x01 -> tnw
kwq OR kpj -> z05
x00 OR x03 -> fst
tgd XOR rvg -> z01
vdt OR tnw -> bfw
bfw AND frj -> z10
ffh OR nrd -> bqk
y00 AND y03 -> djm
y03 OR y00 -> psh
bqk OR frj -> z08
tnw OR fst -> frj
gnj AND tgd -> z11
bfw XOR mjb -> z00
x03 OR x00 -> vdt
gnj AND wpb -> z02
x04 AND y00 -> kjc
djm OR pbm -> qhw
nrd AND vdt -> hwm
kjc AND fst -> rvg
y04 OR y02 -> fgs
y01 AND x02 -> pbm
ntg OR kjc -> kwq
psh XOR fgs -> tgd
qhw XOR tgd -> z09
pbm OR djm -> kpj
x03 XOR y03 -> ffh
x00 XOR y04 -> ntg
bfw OR bqk -> z06
nrd XOR fgs -> wpb
frj XOR qhw -> z04
bqk OR frj -> z07
y03 OR x01 -> nrd
hwm AND bqk -> z03
tgd XOR rvg -> z12
tnw OR pbm -> gnj
</code></pre>
After waiting for values on all wires starting with <code>z</code>, the wires in this system have the following values:


<pre><code>bfw: 1
bqk: 1
djm: 1
ffh: 0
fgs: 1
frj: 1
fst: 1
gnj: 1
hwm: 1
kjc: 0
kpj: 1
kwq: 0
mjb: 1
nrd: 1
ntg: 0
pbm: 1
psh: 1
qhw: 1
rvg: 0
tgd: 0
tnw: 1
vdt: 1
wpb: 0
z00: 0
z01: 0
z02: 0
z03: 1
z04: 0
z05: 1
z06: 1
z07: 1
z08: 1
z09: 1
z10: 1
z11: 0
z12: 0
</code></pre>
Combining the bits from all wires starting with <code>z</code> produces the binary number <code>0011111101000</code>. Converting this number to decimal produces <code><em><b>2024</b></em></code>.


Simulate the system of gates and wires. <em><b>What decimal number does it output on the wires starting with <code>z</code>?</b></em>


