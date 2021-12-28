# --- Day 8: Seven Segment Search ---

You barely reach the safety of the cave when the whale smashes into the cave mouth, collapsing it. Sensors indicate another exit to this cave at a much greater depth, so you have no choice but to press on.


As your submarine slowly makes its way through the cave system, you notice that the four-digit [https://en.wikipedia.org/wiki/Seven-segment_display](seven-segment displays) in your submarine are malfunctioning; <span title="Yes, just the four-digit seven-segment ones. Whole batch must have been faulty.">they must have been damaged</span> during the escape. You'll be in a lot of trouble without them, so you'd better figure out what's wrong.


Each digit of a seven-segment display is rendered by turning on or off any of seven segments named <code>a</code> through <code>g</code>:


<pre><code>  0:      1:      2:      3:      4:
 <em><b>aaaa</b></em>    ....    <em><b>aaaa    aaaa</b></em>    ....
<em><b>b    c</b></em>  .    <em><b>c</b></em>  .    <em><b>c</b></em>  .    <em><b>c  b    c</b></em>
<em><b>b    c</b></em>  .    <em><b>c</b></em>  .    <em><b>c</b></em>  .    <em><b>c  b    c</b></em>
 ....    ....    <em><b>dddd    dddd    dddd</b></em>
<em><b>e    f</b></em>  .    <em><b>f  e</b></em>    .  .    <em><b>f</b></em>  .    <em><b>f</b></em>
<em><b>e    f</b></em>  .    <em><b>f  e</b></em>    .  .    <em><b>f</b></em>  .    <em><b>f</b></em>
 <em><b>gggg</b></em>    ....    <em><b>gggg    gggg</b></em>    ....

  5:      6:      7:      8:      9:
 <em><b>aaaa    aaaa    aaaa    aaaa    aaaa</b></em>
<em><b>b</b></em>    .  <em><b>b</b></em>    .  .    <em><b>c  b    c  b    c</b></em>
<em><b>b</b></em>    .  <em><b>b</b></em>    .  .    <em><b>c  b    c  b    c</b></em>
 <em><b>dddd    dddd</b></em>    ....    <em><b>dddd    dddd</b></em>
.    <em><b>f  e    f</b></em>  .    <em><b>f  e    f</b></em>  .    <em><b>f</b></em>
.    <em><b>f  e    f</b></em>  .    <em><b>f  e    f</b></em>  .    <em><b>f</b></em>
 <em><b>gggg    gggg</b></em>    ....    <em><b>gggg    gggg</b></em>
</code></pre>
So, to render a <code>1</code>, only segments <code>c</code> and <code>f</code> would be turned on; the rest would be off. To render a <code>7</code>, only segments <code>a</code>, <code>c</code>, and <code>f</code> would be turned on.


The problem is that the signals which control the segments have been mixed up on each display. The submarine is still trying to display numbers by producing output on signal wires <code>a</code> through <code>g</code>, but those wires are connected to segments <em><b>randomly</b></em>. Worse, the wire/segment connections are mixed up separately for each four-digit display! (All of the digits <em><b>within</b></em> a display use the same connections, though.)


So, you might know that only signal wires <code>b</code> and <code>g</code> are turned on, but that doesn't mean <em><b>segments</b></em> <code>b</code> and <code>g</code> are turned on: the only digit that uses two segments is <code>1</code>, so it must mean segments <code>c</code> and <code>f</code> are meant to be on. With just that information, you still can't tell which wire (<code>b</code>/<code>g</code>) goes to which segment (<code>c</code>/<code>f</code>). For that, you'll need to collect more information.


For each display, you watch the changing signals for a while, make a note of <em><b>all ten unique signal patterns</b></em> you see, and then write down a single <em><b>four digit output value</b></em> (your puzzle input). Using the signal patterns, you should be able to work out which pattern corresponds to which digit.


For example, here is what you might see in a single entry in your notes:


<pre><code>acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |
cdfeb fcadb cdfeb cdbaf</code></pre>
(The entry is wrapped here to two lines so it fits; in your notes, it will all be on a single line.)


Each entry consists of ten <em><b>unique signal patterns</b></em>, a <code>|</code> delimiter, and finally the <em><b>four digit output value</b></em>. Within an entry, the same wire/segment connections are used (but you don't know what the connections actually are). The unique signal patterns correspond to the ten different ways the submarine tries to render a digit using the current wire/segment connections. Because <code>7</code> is the only digit that uses three segments, <code>dab</code> in the above example means that to render a <code>7</code>, signal lines <code>d</code>, <code>a</code>, and <code>b</code> are on. Because <code>4</code> is the only digit that uses four segments, <code>eafb</code> means that to render a <code>4</code>, signal lines <code>e</code>, <code>a</code>, <code>f</code>, and <code>b</code> are on.


Using this information, you should be able to work out which combination of signal wires corresponds to each of the ten digits. Then, you can decode the four digit output value. Unfortunately, in the above example, all of the digits in the output value (<code>cdfeb fcadb cdfeb cdbaf</code>) use five segments and are more difficult to deduce.


For now, <em><b>focus on the easy digits</b></em>. Consider this larger example:


<pre><code>be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb |
<em><b>fdgacbe</b></em> cefdb cefbgd <em><b>gcbe</b></em>
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec |
fcgedb <em><b>cgb</b></em> <em><b>dgebacf</b></em> <em><b>gc</b></em>
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef |
<em><b>cg</b></em> <em><b>cg</b></em> fdcagb <em><b>cbg</b></em>
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega |
efabcd cedba gadfec <em><b>cb</b></em>
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga |
<em><b>gecf</b></em> <em><b>egdcabf</b></em> <em><b>bgf</b></em> bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf |
<em><b>gebdcfa</b></em> <em><b>ecba</b></em> <em><b>ca</b></em> <em><b>fadegcb</b></em>
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf |
<em><b>cefg</b></em> dcbef <em><b>fcge</b></em> <em><b>gbcadfe</b></em>
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd |
<em><b>ed</b></em> bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg |
<em><b>gbdfcae</b></em> <em><b>bgc</b></em> <em><b>cg</b></em> <em><b>cgb</b></em>
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc |
<em><b>fgae</b></em> cfgab <em><b>fg</b></em> bagce
</code></pre>
Because the digits <code>1</code>, <code>4</code>, <code>7</code>, and <code>8</code> each use a unique number of segments, you should be able to tell which combinations of signals correspond to those digits. Counting <em><b>only digits in the output values</b></em> (the part after <code>|</code> on each line), in the above example, there are <code><em><b>26</b></em></code> instances of digits that use a unique number of segments (highlighted above).


<em><b>In the output values, how many times do digits <code>1</code>, <code>4</code>, <code>7</code>, or <code>8</code> appear?</b></em>


