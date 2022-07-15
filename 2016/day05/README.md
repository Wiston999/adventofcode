# --- Day 5: How About a Nice Game of Chess? ---

You are faced with a security door designed by Easter Bunny engineers that seem to have acquired most of their security knowledge by watching [https://en.wikipedia.org/wiki/Hackers_(film)](hacking) [https://en.wikipedia.org/wiki/WarGames](movies).


The <em><b>eight-character password</b></em> for the door is generated one character at a time by finding the [https://en.wikipedia.org/wiki/MD5](MD5) hash of some Door ID (your puzzle input) and an increasing integer index (starting with <code>0</code>).


A hash indicates the <em><b>next character</b></em> in the password if its [https://en.wikipedia.org/wiki/Hexadecimal](hexadecimal) representation starts with <em><b>five zeroes</b></em>. If it does, the sixth character in the hash is the next character of the password.


For example, if the Door ID is <code>abc</code>:


<ul>
<li>The first index which produces a hash that starts with five zeroes is <code>3231929</code>, which we find by hashing <code>abc3231929</code>; the sixth character of the hash, and thus the first character of the password, is <code>1</code>.</li>
<li><code>5017308</code> produces the next interesting hash, which starts with <code>000008f82...</code>, so the second character of the password is <code>8</code>.</li>
<li>The third time a hash starts with five zeroes is for <code>abc5278568</code>, discovering the character <code>f</code>.</li>
</ul>
In this example, after continuing this search a total of eight times, the password is <code>18f47a30</code>.


Given the actual Door ID, <em><b>what is the password</b></em>?


