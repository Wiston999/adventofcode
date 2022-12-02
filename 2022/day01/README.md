# --- Day 1: Calorie Counting ---

Santa's reindeer typically eat regular reindeer food, but they need a lot of [/2018/day/25](magical energy) to deliver presents on Christmas. For that, their favorite snack is a special type of <em class="star">star</em> fruit that only grows deep in the jungle. The Elves have brought you on their annual expedition to the grove where the fruit grows.


To supply enough magical energy, the expedition needs to retrieve a minimum of <em class="star">fifty stars</em> by December 25th. Although the Elves assure you that the grove has plenty of fruit, you decide to grab any fruit you see along the way, just in case.


Collect stars by solving puzzles.  Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first.  Each puzzle grants <em class="star">one star</em>. Good luck!


The jungle must be too overgrown and difficult to navigate in vehicles or access from the air; the Elves' expedition traditionally goes on foot. As your boats approach land, the Elves begin taking inventory of their supplies. One important consideration is food - in particular, the number of <em><b>Calories</b></em> each Elf is carrying (your puzzle input).


The Elves take turns writing down the number of Calories contained by the various meals, snacks, rations, <span title="By &quot;etc&quot;, you're pretty sure they just mean &quot;more snacks&quot;.">etc.</span> that they've brought with them, one item per line. Each Elf separates their own inventory from the previous Elf's inventory (if any) by a blank line.


For example, suppose the Elves finish writing their items' Calories and end up with the following list:


<pre><code>1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
</code></pre>
This list represents the Calories of the food carried by five Elves:


<ul>
<li>The first Elf is carrying food with <code>1000</code>, <code>2000</code>, and <code>3000</code> Calories, a total of <code><em><b>6000</b></em></code> Calories.</li>
<li>The second Elf is carrying one food item with <code><em><b>4000</b></em></code> Calories.</li>
<li>The third Elf is carrying food with <code>5000</code> and <code>6000</code> Calories, a total of <code><em><b>11000</b></em></code> Calories.</li>
<li>The fourth Elf is carrying food with <code>7000</code>, <code>8000</code>, and <code>9000</code> Calories, a total of <code><em><b>24000</b></em></code> Calories.</li>
<li>The fifth Elf is carrying one food item with <code><em><b>10000</b></em></code> Calories.</li>
</ul>
In case the Elves get hungry and need extra snacks, they need to know which Elf to ask: they'd like to know how many Calories are being carried by the Elf carrying the <em><b>most</b></em> Calories. In the example above, this is <em><b><code>24000</code></b></em> (carried by the fourth Elf).


Find the Elf carrying the most Calories. <em><b>How many total Calories is that Elf carrying?</b></em>


