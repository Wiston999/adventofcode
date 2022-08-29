# --- Day 24: Electromagnetic Moat ---

The CPU itself is a large, black building surrounded by a bottomless pit. Enormous metal tubes extend outward from the side of the building at regular intervals and descend down into the void. There's no way to cross, but you need to get inside.


No way, of course, other than building a <em><b>bridge</b></em> out of the magnetic components strewn about nearby.


Each component has two <em><b>ports</b></em>, one on each end.  The ports come in all different types, and only matching types can be connected.  You take an inventory of the components by their port types (your puzzle input). Each port is identified by the number of <em><b>pins</b></em> it uses; more pins mean a stronger connection for your bridge. A <code>3/7</code> component, for example, has a type-<code>3</code> port on one side, and a type-<code>7</code> port on the other.


Your side of the pit is metallic; a perfect surface to connect a magnetic, <em><b>zero-pin port</b></em>. Because of this, the first port you use must be of type <code>0</code>. It doesn't matter what type of port you end with; your goal is just to make the bridge as strong as possible.


The <em><b>strength</b></em> of a bridge is the sum of the port types in each component. For example, if your bridge is made of components <code>0/3</code>, <code>3/7</code>, and <code>7/4</code>, your bridge has a strength of <code>0+3 + 3+7 + 7+4 = 24</code>.


For example, suppose you had the following components:


<pre><code>0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10
</code></pre>
With them, you could make the following valid bridges:


<ul>
<li><code>0/1</code></li>
<li><code>0/1</code>--<code>10/1</code></li>
<li><code>0/1</code>--<code>10/1</code>--<code>9/10</code></li>
<li><code>0/2</code></li>
<li><code>0/2</code>--<code>2/3</code></li>
<li><code>0/2</code>--<code>2/3</code>--<code>3/4</code></li>
<li><code>0/2</code>--<code>2/3</code>--<code>3/5</code></li>
<li><code>0/2</code>--<code>2/2</code></li>
<li><code>0/2</code>--<code>2/2</code>--<code>2/3</code></li>
<li><code>0/2</code>--<code>2/2</code>--<code>2/3</code>--<code>3/4</code></li>
<li><code>0/2</code>--<code>2/2</code>--<code>2/3</code>--<code>3/5</code></li>
</ul>
(Note how, as shown by <code>10/1</code>, order of ports within a component doesn't matter. However, you may only use each port on a component once.)


Of these bridges, the <em><b>strongest</b></em> one is <code>0/1</code>--<code>10/1</code>--<code>9/10</code>; it has a strength of <code>0+1 + 1+10 + 10+9 = <em><b>31</b></em></code>.


<em><b>What is the strength of the strongest bridge you can make</b></em> with the components you have available?


