# --- Day 19: A Series of Tubes ---

Somehow, a network packet got <span title="I know how fast it's going, but I don't know where it is.">lost</span> and ended up here.  It's trying to follow a routing diagram (your puzzle input), but it's confused about where to go.


Its starting point is just off the top of the diagram. Lines (drawn with <code>|</code>, <code>-</code>, and <code>+</code>) show the path it needs to take, starting by going down onto the only line connected to the top of the diagram. It needs to follow this path until it reaches the end (located somewhere within the diagram) and stop there.


Sometimes, the lines cross over each other; in these cases, it needs to continue going the same direction, and only turn left or right when there's no other option.  In addition, someone has left <em><b>letters</b></em> on the line; these also don't change its direction, but it can use them to keep track of where it's been. For example:


<pre><code>     |          
     |  +--+    
     A  |  C    
 F---|----E|--+ 
     |  |  |  D 
     +B-+  +--+ 

</code></pre>
Given this diagram, the packet needs to take the following path:


<ul>
<li>Starting at the only line touching the top of the diagram, it must go down, pass through <code>A</code>, and continue onward to the first <code>+</code>.</li>
<li>Travel right, up, and right, passing through <code>B</code> in the process.</li>
<li>Continue down (collecting <code>C</code>), right, and up (collecting <code>D</code>).</li>
<li>Finally, go all the way left through <code>E</code> and stopping at <code>F</code>.</li>
</ul>
Following the path to the end, the letters it sees on its path are <code>ABCDEF</code>.


The little packet looks up at you, hoping you can help it find the way.  <em><b>What letters will it see</b></em> (in the order it would see them) if it follows the path? (The routing diagram is very wide; make sure you view it without line wrapping.)


