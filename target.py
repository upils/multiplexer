import sys
for i in range(1, 43):
    if i & 1:
        print(f"{i} stderr stderr stderr stderr stderr stderr stderr stderr", file=sys.stderr)
    else:
        print(f"{i} stdout stdout stdout stdout")
