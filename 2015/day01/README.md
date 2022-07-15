# --- Day 1: Not Quite Lisp ---

Santa was hoping for a white Christmas, but his weather machine's "snow" function is powered by stars, and he's fresh out!  To save Christmas, he needs you to collect <em class="star">fifty stars</em> by December 25th.


Collect stars by helping Santa solve puzzles.  Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first.  Each puzzle grants <em class="star">one star</em>. <span title="Also, some puzzles contain Easter eggs like this one. Yes, I know it's not traditional to do Advent calendars for Easter.">Good luck!</span>


Here's an easy puzzle to warm you up.


Santa is trying to deliver presents in a large apartment building, but he can't find the right floor - the directions he got are a little confusing. He starts on the ground floor (floor <code>0</code>) and then follows the instructions one character at a time.


An opening parenthesis, <code>(</code>, means he should go up one floor, and a closing parenthesis, <code>)</code>, means he should go down one floor.


The apartment building is very tall, and the basement is very deep; he will never find the top or bottom floors.


For example:


<ul>
<li><code>(())</code> and <code>()()</code> both result in floor <code>0</code>.</li>
<li><code>(((</code> and <code>(()(()(</code> both result in floor <code>3</code>.</li>
<li><code>))(((((</code> also results in floor <code>3</code>.</li>
<li><code>())</code> and <code>))(</code> both result in floor <code>-1</code> (the first basement level).</li>
<li><code>)))</code> and <code>)())())</code> both result in floor <code>-3</code>.</li>
</ul>
To <em><b>what floor</b></em> do the instructions take Santa?

