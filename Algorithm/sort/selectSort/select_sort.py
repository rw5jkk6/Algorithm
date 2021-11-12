data=[3,1,4,2]

for i in range(len(data)):
  tmp = i
  for j in range(i+1,len(data)):
    if data[j] < data[tmp]:
      tmp=j
  data[i],data[tmp] = data[tmp],data[i]

print(data)
