# --- Day 4: The Ideal Stocking Stuffer ---

Santa needs help [https://en.wikipedia.org/wiki/Bitcoin#Mining](mining) some <span title="Hey, mined your own business!">AdventCoins</span> (very similar to [https://en.wikipedia.org/wiki/Bitcoin](bitcoins)) to use as gifts for all the economically forward-thinking little girls and boys.


To do this, he needs to find [https://en.wikipedia.org/wiki/MD5](MD5) hashes which, in [https://en.wikipedia.org/wiki/Hexadecimal](hexadecimal), start with at least <em><b>five zeroes</b></em>.  The input to the MD5 hash is some secret key (your puzzle input, given below) followed by a number in decimal. To mine AdventCoins, you must find Santa the lowest positive number (no leading zeroes: <code>1</code>, <code>2</code>, <code>3</code>, ...) that produces such a hash.


For example:


<ul>
<li>If your secret key is <code>abcdef</code>, the answer is <code>609043</code>, because the MD5 hash of <code>abcdef609043</code> starts with five zeroes (<code>000001dbbfa...</code>), and it is the lowest such number to do so.</li>
<li>If your secret key is <code>pqrstuv</code>, the lowest number it combines with to make an MD5 hash starting with five zeroes is <code>1048970</code>; that is, the MD5 hash of <code>pqrstuv1048970</code> looks like <code>000006136ef...</code>.</li>
</ul>
