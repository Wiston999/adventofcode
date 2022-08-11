# --- Day 2: Corruption Checksum ---

As you walk through the door, a glowing humanoid shape yells in your direction. "You there! Your state appears to be idle. Come help us repair the corruption in this spreadsheet - if we take another millisecond, we'll have to display an hourglass cursor!"


The spreadsheet consists of rows of apparently-random numbers. To make sure the recovery process is on the right track, they need you to calculate the spreadsheet's <em><b>checksum</b></em>. For each row, determine the difference between the largest value and the smallest value; the checksum is the sum of all of these differences.


For example, given the following spreadsheet:


<pre><code>5 1 9 5
7 5 3
2 4 6 8</code></pre>
<ul>
<li>The first row's largest and smallest values are <code>9</code> and <code>1</code>, and their difference is <code>8</code>.</li>
<li>The second row's largest and smallest values are <code>7</code> and <code>3</code>, and their difference is <code>4</code>.</li>
<li>The third row's difference is <code>6</code>.</li>
</ul>
In this example, the spreadsheet's checksum would be <code>8 + 4 + 6 = 18</code>.


<em><b>What is the checksum</b></em> for the spreadsheet in your puzzle input?


