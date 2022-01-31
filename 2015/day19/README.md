# --- Day 19: Medicine for Rudolph ---

Rudolph the Red-Nosed Reindeer is sick!  His nose isn't shining very brightly, and he needs medicine.


Red-Nosed Reindeer biology isn't similar to regular reindeer biology; Rudolph is going to need custom-made medicine.  Unfortunately, Red-Nosed Reindeer chemistry isn't similar to regular reindeer chemistry, either.


The North Pole is equipped with a Red-Nosed Reindeer nuclear fusion/fission plant, capable of constructing any Red-Nosed Reindeer molecule you need.  It works by starting with some input molecule and then doing a series of <em><b>replacements</b></em>, one per step, until it has the right molecule.


However, the machine has to be calibrated before it can be used.  Calibration involves determining the number of molecules that can be generated in one step from a given starting point.


For example, imagine a simpler machine that supports only the following replacements:


<pre><code>H => HO
H => OH
O => HH
</code></pre>
Given the replacements above and starting with <code>HOH</code>, the following molecules could be generated:


<ul>
<li><code>HOOH</code> (via <code>H => HO</code> on the first <code>H</code>).</li>
<li><code>HOHO</code> (via <code>H => HO</code> on the second <code>H</code>).</li>
<li><code>OHOH</code> (via <code>H => OH</code> on the first <code>H</code>).</li>
<li><code>HOOH</code> (via <code>H => OH</code> on the second <code>H</code>).</li>
<li><code>HHHH</code> (via <code>O => HH</code>).</li>
</ul>
So, in the example above, there are <code>4</code> <em><b>distinct</b></em> molecules (not five, because <code>HOOH</code> appears twice) after one replacement from <code>HOH</code>. Santa's favorite molecule, <code>HOHOHO</code>, can become <code>7</code> <em><b>distinct</b></em> molecules (over nine replacements: six from <code>H</code>, and three from <code>O</code>).


The machine replaces without regard for the surrounding characters.  For example, given the string <code>H2O</code>, the transition <code>H => OO</code> would result in <code>OO2O</code>.


Your puzzle input describes all of the possible replacements and, at the bottom, the medicine molecule for which you need to calibrate the machine.  <em><b>How many distinct molecules can be created</b></em> after all the different ways you can do one replacement on the medicine molecule?


