def pattern_to_stars(s, p):
    i = 0        # Index of the character 
    count = 1    # How many stars to add
    string = ""  # How many 
    
    while i < len(s):
        if s[i:i+len(p)] == p:
            # Found the pattern
            string += "*"*count
            count += 1
            i += len(p)
        
        else:
            string += s[i]
            i += 1

    return string


print(f("abc123123abc123", "123"))
