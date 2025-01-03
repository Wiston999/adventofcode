# --- Day 7: Recursive Circus ---

Wandering further through the circuits of the computer, you come upon a tower of <span title="Turtles, all the way down.">programs</span> that have gotten themselves into a bit of trouble.  A recursive algorithm has gotten out of hand, and now they're balanced precariously in a large tower.


One program at the bottom supports the entire tower. It's holding a large disc, and on the disc are balanced several more sub-towers. At the bottom of these sub-towers, standing on the bottom disc, are other programs, each holding <em><b>their</b></em> own disc, and so on. At the very tops of these sub-sub-sub-...-towers, many programs stand simply keeping the disc below them balanced but with no disc of their own.


You offer to help, but first you need to understand the structure of these towers.  You ask each program to yell out their <em><b>name</b></em>, their <em><b>weight</b></em>, and (if they're holding a disc) the <em><b>names of the programs immediately above them</b></em> balancing on that disc. You write this information down (your puzzle input). Unfortunately, in their panic, they don't do this in an orderly fashion; by the time you're done, you're not sure which program gave which information.


For example, if your list is the following:


<pre><code>pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)
</code></pre>
...then you would be able to recreate the structure of the towers that looks like this:


<pre><code>                gyxo
              /     
         ugml - ebii
       /      \     
      |         jptl
      |        
      |         pbga
     /        /
tknk --- padx - havc
     \        \
      |         qoyq
      |             
      |         ktlj
       \      /     
         fwft - cntj
              \     
                xhth
</code></pre>
In this example, <code>tknk</code> is at the bottom of the tower (the <em><b>bottom program</b></em>), and is holding up <code>ugml</code>, <code>padx</code>, and <code>fwft</code>.  Those programs are, in turn, holding up other programs; in this example, none of those programs are holding up any other programs, and are all the tops of their own towers. (The actual tower balancing in front of you is much larger.)


Before you're ready to help them, you need to make sure your information is correct.  <em><b>What is the name of the bottom program?</b></em>


