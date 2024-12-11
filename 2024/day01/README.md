# --- Day 1: Historian Hysteria ---

The <em><b>Chief Historian</b></em> is always present for the big Christmas sleigh launch, but nobody has seen him in months! Last anyone heard, he was visiting locations that are historically significant to the North Pole; a group of Senior Historians has asked you to accompany them as they check the places they think he was most likely to visit.


As each location is checked, they will mark it on their list with a <em class="star">star</em>. They figure the Chief Historian <em><b>must</b></em> be in one of the first fifty places they'll look, so in order to save Christmas, you need to help them get <em class="star">fifty stars</em> on their list before Santa takes off on December 25th.


Collect stars by solving puzzles.  Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first.  Each puzzle grants <em class="star">one star</em>. Good luck!


You haven't even left yet and the group of Elvish Senior Historians has already hit a problem: their list of locations to check is currently <em><b>empty</b></em>. Eventually, someone decides that the best place to check first would be the Chief Historian's office.


Upon pouring into the office, everyone confirms that the Chief Historian is indeed nowhere to be found. Instead, the Elves discover an assortment of notes and lists of historically significant locations! This seems to be the planning the Chief Historian was doing before he left. Perhaps these notes can be used to determine which locations to search?


Throughout the Chief's office, the historically significant locations are listed not by name but by a unique number called the <em><b>location ID</b></em>. To make sure they don't miss anything, The Historians split into two groups, each searching the office and trying to create their own complete list of location IDs.


There's just one problem: by holding the two lists up <em><b>side by side</b></em> (your puzzle input), it quickly becomes clear that the lists aren't very similar. Maybe you can help The Historians reconcile their lists?


For example:


<pre><code>3   4
4   3
2   5
1   3
3   9
3   3
</code></pre>
Maybe the lists are only off by a small amount! To find out, pair up the numbers and measure how far apart they are. Pair up the <em><b>smallest number in the left list</b></em> with the <em><b>smallest number in the right list</b></em>, then the <em><b>second-smallest left number</b></em> with the <em><b>second-smallest right number</b></em>, and so on.


Within each pair, figure out <em><b>how far apart</b></em> the two numbers are; you'll need to <em><b>add up all of those distances</b></em>. For example, if you pair up a <code>3</code> from the left list with a <code>7</code> from the right list, the distance apart is <code>4</code>; if you pair up a <code>9</code> with a <code>3</code>, the distance apart is <code>6</code>.


In the example list above, the pairs and distances would be as follows:


<ul>
<li>The smallest number in the left list is <code>1</code>, and the smallest number in the right list is <code>3</code>. The distance between them is <code><em><b>2</b></em></code>.</li>
<li>The second-smallest number in the left list is <code>2</code>, and the second-smallest number in the right list is another <code>3</code>. The distance between them is <code><em><b>1</b></em></code>.</li>
<li>The third-smallest number in both lists is <code>3</code>, so the distance between them is <code><em><b>0</b></em></code>.</li>
<li>The next numbers to pair up are <code>3</code> and <code>4</code>, a distance of <code><em><b>1</b></em></code>.</li>
<li>The fifth-smallest numbers in each list are <code>3</code> and <code>5</code>, a distance of <code><em><b>2</b></em></code>.</li>
<li>Finally, the largest number in the left list is <code>4</code>, while the largest number in the right list is <code>9</code>; these are a distance <code><em><b>5</b></em></code> apart.</li>
</ul>
To find the <em><b>total distance</b></em> between the left list and the right list, add up the distances between all of the pairs you found. In the example above, this is <code>2 + 1 + 0 + 1 + 2 + 5</code>, a total distance of <code><em><b>11</b></em></code>!


Your actual left and right lists contain many location IDs. <em><b>What is the total distance between your lists?</b></em>


