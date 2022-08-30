# --- Day 1: Chronal Calibration ---

"We've detected some temporal anomalies," one of Santa's Elves at the <span title="It's about as big on the inside as you expected.">Temporal Anomaly Research and Detection Instrument Station</span> tells you. She sounded pretty worried when she called you down here. "At 500-year intervals into the past, someone has been changing Santa's history!"


"The good news is that the changes won't propagate to our time stream for another 25 days, and we have a device" - she attaches something to your wrist - "that will let you fix the changes with no such propagation delay. It's configured to send you 500 years further into the past every few days; that was the best we could do on such short notice."


"The bad news is that we are detecting roughly <em><b>fifty</b></em> anomalies throughout time; the device will indicate fixed anomalies with <em class="star">stars</em>. The other bad news is that we only have one device and you're the best person for the job! Good lu--" She taps a button on the device and you suddenly feel like you're falling. To save Christmas, you need to get all <em class="star">fifty stars</em> by December 25th.


Collect stars by solving puzzles.  Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first.  Each puzzle grants <em class="star">one star</em>. Good luck!


After feeling like you've been falling for a few minutes, you look at the device's tiny screen. "Error: Device must be calibrated before first use. Frequency drift detected. Cannot maintain destination lock." Below the message, the device shows a sequence of changes in frequency (your puzzle input). A value like <code>+6</code> means the current frequency increases by <code>6</code>; a value like <code>-3</code> means the current frequency decreases by <code>3</code>.


For example, if the device displays frequency changes of <code>+1, -2, +3, +1</code>, then starting from a frequency of zero, the following changes would occur:


<ul>
<li>Current frequency <code>&nbsp;0</code>, change of <code>+1</code>; resulting frequency <code>&nbsp;1</code>.</li>
<li>Current frequency <code>&nbsp;1</code>, change of <code>-2</code>; resulting frequency <code>-1</code>.</li>
<li>Current frequency <code>-1</code>, change of <code>+3</code>; resulting frequency <code>&nbsp;2</code>.</li>
<li>Current frequency <code>&nbsp;2</code>, change of <code>+1</code>; resulting frequency <code>&nbsp;3</code>.</li>
</ul>
In this example, the resulting frequency is <code>3</code>.


Here are other example situations:


<ul>
<li><code>+1, +1, +1</code> results in <code>&nbsp;3</code></li>
<li><code>+1, +1, -2</code> results in <code>&nbsp;0</code></li>
<li><code>-1, -2, -3</code> results in <code>-6</code></li>
</ul>
Starting with a frequency of zero, <em><b>what is the resulting frequency</b></em> after all of the changes in frequency have been applied?


