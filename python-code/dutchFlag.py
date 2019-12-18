def sortArray(A, index):
    pivot = A[index]
    smaller, equal, larger = 0, 0, len(A)

    while equal < larger:
        if A[equal] < pivot:
            A[equal], A[smaller] = A[smaller], A[equal]
            smaller, equal = smaller + 1, equal + 1
            continue
        
        if A[equal] == pivot:
            equal += 1
            continue
        
        larger -= 1
        A[equal], A[larger] = A[larger], A[equal]
        print(A)
        equal += 1

def sort():
    A = [8, 7, 6, 5]
    sortArray(A, 3)
    print(A)
    print(reversed(range(1, len(A))))
sort()