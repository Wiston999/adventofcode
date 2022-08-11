# --- Day 25: Clock Signal ---

You open the door and find yourself on the roof. The city sprawls away from you for miles and miles.


There's not much time now - it's already Christmas, but you're nowhere near the North Pole, much too far to deliver these stars to the sleigh in time.


However, maybe the <em><b>huge antenna</b></em> up here can offer a solution. After all, the sleigh doesn't need the stars, exactly; it needs the timing data they provide, and you happen to have a massive signal generator right here.


You connect the stars you have to your prototype computer, connect that to the antenna, and begin the transmission.


<span title="Then again, if something ever works on the first try, you should be *very* suspicious.">Nothing happens.</span>


You call the service number printed on the side of the antenna and quickly explain the situation. "I'm not sure what kind of equipment you have connected over there," he says, "but you need a clock signal." You try to explain that this is a signal for a clock.


"No, no, a [https://en.wikipedia.org/wiki/Clock_signal](clock signal) - timing information so the antenna computer knows how to read the data you're sending it. An endless, alternating pattern of <code>0</code>, <code>1</code>, <code>0</code>, <code>1</code>, <code>0</code>, <code>1</code>, <code>0</code>, <code>1</code>, <code>0</code>, <code>1</code>...." He trails off.


You ask if the antenna can handle a clock signal at the frequency you would need to use for the data from the stars. "There's <em><b>no way</b></em> it can! The only antenna we've installed capable of <em><b>that</b></em> is on top of a top-secret Easter Bunny installation, and you're <em><b>definitely</b></em> not-" You hang up the phone.


You've extracted the antenna's clock signal generation [12](assembunny) code (your puzzle input); it looks mostly compatible with code you worked on [23](just recently).


This antenna code, being a signal generator, uses one extra instruction:


<ul>
<li><code>out x</code> <em><b>transmits</b></em> <code>x</code> (either an integer or the <em><b>value</b></em> of a register) as the next value for the clock signal.</li>
</ul>
The code takes a value (via register <code>a</code>) that describes the signal to generate, but you're not sure how it's used. You'll have to find the input to produce the right signal through experimentation.


<em><b>What is the lowest positive integer</b></em> that can be used to initialize register <code>a</code> and cause the code to output a clock signal of <code>0</code>, <code>1</code>, <code>0</code>, <code>1</code>... repeating forever?


