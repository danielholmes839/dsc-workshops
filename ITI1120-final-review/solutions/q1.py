def mekong(l, div):
    """(list of int, list of int)-> list of tuples"""
    final = []
    for i in range(0, len(l)):
        for j in range(i+1, len(l)):
            for k in range(0, len(div)):
                if l[i] % div[k] == 0 and l[j] % div[k] == 0:
                    t = l[i], l[j]
                    if t not in final:
                        final.append(t)
    return final