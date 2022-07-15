# --- Day 6: Signals and Noise ---

Something is jamming your communications with Santa. Fortunately, your signal is only partially jammed, and protocol in situations like this is to switch to a simple [https://en.wikipedia.org/wiki/Repetition_code](repetition code) to get the message through.


In this model, the same message is sent repeatedly.  You've recorded the repeating message signal (your puzzle input), but the data seems quite corrupted - almost too badly to recover. <em><b>Almost</b></em>.


All you need to do is figure out which character is most frequent for each position. For example, suppose you had recorded the following messages:


<pre><code>eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar
</code></pre>
The most common character in the first column is <code>e</code>; in the second, <code>a</code>; in the third, <code>s</code>, and so on. Combining these characters returns the error-corrected message, <code>easter</code>.


Given the recording in your puzzle input, <em><b>what is the error-corrected version</b></em> of the message being sent?


