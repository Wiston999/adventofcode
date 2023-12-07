# --- Day 6: Wait For It ---

The ferry quickly brings you across Island Island. After asking around, you discover that there is indeed normally a large pile of sand somewhere near here, but you don't see anything besides lots of water and the small island where the ferry has docked.


As you try to figure out what to do next, you notice a poster on a wall near the ferry dock. "Boat races! Open to the public! Grand prize is an all-expenses-paid trip to <em><b>Desert Island</b></em>!" That must be where the sand comes from! Best of all, the boat races are starting in just a few minutes.


You manage to sign up as a competitor in the boat races just in time. The organizer explains that it's not really a traditional race - instead, you will get a fixed amount of time during which your boat has to travel as far as it can, and you win if your boat goes the farthest.


As part of signing up, you get a sheet of paper (your puzzle input) that lists the <em><b>time</b></em> allowed for each race and also the best <em><b>distance</b></em> ever recorded in that race. To guarantee you win the grand prize, you need to make sure you <em><b>go farther in each race</b></em> than the current record holder.


The organizer brings you over to the area where the boat races are held. The boats are much smaller than you expected - they're actually <em><b>toy boats</b></em>, each with a big button on top. Holding down the button <em><b>charges the boat</b></em>, and releasing the button <em><b>allows the boat to move</b></em>. Boats move faster if their button was held longer, but time spent holding the button counts against the total race time. You can only hold the button at the start of the race, and boats don't move until the button is released.


For example:


<pre><code>Time:      7  15   30
Distance:  9  40  200
</code></pre>
This document describes three races:


<ul>
<li>The first race lasts 7 milliseconds. The record distance in this race is 9 millimeters.</li>
<li>The second race lasts 15 milliseconds. The record distance in this race is 40 millimeters.</li>
<li>The third race lasts 30 milliseconds. The record distance in this race is 200 millimeters.</li>
</ul>
Your toy boat has a starting speed of <em><b>zero millimeters per millisecond</b></em>. For each whole millisecond you spend at the beginning of the race holding down the button, the boat's speed increases by <em><b>one millimeter per millisecond</b></em>.


So, because the first race lasts 7 milliseconds, you only have a few options:


<ul>
<li>Don't hold the button at all (that is, hold it for <em><b><code>0</code> milliseconds</b></em>) at the start of the race. The boat won't move; it will have traveled <em><b><code>0</code> millimeters</b></em> by the end of the race.</li>
<li>Hold the button for <em><b><code>1</code> millisecond</b></em> at the start of the race. Then, the boat will travel at a speed of <code>1</code> millimeter per millisecond for 6 milliseconds, reaching a total distance traveled of <em><b><code>6</code> millimeters</b></em>.</li>
<li>Hold the button for <em><b><code>2</code> milliseconds</b></em>, giving the boat a speed of <code>2</code> millimeters per millisecond. It will then get 5 milliseconds to move, reaching a total distance of <em><b><code>10</code> millimeters</b></em>.</li>
<li>Hold the button for <em><b><code>3</code> milliseconds</b></em>. After its remaining 4 milliseconds of travel time, the boat will have gone <em><b><code>12</code> millimeters</b></em>.</li>
<li>Hold the button for <em><b><code>4</code> milliseconds</b></em>. After its remaining 3 milliseconds of travel time, the boat will have gone <em><b><code>12</code> millimeters</b></em>.</li>
<li>Hold the button for <em><b><code>5</code> milliseconds</b></em>, causing the boat to travel a total of <em><b><code>10</code> millimeters</b></em>.</li>
<li>Hold the button for <em><b><code>6</code> milliseconds</b></em>, causing the boat to travel a total of <em><b><code>6</code> millimeters</b></em>.</li>
<li>Hold the button for <em><b><code>7</code> milliseconds</b></em>. That's the entire duration of the race. You never let go of the button. The boat can't move until you let go of the button. Please make sure you let go of the button so the boat gets to move. <em><b><code>0</code> millimeters</b></em>.</li>
</ul>
Since the current record for this race is <code>9</code> millimeters, there are actually <code><em><b>4</b></em></code> different ways you could win: you could hold the button for <code>2</code>, <code>3</code>, <code>4</code>, or <code>5</code> milliseconds at the start of the race.


In the second race, you could hold the button for at least <code>4</code> milliseconds and at most <code>11</code> milliseconds and beat the record, a total of <code><em><b>8</b></em></code> different ways to win.


In the third race, you could hold the button for at least <code>11</code> milliseconds and no more than <code>19</code> milliseconds and still beat the record, a total of <code><em><b>9</b></em></code> ways you could win.


To see how much margin of error you have, determine the <em><b>number of ways you can beat the record</b></em> in each race; in this example, if you multiply these values together, you get <code><em><b>288</b></em></code> (<code>4</code> * <code>8</code> * <code>9</code>).


Determine the number of ways you could beat the record in each race. <em><b>What do you get if you multiply these numbers together?</b></em>


