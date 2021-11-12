def sum(num_list):
    sum = 0
    for i in num_list:
        sum += i
    return sum

print(sum([i for i in range(100000000)]))
