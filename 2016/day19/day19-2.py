import collections
import sys
def solve_parttwo():
    left = collections.deque()
    right = collections.deque()
    for i in range(1, ELF_COUNT+1):
        if i < (ELF_COUNT // 2) + 1:
            left.append(i)
        else:
            right.appendleft(i)

    print("L", left)
    print("R", right)
    while left and right:
        if len(left) > len(right):
            left.pop()
        else:
            right.pop()

        # rotate
        right.appendleft(left.popleft())
        left.append(right.pop())
        print("L", left)
        print("R", right)
    return left[0] or right[0]
ELF_COUNT=int(sys.argv[1])
solve_parttwo()
