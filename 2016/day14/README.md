# --- Day 14: One-Time Pad ---

In order to communicate securely with Santa while you're on this mission, you've been using a [https://en.wikipedia.org/wiki/One-time_pad](one-time pad) that you [https://en.wikipedia.org/wiki/Security_through_obscurity](generate) using a <span title="This also happens to be the plot of World War II.">pre-agreed algorithm</span>. Unfortunately, you've run out of keys in your one-time pad, and so you need to generate some more.


To generate keys, you first get a stream of random data by taking the [https://en.wikipedia.org/wiki/MD5](MD5) of a pre-arranged [https://en.wikipedia.org/wiki/Salt_(cryptography)](salt) (your puzzle input) and an increasing integer index (starting with <code>0</code>, and represented in decimal); the resulting MD5 hash should be represented as a string of <em><b>lowercase</b></em> hexadecimal digits.


However, not all of these MD5 hashes are <em><b>keys</b></em>, and you need <code>64</code> new keys for your one-time pad.  A hash is a key <em><b>only if</b></em>:


<ul>
<li>It contains <em><b>three</b></em> of the same character in a row, like <code>777</code>. Only consider the first such triplet in a hash.</li>
<li>One of the next <code>1000</code> hashes in the stream contains that same character <em><b>five</b></em> times in a row, like <code>77777</code>.</li>
</ul>
Considering future hashes for five-of-a-kind sequences does not cause those hashes to be skipped; instead, regardless of whether the current hash is a key, always resume testing for keys starting with the very next hash.


For example, if the pre-arranged salt is <code>abc</code>:


<ul>
<li>The first index which produces a triple is <code>18</code>, because the MD5 hash of <code>abc18</code> contains <code>...cc38887a5...</code>. However, index <code>18</code> does not count as a key for your one-time pad, because none of the next thousand hashes (index <code>19</code> through index <code>1018</code>) contain <code>88888</code>.</li>
<li>The next index which produces a triple is <code>39</code>; the hash of <code>abc39</code> contains <code>eee</code>. It is also the first key: one of the next thousand hashes (the one at index 816) contains <code>eeeee</code>.</li>
<li>None of the next six triples are keys, but the one after that, at index <code>92</code>, is: it contains <code>999</code> and index <code>200</code> contains <code>99999</code>.</li>
<li>Eventually, index <code>22728</code> meets all of the criteria to generate the <code>64</code>th key.</li>
</ul>
So, using our example salt of <code>abc</code>, index <code>22728</code> produces the <code>64</code>th key.


Given the actual salt in your puzzle input, <em><b>what index</b></em> produces your <code>64</code>th one-time pad key?

