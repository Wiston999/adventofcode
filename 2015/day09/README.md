# --- Day 9: All in a Single Night ---

Every year, Santa manages to deliver all of his presents in a single night.


This year, however, he has some <span title="Bonus points if you recognize all of the locations.">new locations</span> to visit; his elves have provided him the distances between every pair of locations.  He can start and end at any two (different) locations he wants, but he must visit each location exactly once.  What is the <em><b>shortest distance</b></em> he can travel to achieve this?


For example, given the following distances:


<pre><code>London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141
</code></pre>
The possible routes are therefore:


<pre><code>Dublin -> London -> Belfast = 982
London -> Dublin -> Belfast = 605
London -> Belfast -> Dublin = 659
Dublin -> Belfast -> London = 659
Belfast -> Dublin -> London = 605
Belfast -> London -> Dublin = 982
</code></pre>
The shortest of these is <code>London -> Dublin -> Belfast = 605</code>, and so the answer is <code>605</code> in this example.


What is the distance of the shortest route?


