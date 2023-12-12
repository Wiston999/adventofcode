# --- Day 12: Hot Springs ---

You finally reach the hot springs! You can see steam rising from secluded areas attached to the primary, ornate building.


As you turn to enter, the [11](researcher) stops you. "Wait - I thought you were looking for the hot springs, weren't you?" You indicate that this definitely looks like hot springs to you.


"Oh, sorry, common mistake! This is actually the [https://en.wikipedia.org/wiki/Onsen](onsen)! The hot springs are next door."


You look in the direction the researcher is pointing and suddenly notice the <span title="I love this joke. I'm not sorry.">massive metal helixes</span> towering overhead. "This way!"


It only takes you a few more steps to reach the main gate of the massive fenced-off area containing the springs. You go through the gate and into a small administrative building.


"Hello! What brings you to the hot springs today? Sorry they're not very hot right now; we're having a <em><b>lava shortage</b></em> at the moment." You ask about the missing machine parts for Desert Island.


"Oh, all of Gear Island is currently offline! Nothing is being manufactured at the moment, not until we get more lava to heat our forges. And our springs. The springs aren't very springy unless they're hot!"


"Say, could you go up and see why the lava stopped flowing? The springs are too cold for normal operation, but we should be able to find one springy enough to launch <em><b>you</b></em> up there!"


There's just one problem - many of the springs have fallen into disrepair, so they're not actually sure which springs would even be <em><b>safe</b></em> to use! Worse yet, their <em><b>condition records of which springs are damaged</b></em> (your puzzle input) are also damaged! You'll need to help them repair the damaged records.


In the giant field just outside, the springs are arranged into <em><b>rows</b></em>. For each row, the condition records show every spring and whether it is <em><b>operational</b></em> (<code>.</code>) or <em><b>damaged</b></em> (<code>#</code>). This is the part of the condition records that is itself damaged; for some springs, it is simply <em><b>unknown</b></em> (<code>?</code>) whether the spring is operational or damaged.


However, the engineer that produced the condition records also duplicated some of this information in a different format! After the list of springs for a given row, the size of each <em><b>contiguous group of damaged springs</b></em> is listed in the order those groups appear in the row. This list always accounts for every damaged spring, and each number is the entire size of its contiguous group (that is, groups are always separated by at least one operational spring: <code>####</code> would always be <code>4</code>, never <code>2,2</code>).


So, condition records with no unknown spring conditions might look like this:


<pre><code>#.#.### 1,1,3
.#...#....###. 1,1,3
.#.###.#.###### 1,3,1,6
####.#...#... 4,1,1
#....######..#####. 1,6,5
.###.##....# 3,2,1
</code></pre>
However, the condition records are partially damaged; some of the springs' conditions are actually <em><b>unknown</b></em> (<code>?</code>). For example:


<pre><code>???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
</code></pre>
Equipped with this information, it is your job to figure out <em><b>how many different arrangements</b></em> of operational and broken springs fit the given criteria in each row.


In the first line (<code>???.### 1,1,3</code>), there is exactly <em><b>one</b></em> way separate groups of one, one, and three broken springs (in that order) can appear in that row: the first three unknown springs must be broken, then operational, then broken (<code>#.#</code>), making the whole row <code>#.#.###</code>.


The second line is more interesting: <code>.??..??...?##. 1,1,3</code> could be a total of <em><b>four</b></em> different arrangements. The last <code>?</code> must always be broken (to satisfy the final contiguous group of three broken springs), and each <code>??</code> must hide exactly one of the two broken springs. (Neither <code>??</code> could be both broken springs or they would form a single contiguous group of two; if that were true, the numbers afterward would have been <code>2,3</code> instead.) Since each <code>??</code> can either be <code>#.</code> or <code>.#</code>, there are four possible arrangements of springs.


The last line is actually consistent with <em><b>ten</b></em> different arrangements! Because the first number is <code>3</code>, the first and second <code>?</code> must both be <code>.</code> (if either were <code>#</code>, the first number would have to be <code>4</code> or higher). However, the remaining run of unknown spring conditions have many different ways they could hold groups of two and one broken springs:


<pre><code>?###???????? 3,2,1
.###.##.#...
.###.##..#..
.###.##...#.
.###.##....#
.###..##.#..
.###..##..#.
.###..##...#
.###...##.#.
.###...##..#
.###....##.#
</code></pre>
In this example, the number of possible arrangements for each row is:


<ul>
<li><code>???.### 1,1,3</code> - <code><em><b>1</b></em></code> arrangement</li>
<li><code>.??..??...?##. 1,1,3</code> - <code><em><b>4</b></em></code> arrangements</li>
<li><code>?#?#?#?#?#?#?#? 1,3,1,6</code> - <code><em><b>1</b></em></code> arrangement</li>
<li><code>????.#...#... 4,1,1</code> - <code><em><b>1</b></em></code> arrangement</li>
<li><code>????.######..#####. 1,6,5</code> - <code><em><b>4</b></em></code> arrangements</li>
<li><code>?###???????? 3,2,1</code> - <code><em><b>10</b></em></code> arrangements</li>
</ul>
Adding all of the possible arrangement counts together produces a total of <code><em><b>21</b></em></code> arrangements.


For each row, count all of the different arrangements of operational and broken springs that meet the given criteria. <em><b>What is the sum of those counts?</b></em>


