import math


def atlantic(sn):
    if len(sn) >= 9:
        return "new"
    
    return  "old"


def southern(n):
    n = int(n)
    if n == 1:
        pounds = float(input("Enter number of weight in pounds: "))
        ounces = float(input("Enter number of weight in ounces: "))
        kilograms = (pounds * 16 + ounces) * 0.02835
        print(f'{pounds} pounds and {ounces} ounces is approximately {kilograms} kilograms')
        
    elif n == 2:
        name = input("What is your name? ")
        student_number = input("What is your student number? ")
        print(f'{name} you have a {atlantic(student_number)} student number')
    else:
        print("Invalid input")


def pacific(g1, g2, g3):
    return (g1 >= 50 and g2 >= 50 and g3 >= 50) and (g1 >= 80 or g2 >= 80 or g3 >= 80)


def arctic(n):
     if n // 100000 !=0: # 6 digit number testing. Alternative solution: if n>=100000 and n<=999999
          d1 = n//100000
          d2 = (n %100000)//10000
          d3 = (n %10000)//1000
          d4 = (n %1000)//100
          d5 = (n %100)//10
          d6 = n % 10
          return d1==d6 and d2==d5 and d3==d4
    
     else: # 4 digit number testing. Alternative solution: elif n>=1000 and n<=9999
          d1 = n//1000
          d2 = (n %1000)//100
          d3 = (n %100)//10
          d4 = n % 10
          return d1==d4 and d2==d3


def arctic2(n):
    """
    (int) -> bool
    """
    palindrome = False
    n = abs(int(n))
    length = math.floor(math.log(n, 10 )+1)
    
    if length == 4:
        n1 = n % 10000 // 1000
        n2 = n % 1000 // 100
        n3 = n % 100 // 10
        n4 = n % 10
        
        if (n1 == n4 and n2 == n3):
            palindrome = True
        
    if length == 6:
        n1 = n % 1000000//100000
        n2 = n % 100000 // 10000
        n3 = n % 10000 // 1000
        n4 = n % 1000 // 100
        n5 = n % 100 // 10
        n6 = n % 10
        
        if (n1 == n6 and n2 == n5 and n3 == n4):
            palindrome = True

    return palindrome
