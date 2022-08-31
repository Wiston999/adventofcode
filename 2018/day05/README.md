# --- Day 5: Alchemical Reduction ---

You've managed to sneak in to the prototype suit manufacturing lab.  The Elves are making decent progress, but are still struggling with the suit's size reduction capabilities.


While the very latest in 1518 alchemical technology might have solved their problem eventually, you can do better.  You scan the chemical composition of the suit's material and discover that it is formed by extremely long [https://en.wikipedia.org/wiki/Polymer](polymers) (one of which is <span title="I've always wanted a polymer!">available</span> as your puzzle input).


The polymer is formed by smaller <em><b>units</b></em> which, when triggered, react with each other such that two adjacent units of the same type and opposite polarity are destroyed. Units' types are represented by letters; units' polarity is represented by capitalization.  For instance, <code>r</code> and <code>R</code> are units with the same type but opposite polarity, whereas <code>r</code> and <code>s</code> are entirely different types and do not react.


For example:


<ul>
<li>In <code>aA</code>, <code>a</code> and <code>A</code> react, leaving nothing behind.</li>
<li>In <code>abBA</code>, <code>bB</code> destroys itself, leaving <code>aA</code>.  As above, this then destroys itself, leaving nothing.</li>
<li>In <code>abAB</code>, no two adjacent units are of the same type, and so nothing happens.</li>
<li>In <code>aabAAB</code>, even though <code>aa</code> and <code>AA</code> are of the same type, their polarities match, and so nothing happens.</li>
</ul>
Now, consider a larger example, <code>dabAcCaCBAcCcaDA</code>:


<pre><code>dabA<em><b>cC</b></em>aCBAcCcaDA  The first 'cC' is removed.
dab<em><b>Aa</b></em>CBAcCcaDA    This creates 'Aa', which is removed.
dabCBA<em><b>cCc</b></em>aDA      Either 'cC' or 'Cc' are removed (the result is the same).
dabCBAcaDA        No further actions can be taken.
</code></pre>
After all possible reactions, the resulting polymer contains <em><b>10 units</b></em>.


<em><b>How many units remain after fully reacting the polymer you scanned?</b></em> <span class="quiet">(Note: in this puzzle and others, the input is large; if you copy/paste your input, make sure you get the whole thing.)</span>


