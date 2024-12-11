# --- Day 7: Bridge Repair ---

The Historians take you to a familiar [/2022/day/9](rope bridge) over a river in the middle of a jungle. The Chief isn't on this side of the bridge, though; maybe he's on the other side?


When you go to cross the bridge, you notice a group of engineers trying to repair it. (Apparently, it breaks pretty frequently.) You won't be able to cross until it's fixed.


You ask how long it'll take; the engineers tell you that it only needs final calibrations, but some young elephants were playing nearby and <em><b>stole all the operators</b></em> from their calibration equations! They could finish the calibrations if only someone could determine which test values could possibly be produced by placing any combination of operators into their calibration equations (your puzzle input).


For example:


<pre><code>190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
</code></pre>
Each line represents a single equation. The test value appears before the colon on each line; it is your job to determine whether the remaining numbers can be combined with operators to produce the test value.


Operators are <em><b>always evaluated left-to-right</b></em>, <em><b>not</b></em> according to precedence rules. Furthermore, numbers in the equations cannot be rearranged. Glancing into the jungle, you can see elephants holding two different types of operators: <em><b>add</b></em> (<code>+</code>) and <em><b>multiply</b></em> (<code>*</code>).


Only three of the above equations can be made true by inserting operators:


<ul>
<li><code>190: 10 19</code> has only one position that accepts an operator: between <code>10</code> and <code>19</code>. Choosing <code>+</code> would give <code>29</code>, but choosing <code>*</code> would give the test value (<code>10 * 19 = 190</code>).</li>
<li><code>3267: 81 40 27</code> has two positions for operators. Of the four possible configurations of the operators, <em><b>two</b></em> cause the right side to match the test value: <code>81 + 40 * 27</code> and <code>81 * 40 + 27</code> both equal <code>3267</code> (when evaluated left-to-right)!</li>
<li><code>292: 11 6 16 20</code> can be solved in exactly one way: <code>11 + 6 * 16 + 20</code>.</li>
</ul>
The engineers just need the <em><b>total calibration result</b></em>, which is the sum of the test values from just the equations that could possibly be true. In the above example, the sum of the test values for the three equations listed above is <code><em><b>3749</b></em></code>.


Determine which equations could possibly be true. <em><b>What is their total calibration result?</b></em>


