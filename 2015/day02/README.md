# --- Day 2: I Was Told There Would Be No Math ---

The elves are running low on wrapping paper, and so they need to submit an order for more.  They have a list of the dimensions (length <code>l</code>, width <code>w</code>, and height <code>h</code>) of each present, and only want to order exactly as much as they need.


Fortunately, every present is a box (a perfect [https://en.wikipedia.org/wiki/Cuboid#Rectangular_cuboid](right rectangular prism)), which makes calculating the required wrapping paper for each gift a little easier: find the surface area of the box, which is <code>2*l*w + 2*w*h + 2*h*l</code>.  The elves also need a little extra paper for each present: the area of the smallest side.


For example:


<ul>
<li>A present with dimensions <code>2x3x4</code> requires <code>2*6 + 2*12 + 2*8 = 52</code> square feet of wrapping paper plus <code>6</code> square feet of slack, for a total of <code>58</code> square feet.</li>
<li>A present with dimensions <code>1x1x10</code> requires <code>2*1 + 2*10 + 2*10 = 42</code> square feet of wrapping paper plus <code>1</code> square foot of slack, for a total of <code>43</code> square feet.</li>
</ul>
All numbers in the elves' list are in <span title="Yes, I realize most of these presents are luxury yachts.">feet</span>.  How many total <em><b>square feet of wrapping paper</b></em> should they order?


