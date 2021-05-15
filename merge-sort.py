# Introduction to Algorithms - 3rd Edition
# page 34

def merge(arr, p, q, r):
    n = q - p + 1
    m = r - q

    left_arr = []
    right_arr = []

    for i in range(0, n):
        left_arr.append(arr[p + i])

    for j in range(0, m):
        right_arr.append(arr[q + j + 1])

    i = 0
    j = 0
    k = p
    while i < len(left_arr) and j < len(right_arr) and k <= r:
        if left_arr[i] <= right_arr[j]:
            arr[k] = left_arr[i]
            k += 1
            i += 1
        else:
            arr[k] = right_arr[j]
            k += 1
            j += 1

    while i < len(left_arr):
        arr[k] = left_arr[i]
        k += 1
        i += 1

    while j < len(right_arr):
        arr[k] = right_arr[j]
        k += 1
        j += 1

    return arr


# p is the first position of the array, starting from 0
# q is the middle of the array
# r is the last position of the array, max n-1 where n = len(arr)
# which means -> [1, 5, 3, 7, 8] -> p = 0 -> q = 2 -> r = 4

# actual   = [1, 9, 3, 15, -8]
# expected = [-8, 1, 3, 9, 15]

def merge_sort(arr, p, r):
    if p < r:
        q = (p + r) // 2
        merge_sort(arr, p, q)
        merge_sort(arr, q + 1, r)
        merge(arr, p, q, r)

if __name__ == '__main__':
    arr = [1, 9, 3, 15, -8]
    merge_sort(arr, 0, len(arr) - 1)
    print(arr)