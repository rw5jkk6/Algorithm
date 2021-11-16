func sum(n :[Int]) -> Int{
    var s:Int = 0
    for i in n{
        s += i
    }
    return s
}


var n:[Int] = []
for i in 0..<100000000{
    n.append(i)
}

print(sum(n:n))
