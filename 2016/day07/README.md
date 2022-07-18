# --- Day 7: Internet Protocol Version 7 ---

While snooping around the local network of EBHQ, you compile a list of [https://en.wikipedia.org/wiki/IP_address](IP addresses) (they're IPv7, of course; [https://en.wikipedia.org/wiki/IPv6](IPv6) is much too limited). You'd like to figure out which IPs support <em><b>TLS</b></em> (transport-layer snooping).


An IP supports TLS if it has an Autonomous Bridge Bypass Annotation, or <span title="Any similarity to the pattern it describes is purely coincidental."><em><b>ABBA</b></em></span>.  An ABBA is any four-character sequence which consists of a pair of two different characters followed by the reverse of that pair, such as <code>xyyx</code> or <code>abba</code>.  However, the IP also must not have an ABBA within any hypernet sequences, which are contained by <em><b>square brackets</b></em>.


For example:


<ul>
<li><code>abba[mnop]qrst</code> supports TLS (<code>abba</code> outside square brackets).</li>
<li><code>abcd[bddb]xyyx</code> does <em><b>not</b></em> support TLS (<code>bddb</code> is within square brackets, even though <code>xyyx</code> is outside square brackets).</li>
<li><code>aaaa[qwer]tyui</code> does <em><b>not</b></em> support TLS (<code>aaaa</code> is invalid; the interior characters must be different).</li>
<li><code>ioxxoj[asdfgh]zxcvbn</code> supports TLS (<code>oxxo</code> is outside square brackets, even though it's within a larger string).</li>
</ul>
<em><b>How many IPs</b></em> in your puzzle input support TLS?


