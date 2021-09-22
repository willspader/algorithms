# m   -> the value to search
# arr -> the numbers provided
def resolve(m, arr):
    hashTable = {}

    i = 0
    for cost in arr:
        sub = m - cost

        if sub in hashTable:
            return [hashTable[sub] + 1, i + 1]

        hashTable[cost] = i

        i += 1
