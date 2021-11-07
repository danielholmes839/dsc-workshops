# Q1
def list_work(l):
    l2 = []  # new list
    if len(l) == 1:
        return l
    
    for i in range(0, len(l)-1):
        first = l[i]
        second = l[i+1]
        
        # both even or odd, then add their sum to new list
        if (first % 2 == 0 and second % 2 == 0) or (first % 2 != 0 and second % 2 != 0):
            l2.append(first + second)
        else:
            l2.append(0)
            
    return l2

# Q2
def get_list_ofint(s):  # 5 points
    """
    (str) -> list of int
    """
    # return [int(section.strip(',')) for section in s.split(" ")]
    
    # Remove commas
    string = ""
    for character in s:
        if character != ",":
            string += character
    
    # Split by spaces
    numbers = []
    for number_string in string.split(" "):
        numbers.append(int(number_string))
    
    return numbers

# Q5
def diff_start_end(l, q):
    first_diff = None
    last_diff = None
    
    for i in range(len(l)):
        if first_diff is None and l[i] != q[i]:
            first_diff = i
            
        if last_diff is None and l[-i-1] != q[-i-1]:
            last_diff = -i-1
        
    return first_diff, last_diff

def diff_start_end(l, q):
    first_diff = None
    for i in range(0, len(l)):
        if l[i] != q[i]:
            first_diff = i
            break
            
    last_diff = None
    for i in range(-1, -len(l)-1, -1):
        if l[i] != q[i]:
            last_diff = i
            break
            
    return first_diff, last_diff

# Q6
def make_teams(players, num_teams):
    teams = []
    for i in range(num_teams):
        teams.append(players[i::num_teams])
    return teams

# Q7
def draw_w_stars(n):
    for spaces, i in enumerate(range(n, 0, -2)):
        print(" "*spaces + "*"*i)