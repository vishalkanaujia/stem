def find_lexical_smaller_string(str1, str2):
    pass
    # check 1: proper prefix
    if str2.find(str1) == 0 and len(str1) != len(str2):
        return "proper prefix"
   
    # check 2: string2 is lexicographically smaller
    is_impossible = True
    for c,d in zip(str1, str2):
        # print(c,d)
        if c < d:
            is_impossible = False
            break
        
    if is_impossible:
        return "---"
    
    
    return find_and_swap_pair(str1, str2)


def find_and_swap_pair(str1, str2):
    char_set = {}
    index = 0
    did_swap = False
    
    for c in str1:
        if c in char_set:
            char_set[c].append(index)
        else:
            char_set[c] = [index]
        index += 1
    
    print("char_set={}".format(char_set))
    index = 0

    s1 = list(str1)
    s2 = list(str2)
    
    for c,d in zip(s1, s2):
        if c == d:
            index += 1
            continue
        
        if c > d:
            ch = find_char_in_range(char_set, index, d)
            print("ch={}".format(ch))
            s1[index], s1[ch] = s1[ch], s1[index]
            did_swap = True
            break
        
        index += 1
        
    if did_swap:
        print("s1={}".format(s1))
        new_str1 = ''.join(s1)
        return new_str1
    
    return None


def find_char_in_range(char_set, index, high_limit):
    ascii_target = ord(high_limit)
    print("ascii_target={}-{}".format(high_limit, ascii_target))
    
    ascii_target -= 1
    
    while ascii_target >= 0:
        c = chr(ascii_target)
        ascii_target -= 1

        if c in char_set:
            index_list = char_set[c]
            min_index = find_minimum_index(index_list, index)
            if min_index < 0:
                continue
            else:
                break

    if min_index >= 0:
        print("min_index={}".format(min_index))
    return min_index
    
        
def find_minimum_index(index_list, index):
    if len(index_list) > 0:
        print("index_list={}".format(index_list))
        min_index = index_list[0]
        del index_list[0]
        return min_index
    return -1
        
print (find_lexical_smaller_string('AON', 'AONE'))
print (find_lexical_smaller_string('AZAMON', 'APPLE'))
print (find_lexical_smaller_string('AZAMON', 'AAAAAAAAAAPPLE'))