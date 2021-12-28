# --- Day 23: Amphipod ---

A group of [https://en.wikipedia.org/wiki/Amphipoda](amphipods) notice your fancy submarine and flag you down. "With such an impressive shell," one amphipod <span title="What? You didn't know amphipods can talk?">says</span>, "surely you can help us with a question that has stumped our best scientists."


They go on to explain that a group of timid, stubborn amphipods live in a nearby burrow. Four types of amphipods live there: <em><b>Amber</b></em> (<code>A</code>), <em><b>Bronze</b></em> (<code>B</code>), <em><b>Copper</b></em> (<code>C</code>), and <em><b>Desert</b></em> (<code>D</code>). They live in a burrow that consists of a <em><b>hallway</b></em> and four <em><b>side rooms</b></em>. The side rooms are initially full of amphipods, and the hallway is initially empty.


They give you a <em><b>diagram of the situation</b></em> (your puzzle input), including locations of each amphipod (<code>A</code>, <code>B</code>, <code>C</code>, or <code>D</code>, each of which is occupying an otherwise open space), walls (<code>#</code>), and open space (<code>.</code>).


For example:


<pre><code>#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########
</code></pre>
The amphipods would like a method to organize every amphipod into side rooms so that each side room contains one type of amphipod and the types are sorted <code>A</code>-<code>D</code> going left to right, like this:


<pre><code>#############
#...........#
###A#B#C#D###
  #A#B#C#D#
  #########
</code></pre>
Amphipods can move up, down, left, or right so long as they are moving into an unoccupied open space. Each type of amphipod requires a different amount of <em><b>energy</b></em> to move one step: Amber amphipods require <code>1</code> energy per step, Bronze amphipods require <code>10</code> energy, Copper amphipods require <code>100</code>, and Desert ones require <code>1000</code>. The amphipods would like you to find a way to organize the amphipods that requires the <em><b>least total energy</b></em>.


However, because they are timid and stubborn, the amphipods have some extra rules:


<ul>
<li>Amphipods will never <em><b>stop on the space immediately outside any room</b></em>. They can move into that space so long as they immediately continue moving. (Specifically, this refers to the four open spaces in the hallway that are directly above an amphipod starting position.)</li>
<li>Amphipods will never <em><b>move from the hallway into a room</b></em> unless that room is their destination room <em><b>and</b></em> that room contains no amphipods which do not also have that room as their own destination. If an amphipod's starting room is not its destination room, it can stay in that room until it leaves the room. (For example, an Amber amphipod will not move from the hallway into the right three rooms, and will only move into the leftmost room if that room is empty or if it only contains other Amber amphipods.)</li>
<li>Once an amphipod stops moving in the hallway, <em><b>it will stay in that spot until it can move into a room</b></em>. (That is, once any amphipod starts moving, any other amphipods currently in the hallway are locked in place and will not move again until they can move fully into a room.)</li>
</ul>
In the above example, the amphipods can be organized using a minimum of <code><em><b>12521</b></em></code> energy. One way to do this is shown below.


Starting configuration:


<pre><code>#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########
</code></pre>
One Bronze amphipod moves into the hallway, taking 4 steps and using <code>40</code> energy:


<pre><code>#############
#...B.......#
###B#C#.#D###
  #A#D#C#A#
  #########
</code></pre>
The only Copper amphipod not in its side room moves there, taking 4 steps and using <code>400</code> energy:


<pre><code>#############
#...B.......#
###B#.#C#D###
  #A#D#C#A#
  #########
</code></pre>
A Desert amphipod moves out of the way, taking 3 steps and using <code>3000</code> energy, and then the Bronze amphipod takes its place, taking 3 steps and using <code>30</code> energy:


<pre><code>#############
#.....D.....#
###B#.#C#D###
  #A#B#C#A#
  #########
</code></pre>
The leftmost Bronze amphipod moves to its room using <code>40</code> energy:


<pre><code>#############
#.....D.....#
###.#B#C#D###
  #A#B#C#A#
  #########
</code></pre>
Both amphipods in the rightmost room move into the hallway, using <code>2003</code> energy in total:


<pre><code>#############
#.....D.D.A.#
###.#B#C#.###
  #A#B#C#.#
  #########
</code></pre>
Both Desert amphipods move into the rightmost room using <code>7000</code> energy:


<pre><code>#############
#.........A.#
###.#B#C#D###
  #A#B#C#D#
  #########
</code></pre>
Finally, the last Amber amphipod moves into its room, using <code>8</code> energy:


<pre><code>#############
#...........#
###A#B#C#D###
  #A#B#C#D#
  #########
</code></pre>
<em><b>What is the least energy required to organize the amphipods?</b></em>


