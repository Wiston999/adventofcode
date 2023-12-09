# --- Day 9: Mirage Maintenance ---

You ride the camel through the sandstorm and stop where the ghost's maps told you to stop. <span title="The sound of a sandstorm slowly settling.">The sandstorm subsequently subsides, somehow seeing you standing at an <em><b>oasis</b></em>!</span>


The camel goes to get some water and you stretch your neck. As you look up, you discover what must be yet another giant floating island, this one made of metal! That must be where the <em><b>parts to fix the sand machines</b></em> come from.


There's even a [https://en.wikipedia.org/wiki/Hang_gliding](hang glider) partially buried in the sand here; once the sun rises and heats up the sand, you might be able to use the glider and the hot air to get all the way up to the metal island!


While you wait for the sun to rise, you admire the oasis hidden here in the middle of Desert Island. It must have a delicate ecosystem; you might as well take some ecological readings while you wait. Maybe you can report any environmental instabilities you find to someone so the oasis can be around for the next sandstorm-worn traveler.


You pull out your handy <em><b>Oasis And Sand Instability Sensor</b></em> and analyze your surroundings. The OASIS produces a report of many values and how they are changing over time (your puzzle input). Each line in the report contains the <em><b>history</b></em> of a single value. For example:


<pre><code>0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
</code></pre>
To best protect the oasis, your environmental report should include a <em><b>prediction of the next value</b></em> in each history. To do this, start by making a new sequence from the <em><b>difference at each step</b></em> of your history. If that sequence is <em><b>not</b></em> all zeroes, repeat this process, using the sequence you just generated as the input sequence. Once all of the values in your latest sequence are zeroes, you can extrapolate what the next value of the original history should be.


In the above dataset, the first history is <code>0 3 6 9 12 15</code>. Because the values increase by <code>3</code> each step, the first sequence of differences that you generate will be <code>3 3 3 3 3</code>. Note that this sequence has one fewer value than the input sequence because at each step it considers two numbers from the input. Since these values aren't <em><b>all zero</b></em>, repeat the process: the values differ by <code>0</code> at each step, so the next sequence is <code>0 0 0 0</code>. This means you have enough information to extrapolate the history! Visually, these sequences can be arranged like this:


<pre><code>0   3   6   9  12  15
  3   3   3   3   3
    0   0   0   0
</code></pre>
To extrapolate, start by adding a new zero to the end of your list of zeroes; because the zeroes represent differences between the two values above them, this also means there is now a placeholder in every sequence above it:<p>
<pre><code>0   3   6   9  12  15   <em><b>B</b></em>
  3   3   3   3   3   <em><b>A</b></em>
    0   0   0   0   <em><b>0</b></em>
</code></pre>
<p>You can then start filling in placeholders from the bottom up. <code>A</code> needs to be the result of increasing <code>3</code> (the value to its left) by <code>0</code> (the value below it); this means <code>A</code> must be <code><em><b>3</b></em></code>:


<pre><code>0   3   6   9  12  15   B
  3   3   3   3   <em><b>3</b></em>   <em><b>3</b></em>
    0   0   0   0   <em><b>0</b></em>
</code></pre>
Finally, you can fill in <code>B</code>, which needs to be the result of increasing <code>15</code> (the value to its left) by <code>3</code> (the value below it), or <code><em><b>18</b></em></code>:


<pre><code>0   3   6   9  12  <em><b>15</b></em>  <em><b>18</b></em>
  3   3   3   3   3   <em><b>3</b></em>
    0   0   0   0   0
</code></pre>
So, the next value of the first history is <code><em><b>18</b></em></code>.


Finding all-zero differences for the second history requires an additional sequence:


<pre><code>1   3   6  10  15  21
  2   3   4   5   6
    1   1   1   1
      0   0   0
</code></pre>
Then, following the same process as before, work out the next value in each sequence from the bottom up:


<pre><code>1   3   6  10  15  21  <em><b>28</b></em>
  2   3   4   5   6   <em><b>7</b></em>
    1   1   1   1   <em><b>1</b></em>
      0   0   0   <em><b>0</b></em>
</code></pre>
So, the next value of the second history is <code><em><b>28</b></em></code>.


The third history requires even more sequences, but its next value can be found the same way:


<pre><code>10  13  16  21  30  45  <em><b>68</b></em>
   3   3   5   9  15  <em><b>23</b></em>
     0   2   4   6   <em><b>8</b></em>
       2   2   2   <em><b>2</b></em>
         0   0   <em><b>0</b></em>
</code></pre>
So, the next value of the third history is <code><em><b>68</b></em></code>.


If you find the next value for each history in this example and add them together, you get <code><em><b>114</b></em></code>.


Analyze your OASIS report and extrapolate the next value for each history. <em><b>What is the sum of these extrapolated values?</b></em>


