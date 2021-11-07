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


def main():
    print("Enter A or B to choose one of the following two options:\n")
    print("A: Would work with numbers?")
    print("B: Would like to replace a pattern with stars?\n")

    answer = input("Enter A or B (lower or upper case accepted): ").strip().lower()

    while answer != "a" and answer != "b":
        print("Invalid input")
        answer = input("Enter A or B (lower or upper case accepted): ").strip().lower()

    while True:
        answer = input("Enter A or B (lower or upper case accepted): ").strip().lower()
        if answer == 'a' or answer == 'b':
            break
            
def main():
    print('Enter A or B to choose one of the following two options:')
    print('A: Would work with numbers?')
    print('B: Would like to replace a pattern with stars?')

    while True:
        answer = input("Enter A or B (lower or upper case accepted): ").strip().lower()
        if answer == 'a' or answer == 'b':
            break