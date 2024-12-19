# --- Day 19: Linen Layout ---

Today, The Historians take you up to the [/2023/day/12](hot springs) on Gear Island! Very [https://www.youtube.com/watch?v=ekL881PJMjI](suspiciously), absolutely nothing goes wrong as they begin their careful search of the vast field of helixes.


Could this <em><b>finally</b></em> be your chance to visit the [https://en.wikipedia.org/wiki/Onsen](onsen) next door? Only one way to find out.


After a brief conversation with the reception staff at the onsen front desk, you discover that you don't have the right kind of money to pay the admission fee. However, before you can leave, the staff get your attention. Apparently, they've heard about how you helped at the hot springs, and they're willing to make a deal: if you can simply help them <em><b>arrange their towels</b></em>, they'll let you in for <em><b>free</b></em>!


Every towel at this onsen is marked with a <em><b>pattern of colored stripes</b></em>. There are only a few patterns, but for any particular pattern, the staff can get you as many towels with that pattern as you need. Each <span title="It really seems like they've gathered a lot of magic into the towel colors.">stripe</span> can be <em><b>white</b></em> (<code>w</code>), <em><b>blue</b></em> (<code>u</code>), <em><b>black</b></em> (<code>b</code>), <em><b>red</b></em> (<code>r</code>), or <em><b>green</b></em> (<code>g</code>). So, a towel with the pattern <code>ggr</code> would have a green stripe, a green stripe, and then a red stripe, in that order. (You can't reverse a pattern by flipping a towel upside-down, as that would cause the onsen logo to face the wrong way.)


The Official Onsen Branding Expert has produced a list of <em><b>designs</b></em> - each a long sequence of stripe colors - that they would like to be able to display. You can use any towels you want, but all of the towels' stripes must exactly match the desired design. So, to display the design <code>rgrgr</code>, you could use two <code>rg</code> towels and then an <code>r</code> towel, an <code>rgr</code> towel and then a <code>gr</code> towel, or even a single massive <code>rgrgr</code> towel (assuming such towel patterns were actually available).


To start, collect together all of the available towel patterns and the list of desired designs (your puzzle input). For example:


<pre><code>r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb
</code></pre>
The first line indicates the available towel patterns; in this example, the onsen has unlimited towels with a single red stripe (<code>r</code>), unlimited towels with a white stripe and then a red stripe (<code>wr</code>), and so on.


After the blank line, the remaining lines each describe a design the onsen would like to be able to display. In this example, the first design (<code>brwrr</code>) indicates that the onsen would like to be able to display a black stripe, a red stripe, a white stripe, and then two red stripes, in that order.


Not all designs will be possible with the available towels. In the above example, the designs are possible or impossible as follows:


<ul>
<li><code>brwrr</code> can be made with a <code>br</code> towel, then a <code>wr</code> towel, and then finally an <code>r</code> towel.</li>
<li><code>bggr</code> can be made with a <code>b</code> towel, two <code>g</code> towels, and then an <code>r</code> towel.</li>
<li><code>gbbr</code> can be made with a <code>gb</code> towel and then a <code>br</code> towel.</li>
<li><code>rrbgbr</code> can be made with <code>r</code>, <code>rb</code>, <code>g</code>, and <code>br</code>.</li>
<li><code>ubwu</code> is <em><b>impossible</b></em>.</li>
<li><code>bwurrg</code> can be made with <code>bwu</code>, <code>r</code>, <code>r</code>, and <code>g</code>.</li>
<li><code>brgr</code> can be made with <code>br</code>, <code>g</code>, and <code>r</code>.</li>
<li><code>bbrgwb</code> is <em><b>impossible</b></em>.</li>
</ul>
In this example, <code><em><b>6</b></em></code> of the eight designs are possible with the available towel patterns.


To get into the onsen as soon as possible, consult your list of towel patterns and desired designs carefully. <em><b>How many designs are possible?</b></em>


