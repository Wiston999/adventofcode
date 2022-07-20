# --- Day 10: Balance Bots ---

You come upon a factory in which many robots are [https://www.youtube.com/watch?v=JnkMyfQ5YfY&t=40](zooming around) handing small microchips to each other.


Upon closer examination, you notice that each bot only proceeds when it has <em><b>two</b></em> microchips, and once it does, it gives each one to a different bot or puts it in a marked "output" bin. Sometimes, bots take microchips from "input" bins, too.


Inspecting one of the microchips, it seems like they each contain a single number; the bots must use some logic to decide what to do with each chip. You access the local control computer and download the bots' instructions (your puzzle input).


Some of the instructions specify that a specific-valued microchip should be given to a specific bot; the rest of the instructions indicate what a given bot should do with its <em><b>lower-value</b></em> or <em><b>higher-value</b></em> chip.


For example, consider the following instructions:


<pre><code>value 5 goes to bot 2
bot 2 gives low to bot 1 and high to bot 0
value 3 goes to bot 1
bot 1 gives low to output 1 and high to bot 0
bot 0 gives low to output 2 and high to output 0
value 2 goes to bot 2
</code></pre>
<ul>
<li>Initially, bot <code>1</code> starts with a value-<code>3</code> chip, and bot <code>2</code> starts with a value-<code>2</code> chip and a value-<code>5</code> chip.</li>
<li>Because bot <code>2</code> has two microchips, it gives its lower one (<code>2</code>) to bot <code>1</code> and its higher one (<code>5</code>) to bot <code>0</code>.</li>
<li>Then, bot <code>1</code> has two microchips; it puts the value-<code>2</code> chip in output <code>1</code> and gives the value-<code>3</code> chip to bot <code>0</code>.</li>
<li>Finally, bot <code>0</code> has two microchips; it puts the <code>3</code> in output <code>2</code> and the <code>5</code> in output <code>0</code>.</li>
</ul>
In the end, output bin <code>0</code> contains a value-<code>5</code> microchip, output bin <code>1</code> contains a value-<code>2</code> microchip, and output bin <code>2</code> contains a value-<code>3</code> microchip. In this configuration, bot number <em><b><code>2</code></b></em> is responsible for comparing value-<code>5</code> microchips with value-<code>2</code> microchips.


Based on your instructions, <em><b>what is the number of the bot</b></em> that is responsible for comparing value-<code>61</code> microchips with value-<code>17</code> microchips?


