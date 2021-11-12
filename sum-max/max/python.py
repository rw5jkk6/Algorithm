def max(num_list):
    maxnum = 0 
    for i in num_list:
        if maxnum < i:
           maxnum = i
    return maxnum

print(max([i for i in range(10000000)]))
