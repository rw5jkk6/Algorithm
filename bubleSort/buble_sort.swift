var data = [4, 1, 2, 3]


for _ in 0 ..< data.count{
    for j in 0 ..< data.count - 1{
        if data[j] > data[j + 1]{
            (data[j], data[j + 1]) = (data[j + 1], data[j])
        }
    }
}
print(data)
