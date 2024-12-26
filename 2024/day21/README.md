# --- Day 21: Keypad Conundrum ---

As you teleport onto Santa's [/2019/day/25](Reindeer-class starship), The Historians begin to panic: someone from their search party is <em><b>missing</b></em>. A quick life-form scan by the ship's computer reveals that when the missing Historian teleported, he arrived in another part of the ship.


The door to that area is locked, but the computer can't open it; it can only be opened by <em><b>physically typing</b></em> the door codes (your puzzle input) on the numeric keypad on the door.


The numeric keypad has four rows of buttons: <code>789</code>, <code>456</code>, <code>123</code>, and finally an empty gap followed by <code>0A</code>. Visually, they are arranged like this:


<pre><code>+---+---+---+
| 7 | 8 | 9 |
+---+---+---+
| 4 | 5 | 6 |
+---+---+---+
| 1 | 2 | 3 |
+---+---+---+
    | 0 | A |
    +---+---+
</code></pre>
Unfortunately, the area outside the door is currently <em><b>depressurized</b></em> and nobody can go near the door. A robot needs to be sent instead.


The robot has no problem navigating the ship and finding the numeric keypad, but it's not designed for button pushing: it can't be told to push a specific button directly. Instead, it has a robotic arm that can be controlled remotely via a <em><b>directional keypad</b></em>.


The directional keypad has two rows of buttons: a gap / <code>^</code> (up) / <code>A</code> (activate) on the first row and <code>&lt;</code> (left) / <code>v</code> (down) / <code>&gt;</code> (right) on the second row. Visually, they are arranged like this:


<pre><code>    +---+---+
    | ^ | A |
+---+---+---+
| &lt; | v | &gt; |
+---+---+---+
</code></pre>
When the robot arrives at the numeric keypad, its robotic arm is pointed at the <code>A</code> button in the bottom right corner. After that, this directional keypad remote control must be used to maneuver the robotic arm: the up / down / left / right buttons cause it to move its arm one button in that direction, and the <code>A</code> button causes the robot to briefly move forward, pressing the button being aimed at by the robotic arm.


For example, to make the robot type <code>029A</code> on the numeric keypad, one sequence of inputs on the directional keypad you could use is:


<ul>
<li><code>&lt;</code> to move the arm from <code>A</code> (its initial position) to <code>0</code>.</li>
<li><code>A</code> to push the <code>0</code> button.</li>
<li><code>^A</code> to move the arm to the <code>2</code> button and push it.</li>
<li><code>&gt;^^A</code> to move the arm to the <code>9</code> button and push it.</li>
<li><code>vvvA</code> to move the arm to the <code>A</code> button and push it.</li>
</ul>
In total, there are three shortest possible sequences of button presses on this directional keypad that would cause the robot to type <code>029A</code>: <code>&lt;A^A&gt;^^AvvvA</code>, <code>&lt;A^A^&gt;^AvvvA</code>, and <code>&lt;A^A^^&gt;AvvvA</code>.


Unfortunately, the area containing this directional keypad remote control is currently experiencing <em><b>high levels of radiation</b></em> and nobody can go near it. A robot needs to be sent instead.


When the robot arrives at the directional keypad, its robot arm is pointed at the <code>A</code> button in the upper right corner. After that, a <em><b>second, different</b></em> directional keypad remote control is used to control this robot (in the same way as the first robot, except that this one is typing on a directional keypad instead of a numeric keypad).


There are multiple shortest possible sequences of directional keypad button presses that would cause this robot to tell the first robot to type <code>029A</code> on the door. One such sequence is <code>v&lt;&lt;A&gt;&gt;^A&lt;A&gt;AvA&lt;^AA&gt;A&lt;vAAA&gt;^A</code>.


Unfortunately, the area containing this second directional keypad remote control is currently <em><b><code>-40</code> degrees</b></em>! Another robot will need to be sent to type on that directional keypad, too.


There are many shortest possible sequences of directional keypad button presses that would cause this robot to tell the second robot to tell the first robot to eventually type <code>029A</code> on the door. One such sequence is <code>&lt;vA&lt;AA&gt;&gt;^AvAA&lt;^A&gt;A&lt;v&lt;A&gt;&gt;^AvA^A&lt;vA&gt;^A&lt;v&lt;A&gt;^A&gt;AAvA^A&lt;v&lt;A&gt;A&gt;^AAAvA&lt;^A&gt;A</code>.


Unfortunately, the area containing this third directional keypad remote control is currently <em><b>full of Historians</b></em>, so no robots can find a clear path there. Instead, <em><b>you</b></em> will have to type this sequence yourself.


Were you to choose this sequence of button presses, here are all of the buttons that would be pressed on your directional keypad, the two robots' directional keypads, and the numeric keypad:


<pre><code>&lt;vA&lt;AA&gt;&gt;^AvAA&lt;^A&gt;A&lt;v&lt;A&gt;&gt;^AvA^A&lt;vA&gt;^A&lt;v&lt;A&gt;^A&gt;AAvA^A&lt;v&lt;A&gt;A&gt;^AAAvA&lt;^A&gt;A
v&lt;&lt;A&gt;&gt;^A&lt;A&gt;AvA&lt;^AA&gt;A&lt;vAAA&gt;^A
&lt;A^A&gt;^^AvvvA
029A
</code></pre>
In summary, there are the following keypads:


<ul>
<li>One directional keypad that <em><b>you</b></em> are using.</li>
<li>Two directional keypads that <em><b>robots</b></em> are using.</li>
<li>One numeric keypad (on a door) that a <em><b>robot</b></em> is using.</li>
</ul>
It is important to remember that these robots are not designed for button pushing. In particular, if a robot arm is ever aimed at a <em><b>gap</b></em> where no button is present on the keypad, even for an instant, the robot will <em><b>panic</b></em> unrecoverably. So, don't do that. All robots will initially aim at the keypad's <code>A</code> key, wherever it is.


To unlock the door, <em><b>five</b></em> codes will need to be typed on its numeric keypad. For example:


<pre><code>029A
980A
179A
456A
379A
</code></pre>
For each of these, here is a shortest sequence of button presses you could type to cause the desired code to be typed on the numeric keypad:


<pre><code>029A: &lt;vA&lt;AA&gt;&gt;^AvAA&lt;^A&gt;A&lt;v&lt;A&gt;&gt;^AvA^A&lt;vA&gt;^A&lt;v&lt;A&gt;^A&gt;AAvA^A&lt;v&lt;A&gt;A&gt;^AAAvA&lt;^A&gt;A
980A: &lt;v&lt;A&gt;&gt;^AAAvA^A&lt;vA&lt;AA&gt;&gt;^AvAA&lt;^A&gt;A&lt;v&lt;A&gt;A&gt;^AAAvA&lt;^A&gt;A&lt;vA&gt;^A&lt;A&gt;A
179A: &lt;v&lt;A&gt;&gt;^A&lt;vA&lt;A&gt;&gt;^AAvAA&lt;^A&gt;A&lt;v&lt;A&gt;&gt;^AAvA^A&lt;vA&gt;^AA&lt;A&gt;A&lt;v&lt;A&gt;A&gt;^AAAvA&lt;^A&gt;A
456A: &lt;v&lt;A&gt;&gt;^AA&lt;vA&lt;A&gt;&gt;^AAvAA&lt;^A&gt;A&lt;vA&gt;^A&lt;A&gt;A&lt;vA&gt;^A&lt;A&gt;A&lt;v&lt;A&gt;A&gt;^AAvA&lt;^A&gt;A
379A: &lt;v&lt;A&gt;&gt;^AvA^A&lt;vA&lt;AA&gt;&gt;^AAvA&lt;^A&gt;AAvA^A&lt;vA&gt;^AA&lt;A&gt;A&lt;v&lt;A&gt;A&gt;^AAAvA&lt;^A&gt;A
</code></pre>
The Historians are getting nervous; the ship computer doesn't remember whether the missing Historian is trapped in the area containing a <em><b>giant electromagnet</b></em> or <em><b>molten lava</b></em>. You'll need to make sure that for each of the five codes, you find the <em><b>shortest sequence</b></em> of button presses necessary.


The <em><b>complexity</b></em> of a single code (like <code>029A</code>) is equal to the result of multiplying these two values:


<ul>
<li>The <em><b>length of the shortest sequence</b></em> of button presses you need to type on your directional keypad in order to cause the code to be typed on the numeric keypad; for <code>029A</code>, this would be <code>68</code>.</li>
<li>The <em><b>numeric part of the code</b></em> (ignoring leading zeroes); for <code>029A</code>, this would be <code>29</code>.</li>
</ul>
In the above example, complexity of the five codes can be found by calculating <code>68 * 29</code>, <code>60 * 980</code>, <code>68 * 179</code>, <code>64 * 456</code>, and <code>64 * 379</code>. Adding these together produces <code><em><b>126384</b></em></code>.


Find the fewest number of button presses you'll need to perform in order to cause the robot in front of the door to type each code. <em><b>What is the sum of the complexities of the five codes on your list?</b></em>


