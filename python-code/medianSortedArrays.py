# Program to find median of two sorted arrays
start_a = 0
start_b = 0
end_a = 0
end = 0

def findMedianFaster(arr_a, arr_b, start_a, end_a, start_b, end_b):
    global end
    if end == 20:
        return
    end += 1    
    # first base case of both arrays having 2 elements
    if (end_a - start_a) == 1 and (end_b - start_b) == 1:
        m1 = max(arr_a[start_a], arr_b[start_b])
        m2 = min(arr_a[end_a], arr_b[end_b])
        return (m1+m2)/2
    # Second base case
    # Medians are equal in both the arrays
    m1_idx = (end_a + start_a)/2
    m2_idx = (end_b + start_b)/2
    m1 = arr_a[m1_idx]
    m2 = arr_b[m2_idx]
    print("m1={} m2={}".format(m1, m2))
    if m1 == m2:
        return m2

    if m1 < m2:
        start_a = m1_idx
        end_b = m2_idx
    if m2 < m1:
        start_b = m2_idx
        end_a = m1_idx
    print("start_a={} end_a={} start_b={} end_b={}".format(start_a, end_a, start_b, end_b))    
    return findMedianFaster(arr_a, arr_b, start_a, end_a, start_b, end_b)

def findMedian(arr1, arr2):
    len1 = len(arr1)
    len2 = len(arr2)

    i = 0
    j = 0
    count = 0
    m1 = 0
    m2 = 0

    # Perform the merge phase of mergesort
    while 1:
        if i == len1:
            m1 = m2
            m2 = arr2[0]
            break

        if j == len2:
            m1 = m2
            m2 = arr2[1]
            break

        if count == (len1 + len2) /2:
            break

        if arr1[i] < arr2[j]:
            count += 1
            m1 = m2
            m2= arr1[i]
            i += 1
        elif arr2[j] < arr1[i]:
            count += 1
            m1 = m2
            m2= arr2[j]
            j += 1
    print("m1={} m2={}".format(m1, m2))        
    return m1, m2
        
arr1 = [1, 2, 3, 4, 5]
arr2 = [4, 7, 11, 12, 100]
# 1 4 7 8 10 11
#print(findMedian(arr1, arr2))
print(findMedianFaster(arr1, arr2, 0, 4, 0, 4))