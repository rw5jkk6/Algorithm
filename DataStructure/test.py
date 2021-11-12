from collections import Counter

c = Counter()
i1 = {'sword':1, 'portion':2}
i2 = {'apple':3, 'portion':2}
i3 = {'bread':1, 'sword':2}

c.update(i1)
c.update(i2)
c.update(i3)

print(c)
# Counter({'portion': 4, 'sword': 3, 'apple': 3, 'bread': 1})

print(c['sword'])
# 3

# 上位２つをとってくるが、ここではappleが出てこない
print(c.most_common(2))
# [('portion', 4), ('sword', 3)]