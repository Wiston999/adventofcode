# --- Day 2: Inventory Management System ---

You stop falling through time, catch your breath, and check the screen on the device. "Destination reached. Current Year: 1518. Current Location: North Pole Utility Closet 83N10." You made it! Now, to find those anomalies.


Outside the utility closet, you hear footsteps and a voice. "...I'm not sure either. But now that <span title="This is, in fact, roughly when chimneys became common in houses.">so many people have chimneys</span>, maybe he could sneak in that way?" Another voice responds, "Actually, we've been working on a new kind of <em><b>suit</b></em> that would let him fit through tight spaces like that. But, I heard that a few days ago, they lost the prototype fabric, the design plans, everything! Nobody on the team can even seem to remember important details of the project!"


"Wouldn't they have had enough fabric to fill several boxes in the warehouse? They'd be stored together, so the box IDs should be similar. Too bad it would take forever to search the warehouse for <em><b>two similar box IDs</b></em>..." They walk too far away to hear any more.


Late at night, you sneak to the warehouse - who knows what kinds of paradoxes you could cause if you were discovered - and use your fancy wrist device to quickly scan every box and produce a list of the likely candidates (your puzzle input).


To make sure you didn't miss any, you scan the likely candidate boxes again, counting the number that have an ID containing <em><b>exactly two of any letter</b></em> and then separately counting those with <em><b>exactly three of any letter</b></em>. You can multiply those two counts together to get a rudimentary [https://en.wikipedia.org/wiki/Checksum](checksum) and compare it to what your device predicts.


For example, if you see the following box IDs:


<ul>
<li><code>abcdef</code> contains no letters that appear exactly two or three times.</li>
<li><code>bababc</code> contains two <code>a</code> and three <code>b</code>, so it counts for both.</li>
<li><code>abbcde</code> contains two <code>b</code>, but no letter appears exactly three times.</li>
<li><code>abcccd</code> contains three <code>c</code>, but no letter appears exactly two times.</li>
<li><code>aabcdd</code> contains two <code>a</code> and two <code>d</code>, but it only counts once.</li>
<li><code>abcdee</code> contains two <code>e</code>.</li>
<li><code>ababab</code> contains three <code>a</code> and three <code>b</code>, but it only counts once.</li>
</ul>
Of these box IDs, four of them contain a letter which appears exactly twice, and three of them contain a letter which appears exactly three times. Multiplying these together produces a checksum of <code>4 * 3 = 12</code>.


<em><b>What is the checksum</b></em> for your list of box IDs?


