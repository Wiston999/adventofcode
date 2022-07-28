# --- Day 20: Firewall Rules ---

You'd like to set up a small hidden computer here so you can use it to <span title="I'll create a GUI interface using Visual Basic... see if I can track an IP address.">get back into the network</span> later. However, the corporate firewall only allows communication with certain external [https://en.wikipedia.org/wiki/IPv4#Addressing](IP addresses).


You've retrieved the list of blocked IPs from the firewall, but the list seems to be messy and poorly maintained, and it's not clear which IPs are allowed. Also, rather than being written in [https://en.wikipedia.org/wiki/Dot-decimal_notation](dot-decimal) notation, they are written as plain [https://en.wikipedia.org/wiki/32-bit](32-bit integers), which can have any value from <code>0</code> through <code>4294967295</code>, inclusive.


For example, suppose only the values <code>0</code> through <code>9</code> were valid, and that you retrieved the following blacklist:


<pre><code>5-8
0-2
4-7
</code></pre>
The blacklist specifies ranges of IPs (inclusive of both the start and end value) that are <em><b>not</b></em> allowed. Then, the only IPs that this firewall allows are <code>3</code> and <code>9</code>, since those are the only numbers not in any range.


Given the list of blocked IPs you retrieved from the firewall (your puzzle input), <em><b>what is the lowest-valued IP</b></em> that is not blocked?


