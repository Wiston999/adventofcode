# --- Day 17: Two Steps Forward ---

You're trying to access a secure vault protected by a <code>4x4</code> grid of small rooms connected by doors. You start in the top-left room (marked <code>S</code>), and you can access the vault (marked <code>V</code>) once you reach the bottom-right room:


<pre><code>#########
#S| | | #
#-#-#-#-#
# | | | #
#-#-#-#-#
# | | | #
#-#-#-#-#
# | | |  
####### V
</code></pre>
Fixed walls are marked with <code>#</code>, and doors are marked with <code>-</code> or <code>|</code>.


The doors in your <em><b>current room</b></em> are either open or closed (and locked) based on the hexadecimal [https://en.wikipedia.org/wiki/MD5](MD5) hash of a passcode (your puzzle input) followed by a sequence of uppercase characters representing the <em><b>path you have taken so far</b></em> (<code>U</code> for up, <code>D</code> for down, <code>L</code> for left, and <code>R</code> for right).


Only the first four characters of the hash are used; they represent, respectively, the doors <em><b>up, down, left, and right</b></em> from your current position. Any <code>b</code>, <code>c</code>, <code>d</code>, <code>e</code>, or <code>f</code> means that the corresponding door is <em><b>open</b></em>; any other character (any number or <code>a</code>) means that the corresponding door is <em><b>closed and locked</b></em>.


To access the vault, all you need to do is reach the bottom-right room; reaching this room opens the vault and all doors in the maze.


For example, suppose the passcode is <code>hijkl</code>. Initially, you have taken no steps, and so your path is empty: you simply find the MD5 hash of <code>hijkl</code> alone. The first four characters of this hash are <code>ced9</code>, which indicate that up is open (<code>c</code>), down is open (<code>e</code>), left is open (<code>d</code>), and right is closed and locked (<code>9</code>). Because you start in the top-left corner, there are no "up" or "left" doors to be open, so your only choice is <em><b>down</b></em>.


Next, having gone only one step (down, or <code>D</code>), you find the hash of <code>hijkl<em><b>D</b></em></code>. This produces <code>f2bc</code>, which indicates that you can go back up, left (but that's a wall), or right. Going right means hashing <code>hijkl<em><b>DR</b></em></code> to get <code>5745</code> - all doors closed and locked. However, going <em><b>up</b></em> instead is worthwhile: even though it returns you to the room you started in, your path would then be <code>DU</code>, opening a <em><b>different set of doors</b></em>.


After going <code>DU</code> (and then hashing <code>hijkl<em><b>DU</b></em></code> to get <code>528e</code>), only the right door is open; after going <code>DUR</code>, all doors lock. (Fortunately, your actual passcode is <span title="It took four days to rescue the engineer that tried this.">not <code>hijkl</code></span>).


Passcodes actually used by Easter Bunny Vault Security do allow access to the vault if you know the right path.  For example:


<ul>
<li>If your passcode were <code>ihgpwlah</code>, the shortest path would be <code>DDRRRD</code>.</li>
<li>With <code>kglvqrro</code>, the shortest path would be <code>DDUDRLRRUDRD</code>.</li>
<li>With <code>ulqzkmiv</code>, the shortest would be <code>DRURDRUDDLLDLUURRDULRLDUUDDDRR</code>.</li>
</ul>
Given your vault's passcode, <em><b>what is the shortest path</b></em> (the actual path, not just the length) to reach the vault?


