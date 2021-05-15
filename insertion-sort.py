# Introduction to Algorithms - 3rd Edition
# page 16

# actual   = [11, 10, -1, 5, 15, 3, 5, 13]
# expected = [-1, 3, 5, 5, 10, 11, 13, 15]
def insertion_sort(arr):
    for i in range(1, len(arr)):
        key = arr[i]
        j = i - 1
        while j >= 0 and arr[j] > key:
            arr[j + 1] = arr[j]
            j -= 1
        arr[j + 1] = key

    return arr

# actual   = [11, 10, -1, 5, 15, 3, 5, 13]
# expected = [15, 13, 11, 10, 5, 5, 3, -1]
def insertion_sort_non_increasing(arr):
    for i in range(1, len(arr)):
        key = arr[i]
        j = i - 1
        while j >= 0 and arr[j] < key:
            arr[j + 1] = arr[j]
            j -= 1
        arr[j + 1] = key

    return arr

if __name__ == '__main__':
    arr = [11, 10, -1, 5, 15, 3, 5, 13]
    print('increasing')
    print(insertion_sort(arr))

    arr2 = [11, 10, -1, 5, 15, 3, 5, 13]
    print('decreasing')
    print(insertion_sort_non_increasing(arr2))