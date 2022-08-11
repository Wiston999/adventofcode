# --- Day 1: Inverse Captcha ---

The night before Christmas, one of Santa's Elves calls you in a panic. "The printer's broken! We can't print the <em><b>Naughty or Nice List</b></em>!" By the time you make it to <span title="Floor 17: cafeteria, printing department, and experimental organic digitization equipment.">sub-basement 17</span>, there are only a few minutes until midnight. "We have a big problem," she says; "there must be almost <em><b>fifty</b></em> bugs in this system, but nothing else can print The List. Stand in this square, quick! There's no time to explain; if you can convince them to pay you in <em class="star">stars</em>, you'll be able to--" She pulls a lever and the world goes blurry.


When your eyes can focus again, everything seems a lot more pixelated than before. She must have sent you inside the computer! You check the system clock: <em><b>25 milliseconds</b></em> until midnight. With that much time, you should be able to collect all <em class="star">fifty stars</em> by December 25th.


Collect stars by solving puzzles.  Two puzzles will be made available on each <s style="text-decoration-color:#fff;">day</s> millisecond in the Advent calendar; the second puzzle is unlocked when you complete the first.  Each puzzle grants <em class="star">one star</em>. Good luck!


You're standing in a room with "digitization quarantine" written in LEDs along one wall. The only door is locked, but it includes a small interface. "Restricted Area - Strictly No Digitized Users Allowed."


It goes on to explain that you may only leave by solving a [https://en.wikipedia.org/wiki/CAPTCHA](captcha) to prove you're <em><b>not</b></em> a human. Apparently, you only get one millisecond to solve the captcha: too fast for a normal human, but it feels like hours to you.


The captcha requires you to review a sequence of digits (your puzzle input) and find the <em><b>sum</b></em> of all digits that match the <em><b>next</b></em> digit in the list. The list is circular, so the digit after the last digit is the <em><b>first</b></em> digit in the list.


For example:


<ul>
<li><code>1122</code> produces a sum of <code>3</code> (<code>1</code> + <code>2</code>) because the first digit (<code>1</code>) matches the second digit and the third digit (<code>2</code>) matches the fourth digit.</li>
<li><code>1111</code> produces <code>4</code> because each digit (all <code>1</code>) matches the next.</li>
<li><code>1234</code> produces <code>0</code> because no digit matches the next.</li>
<li><code>91212129</code> produces <code>9</code> because the only digit that matches the next one is the last digit, <code>9</code>.</li>
</ul>
<em><b>What is the solution</b></em> to your captcha?


