# --- Day 13: Claw Contraption ---

Next up: the [/2020/day/24](lobby) of a resort on a tropical island. The Historians take a moment to admire the hexagonal floor tiles before spreading out.


Fortunately, it looks like the resort has a new [https://en.wikipedia.org/wiki/Amusement_arcade](arcade)! Maybe you can win some prizes from the [https://en.wikipedia.org/wiki/Claw_machine](claw machines)?


The claw machines here are a little unusual. Instead of a joystick or directional buttons to control the claw, these machines have two buttons labeled <code>A</code> and <code>B</code>. Worse, you can't just put in a token and play; it costs <em><b>3 tokens</b></em> to push the <code>A</code> button and <em><b>1 token</b></em> to push the <code>B</code> button.


With a little experimentation, you figure out that each machine's buttons are configured to move the claw a specific amount to the <em><b>right</b></em> (along the <code>X</code> axis) and a specific amount <em><b>forward</b></em> (along the <code>Y</code> axis) each time that button is pressed.


Each machine contains one <em><b>prize</b></em>; to win the prize, the claw must be positioned <em><b>exactly</b></em> above the prize on both the <code>X</code> and <code>Y</code> axes.


You wonder: what is the smallest number of tokens you would have to spend to win as many prizes as possible? You assemble a list of every machine's button behavior and prize location (your puzzle input). For example:


<pre><code>Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279
</code></pre>
This list describes the button configuration and prize location of four different claw machines.


For now, consider just the first claw machine in the list:


<ul>
<li>Pushing the machine's <code>A</code> button would move the claw <code>94</code> units along the <code>X</code> axis and <code>34</code> units along the <code>Y</code> axis.</li>
<li>Pushing the <code>B</code> button would move the claw <code>22</code> units along the <code>X</code> axis and <code>67</code> units along the <code>Y</code> axis.</li>
<li>The prize is located at <code>X=8400</code>, <code>Y=5400</code>; this means that from the claw's initial position, it would need to move exactly <code>8400</code> units along the <code>X</code> axis and exactly <code>5400</code> units along the <code>Y</code> axis to be perfectly aligned with the prize in this machine.</li>
</ul>
The cheapest way to win the prize is by pushing the <code>A</code> button <code>80</code> times and the <code>B</code> button <code>40</code> times. This would line up the claw along the <code>X</code> axis (because <code>80*94 + 40*22 = 8400</code>) and along the <code>Y</code> axis (because <code>80*34 + 40*67 = 5400</code>). Doing this would cost <code>80*3</code> tokens for the <span title="Half A presses are not allowed."><code>A</code> presses</span> and <code>40*1</code> for the <code>B</code> presses, a total of <code><em><b>280</b></em></code> tokens.


For the second and fourth claw machines, there is no combination of A and B presses that will ever win a prize.


For the third claw machine, the cheapest way to win the prize is by pushing the <code>A</code> button <code>38</code> times and the <code>B</code> button <code>86</code> times. Doing this would cost a total of <code><em><b>200</b></em></code> tokens.


So, the most prizes you could possibly win is two; the minimum tokens you would have to spend to win all (two) prizes is <code><em><b>480</b></em></code>.


You estimate that each button would need to be pressed <em><b>no more than <code>100</code> times</b></em> to win a prize. How else would someone be expected to play?


Figure out how to win as many prizes as possible. <em><b>What is the fewest tokens you would have to spend to win all possible prizes?</b></em>


