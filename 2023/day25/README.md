# --- Day 25: Snowverload ---

<em><b>Still</b></em> somehow without snow, you go to the last place you haven't checked: the center of Snow Island, directly below the waterfall.


Here, someone has clearly been trying to fix the problem. Scattered everywhere are hundreds of weather machines, almanacs, communication modules, hoof prints, machine parts, mirrors, lenses, and so on.


Somehow, everything has been <em><b>wired together</b></em> into a massive snow-producing apparatus, but nothing seems to be running. You check a tiny screen on one of the communication modules: <code>Error 2023</code>. It doesn't say what <code>Error 2023</code> means, but it <em><b>does</b></em> have the phone number for a support line printed on it.


"Hi, you've reached Weather Machines And So On, Inc. How can I help you?" You explain the situation.


"Error 2023, you say? Why, that's a power overload error, of course! It means you have too many components plugged in. Try unplugging some components and--" You explain that there are hundreds of components here and you're in a bit of a hurry.


"Well, let's see how bad it is; do you see a <em><b>big red reset button</b></em> somewhere? It should be on its own module. If you push it, it probably won't fix anything, but it'll report how overloaded things are." After a minute or two, you find the reset button; it's so big that it takes two hands just to get enough leverage to push it. Its screen then displays:


<pre><code>SYSTEM OVERLOAD!

Connected components would require
power equal to at least <em class="star">100 stars</em>!
</code></pre>
"Wait, <em><b>how</b></em> many components did you say are plugged in? With that much equipment, you could produce snow for an <em><b>entire</b></em>--" You disconnect the call.


You have nowhere near that many stars - you need to find a way to disconnect at least half of the equipment here, but it's already Christmas! You only have time to disconnect <em><b>three wires</b></em>.


Fortunately, someone left a wiring diagram (your puzzle input) that shows <em><b>how the components are connected</b></em>. For example:


<pre><code>jqt: rhn xhk nvd
rsh: frs pzl lsr
xhk: hfx
cmg: qnr nvd lhk bvb
rhn: xhk bvb hfx
bvb: xhk hfx
pzl: lsr hfx nvd
qnr: nvd
ntq: jqt hfx bvb xhk
nvd: lhk
lsr: lhk
rzs: qnr cmg lsr rsh
frs: qnr lhk lsr
</code></pre>
Each line shows the <em><b>name of a component</b></em>, a colon, and then <em><b>a list of other components</b></em> to which that component is connected. Connections aren't directional; <code>abc: xyz</code> and <code>xyz: abc</code> both represent the same configuration. Each connection between two components is represented only once, so some components might only ever appear on the left or right side of a colon.


In this example, if you disconnect the wire between <code>hfx</code>/<code>pzl</code>, the wire between <code>bvb</code>/<code>cmg</code>, and the wire between <code>nvd</code>/<code>jqt</code>, you will <em><b>divide the components into two separate, disconnected groups</b></em>:


<ul>
<li><code><em><b>9</b></em></code> components: <code>cmg</code>, <code>frs</code>, <code>lhk</code>, <code>lsr</code>, <code>nvd</code>, <code>pzl</code>, <code>qnr</code>, <code>rsh</code>, and <code>rzs</code>.</li>
<li><code><em><b>6</b></em></code> components: <code>bvb</code>, <code>hfx</code>, <code>jqt</code>, <code>ntq</code>, <code>rhn</code>, and <code>xhk</code>.</li>
</ul>
Multiplying the sizes of these groups together produces <code><em><b>54</b></em></code>.


Find the three wires you need to disconnect in order to divide the components into two separate groups. <em><b>What do you get if you multiply the sizes of these two groups together?</b></em>


