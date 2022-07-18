# --- Day 8: Two-Factor Authentication ---

You come across a door implementing what you can only assume is an implementation of [https://en.wikipedia.org/wiki/Multi-factor_authentication](two-factor authentication) after a long game of [https://en.wikipedia.org/wiki/Requirement](requirements) [https://en.wikipedia.org/wiki/Chinese_whispers](telephone).


To get past the door, you first swipe a keycard (no problem; there was one on a nearby desk). Then, it displays a code on a [https://www.google.com/search?q=tiny+lcd&tbm=isch](little screen), and you type that code on a keypad. Then, presumably, the door unlocks.


Unfortunately, the screen has been <span title="BUT BY WHOM?!">smashed</span>. After a few minutes, you've taken everything apart and figured out how it works. Now you just have to work out what the screen <em><b>would</b></em> have displayed.


The magnetic strip on the card you swiped encodes a series of instructions for the screen; these instructions are your puzzle input. The screen is <em><b><code>50</code> pixels wide and <code>6</code> pixels tall</b></em>, all of which start <em><b>off</b></em>, and is capable of three somewhat peculiar operations:


<ul>
<li><code>rect AxB</code> turns <em><b>on</b></em> all of the pixels in a rectangle at the top-left of the screen which is <code>A</code> wide and <code>B</code> tall.</li>
<li><code>rotate row y=A by B</code> shifts all of the pixels in row <code>A</code> (0 is the top row) <em><b>right</b></em> by <code>B</code> pixels. Pixels that would fall off the right end appear at the left end of the row.</li>
<li><code>rotate column x=A by B</code> shifts all of the pixels in column <code>A</code> (0 is the left column) <em><b>down</b></em> by <code>B</code> pixels. Pixels that would fall off the bottom appear at the top of the column.</li>
</ul>
For example, here is a simple sequence on a smaller screen:


<ul>
<li><code>rect 3x2</code> creates a small rectangle in the top-left corner:

<pre><code>###....
###....
.......</code></pre></li>
<li><code>rotate column x=1 by 1</code> rotates the second column down by one pixel:

<pre><code>#.#....
###....
.#.....</code></pre></li>
<li><code>rotate row y=0 by 4</code> rotates the top row right by four pixels:

<pre><code>....#.#
###....
.#.....</code></pre></li>
<li><code>rotate column x=1 by 1</code> again rotates the second column down by one pixel, causing the bottom pixel to wrap back to the top:

<pre><code>.#..#.#
#.#....
.#.....</code></pre></li>
</ul>
As you can see, this display technology is extremely powerful, and will soon dominate the tiny-code-displaying-screen market.  That's what the advertisement on the back of the display tries to convince you, anyway.


There seems to be an intermediate check of the voltage used by the display: after you swipe your card, if the screen did work, <em><b>how many pixels should be lit?</b></em>


