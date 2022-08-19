# --- Day 16: Permutation Promenade ---

You come upon a very unusual sight; a group of programs here appear to be [https://www.youtube.com/watch?v=lyZQPjUT5B4&t=53](dancing).


There are sixteen programs in total, named <code>a</code> through <code>p</code>. They start by standing in a <span title="This is called a 'newline'.">line</span>: <code>a</code> stands in position <code>0</code>, <code>b</code> stands in position <code>1</code>, and so on until <code>p</code>, which stands in position <code>15</code>.


The programs' <em><b>dance</b></em> consists of a sequence of <em><b>dance moves</b></em>:


<ul>
<li><em><b>Spin</b></em>, written <code>sX</code>, makes <code>X</code> programs move from the end to the front, but maintain their order otherwise. (For example, <code>s3</code> on <code>abcde</code> produces <code>cdeab</code>).</li>
<li><em><b>Exchange</b></em>, written <code>xA/B</code>, makes the programs at positions <code>A</code> and <code>B</code> swap places.</li>
<li><em><b>Partner</b></em>, written <code>pA/B</code>, makes the programs named <code>A</code> and <code>B</code> swap places.</li>
</ul>
For example, with only five programs standing in a line (<code>abcde</code>), they could do the following dance:


<ul>
<li><code>s1</code>, a spin of size <code>1</code>: <code>eabcd</code>.</li>
<li><code>x3/4</code>, swapping the last two programs: <code>eabdc</code>.</li>
<li><code>pe/b</code>, swapping programs <code>e</code> and <code>b</code>: <code>baedc</code>.</li>
</ul>
After finishing their dance, the programs end up in order <code>baedc</code>.<p>
<p>You watch the dance for a while and record their dance moves (your puzzle input). <em><b>In what order are the programs standing</b></em> after their dance?


