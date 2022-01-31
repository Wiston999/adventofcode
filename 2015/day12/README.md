# --- Day 12: JSAbacusFramework.io ---

Santa's Accounting-Elves need help balancing the books after a recent order.  Unfortunately, their accounting software uses a peculiar storage format.  That's where you come in.


They have a [http://json.org/](JSON) document which contains a variety of things: arrays (<code>[1,2,3]</code>), objects (<code>{"a":1, "b":2}</code>), numbers, and strings.  Your first job is to simply find all of the <em><b>numbers</b></em> throughout the document and add them together.


For example:


<ul>
<li><code>[1,2,3]</code> and <code>{"a":2,"b":4}</code> both have a sum of <code>6</code>.</li>
<li><code>[[[3]]]</code> and <code>{"a":{"b":4},"c":-1}</code> both have a sum of <code>3</code>.</li>
<li><code>{"a":[-1,1]}</code> and <code>[-1,{"a":1}]</code> both have a sum of <code>0</code>.</li>
<li><code>[]</code> and <code>{}</code> both have a sum of <code>0</code>.</li>
</ul>
You will not <span title="Nor are you likely to be eaten by a grue... during *this* puzzle, anyway.">encounter</span> any strings containing numbers.


What is the <em><b>sum of all numbers</b></em> in the document?


