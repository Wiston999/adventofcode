# --- Day 6: Tuning Trouble ---

The preparations are finally complete; you and the Elves leave camp on foot and begin to make your way toward the <em class="star">star</em> fruit grove.


As you move through the dense undergrowth, one of the Elves gives you a handheld <em><b>device</b></em>. He says that it has many fancy features, but the most important one to set up right now is the <em><b>communication system</b></em>.


However, because he's heard you have [/2016/day/6](significant) [/2016/day/25](experience) [/2019/day/7](dealing) [/2019/day/9](with) [/2019/day/16](signal-based) [/2021/day/25](systems), he convinced the other Elves that it would be okay to give you their one malfunctioning device - surely you'll have no problem fixing it.


As if inspired by comedic timing, the device emits a few <span title="The magic smoke, on the other hand, seems to be contained... FOR NOW!">colorful sparks</span>.


To be able to communicate with the Elves, the device needs to <em><b>lock on to their signal</b></em>. The signal is a series of seemingly-random characters that the device receives one at a time.


To fix the communication system, you need to add a subroutine to the device that detects a <em><b>start-of-packet marker</b></em> in the datastream. In the protocol being used by the Elves, the start of a packet is indicated by a sequence of <em><b>four characters that are all different</b></em>.


The device will send your subroutine a datastream buffer (your puzzle input); your subroutine needs to identify the first position where the four most recently received characters were all different. Specifically, it needs to report the number of characters from the beginning of the buffer to the end of the first such four-character marker.


For example, suppose you receive the following datastream buffer:


<pre><code>mjqjpqmgbljsphdztnvjfqwrcgsmlb</code></pre>
After the first three characters (<code>mjq</code>) have been received, there haven't been enough characters received yet to find the marker. The first time a marker could occur is after the fourth character is received, making the most recent four characters <code>mjqj</code>. Because <code>j</code> is repeated, this isn't a marker.


The first time a marker appears is after the <em><b>seventh</b></em> character arrives. Once it does, the last four characters received are <code>jpqm</code>, which are all different. In this case, your subroutine should report the value <code><em><b>7</b></em></code>, because the first start-of-packet marker is complete after 7 characters have been processed.


Here are a few more examples:


<ul>
<li><code>bvwbjplbgvbhsrlpgdmjqwftvncz</code>: first marker after character <code><em><b>5</b></em></code></li>
<li><code>nppdvjthqldpwncqszvftbrmjlhg</code>: first marker after character <code><em><b>6</b></em></code></li>
<li><code>nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg</code>: first marker after character <code><em><b>10</b></em></code></li>
<li><code>zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw</code>: first marker after character <code><em><b>11</b></em></code></li>
</ul>
<em><b>How many characters need to be processed before the first start-of-packet marker is detected?</b></em>


