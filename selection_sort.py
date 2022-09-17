# Time: O(n^2) | Space: O(1)
def selectionSort(array):
    for currentIdx in range(len(array)):
        smallestIdx = currentIdx
        for j in range(currentIdx+1, len(array)):
            if array[smallestIdx] > array[j]:
                smallestIdx = j
        swap(currentIdx, smallestIdx, array)

    return array

def swap(i, j, array):
    array[i], array[j] = array[j], array[i] 

if __name__ == "__main__":
    array = [8, 2, 5, 4, 1]
    print(selectionSort(array))

# output: [1, 2, 4, 5, 8]