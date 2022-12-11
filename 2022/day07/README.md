# --- Day 7: No Space Left On Device ---

You can hear birds chirping and raindrops hitting leaves as the expedition proceeds. Occasionally, you can even hear much louder sounds in the distance; how big do the animals get out here, anyway?


The device the Elves gave you has problems with more than just its communication system. You try to run a system update:


<pre><code>$ system-update --please --pretty-please-with-sugar-on-top
<span title="E099 PROGRAMMER IS OVERLY POLITE">Error</span>: No space left on device
</code></pre>
Perhaps you can delete some files to make space for the update?


You browse around the filesystem to assess the situation and save the resulting terminal output (your puzzle input). For example:


<pre><code>$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
</code></pre>
The filesystem consists of a tree of files (plain data) and directories (which can contain other directories or files). The outermost directory is called <code>/</code>. You can navigate around the filesystem, moving into or out of directories and listing the contents of the directory you're currently in.


Within the terminal output, lines that begin with <code>$</code> are <em><b>commands you executed</b></em>, very much like some modern computers:


<ul>
<li><code>cd</code> means <em><b>change directory</b></em>. This changes which directory is the current directory, but the specific result depends on the argument:
  <ul>
  <li><code>cd x</code> moves <em><b>in</b></em> one level: it looks in the current directory for the directory named <code>x</code> and makes it the current directory.</li>
  <li><code>cd ..</code> moves <em><b>out</b></em> one level: it finds the directory that contains the current directory, then makes that directory the current directory.</li>
  <li><code>cd /</code> switches the current directory to the outermost directory, <code>/</code>.</li>
  </ul>
</li>
<li><code>ls</code> means <em><b>list</b></em>. It prints out all of the files and directories immediately contained by the current directory:
  <ul>
  <li><code>123 abc</code> means that the current directory contains a file named <code>abc</code> with size <code>123</code>.</li>
  <li><code>dir xyz</code> means that the current directory contains a directory named <code>xyz</code>.</li>
  </ul>
</li>
</ul>
Given the commands and output in the example above, you can determine that the filesystem looks visually like this:


<pre><code>- / (dir)
  - a (dir)
    - e (dir)
      - i (file, size=584)
    - f (file, size=29116)
    - g (file, size=2557)
    - h.lst (file, size=62596)
  - b.txt (file, size=14848514)
  - c.dat (file, size=8504156)
  - d (dir)
    - j (file, size=4060174)
    - d.log (file, size=8033020)
    - d.ext (file, size=5626152)
    - k (file, size=7214296)
</code></pre>
Here, there are four directories: <code>/</code> (the outermost directory), <code>a</code> and <code>d</code> (which are in <code>/</code>), and <code>e</code> (which is in <code>a</code>). These directories also contain files of various sizes.


Since the disk is full, your first step should probably be to find directories that are good candidates for deletion. To do this, you need to determine the <em><b>total size</b></em> of each directory. The total size of a directory is the sum of the sizes of the files it contains, directly or indirectly. (Directories themselves do not count as having any intrinsic size.)


The total sizes of the directories above can be found as follows:


<ul>
<li>The total size of directory <code>e</code> is <em><b>584</b></em> because it contains a single file <code>i</code> of size 584 and no other directories.</li>
<li>The directory <code>a</code> has total size <em><b>94853</b></em> because it contains files <code>f</code> (size 29116), <code>g</code> (size 2557), and <code>h.lst</code> (size 62596), plus file <code>i</code> indirectly (<code>a</code> contains <code>e</code> which contains <code>i</code>).</li>
<li>Directory <code>d</code> has total size <em><b>24933642</b></em>.</li>
<li>As the outermost directory, <code>/</code> contains every file. Its total size is <em><b>48381165</b></em>, the sum of the size of every file.</li>
</ul>
To begin, find all of the directories with a total size of <em><b>at most 100000</b></em>, then calculate the sum of their total sizes. In the example above, these directories are <code>a</code> and <code>e</code>; the sum of their total sizes is <code><em><b>95437</b></em></code> (94853 + 584). (As in this example, this process can count files more than once!)


Find all of the directories with a total size of at most 100000. <em><b>What is the sum of the total sizes of those directories?</b></em>


