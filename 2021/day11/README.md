# --- Day 11: Dumbo Octopus ---

You enter a large cavern full of rare bioluminescent [https://www.youtube.com/watch?v=eih-VSaS2g0](dumbo octopuses)! They seem to not like the Christmas lights on your submarine, so you turn them off for now.


There are 100 <span title="I know it's weird; I grew saying 'octopi' too.">octopuses</span> arranged neatly in a 10 by 10 grid. Each octopus slowly gains <em><b>energy</b></em> over time and <em><b>flashes</b></em> brightly for a moment when its energy is full. Although your lights are off, maybe you could navigate through the cave without disturbing the octopuses if you could predict when the flashes of light will happen.


Each octopus has an <em><b>energy level</b></em> - your submarine can remotely measure the energy level of each octopus (your puzzle input). For example:


<pre><code>5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
</code></pre>
The energy level of each octopus is a value between <code>0</code> and <code>9</code>. Here, the top-left octopus has an energy level of <code>5</code>, the bottom-right one has an energy level of <code>6</code>, and so on.


You can model the energy levels and flashes of light in <em><b>steps</b></em>. During a single step, the following occurs:


<ul>
<li>First, the energy level of each octopus increases by <code>1</code>.</li>
<li>Then, any octopus with an energy level greater than <code>9</code> <em><b>flashes</b></em>. This increases the energy level of all adjacent octopuses by <code>1</code>, including octopuses that are diagonally adjacent. If this causes an octopus to have an energy level greater than <code>9</code>, it <em><b>also flashes</b></em>. This process continues as long as new octopuses keep having their energy level increased beyond <code>9</code>. (An octopus can only flash <em><b>at most once per step</b></em>.)</li>
<li>Finally, any octopus that flashed during this step has its energy level set to <code>0</code>, as it used all of its energy to flash.</li>
</ul>
Adjacent flashes can cause an octopus to flash on a step even if it begins that step with very little energy. Consider the middle octopus with <code>1</code> energy in this situation:


<pre><code>Before any steps:
11111
19991
19191
19991
11111

After step 1:
34543
4<em><b>000</b></em>4
5<em><b>000</b></em>5
4<em><b>000</b></em>4
34543

After step 2:
45654
51115
61116
51115
45654
</code></pre>
An octopus is <em><b>highlighted</b></em> when it flashed during the given step.


Here is how the larger example above progresses:


<pre><code>Before any steps:
5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526

After step 1:
6594254334
3856965822
6375667284
7252447257
7468496589
5278635756
3287952832
7993992245
5957959665
6394862637

After step 2:
88<em><b>0</b></em>7476555
5<em><b>0</b></em>89<em><b>0</b></em>87<em><b>0</b></em>54
85978896<em><b>0</b></em>8
84857696<em><b>00</b></em>
87<em><b>00</b></em>9<em><b>0</b></em>88<em><b>00</b></em>
66<em><b>000</b></em>88989
68<em><b>0000</b></em>5943
<em><b>000000</b></em>7456
9<em><b>000000</b></em>876
87<em><b>0000</b></em>6848

After step 3:
<em><b>00</b></em>5<em><b>0</b></em>9<em><b>00</b></em>866
85<em><b>00</b></em>8<em><b>00</b></em>575
99<em><b>000000</b></em>39
97<em><b>000000</b></em>41
9935<em><b>0</b></em>8<em><b>00</b></em>63
77123<em><b>00000</b></em>
791125<em><b>000</b></em>9
221113<em><b>0000</b></em>
<em><b>0</b></em>421125<em><b>000</b></em>
<em><b>00</b></em>21119<em><b>000</b></em>

After step 4:
2263<em><b>0</b></em>31977
<em><b>0</b></em>923<em><b>0</b></em>31697
<em><b>00</b></em>3222115<em><b>0</b></em>
<em><b>00</b></em>41111163
<em><b>00</b></em>76191174
<em><b>00</b></em>53411122
<em><b>00</b></em>4236112<em><b>0</b></em>
5532241122
1532247211
113223<em><b>0</b></em>211

After step 5:
4484144<em><b>000</b></em>
2<em><b>0</b></em>44144<em><b>000</b></em>
2253333493
1152333274
11873<em><b>0</b></em>3285
1164633233
1153472231
6643352233
2643358322
2243341322

After step 6:
5595255111
3155255222
33644446<em><b>0</b></em>5
2263444496
2298414396
2275744344
2264583342
7754463344
3754469433
3354452433

After step 7:
67<em><b>0</b></em>7366222
4377366333
4475555827
34966557<em><b>0</b></em>9
35<em><b>00</b></em>6256<em><b>0</b></em>9
35<em><b>0</b></em>9955566
3486694453
8865585555
486558<em><b>0</b></em>644
4465574644

After step 8:
7818477333
5488477444
5697666949
46<em><b>0</b></em>876683<em><b>0</b></em>
473494673<em><b>0</b></em>
474<em><b>00</b></em>97688
69<em><b>0000</b></em>7564
<em><b>000000</b></em>9666
8<em><b>00000</b></em>4755
68<em><b>0000</b></em>7755

After step 9:
9<em><b>0</b></em>6<em><b>0000</b></em>644
78<em><b>00000</b></em>976
69<em><b>000000</b></em>8<em><b>0</b></em>
584<em><b>00000</b></em>82
5858<em><b>0000</b></em>93
69624<em><b>00000</b></em>
8<em><b>0</b></em>2125<em><b>000</b></em>9
222113<em><b>000</b></em>9
9111128<em><b>0</b></em>97
7911119976

After step 10:
<em><b>0</b></em>481112976
<em><b>00</b></em>31112<em><b>00</b></em>9
<em><b>00</b></em>411125<em><b>0</b></em>4
<em><b>00</b></em>811114<em><b>0</b></em>6
<em><b>00</b></em>991113<em><b>0</b></em>6
<em><b>00</b></em>93511233
<em><b>0</b></em>44236113<em><b>0</b></em>
553225235<em><b>0</b></em>
<em><b>0</b></em>53225<em><b>0</b></em>6<em><b>00</b></em>
<em><b>00</b></em>3224<em><b>0000</b></em>
</code></pre>

After step 10, there have been a total of <code>204</code> flashes. Fast forwarding, here is the same configuration every 10 steps:



<pre><code>After step 20:
3936556452
56865568<em><b>0</b></em>6
449655569<em><b>0</b></em>
444865558<em><b>0</b></em>
445686557<em><b>0</b></em>
568<em><b>00</b></em>86577
7<em><b>00000</b></em>9896
<em><b>0000000</b></em>344
6<em><b>000000</b></em>364
46<em><b>0000</b></em>9543

After step 30:
<em><b>0</b></em>643334118
4253334611
3374333458
2225333337
2229333338
2276733333
2754574565
5544458511
9444447111
7944446119

After step 40:
6211111981
<em><b>0</b></em>421111119
<em><b>00</b></em>42111115
<em><b>000</b></em>3111115
<em><b>000</b></em>3111116
<em><b>00</b></em>65611111
<em><b>0</b></em>532351111
3322234597
2222222976
2222222762

After step 50:
9655556447
48655568<em><b>0</b></em>5
448655569<em><b>0</b></em>
445865558<em><b>0</b></em>
457486557<em><b>0</b></em>
57<em><b>000</b></em>86566
6<em><b>00000</b></em>9887
8<em><b>000000</b></em>533
68<em><b>00000</b></em>633
568<em><b>0000</b></em>538

After step 60:
25333342<em><b>00</b></em>
274333464<em><b>0</b></em>
2264333458
2225333337
2225333338
2287833333
3854573455
1854458611
1175447111
1115446111

After step 70:
8211111164
<em><b>0</b></em>421111166
<em><b>00</b></em>42111114
<em><b>000</b></em>4211115
<em><b>0000</b></em>211116
<em><b>00</b></em>65611111
<em><b>0</b></em>532351111
7322235117
5722223475
4572222754

After step 80:
1755555697
59655556<em><b>0</b></em>9
448655568<em><b>0</b></em>
445865558<em><b>0</b></em>
457<em><b>0</b></em>86557<em><b>0</b></em>
57<em><b>000</b></em>86566
7<em><b>00000</b></em>8666
<em><b>0000000</b></em>99<em><b>0</b></em>
<em><b>0000000</b></em>8<em><b>00</b></em>
<em><b>0000000000</b></em>

After step 90:
7433333522
2643333522
2264333458
2226433337
2222433338
2287833333
2854573333
4854458333
3387779333
3333333333

After step 100:
<em><b>0</b></em>397666866
<em><b>0</b></em>749766918
<em><b>00</b></em>53976933
<em><b>000</b></em>4297822
<em><b>000</b></em>4229892
<em><b>00</b></em>53222877
<em><b>0</b></em>532222966
9322228966
7922286866
6789998766
</code></pre>
After 100 steps, there have been a total of <code><em><b>1656</b></em></code> flashes.


Given the starting energy levels of the dumbo octopuses in your cavern, simulate 100 steps. <em><b>How many total flashes are there after 100 steps?</b></em>


