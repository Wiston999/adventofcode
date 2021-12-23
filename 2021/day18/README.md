# --- Day 18: Snailfish ---

You descend into the ocean trench and encounter some [https://en.wikipedia.org/wiki/Snailfish](snailfish). They say they saw the sleigh keys! They'll even tell you which direction the keys went if you help one of the smaller snailfish with his <em><b><span title="Or 'maths', if you have more than one.">math</span> homework</b></em>.


Snailfish numbers aren't like regular numbers. Instead, every snailfish number is a <em><b>pair</b></em> - an ordered list of two elements. Each element of the pair can be either a regular number or another pair.


Pairs are written as <code>[x,y]</code>, where <code>x</code> and <code>y</code> are the elements within the pair. Here are some example snailfish numbers, one snailfish number per line:


<pre><code>[1,2]
[[1,2],3]
[9,[8,7]]
[[1,9],[8,5]]
[[[[1,2],[3,4]],[[5,6],[7,8]]],9]
[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]
[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]
</code></pre>
This snailfish homework is about <em><b>addition</b></em>. To add two snailfish numbers, form a pair from the left and right parameters of the addition operator. For example, <code>[1,2]</code> + <code>[[3,4],5]</code> becomes <code>[[1,2],[[3,4],5]]</code>.


There's only one problem: <em><b>snailfish numbers must always be reduced</b></em>, and the process of adding two snailfish numbers can result in snailfish numbers that need to be reduced.


To <em><b>reduce a snailfish number</b></em>, you must repeatedly do the first action in this list that applies to the snailfish number:


<ul>
<li>If any pair is <em><b>nested inside four pairs</b></em>, the leftmost such pair <em><b>explodes</b></em>.</li>
<li>If any regular number is <em><b>10 or greater</b></em>, the leftmost such regular number <em><b>splits</b></em>.</li>
</ul>
Once no action in the above list applies, the snailfish number is reduced.


During reduction, at most one action applies, after which the process returns to the top of the list of actions. For example, if <em><b>split</b></em> produces a pair that meets the <em><b>explode</b></em> criteria, that pair <em><b>explodes</b></em> before other <em><b>splits</b></em> occur.


To <em><b>explode</b></em> a pair, the pair's left value is added to the first regular number to the left of the exploding pair (if any), and the pair's right value is added to the first regular number to the right of the exploding pair (if any). Exploding pairs will always consist of two regular numbers. Then, the entire exploding pair is replaced with the regular number <code>0</code>.


Here are some examples of a single explode action:


<ul>
<li><code>[[[[<em><b>[9,8]</b></em>,1],2],3],4]</code> becomes <code>[[[[<em><b>0</b></em>,<em><b>9</b></em>],2],3],4]</code> (the <code>9</code> has no regular number to its left, so it is not added to any regular number).</li>
<li><code>[7,[6,[5,[4,<em><b>[3,2]</b></em>]]]]</code> becomes <code>[7,[6,[5,[<em><b>7</b></em>,<em><b>0</b></em>]]]]</code> (the <code>2</code> has no regular number to its right, and so it is not added to any regular number).</li>
<li><code>[[6,[5,[4,<em><b>[3,2]</b></em>]]],1]</code> becomes <code>[[6,[5,[<em><b>7</b></em>,<em><b>0</b></em>]]],<em><b>3</b></em>]</code>.</li>
<li><code>[[3,[2,[1,<em><b>[7,3]</b></em>]]],[6,[5,[4,[3,2]]]]]</code> becomes <code>[[3,[2,[<em><b>8</b></em>,<em><b>0</b></em>]]],[<em><b>9</b></em>,[5,[4,[3,2]]]]]</code> (the pair <code>[3,2]</code> is unaffected because the pair <code>[7,3]</code> is further to the left; <code>[3,2]</code> would explode on the next action).</li>
<li><code>[[3,[2,[8,0]]],[9,[5,[4,<em><b>[3,2]</b></em>]]]]</code> becomes <code>[[3,[2,[8,0]]],[9,[5,[<em><b>7</b></em>,<em><b>0</b></em>]]]]</code>.</li>
</ul>
To <em><b>split</b></em> a regular number, replace it with a pair; the left element of the pair should be the regular number divided by two and rounded <em><b>down</b></em>, while the right element of the pair should be the regular number divided by two and rounded <em><b>up</b></em>. For example, <code>10</code> becomes <code>[5,5]</code>, <code>11</code> becomes <code>[5,6]</code>, <code>12</code> becomes <code>[6,6]</code>, and so on.


Here is the process of finding the reduced result of <code>[[[[4,3],4],4],[7,[[8,4],9]]]</code> + <code>[1,1]</code>:


<pre><code>after addition: [[[[<em><b>[4,3]</b></em>,4],4],[7,[[8,4],9]]],[1,1]]
after explode:  [[[[0,7],4],[7,[<em><b>[8,4]</b></em>,9]]],[1,1]]
after explode:  [[[[0,7],4],[<em><b>15</b></em>,[0,13]]],[1,1]]
after split:    [[[[0,7],4],[[7,8],[0,<em><b>13</b></em>]]],[1,1]]
after split:    [[[[0,7],4],[[7,8],[0,<em><b>[6,7]</b></em>]]],[1,1]]
after explode:  [[[[0,7],4],[[7,8],[6,0]]],[8,1]]
</code></pre>
Once no reduce actions apply, the snailfish number that remains is the actual result of the addition operation: <code>[[[[0,7],4],[[7,8],[6,0]]],[8,1]]</code>.


The homework assignment involves adding up a <em><b>list of snailfish numbers</b></em> (your puzzle input). The snailfish numbers are each listed on a separate line. Add the first snailfish number and the second, then add that result and the third, then add that result and the fourth, and so on until all numbers in the list have been used once.


For example, the final sum of this list is <code>[[[[1,1],[2,2]],[3,3]],[4,4]]</code>:


<pre><code>[1,1]
[2,2]
[3,3]
[4,4]
</code></pre>
The final sum of this list is <code>[[[[3,0],[5,3]],[4,4]],[5,5]]</code>:


<pre><code>[1,1]
[2,2]
[3,3]
[4,4]
[5,5]
</code></pre>
The final sum of this list is <code>[[[[5,0],[7,4]],[5,5]],[6,6]]</code>:


<pre><code>[1,1]
[2,2]
[3,3]
[4,4]
[5,5]
[6,6]
</code></pre>
Here's a slightly larger example:


<pre><code>[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]
</code></pre>
The final sum <code>[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]</code> is found after adding up the above snailfish numbers:


<pre><code>  [[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
+ [7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
= [[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]

  [[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]
+ [[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
= [[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]

  [[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]
+ [[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
= [[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]

  [[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]
+ [7,[5,[[3,8],[1,4]]]]
= [[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]

  [[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]
+ [[2,[2,2]],[8,[8,1]]]
= [[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]

  [[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]
+ [2,9]
= [[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]

  [[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]
+ [1,[[[9,3],9],[[9,0],[0,7]]]]
= [[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]

  [[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]
+ [[[5,[7,4]],7],1]
= [[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]

  [[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]
+ [[[[4,2],2],6],[8,7]]
= [[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]
</code></pre>
To check whether it's the right answer, the snailfish teacher only checks the <em><b>magnitude</b></em> of the final sum. The magnitude of a pair is 3 times the magnitude of its left element plus 2 times the magnitude of its right element. The magnitude of a regular number is just that number.


For example, the magnitude of <code>[9,1]</code> is <code>3*9 + 2*1 = <em><b>29</b></em></code>; the magnitude of <code>[1,9]</code> is <code>3*1 + 2*9 = <em><b>21</b></em></code>. Magnitude calculations are recursive: the magnitude of <code>[[9,1],[1,9]]</code> is <code>3*29 + 2*21 = <em><b>129</b></em></code>.


Here are a few more magnitude examples:


<ul>
<li><code>[[1,2],[[3,4],5]]</code> becomes <code><em><b>143</b></em></code>.</li>
<li><code>[[[[0,7],4],[[7,8],[6,0]]],[8,1]]</code> becomes <code><em><b>1384</b></em></code>.</li>
<li><code>[[[[1,1],[2,2]],[3,3]],[4,4]]</code> becomes <code><em><b>445</b></em></code>.</li>
<li><code>[[[[3,0],[5,3]],[4,4]],[5,5]]</code> becomes <code><em><b>791</b></em></code>.</li>
<li><code>[[[[5,0],[7,4]],[5,5]],[6,6]]</code> becomes <code><em><b>1137</b></em></code>.</li>
<li><code>[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]</code> becomes <code><em><b>3488</b></em></code>.</li>
</ul>
So, given this example homework assignment:


<pre><code>[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]
</code></pre>
The final sum is:


<pre><code>[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]</code></pre>
The magnitude of this final sum is <code><em><b>4140</b></em></code>.


Add up all of the snailfish numbers from the homework assignment in the order they appear. <em><b>What is the magnitude of the final sum?</b></em>


