#!/usr/bin/python
from sys import argv
import heapq

if len(argv) <= 2 or argv[2] not in ["1", "2"]:
    print("Please specify input file and part")
    exit()

with open(argv[1], "r") as f:
    line = f.readline().strip()

disk = list()
empty_index = [list() for i in range(10)]
file_index = list()
for index, size in enumerate(line):
    s = int(size)
    if index%2 != 0 and s> 0:
        heapq.heappush(empty_index[s], len(disk))
    elif index%2 == 0:
        file_index.append((len(disk), s))
        assert len(file_index) == index//2+1
    disk += [index//2 if index%2==0 else -1 for i in range(s)]

# part 1
if argv[2] == "1":
    i, j = (0, len(disk)-1)
    while i < j:
        while disk[i] != -1:
            i+=1
        while disk[j] == -1:
            j-=1
        if i < j:
            disk[i] = disk[j]
            disk[j] = -1
        else:
            break

# part 2 with priority queue
elif argv[2] == "2" and len(argv) == 3 :
    for (index, length) in reversed(file_index):
        min_index = len(disk)
        min_size = 10
        for l in range(length, 10):
            if empty_index[l] and empty_index[l][0] < min_index:
                min_size = l
                min_index = empty_index[l][0]
        if min_index > index: continue
        disk[min_index:min_index+length] = [disk[index]]*length
        disk[index:index+length] = [-1]*length
        heapq.heappop(empty_index[min_size])
        if min_size > length:
            heapq.heappush(empty_index[min_size-length], min_index+length)

# part 2 with brute force
elif argv[3] == "-b":
    i, j = (0, len(disk)-1)
    ilen, jlen = (0,0)
    fileno = disk[j]
    while fileno > 0:
        while disk[j] != fileno:
            j-=1
        while disk[j-jlen] == fileno:
            jlen += 1
        i,ilen = (0,0)
        while i <= j-jlen:
            while i < j and disk[i] != -1:
                i+=1
            if i >= j:
                break
            while disk[i+ilen] == -1:
                ilen += 1
            if ilen < jlen:
                i = i+ilen
                ilen=0
                continue
            disk[i:i+jlen] = [fileno]*jlen
            disk[j-jlen+1:j+1] = [-1]*jlen
            break
        j -= jlen
        jlen = 0
        fileno -= 1

checksum = sum(map(lambda x,y: x*y if y>0 else 0, *zip(*enumerate(disk))))
print(checksum)
print(disk)
