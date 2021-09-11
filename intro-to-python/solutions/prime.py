def is_prime(n):
    for factor in range(2, n):
        if n % factor == 0:
            return False
        
    return True

# Primes
for prime in range(0, 100):
    if is_prime(prime):
        print(prime)