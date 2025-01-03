# --- Day 16: Proboscidea Volcanium ---

The sensors have led you to the origin of the distress signal: yet another handheld device, just like the one the Elves gave you. However, you don't see any Elves around; instead, the device is surrounded by elephants! They must have gotten lost in these tunnels, and one of the elephants apparently figured out how to turn on the distress signal.


The ground rumbles again, much stronger this time. What kind of cave is this, exactly? You scan the cave with your handheld device; it reports mostly igneous rock, some ash, pockets of pressurized gas, magma... this isn't just a cave, it's a volcano!


You need to get the elephants out of here, quickly. Your device estimates that you have <em><b>30 minutes</b></em> before the volcano erupts, so you don't have time to go back out the way you came in.


You scan the cave for other options and discover a network of pipes and pressure-release <em><b>valves</b></em>. You aren't sure how such a system got into a volcano, but you don't have time to complain; your device produces a report (your puzzle input) of each valve's <em><b>flow rate</b></em> if it were opened (in pressure per minute) and the tunnels you could use to move between the valves.


There's even a valve in the room you and the elephants are currently standing in labeled <code>AA</code>. You estimate it will take you one minute to open a single valve and one minute to follow any tunnel from one valve to another. What is the most pressure you could release?


For example, suppose you had the following scan output:


<pre><code>Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II
</code></pre>
All of the valves begin <em><b>closed</b></em>. You start at valve <code>AA</code>, but it must be damaged or <span title="Wait, sir! The valve, sir! it appears to be... jammed!">jammed</span> or something: its flow rate is <code>0</code>, so there's no point in opening it. However, you could spend one minute moving to valve <code>BB</code> and another minute opening it; doing so would release pressure during the remaining <em><b>28 minutes</b></em> at a flow rate of <code>13</code>, a total eventual pressure release of <code>28 * 13 = <em><b>364</b></em></code>. Then, you could spend your third minute moving to valve <code>CC</code> and your fourth minute opening it, providing an additional <em><b>26 minutes</b></em> of eventual pressure release at a flow rate of <code>2</code>, or <code><em><b>52</b></em></code> total pressure released by valve <code>CC</code>.


Making your way through the tunnels like this, you could probably open many or all of the valves by the time 30 minutes have elapsed. However, you need to release as much pressure as possible, so you'll need to be methodical. Instead, consider this approach:


<pre><code>== Minute 1 ==
No valves are open.
You move to valve DD.

== Minute 2 ==
No valves are open.
You open valve DD.

== Minute 3 ==
Valve DD is open, releasing <em><b>20</b></em> pressure.
You move to valve CC.

== Minute 4 ==
Valve DD is open, releasing <em><b>20</b></em> pressure.
You move to valve BB.

== Minute 5 ==
Valve DD is open, releasing <em><b>20</b></em> pressure.
You open valve BB.

== Minute 6 ==
Valves BB and DD are open, releasing <em><b>33</b></em> pressure.
You move to valve AA.

== Minute 7 ==
Valves BB and DD are open, releasing <em><b>33</b></em> pressure.
You move to valve II.

== Minute 8 ==
Valves BB and DD are open, releasing <em><b>33</b></em> pressure.
You move to valve JJ.

== Minute 9 ==
Valves BB and DD are open, releasing <em><b>33</b></em> pressure.
You open valve JJ.

== Minute 10 ==
Valves BB, DD, and JJ are open, releasing <em><b>54</b></em> pressure.
You move to valve II.

== Minute 11 ==
Valves BB, DD, and JJ are open, releasing <em><b>54</b></em> pressure.
You move to valve AA.

== Minute 12 ==
Valves BB, DD, and JJ are open, releasing <em><b>54</b></em> pressure.
You move to valve DD.

== Minute 13 ==
Valves BB, DD, and JJ are open, releasing <em><b>54</b></em> pressure.
You move to valve EE.

== Minute 14 ==
Valves BB, DD, and JJ are open, releasing <em><b>54</b></em> pressure.
You move to valve FF.

== Minute 15 ==
Valves BB, DD, and JJ are open, releasing <em><b>54</b></em> pressure.
You move to valve GG.

== Minute 16 ==
Valves BB, DD, and JJ are open, releasing <em><b>54</b></em> pressure.
You move to valve HH.

== Minute 17 ==
Valves BB, DD, and JJ are open, releasing <em><b>54</b></em> pressure.
You open valve HH.

== Minute 18 ==
Valves BB, DD, HH, and JJ are open, releasing <em><b>76</b></em> pressure.
You move to valve GG.

== Minute 19 ==
Valves BB, DD, HH, and JJ are open, releasing <em><b>76</b></em> pressure.
You move to valve FF.

== Minute 20 ==
Valves BB, DD, HH, and JJ are open, releasing <em><b>76</b></em> pressure.
You move to valve EE.

== Minute 21 ==
Valves BB, DD, HH, and JJ are open, releasing <em><b>76</b></em> pressure.
You open valve EE.

== Minute 22 ==
Valves BB, DD, EE, HH, and JJ are open, releasing <em><b>79</b></em> pressure.
You move to valve DD.

== Minute 23 ==
Valves BB, DD, EE, HH, and JJ are open, releasing <em><b>79</b></em> pressure.
You move to valve CC.

== Minute 24 ==
Valves BB, DD, EE, HH, and JJ are open, releasing <em><b>79</b></em> pressure.
You open valve CC.

== Minute 25 ==
Valves BB, CC, DD, EE, HH, and JJ are open, releasing <em><b>81</b></em> pressure.

== Minute 26 ==
Valves BB, CC, DD, EE, HH, and JJ are open, releasing <em><b>81</b></em> pressure.

== Minute 27 ==
Valves BB, CC, DD, EE, HH, and JJ are open, releasing <em><b>81</b></em> pressure.

== Minute 28 ==
Valves BB, CC, DD, EE, HH, and JJ are open, releasing <em><b>81</b></em> pressure.

== Minute 29 ==
Valves BB, CC, DD, EE, HH, and JJ are open, releasing <em><b>81</b></em> pressure.

== Minute 30 ==
Valves BB, CC, DD, EE, HH, and JJ are open, releasing <em><b>81</b></em> pressure.
</code></pre>
This approach lets you release the most pressure possible in 30 minutes with this valve layout, <code><em><b>1651</b></em></code>.


Work out the steps to release the most pressure in 30 minutes. <em><b>What is the most pressure you can release?</b></em>


