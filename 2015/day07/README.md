# --- Day 7: Some Assembly Required ---

This year, Santa brought little Bobby Tables a set of wires and [https://en.wikipedia.org/wiki/Bitwise_operation](bitwise logic gates)!  Unfortunately, little Bobby is a little under the recommended age range, and he needs help <span title="You had one of these as a kid, right?">assembling the circuit</span>.


Each wire has an identifier (some lowercase letters) and can carry a [https://en.wikipedia.org/wiki/16-bit](16-bit) signal (a number from <code>0</code> to <code>65535</code>).  A signal is provided to each wire by a gate, another wire, or some specific value. Each wire can only get a signal from one source, but can provide its signal to multiple destinations.  A gate provides no signal until all of its inputs have a signal.


The included instructions booklet describes how to connect the parts together: <code>x AND y -> z</code> means to connect wires <code>x</code> and <code>y</code> to an AND gate, and then connect its output to wire <code>z</code>.


For example:


<ul>
<li><code>123 -> x</code> means that the signal <code>123</code> is provided to wire <code>x</code>.</li>
<li><code>x AND y -> z</code> means that the [https://en.wikipedia.org/wiki/Bitwise_operation#AND](bitwise AND) of wire <code>x</code> and wire <code>y</code> is provided to wire <code>z</code>.</li>
<li><code>p LSHIFT 2 -> q</code> means that the value from wire <code>p</code> is [https://en.wikipedia.org/wiki/Logical_shift](left-shifted) by <code>2</code> and then provided to wire <code>q</code>.</li>
<li><code>NOT e -> f</code> means that the [https://en.wikipedia.org/wiki/Bitwise_operation#NOT](bitwise complement) of the value from wire <code>e</code> is provided to wire <code>f</code>.</li>
</ul>
Other possible gates include <code>OR</code> ([https://en.wikipedia.org/wiki/Bitwise_operation#OR](bitwise OR)) and <code>RSHIFT</code> ([https://en.wikipedia.org/wiki/Logical_shift](right-shift)).  If, for some reason, you'd like to <em><b>emulate</b></em> the circuit instead, almost all programming languages (for example, [https://en.wikipedia.org/wiki/Bitwise_operations_in_C](C), [https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Bitwise_Operators](JavaScript), or [https://wiki.python.org/moin/BitwiseOperators](Python)) provide operators for these gates.


For example, here is a simple circuit:


<pre><code>123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i
</code></pre>
After it is run, these are the signals on the wires:


<pre><code>d: 72
e: 507
f: 492
g: 114
h: 65412
i: 65079
x: 123
y: 456
</code></pre>
In little Bobby's kit's instructions booklet (provided as your puzzle input), what signal is ultimately provided to <em><b>wire <code>a</code></b></em>?


