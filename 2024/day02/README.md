# --- Day 2: Red-Nosed Reports ---

Fortunately, the first location The Historians want to search isn't a long walk from the Chief Historian's office.


While the [/2015/day/19](Red-Nosed Reindeer nuclear fusion/fission plant) appears to contain no sign of the Chief Historian, the engineers there run up to you as soon as they see you. Apparently, they <em><b>still</b></em> talk about the time Rudolph was saved through molecular synthesis from a single electron.


They're quick to add that - since you're already here - they'd really appreciate your help analyzing some unusual data from the Red-Nosed reactor. You turn to check if The Historians are waiting for you, but they seem to have already divided into groups that are currently searching every corner of the facility. You offer to help with the unusual data.


The unusual data (your puzzle input) consists of many <em><b>reports</b></em>, one report per line. Each report is a list of numbers called <em><b>levels</b></em> that are separated by spaces. For example:


<pre><code>7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
</code></pre>
This example data contains six reports each containing five levels.


The engineers are trying to figure out which reports are <em><b>safe</b></em>. The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report only counts as safe if both of the following are true:


<ul>
<li>The levels are either <em><b>all increasing</b></em> or <em><b>all decreasing</b></em>.</li>
<li>Any two adjacent levels differ by <em><b>at least one</b></em> and <em><b>at most three</b></em>.</li>
</ul>
In the example above, the reports can be found safe or unsafe by checking those rules:


<ul>
<li><code>7 6 4 2 1</code>: <em><b>Safe</b></em> because the levels are all decreasing by 1 or 2.</li>
<li><code>1 2 7 8 9</code>: <em><b>Unsafe</b></em> because <code>2 7</code> is an increase of 5.</li>
<li><code>9 7 6 2 1</code>: <em><b>Unsafe</b></em> because <code>6 2</code> is a decrease of 4.</li>
<li><code>1 3 2 4 5</code>: <em><b>Unsafe</b></em> because <code>1 3</code> is increasing but <code>3 2</code> is decreasing.</li>
<li><code>8 6 4 4 1</code>: <em><b>Unsafe</b></em> because <code>4 4</code> is neither an increase or a decrease.</li>
<li><code>1 3 6 7 9</code>: <em><b>Safe</b></em> because the levels are all increasing by 1, 2, or 3.</li>
</ul>
So, in this example, <code><em><b>2</b></em></code> reports are <em><b>safe</b></em>.


Analyze the unusual data from the engineers. <em><b>How many reports are safe?</b></em>


