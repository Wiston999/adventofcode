# --- Day 20: Pulse Propagation ---

With your help, the Elves manage to find the right parts and fix all of the machines. Now, they just need to send the command to boot up the machines and get the sand flowing again.


The machines are far apart and wired together with long <em><b>cables</b></em>. The cables don't connect to the machines directly, but rather to communication <em><b>modules</b></em> attached to the machines that perform various initialization tasks and also act as communication relays.


Modules communicate using <em><b>pulses</b></em>. Each pulse is either a <em><b>high pulse</b></em> or a <em><b>low pulse</b></em>. When a module sends a pulse, it sends that type of pulse to each module in its list of <em><b>destination modules</b></em>.


There are several different types of modules:


<em><b>Flip-flop</b></em> modules (prefix <code>%</code>) are either <em><b>on</b></em> or <em><b>off</b></em>; they are initially <em><b>off</b></em>. If a flip-flop module receives a high pulse, it is ignored and nothing happens. However, if a flip-flop module receives a low pulse, it <em><b>flips between on and off</b></em>. If it was off, it turns on and sends a high pulse. If it was on, it turns off and sends a low pulse.


<em><b>Conjunction</b></em> modules (prefix <code>&amp;</code>) <em><b>remember</b></em> the type of the most recent pulse received from <em><b>each</b></em> of their connected input modules; they initially default to remembering a <em><b>low pulse</b></em> for each input. When a pulse is received, the conjunction module first updates its memory for that input. Then, if it remembers <em><b>high pulses</b></em> for all inputs, it sends a <em><b>low pulse</b></em>; otherwise, it sends a <em><b>high pulse</b></em>.


There is a single <em><b>broadcast module</b></em> (named <code>broadcaster</code>). When it receives a pulse, it sends the same pulse to all of its destination modules.


Here at Desert Machine Headquarters, there is a module with a single button on it called, aptly, the <em><b>button module</b></em>. When you push the button, a single <em><b>low pulse</b></em> is sent directly to the <code>broadcaster</code> module.


After pushing the button, you must wait until all pulses have been delivered and fully handled before pushing it again. Never push the button if modules are still processing pulses.


Pulses are always processed <em><b>in the order they are sent</b></em>. So, if a pulse is sent to modules <code>a</code>, <code>b</code>, and <code>c</code>, and then module <code>a</code> processes its pulse and sends more pulses, the pulses sent to modules <code>b</code> and <code>c</code> would have to be handled first.


The module configuration (your puzzle input) lists each module. The name of the module is preceded by a symbol identifying its type, if any. The name is then followed by an arrow and a list of its destination modules. For example:


<pre><code>broadcaster -&gt; a, b, c
%a -&gt; b
%b -&gt; c
%c -&gt; inv
&amp;inv -&gt; a
</code></pre>
In this module configuration, the broadcaster has three destination modules named <code>a</code>, <code>b</code>, and <code>c</code>. Each of these modules is a flip-flop module (as indicated by the <code>%</code> prefix). <code>a</code> outputs to <code>b</code> which outputs to <code>c</code> which outputs to another module named <code>inv</code>. <code>inv</code> is a conjunction module (as indicated by the <code>&amp;</code> prefix) which, because it has only one input, acts like an <span title="This puzzle originally had a separate inverter module type until I realized it was just a worse conjunction module.">inverter</span> (it sends the opposite of the pulse type it receives); it outputs to <code>a</code>.


By pushing the button once, the following pulses are sent:


<pre><code>button -low-&gt; broadcaster
broadcaster -low-&gt; a
broadcaster -low-&gt; b
broadcaster -low-&gt; c
a -high-&gt; b
b -high-&gt; c
c -high-&gt; inv
inv -low-&gt; a
a -low-&gt; b
b -low-&gt; c
c -low-&gt; inv
inv -high-&gt; a
</code></pre>
After this sequence, the flip-flop modules all end up <em><b>off</b></em>, so pushing the button again repeats the same sequence.


Here's a more interesting example:


<pre><code>broadcaster -> a
%a -> inv, con
&amp;inv -> b
%b -> con
&amp;con -> output
</code></pre>
This module configuration includes the <code>broadcaster</code>, two flip-flops (named <code>a</code> and <code>b</code>), a single-input conjunction module (<code>inv</code>), a multi-input conjunction module (<code>con</code>), and an untyped module named <code>output</code> (for testing purposes). The multi-input conjunction module <code>con</code> watches the two flip-flop modules and, if they're both on, sends a <em><b>low pulse</b></em> to the <code>output</code> module.


Here's what happens if you push the button once:


<pre><code>button -low-&gt; broadcaster
broadcaster -low-&gt; a
a -high-&gt; inv
a -high-&gt; con
inv -low-&gt; b
con -high-&gt; output
b -high-&gt; con
con -low-&gt; output
</code></pre>
Both flip-flops turn on and a low pulse is sent to <code>output</code>! However, now that both flip-flops are on and <code>con</code> remembers a high pulse from each of its two inputs, pushing the button a second time does something different:


<pre><code>button -low-&gt; broadcaster
broadcaster -low-&gt; a
a -low-&gt; inv
a -low-&gt; con
inv -high-&gt; b
con -high-&gt; output
</code></pre>
Flip-flop <code>a</code> turns off! Now, <code>con</code> remembers a low pulse from module <code>a</code>, and so it sends only a high pulse to <code>output</code>.


Push the button a third time:


<pre><code>button -low-&gt; broadcaster
broadcaster -low-&gt; a
a -high-&gt; inv
a -high-&gt; con
inv -low-&gt; b
con -low-&gt; output
b -low-&gt; con
con -high-&gt; output
</code></pre>
This time, flip-flop <code>a</code> turns on, then flip-flop <code>b</code> turns off. However, before <code>b</code> can turn off, the pulse sent to <code>con</code> is handled first, so it <em><b>briefly remembers all high pulses</b></em> for its inputs and sends a low pulse to <code>output</code>. After that, flip-flop <code>b</code> turns off, which causes <code>con</code> to update its state and send a high pulse to <code>output</code>.


Finally, with <code>a</code> on and <code>b</code> off, push the button a fourth time:


<pre><code>button -low-&gt; broadcaster
broadcaster -low-&gt; a
a -low-&gt; inv
a -low-&gt; con
inv -high-&gt; b
con -high-&gt; output
</code></pre>
This completes the cycle: <code>a</code> turns off, causing <code>con</code> to remember only low pulses and restoring all modules to their original states.


To get the cables warmed up, the Elves have pushed the button <code>1000</code> times. How many pulses got sent as a result (including the pulses sent by the button itself)?


In the first example, the same thing happens every time the button is pushed: <code>8</code> low pulses and <code>4</code> high pulses are sent. So, after pushing the button <code>1000</code> times, <code>8000</code> low pulses and <code>4000</code> high pulses are sent. Multiplying these together gives <code><em><b>32000000</b></em></code>.


In the second example, after pushing the button <code>1000</code> times, <code>4250</code> low pulses and <code>2750</code> high pulses are sent. Multiplying these together gives <code><em><b>11687500</b></em></code>.


Consult your module configuration; determine the number of low pulses and high pulses that would be sent after pushing the button <code>1000</code> times, waiting for all pulses to be fully handled after each push of the button. <em><b>What do you get if you multiply the total number of low pulses sent by the total number of high pulses sent?</b></em>


