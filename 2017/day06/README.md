# --- Day 6: Memory Reallocation ---

A debugger program here is having an issue: it is trying to repair a memory reallocation routine, but it keeps getting stuck in an infinite loop.


In this area, there are <span title="There are also five currency banks, two river banks, three airplanes banking, a banked billards shot, and a left bank.">sixteen memory banks</span>; each memory bank can hold any number of <em><b>blocks</b></em>. The goal of the reallocation routine is to balance the blocks between the memory banks.


The reallocation routine operates in cycles. In each cycle, it finds the memory bank with the most blocks (ties won by the lowest-numbered memory bank) and redistributes those blocks among the banks. To do this, it removes all of the blocks from the selected bank, then moves to the next (by index) memory bank and inserts one of the blocks. It continues doing this until it runs out of blocks; if it reaches the last memory bank, it wraps around to the first one.


The debugger would like to know how many redistributions can be done before a blocks-in-banks configuration is produced that <em><b>has been seen before</b></em>.


For example, imagine a scenario with only four memory banks:


<ul>
<li>The banks start with <code>0</code>, <code>2</code>, <code>7</code>, and <code>0</code> blocks. The third bank has the most blocks, so it is chosen for redistribution.</li>
<li>Starting with the next bank (the fourth bank) and then continuing to the first bank, the second bank, and so on, the <code>7</code> blocks are spread out over the memory banks. The fourth, first, and second banks get two blocks each, and the third bank gets one back. The final result looks like this: <code>2 4 1 2</code>.</li>
<li>Next, the second bank is chosen because it contains the most blocks (four). Because there are four memory banks, each gets one block. The result is: <code>3 1 2 3</code>.</li>
<li>Now, there is a tie between the first and fourth memory banks, both of which have three blocks. The first bank wins the tie, and its three blocks are distributed evenly over the other three banks, leaving it with none: <code>0 2 3 4</code>.</li>
<li>The fourth bank is chosen, and its four blocks are distributed such that each of the four banks receives one: <code>1 3 4 1</code>.</li>
<li>The third bank is chosen, and the same thing happens: <code>2 4 1 2</code>.</li>
</ul>
At this point, we've reached a state we've seen before: <code>2 4 1 2</code> was already seen. The infinite loop is detected after the fifth block redistribution cycle, and so the answer in this example is <code>5</code>.


Given the initial block counts in your puzzle input, <em><b>how many redistribution cycles</b></em> must be completed before a configuration is produced that has been seen before?


