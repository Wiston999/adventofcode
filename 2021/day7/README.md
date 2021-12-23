# --- Day 7: The Treachery of Whales ---

A giant [https://en.wikipedia.org/wiki/Sperm_whale](whale) has decided your submarine is its next meal, and it's much faster than you are. There's nowhere to run!


Suddenly, a swarm of crabs (each in its own tiny submarine - it's too deep for them otherwise) zooms in to rescue you! They seem to be preparing to blast a hole in the ocean floor; sensors indicate a <em><b>massive underground cave system</b></em> just beyond where they're aiming!


The crab submarines all need to be aligned before they'll have enough power to blast a large enough hole for your submarine to get through. However, it doesn't look like they'll be aligned before the whale catches you! Maybe you can help?


There's one major catch - crab submarines can only move horizontally.


You quickly make a list of <em><b>the horizontal position of each crab</b></em> (your puzzle input). Crab submarines have limited fuel, so you need to find a way to make all of their horizontal positions match while requiring them to spend as little fuel as possible.


For example, consider the following horizontal positions:


<pre><code>16,1,2,0,4,2,7,1,2,14</code></pre>
This means there's a crab with horizontal position <code>16</code>, a crab with horizontal position <code>1</code>, and so on.


Each change of 1 step in horizontal position of a single crab costs 1 fuel. You could choose any horizontal position to align them all on, but the one that costs the least fuel is horizontal position <code>2</code>:


<ul>
<li>Move from <code>16</code> to <code>2</code>: <code>14</code> fuel</li>
<li>Move from <code>1</code> to <code>2</code>: <code>1</code> fuel</li>
<li>Move from <code>2</code> to <code>2</code>: <code>0</code> fuel</li>
<li>Move from <code>0</code> to <code>2</code>: <code>2</code> fuel</li>
<li>Move from <code>4</code> to <code>2</code>: <code>2</code> fuel</li>
<li>Move from <code>2</code> to <code>2</code>: <code>0</code> fuel</li>
<li>Move from <code>7</code> to <code>2</code>: <code>5</code> fuel</li>
<li>Move from <code>1</code> to <code>2</code>: <code>1</code> fuel</li>
<li>Move from <code>2</code> to <code>2</code>: <code>0</code> fuel</li>
<li>Move from <code>14</code> to <code>2</code>: <code>12</code> fuel</li>
</ul>
This costs a total of <code><em><b>37</b></em></code> fuel. This is the cheapest possible outcome; more expensive outcomes include aligning at position <code>1</code> (<code>41</code> fuel), position <code>3</code> (<code>39</code> fuel), or position <code>10</code> (<code>71</code> fuel).


Determine the horizontal position that the crabs can align to using the least fuel possible. <em><b>How much fuel must they spend to align to that position?</b></em>


