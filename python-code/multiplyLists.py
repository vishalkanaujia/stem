def multiply(num1, num2):
    sign = -1 if num1[0] < 0 ^ num2[0] < 0 else 1
    
    num1[0], num2[0] = abs(num1[0]), abs(num2[0])
    
    result = [0] * (len(num1) + len(num2))
    
    #for i in reversed(range(len(num1)):


# Multiple the list by a number n
list = [9,2,3]
n = 6
result = [0] * 4
carry = 0
idx = len(result) - 1
for i in reversed(range(len(list))):
    m = list[i] * n  + carry
    digit = m % 10
    carry = m / 10
    result[idx] = digit
    idx -= 1
result[0] = carry
print(result)

print(123//10)
print(3//10)