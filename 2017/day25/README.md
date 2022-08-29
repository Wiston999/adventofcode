# --- Day 25: The Halting Problem ---

Following the twisty passageways deeper and deeper into the CPU, you finally reach the <span title="Get it? CPU core?">core</span> of the computer. Here, in the expansive central chamber, you find a grand apparatus that fills the entire room, suspended nanometers above your head.


You had always imagined CPUs to be noisy, chaotic places, bustling with activity. Instead, the room is quiet, motionless, and dark.


Suddenly, you and the CPU's <em><b>garbage collector</b></em> startle each other. "It's not often we get  many visitors here!", he says. You inquire about the stopped machinery.


"It stopped milliseconds ago; not sure why. I'm a garbage collector, not a doctor." You ask what the machine is for.


"Programs these days, don't know their origins. That's the <em><b>Turing machine</b></em>! It's what makes the whole computer work." You try to explain that Turing machines are merely models of computation, but he cuts you off. "No, see, that's just what they <em><b>want</b></em> you to think. Ultimately, inside every CPU, there's a Turing machine driving the whole thing! Too bad this one's broken. [https://www.youtube.com/watch?v=cTwZZz0HV8I](We're doomed!)"


You ask how you can help. "Well, unfortunately, the only way to get the computer running again would be to create a whole new Turing machine from scratch, but there's no <em><b>way</b></em> you can-" He notices the look on your face, gives you a curious glance, shrugs, and goes back to sweeping the floor.


You find the <em><b>Turing machine blueprints</b></em> (your puzzle input) on a tablet in a nearby pile of debris. Looking back up at the broken Turing machine above, you can start to identify its parts:


<ul>
<li>A <em><b>tape</b></em> which contains <code>0</code> repeated infinitely to the left and right.</li>
<li>A <em><b>cursor</b></em>, which can move left or right along the tape and read or write values at its current position.</li>
<li>A set of <em><b>states</b></em>, each containing rules about what to do based on the current value under the cursor.</li>
</ul>
Each slot on the tape has two possible values: <code>0</code> (the starting value for all slots) and <code>1</code>. Based on whether the cursor is pointing at a <code>0</code> or a <code>1</code>, the current state says <em><b>what value to write</b></em> at the current position of the cursor, whether to <em><b>move the cursor</b></em> left or right one slot, and <em><b>which state to use next</b></em>.


For example, suppose you found the following blueprint:


<pre><code>Begin in state A.
Perform a diagnostic checksum after 6 steps.

In state A:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state B.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the left.
    - Continue with state B.

In state B:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state A.
  If the current value is 1:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state A.
</code></pre>
Running it until the number of steps required to take the listed <em><b>diagnostic checksum</b></em> would result in the following tape configurations (with the <em><b>cursor</b></em> marked in square brackets):


<pre><code>... 0  0  0 [0] 0  0 ... (before any steps; about to run state A)
... 0  0  0  1 [0] 0 ... (after 1 step;     about to run state B)
... 0  0  0 [1] 1  0 ... (after 2 steps;    about to run state A)
... 0  0 [0] 0  1  0 ... (after 3 steps;    about to run state B)
... 0 [0] 1  0  1  0 ... (after 4 steps;    about to run state A)
... 0  1 [1] 0  1  0 ... (after 5 steps;    about to run state B)
... 0  1  1 [0] 1  0 ... (after 6 steps;    about to run state A)
</code></pre>
The CPU can confirm that the Turing machine is working by taking a <em><b>diagnostic checksum</b></em> after a specific number of steps (given in the blueprint). Once the specified number of steps have been executed, the Turing machine should pause; once it does, count the number of times <code>1</code> appears on the tape. In the above example, the <em><b>diagnostic checksum</b></em> is <em><b><code>3</code></b></em>.


Recreate the Turing machine and save the computer! <em><b>What is the diagnostic checksum</b></em> it produces once it's working again?


