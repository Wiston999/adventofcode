# --- Day 1: Sonar Sweep ---

You're minding your own business on a ship at sea when the overboard alarm goes off! You rush to see if you can help. Apparently, one of the Elves tripped and accidentally sent the sleigh keys flying into the ocean!


Before you know it, you're inside a submarine the Elves keep ready for situations like this. It's covered in Christmas lights (because of course it is), and it even has an experimental antenna that should be able to track the keys if you can boost its signal strength high enough; there's a little meter that indicates the antenna's signal strength by displaying 0-50 <em class="star">stars</em>.


Your instincts tell you that in order to save Christmas, you'll need to get all <em class="star">fifty stars</em> by December 25th.


Collect stars by solving puzzles.  Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first.  Each puzzle grants <em class="star">one star</em>. Good luck!


As the submarine drops below the surface of the ocean, it automatically performs a sonar sweep of the nearby sea floor. On a small screen, the sonar sweep report (your puzzle input) appears: each line is a measurement of the sea floor depth as the sweep looks further and further away from the submarine.


For example, suppose you had the following report:


<pre><code>199
200
208
210
200
207
240
269
260
263
</code></pre>
This report indicates that, scanning outward from the submarine, the sonar sweep found depths of <code>199</code>, <code>200</code>, <code>208</code>, <code>210</code>, and so on.


The first order of business is to figure out how quickly the depth increases, just so you know what you're dealing with - you never know if the keys will get <span title="Does this premise seem fishy to you?">carried into deeper water</span> by an ocean current or a fish or something.


To do this, count <em><b>the number of times a depth measurement increases</b></em> from the previous measurement. (There is no measurement before the first measurement.) In the example above, the changes are as follows:


<pre><code>199 (N/A - no previous measurement)
200 (<em><b>increased</b></em>)
208 (<em><b>increased</b></em>)
210 (<em><b>increased</b></em>)
200 (decreased)
207 (<em><b>increased</b></em>)
240 (<em><b>increased</b></em>)
269 (<em><b>increased</b></em>)
260 (decreased)
263 (<em><b>increased</b></em>)
</code></pre>
In this example, there are <em><b><code>7</code></b></em> measurements that are larger than the previous measurement.


<em><b>How many measurements are larger than the previous measurement?</b></em>


