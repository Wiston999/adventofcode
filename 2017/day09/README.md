# --- Day 9: Stream Processing ---

A large stream blocks your path. According to the locals, it's not safe to <span title="&quot;Don't cross the streams!&quot;, they yell, even though there's only one. They seem to think they're hilarious.">cross the stream</span> at the moment because it's full of <em><b>garbage</b></em>. You look down at the stream; rather than water, you discover that it's a <em><b>stream of characters</b></em>.


You sit for a while and record part of the stream (your puzzle input). The characters represent <em><b>groups</b></em> - sequences that begin with <code>{</code> and end with <code>}</code>. Within a group, there are zero or more other things, separated by commas: either another <em><b>group</b></em> or <em><b>garbage</b></em>. Since groups can contain other groups, a <code>}</code> only closes the <em><b>most-recently-opened unclosed group</b></em> - that is, they are nestable. Your puzzle input represents a single, large group which itself contains many smaller ones.


Sometimes, instead of a group, you will find <em><b>garbage</b></em>. Garbage begins with <code>&lt;</code> and ends with <code>&gt;</code>. Between those angle brackets, almost any character can appear, including <code>{</code> and <code>}</code>. <em><b>Within</b></em> garbage, <code>&lt;</code> has no special meaning.


In a futile attempt to clean up the garbage, some program has <em><b>canceled</b></em> some of the characters within it using <code>!</code>: inside garbage, <em><b>any</b></em> character that comes after <code>!</code> should be <em><b>ignored</b></em>, including <code>&lt;</code>, <code>&gt;</code>, and even another <code>!</code>.


You don't see any characters that deviate from these rules.  Outside garbage, you only find well-formed groups, and garbage always terminates according to the rules above.


Here are some self-contained pieces of garbage:


<ul>
<li><code>&lt;&gt;</code>, empty garbage.</li>
<li><code>&lt;random characters&gt;</code>, garbage containing random characters.</li>
<li><code>&lt;&lt;&lt;&lt;&gt;</code>, because the extra <code>&lt;</code> are ignored.</li>
<li><code>&lt;{!&gt;}&gt;</code>, because the first <code>&gt;</code> is canceled.</li>
<li><code>&lt;!!&gt;</code>, because the second <code>!</code> is canceled, allowing the <code>&gt;</code> to terminate the garbage.</li>
<li><code>&lt;!!!&gt;&gt;</code>, because the second <code>!</code> and the first <code>&gt;</code> are canceled.</li>
<li><code>&lt;{o"i!a,&lt;{i&lt;a&gt;</code>, which ends at the first <code>&gt;</code>.</li>
</ul>
Here are some examples of whole streams and the number of groups they contain:


<ul>
<li><code>{}</code>, <code>1</code> group.</li>
<li><code>{{{}}}</code>, <code>3</code> groups.</li>
<li><code>{{},{}}</code>, also <code>3</code> groups.</li>
<li><code>{{{},{},{{}}}}</code>, <code>6</code> groups.</li>
<li><code>{&lt;{},{},{{}}&gt;}</code>, <code>1</code> group (which itself contains garbage).</li>
<li><code>{&lt;a&gt;,&lt;a&gt;,&lt;a&gt;,&lt;a&gt;}</code>, <code>1</code> group.</li>
<li><code>{{&lt;a&gt;},{&lt;a&gt;},{&lt;a&gt;},{&lt;a&gt;}}</code>, <code>5</code> groups.</li>
<li><code>{{&lt;!&gt;},{&lt;!&gt;},{&lt;!&gt;},{&lt;a&gt;}}</code>, <code>2</code> groups (since all but the last <code>&gt;</code> are canceled).</li>
</ul>
Your goal is to find the total score for all groups in your input. Each group is assigned a <em><b>score</b></em> which is one more than the score of the group that immediately contains it. (The outermost group gets a score of <code>1</code>.)


<ul>
<li><code>{}</code>, score of <code>1</code>.</li>
<li><code>{{{}}}</code>, score of <code>1 + 2 + 3 = 6</code>.</li>
<li><code>{{},{}}</code>, score of <code>1 + 2 + 2 = 5</code>.</li>
<li><code>{{{},{},{{}}}}</code>, score of <code>1 + 2 + 3 + 3 + 3 + 4 = 16</code>.</li>
<li><code>{&lt;a&gt;,&lt;a&gt;,&lt;a&gt;,&lt;a&gt;}</code>, score of <code>1</code>.</li>
<li><code>{{&lt;ab&gt;},{&lt;ab&gt;},{&lt;ab&gt;},{&lt;ab&gt;}}</code>, score of <code>1 + 2 + 2 + 2 + 2 = 9</code>.</li>
<li><code>{{&lt;!!&gt;},{&lt;!!&gt;},{&lt;!!&gt;},{&lt;!!&gt;}}</code>, score of <code>1 + 2 + 2 + 2 + 2 = 9</code>.</li>
<li><code>{{&lt;a!&gt;},{&lt;a!&gt;},{&lt;a!&gt;},{&lt;ab&gt;}}</code>, score of <code>1 + 2 = 3</code>.</li>
</ul>
<em><b>What is the total score</b></em> for all groups in your input?


