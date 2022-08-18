# --- Day 13: Packet Scanners ---

You need to cross a vast <em><b>firewall</b></em>. The firewall consists of several layers, each with a <em><b>security scanner</b></em> that moves back and forth across the layer. To succeed, you must not be detected by a scanner.


By studying the firewall briefly, you are able to record (in your puzzle input) the <em><b>depth</b></em> of each layer and the <em><b>range</b></em> of the scanning area for the scanner within it, written as <code>depth: range</code>. Each layer has a thickness of exactly <code>1</code>. A layer at depth <code>0</code> begins immediately inside the firewall; a layer at depth <code>1</code> would start immediately after that.


For example, suppose you've recorded the following:


<pre><code>0: 3
1: 2
4: 4
6: 4
</code></pre>
This means that there is a layer immediately inside the firewall (with range <code>3</code>), a second layer immediately after that (with range <code>2</code>), a third layer which begins at depth <code>4</code> (with range <code>4</code>), and a fourth layer which begins at depth 6 (also with range <code>4</code>). Visually, it might look like this:


<pre><code> 0   1   2   3   4   5   6
[ ] [ ] ... ... [ ] ... [ ]
[ ] [ ]         [ ]     [ ]
[ ]             [ ]     [ ]
                [ ]     [ ]
</code></pre>
Within each layer, a security scanner moves back and forth within its range. Each security scanner starts at the top and moves down until it reaches the bottom, then moves up until it reaches the top, and repeats. A security scanner takes <em><b>one picosecond</b></em> to move one step.  Drawing scanners as <code>S</code>, the first few picoseconds look like this:


<pre><code>
Picosecond 0:
 0   1   2   3   4   5   6
[S] [S] ... ... [S] ... [S]
[ ] [ ]         [ ]     [ ]
[ ]             [ ]     [ ]
                [ ]     [ ]

Picosecond 1:
 0   1   2   3   4   5   6
[ ] [ ] ... ... [ ] ... [ ]
[S] [S]         [S]     [S]
[ ]             [ ]     [ ]
                [ ]     [ ]

Picosecond 2:
 0   1   2   3   4   5   6
[ ] [S] ... ... [ ] ... [ ]
[ ] [ ]         [ ]     [ ]
[S]             [S]     [S]
                [ ]     [ ]

Picosecond 3:
 0   1   2   3   4   5   6
[ ] [ ] ... ... [ ] ... [ ]
[S] [S]         [ ]     [ ]
[ ]             [ ]     [ ]
                [S]     [S]
</code></pre>
Your plan is to hitch a ride on a packet about to move through the firewall.  The packet will travel along the top of each layer, and it moves at <em><b>one layer per picosecond</b></em>. Each picosecond, the packet moves one layer forward (its first move takes it into layer 0), and then the scanners move one step. If there is a scanner at the top of the layer <em><b>as your packet enters it</b></em>, you are <em><b>caught</b></em>. (If a scanner moves into the top of its layer while you are there, you are <em><b>not</b></em> caught: it doesn't have time to notice you before you leave.) If you were to do this in the configuration above, marking your current position with parentheses, your passage through the firewall would look like this:


<pre><code>Initial state:
 0   1   2   3   4   5   6
[S] [S] ... ... [S] ... [S]
[ ] [ ]         [ ]     [ ]
[ ]             [ ]     [ ]
                [ ]     [ ]

Picosecond 0:
 0   1   2   3   4   5   6
(S) [S] ... ... [S] ... [S]
[ ] [ ]         [ ]     [ ]
[ ]             [ ]     [ ]
                [ ]     [ ]

 0   1   2   3   4   5   6
( ) [ ] ... ... [ ] ... [ ]
[S] [S]         [S]     [S]
[ ]             [ ]     [ ]
                [ ]     [ ]


Picosecond 1:
 0   1   2   3   4   5   6
[ ] ( ) ... ... [ ] ... [ ]
[S] [S]         [S]     [S]
[ ]             [ ]     [ ]
                [ ]     [ ]

 0   1   2   3   4   5   6
[ ] (S) ... ... [ ] ... [ ]
[ ] [ ]         [ ]     [ ]
[S]             [S]     [S]
                [ ]     [ ]


Picosecond 2:
 0   1   2   3   4   5   6
[ ] [S] (.) ... [ ] ... [ ]
[ ] [ ]         [ ]     [ ]
[S]             [S]     [S]
                [ ]     [ ]

 0   1   2   3   4   5   6
[ ] [ ] (.) ... [ ] ... [ ]
[S] [S]         [ ]     [ ]
[ ]             [ ]     [ ]
                [S]     [S]


Picosecond 3:
 0   1   2   3   4   5   6
[ ] [ ] ... (.) [ ] ... [ ]
[S] [S]         [ ]     [ ]
[ ]             [ ]     [ ]
                [S]     [S]

 0   1   2   3   4   5   6
[S] [S] ... (.) [ ] ... [ ]
[ ] [ ]         [ ]     [ ]
[ ]             [S]     [S]
                [ ]     [ ]


Picosecond 4:
 0   1   2   3   4   5   6
[S] [S] ... ... ( ) ... [ ]
[ ] [ ]         [ ]     [ ]
[ ]             [S]     [S]
                [ ]     [ ]

 0   1   2   3   4   5   6
[ ] [ ] ... ... ( ) ... [ ]
[S] [S]         [S]     [S]
[ ]             [ ]     [ ]
                [ ]     [ ]


Picosecond 5:
 0   1   2   3   4   5   6
[ ] [ ] ... ... [ ] (.) [ ]
[S] [S]         [S]     [S]
[ ]             [ ]     [ ]
                [ ]     [ ]

 0   1   2   3   4   5   6
[ ] [S] ... ... [S] (.) [S]
[ ] [ ]         [ ]     [ ]
[S]             [ ]     [ ]
                [ ]     [ ]


Picosecond 6:
 0   1   2   3   4   5   6
[ ] [S] ... ... [S] ... (S)
[ ] [ ]         [ ]     [ ]
[S]             [ ]     [ ]
                [ ]     [ ]

 0   1   2   3   4   5   6
[ ] [ ] ... ... [ ] ... ( )
[S] [S]         [S]     [S]
[ ]             [ ]     [ ]
                [ ]     [ ]
</code></pre>
In this situation, you are <em><b>caught</b></em> in layers <code>0</code> and <code>6</code>, because your packet entered the layer when its scanner was at the top when you entered it. You are <em><b>not</b></em> caught in layer <code>1</code>, since the scanner moved into the top of the layer once you were already there.


The <em><b>severity</b></em> of getting caught on a layer is equal to its <em><b>depth</b></em> multiplied by its <em><b>range</b></em>. (Ignore layers in which you do not get caught.) The severity of the whole trip is the sum of these values.  In the example above, the trip severity is <code>0*3 + 6*4 = <em><b>24</b></em></code>.


Given the details of the firewall you've recorded, if you leave immediately, <em><b>what is the severity of your whole trip</b></em>?


