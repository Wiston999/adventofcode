# --- Day 15: Dueling Generators ---

Here, you encounter a pair of dueling <span title="I guess they *are* a little banjo-shaped. Why do you ask?">generators</span>. The generators, called <em><b>generator A</b></em> and <em><b>generator B</b></em>, are trying to agree on a sequence of numbers. However, one of them is malfunctioning, and so the sequences don't always match.


As they do this, a <em><b>judge</b></em> waits for each of them to generate its next value, compares the lowest 16 bits of both values, and keeps track of the number of times those parts of the values match.


The generators both work on the same principle. To create its next value, a generator will take the previous value it produced, multiply it by a <em><b>factor</b></em> (generator A uses <code>16807</code>; generator B uses <code>48271</code>), and then keep the remainder of dividing that resulting product by <code>2147483647</code>. That final remainder is the value it produces next.


To calculate each generator's first value, it instead uses a specific starting value as its "previous value" (as listed in your puzzle input).


For example, suppose that for starting values, generator A uses <code>65</code>, while generator B uses <code>8921</code>. Then, the first five pairs of generated values are:


<pre><code>--Gen. A--  --Gen. B--
   1092455   430625591
1181022009  1233683848
 245556042  1431495498
1744312007   137874439
1352636452   285222916
</code></pre>
In binary, these pairs are (with generator A's value first in each pair):


<pre><code>00000000000100001010101101100111
00011001101010101101001100110111

01000110011001001111011100111001
01001001100010001000010110001000

00001110101000101110001101001010
01010101010100101110001101001010

01100111111110000001011011000111
00001000001101111100110000000111

01010000100111111001100000100100
00010001000000000010100000000100
</code></pre>
Here, you can see that the lowest (here, rightmost) 16 bits of the third value match: <code>1110001101001010</code>. Because of this one match, after processing these five pairs, the judge would have added only <code>1</code> to its total.


To get a significant sample, the judge would like to consider <em><b>40 million</b></em> pairs. (In the example above, the judge would eventually find a total of <code>588</code> pairs that match in their lowest 16 bits.)


After 40 million pairs, <em><b>what is the judge's final count</b></em>?


