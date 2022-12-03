# --- Day 2: Rock Paper Scissors ---

The Elves begin to set up camp on the beach. To decide whose tent gets to be closest to the snack storage, a giant [https://en.wikipedia.org/wiki/Rock_paper_scissors](Rock Paper Scissors) tournament is already in progress.


Rock Paper Scissors is a game between two players. Each game contains many rounds; in each round, the players each simultaneously choose one of Rock, Paper, or Scissors using a hand shape. Then, a winner for that round is selected: Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock. If both players choose the same shape, the round instead ends in a draw.


Appreciative of your help yesterday, one Elf gives you an <em><b>encrypted strategy guide</b></em> (your puzzle input) that they say will be sure to help you win. "The first column is what your opponent is going to play: <code>A</code> for Rock, <code>B</code> for Paper, and <code>C</code> for Scissors. The second column--" Suddenly, the Elf is called away to help with someone's tent.


The second column, <span title="Why do you keep guessing?!">you reason</span>, must be what you should play in response: <code>X</code> for Rock, <code>Y</code> for Paper, and <code>Z</code> for Scissors. Winning every time would be suspicious, so the responses must have been carefully chosen.


The winner of the whole tournament is the player with the highest score. Your <em><b>total score</b></em> is the sum of your scores for each round. The score for a single round is the score for the <em><b>shape you selected</b></em> (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the <em><b>outcome of the round</b></em> (0 if you lost, 3 if the round was a draw, and 6 if you won).


Since you can't be sure if the Elf is trying to help you or trick you, you should calculate the score you would get if you were to follow the strategy guide.


For example, suppose you were given the following strategy guide:


<pre><code>A Y
B X
C Z
</code></pre>
This strategy guide predicts and recommends the following:


<ul>
<li>In the first round, your opponent will choose Rock (<code>A</code>), and you should choose Paper (<code>Y</code>). This ends in a win for you with a score of <em><b>8</b></em> (2 because you chose Paper + 6 because you won).</li>
<li>In the second round, your opponent will choose Paper (<code>B</code>), and you should choose Rock (<code>X</code>). This ends in a loss for you with a score of <em><b>1</b></em> (1 + 0).</li>
<li>The third round is a draw with both players choosing Scissors, giving you a score of 3 + 3 = <em><b>6</b></em>.</li>
</ul>
In this example, if you were to follow the strategy guide, you would get a total score of <code><em><b>15</b></em></code> (8 + 1 + 6).


<em><b>What would your total score be if everything goes exactly according to your strategy guide?</b></em>


