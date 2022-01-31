# --- Day 15: Science for Hungry People ---

Today, you set out on the task of perfecting your milk-dunking cookie recipe.  All you have to do is find the right balance of ingredients.


Your recipe leaves room for exactly <code>100</code> teaspoons of ingredients.  You make a list of the <em><b>remaining ingredients you could use to finish the recipe</b></em> (your puzzle input) and their <em><b>properties per teaspoon</b></em>:


<ul>
<li><code>capacity</code> (how well it helps the cookie absorb milk)</li>
<li><code>durability</code> (how well it keeps the cookie intact when full of milk)</li>
<li><code>flavor</code> (how tasty it makes the cookie)</li>
<li><code>texture</code> (how it improves the feel of the cookie)</li>
<li><code>calories</code> (how many calories it adds to the cookie)</li>
</ul>
You can only measure ingredients in whole-teaspoon amounts accurately, and you have to be accurate so you can reproduce your results in the future.  The <em><b>total score</b></em> of a cookie can be found by adding up each of the properties (negative totals become <code>0</code>) and then multiplying together everything except calories.


For instance, suppose you have <span title="* I know what your preference is, but...">these two ingredients</span>:


<pre><code>Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
</code></pre>
Then, choosing to use <code>44</code> teaspoons of butterscotch and <code>56</code> teaspoons of cinnamon (because the amounts of each ingredient must add up to <code>100</code>) would result in a cookie with the following properties:


<ul>
<li>A <code>capacity</code> of <code>44*-1 + 56*2 = 68</code></li>
<li>A <code>durability</code> of <code>44*-2 + 56*3 = 80</code></li>
<li>A <code>flavor</code> of <code>44*6 + 56*-2 = 152</code></li>
<li>A <code>texture</code> of <code>44*3 + 56*-1 = 76</code></li>
</ul>
Multiplying these together (<code>68 * 80 * 152 * 76</code>, ignoring <code>calories</code> for now) results in a total score of  <code>62842880</code>, which happens to be the best score possible given these ingredients.  If any properties had produced a negative total, it would have instead become zero, causing the whole score to multiply to zero.


Given the ingredients in your kitchen and their properties, what is the <em><b>total score</b></em> of the highest-scoring cookie you can make?


