# --- Day 3: Rucksack Reorganization ---

One Elf has the important job of loading all of the [https://en.wikipedia.org/wiki/Rucksack](rucksacks) with supplies for the <span title="Where there's jungle, there's hijinxs.">jungle</span> journey. Unfortunately, that Elf didn't quite follow the packing instructions, and so a few items now need to be rearranged.


Each rucksack has two large <em><b>compartments</b></em>. All items of a given type are meant to go into exactly one of the two compartments. The Elf that did the packing failed to follow this rule for exactly one item type per rucksack.


The Elves have made a list of all of the items currently in each rucksack (your puzzle input), but they need your help finding the errors. Every item type is identified by a single lowercase or uppercase letter (that is, <code>a</code> and <code>A</code> refer to different types of items).


The list of items for each rucksack is given as characters all on a single line. A given rucksack always has the same number of items in each of its two compartments, so the first half of the characters represent items in the first compartment, while the second half of the characters represent items in the second compartment.


For example, suppose you have the following list of contents from six rucksacks:


<pre><code>vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
</code></pre>
<ul>
<li>The first rucksack contains the items <code>vJrwpWtwJgWrhcsFMMfFFhFp</code>, which means its first compartment contains the items <code>vJrwpWtwJgWr</code>, while the second compartment contains the items <code>hcsFMMfFFhFp</code>. The only item type that appears in both compartments is lowercase <code><em><b>p</b></em></code>.</li>
<li>The second rucksack's compartments contain <code>jqHRNqRjqzjGDLGL</code> and <code>rsFMfFZSrLrFZsSL</code>. The only item type that appears in both compartments is uppercase <code><em><b>L</b></em></code>.</li>
<li>The third rucksack's compartments contain <code>PmmdzqPrV</code> and <code>vPwwTWBwg</code>; the only common item type is uppercase <code><em><b>P</b></em></code>.</li>
<li>The fourth rucksack's compartments only share item type <code><em><b>v</b></em></code>.</li>
<li>The fifth rucksack's compartments only share item type <code><em><b>t</b></em></code>.</li>
<li>The sixth rucksack's compartments only share item type <code><em><b>s</b></em></code>.</li>
</ul>
To help prioritize item rearrangement, every item type can be converted to a <em><b>priority</b></em>:


<ul>
<li>Lowercase item types <code>a</code> through <code>z</code> have priorities 1 through 26.</li>
<li>Uppercase item types <code>A</code> through <code>Z</code> have priorities 27 through 52.</li>
</ul>
In the above example, the priority of the item type that appears in both compartments of each rucksack is 16 (<code>p</code>), 38 (<code>L</code>), 42 (<code>P</code>), 22 (<code>v</code>), 20 (<code>t</code>), and 19 (<code>s</code>); the sum of these is <code><em><b>157</b></em></code>.


Find the item type that appears in both compartments of each rucksack. <em><b>What is the sum of the priorities of those item types?</b></em>


