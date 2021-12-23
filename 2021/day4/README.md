# --- Day 4: Giant Squid ---

You're already almost 1.5km (almost a mile) below the surface of the ocean, already so deep that you can't see any sunlight. What you <em><b>can</b></em> see, however, is a giant squid that has attached itself to the outside of your submarine.


Maybe it wants to play [https://en.wikipedia.org/wiki/Bingo_(American_version)](bingo)?


Bingo is played on a set of boards each consisting of a 5x5 grid of numbers. Numbers are chosen at random, and the chosen number is <em><b>marked</b></em> on all boards on which it appears. (Numbers may not appear on all boards.) If all numbers in any row or any column of a board are marked, that board <em><b>wins</b></em>. (Diagonals don't count.)


The submarine has a <em><b>bingo subsystem</b></em> to help passengers (currently, you and the giant squid) pass the time. It automatically generates a random order in which to draw numbers and a random set of boards (your puzzle input). For example:


<pre><code>7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
</code></pre>
After the first five numbers are drawn (<code>7</code>, <code>4</code>, <code>9</code>, <code>5</code>, and <code>11</code>), there are no winners, but the boards are marked as follows (shown here adjacent to each other to save space):


<pre><code>22 13 17 <em><b>11</b></em>  0         3 15  0  2 22        14 21 17 24  <em><b>4</b></em>
 8  2 23  <em><b>4</b></em> 24         <em><b>9</b></em> 18 13 17  <em><b>5</b></em>        10 16 15  <em><b>9</b></em> 19
21  <em><b>9</b></em> 14 16  <em><b>7</b></em>        19  8  <em><b>7</b></em> 25 23        18  8 23 26 20
 6 10  3 18  <em><b>5</b></em>        20 <em><b>11</b></em> 10 24  <em><b>4</b></em>        22 <em><b>11</b></em> 13  6  <em><b>5</b></em>
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  <em><b>7</b></em>
</code></pre>
After the next six numbers are drawn (<code>17</code>, <code>23</code>, <code>2</code>, <code>0</code>, <code>14</code>, and <code>21</code>), there are still no winners:


<pre><code>22 13 <em><b>17</b></em> <em><b>11</b></em>  <em><b>0</b></em>         3 15  <em><b>0</b></em>  <em><b>2</b></em> 22        <em><b>14</b></em> <em><b>21</b></em> <em><b>17</b></em> 24  <em><b>4</b></em>
 8  <em><b>2</b></em> <em><b>23</b></em>  <em><b>4</b></em> 24         <em><b>9</b></em> 18 13 <em><b>17</b></em>  <em><b>5</b></em>        10 16 15  <em><b>9</b></em> 19
<em><b>21</b></em>  <em><b>9</b></em> <em><b>14</b></em> 16  <em><b>7</b></em>        19  8  <em><b>7</b></em> 25 <em><b>23</b></em>        18  8 <em><b>23</b></em> 26 20
 6 10  3 18  <em><b>5</b></em>        20 <em><b>11</b></em> 10 24  <em><b>4</b></em>        22 <em><b>11</b></em> 13  6  <em><b>5</b></em>
 1 12 20 15 19        <em><b>14</b></em> <em><b>21</b></em> 16 12  6         <em><b>2</b></em>  <em><b>0</b></em> 12  3  <em><b>7</b></em>
</code></pre>
Finally, <code>24</code> is drawn:


<pre><code>22 13 <em><b>17</b></em> <em><b>11</b></em>  <em><b>0</b></em>         3 15  <em><b>0</b></em>  <em><b>2</b></em> 22        <em><b>14</b></em> <em><b>21</b></em> <em><b>17</b></em> <em><b>24</b></em>  <em><b>4</b></em>
 8  <em><b>2</b></em> <em><b>23</b></em>  <em><b>4</b></em> <em><b>24</b></em>         <em><b>9</b></em> 18 13 <em><b>17</b></em>  <em><b>5</b></em>        10 16 15  <em><b>9</b></em> 19
<em><b>21</b></em>  <em><b>9</b></em> <em><b>14</b></em> 16  <em><b>7</b></em>        19  8  <em><b>7</b></em> 25 <em><b>23</b></em>        18  8 <em><b>23</b></em> 26 20
 6 10  3 18  <em><b>5</b></em>        20 <em><b>11</b></em> 10 <em><b>24</b></em>  <em><b>4</b></em>        22 <em><b>11</b></em> 13  6  <em><b>5</b></em>
 1 12 20 15 19        <em><b>14</b></em> <em><b>21</b></em> 16 12  6         <em><b>2</b></em>  <em><b>0</b></em> 12  3  <em><b>7</b></em>
</code></pre>
At this point, the third board <em><b>wins</b></em> because it has at least one complete row or column of marked numbers (in this case, the entire top row is marked: <code><em><b>14 21 17 24  4</b></em></code>).


The <em><b>score</b></em> of the winning board can now be calculated. Start by finding the <em><b>sum of all unmarked numbers</b></em> on that board; in this case, the sum is <code>188</code>. Then, multiply that sum by <em><b>the number that was just called</b></em> when the board won, <code>24</code>, to get the final score, <code>188 * 24 = <em><b>4512</b></em></code>.


To guarantee victory against the giant squid, figure out which board will win first. <em><b>What will your final score be if you choose that board?</b></em>


