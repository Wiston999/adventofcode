# --- Day 15: Lens Library ---

The newly-focused parabolic reflector dish is sending all of the collected light to a point on the side of yet another mountain - the largest mountain on Lava Island. As you approach the mountain, you find that the light is being collected by the wall of a large facility embedded in the mountainside.


You find a door under a large sign that says "Lava Production Facility" and next to a smaller sign that says "Danger - Personal Protective Equipment required beyond this point".


As you step inside, you are immediately greeted by a somewhat panicked <span title="do you like my hard hat">reindeer</span> wearing goggles and a loose-fitting [https://en.wikipedia.org/wiki/Hard_hat](hard hat). The reindeer leads you to a shelf of goggles and hard hats (you quickly find some that fit) and then further into the facility. At one point, you pass a button with a faint snout mark and the label "PUSH FOR HELP". No wonder you were loaded into that [1](trebuchet) so quickly!


You pass through a final set of doors surrounded with even more warning signs and into what must be the room that collects all of the light from outside. As you admire the large assortment of lenses available to further focus the light, the reindeer brings you a book titled "Initialization Manual".


"Hello!", the book cheerfully begins, apparently unaware of the concerned reindeer reading over your shoulder. "This procedure will let you bring the Lava Production Facility online - all without burning or melting anything unintended!"


"Before you begin, please be prepared to use the Holiday ASCII String Helper algorithm (appendix 1A)." You turn to appendix 1A. The reindeer leans closer with interest.


The HASH algorithm is a way to turn any [https://en.wikipedia.org/wiki/String_(computer_science)](string) of characters into a single <em><b>number</b></em> in the range 0 to 255. To run the HASH algorithm on a string, start with a <em><b>current value</b></em> of <code>0</code>. Then, for each character in the string starting from the beginning:


<ul>
<li>Determine the [https://en.wikipedia.org/wiki/ASCII#Printable_characters](ASCII code) for the current character of the string.</li>
<li>Increase the <em><b>current value</b></em> by the ASCII code you just determined.</li>
<li>Set the <em><b>current value</b></em> to itself multiplied by <code>17</code>.</li>
<li>Set the <em><b>current value</b></em> to the [https://en.wikipedia.org/wiki/Modulo](remainder) of dividing itself by <code>256</code>.</li>
</ul>
After following these steps for each character in the string in order, the <em><b>current value</b></em> is the output of the HASH algorithm.


So, to find the result of running the HASH algorithm on the string <code>HASH</code>:


<ul>
<li>The <em><b>current value</b></em> starts at <code>0</code>.</li>
<li>The first character is <code>H</code>; its ASCII code is <code>72</code>.</li>
<li>The <em><b>current value</b></em> increases to <code>72</code>.</li>
<li>The <em><b>current value</b></em> is multiplied by <code>17</code> to become <code>1224</code>.</li>
<li>The <em><b>current value</b></em> becomes <code><em><b>200</b></em></code> (the remainder of <code>1224</code> divided by <code>256</code>).</li>
<li>The next character is <code>A</code>; its ASCII code is <code>65</code>.</li>
<li>The <em><b>current value</b></em> increases to <code>265</code>.</li>
<li>The <em><b>current value</b></em> is multiplied by <code>17</code> to become <code>4505</code>.</li>
<li>The <em><b>current value</b></em> becomes <code><em><b>153</b></em></code> (the remainder of <code>4505</code> divided by <code>256</code>).</li>
<li>The next character is <code>S</code>; its ASCII code is <code>83</code>.</li>
<li>The <em><b>current value</b></em> increases to <code>236</code>.</li>
<li>The <em><b>current value</b></em> is multiplied by <code>17</code> to become <code>4012</code>.</li>
<li>The <em><b>current value</b></em> becomes <code><em><b>172</b></em></code> (the remainder of <code>4012</code> divided by <code>256</code>).</li>
<li>The next character is <code>H</code>; its ASCII code is <code>72</code>.</li>
<li>The <em><b>current value</b></em> increases to <code>244</code>.</li>
<li>The <em><b>current value</b></em> is multiplied by <code>17</code> to become <code>4148</code>.</li>
<li>The <em><b>current value</b></em> becomes <code><em><b>52</b></em></code> (the remainder of <code>4148</code> divided by <code>256</code>).</li>
</ul>
So, the result of running the HASH algorithm on the string <code>HASH</code> is <code><em><b>52</b></em></code>.


The <em><b>initialization sequence</b></em> (your puzzle input) is a comma-separated list of steps to start the Lava Production Facility. <em><b>Ignore newline characters</b></em> when parsing the initialization sequence. To verify that your HASH algorithm is working, the book offers the sum of the result of running the HASH algorithm on each step in the initialization sequence.


For example:


<pre><code>rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7</code></pre>
This initialization sequence specifies 11 individual steps; the result of running the HASH algorithm on each of the steps is as follows:


<ul>
<li><code>rn=1</code> becomes <code><em><b>30</b></em></code>.</li>
<li><code>cm-</code> becomes <code><em><b>253</b></em></code>.</li>
<li><code>qp=3</code> becomes <code><em><b>97</b></em></code>.</li>
<li><code>cm=2</code> becomes <code><em><b>47</b></em></code>.</li>
<li><code>qp-</code> becomes <code><em><b>14</b></em></code>.</li>
<li><code>pc=4</code> becomes <code><em><b>180</b></em></code>.</li>
<li><code>ot=9</code> becomes <code><em><b>9</b></em></code>.</li>
<li><code>ab=5</code> becomes <code><em><b>197</b></em></code>.</li>
<li><code>pc-</code> becomes <code><em><b>48</b></em></code>.</li>
<li><code>pc=6</code> becomes <code><em><b>214</b></em></code>.</li>
<li><code>ot=7</code> becomes <code><em><b>231</b></em></code>.</li>
</ul>
In this example, the sum of these results is <code><em><b>1320</b></em></code>. Unfortunately, the reindeer has stolen the page containing the expected verification number and is currently running around the facility with it excitedly.


Run the HASH algorithm on each step in the initialization sequence. <em><b>What is the sum of the results?</b></em> (The initialization sequence is one long line; be careful when copy-pasting it.)


