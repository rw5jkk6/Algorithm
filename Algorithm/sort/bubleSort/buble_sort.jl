x = [4, 1, 3, 2]

for i in 1:4
    for j in 1:4 - 1
        if x[j] > x[j+1]
            x[j], x[j+1] = x[j+1], x[j]
        end
    end
end

println(x)

# list = [4, 1, 3, 2]

# function swap(a, b)
#     tmp = a
#     a = b
#     b = tmp
# end

# for i = 1:5
#     for j = 1:5
#         if list[j] > list[j+1]
#             swap(list[j], list[j+1])
#         end
#     end
# end
# print(list)
