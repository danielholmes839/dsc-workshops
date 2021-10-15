def main():
    print('Enter A or B to choose one of the following two options:')
    print('A: Would work with numbers?')
    print('B: Would like to replace a pattern with stars?')

    while True:
        answer = input("Enter A or B (lower or upper case accepted): ").strip().lower()
        if answer == 'a' or answer == 'b':
            break