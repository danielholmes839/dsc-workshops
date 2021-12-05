def split_guesses(guesses, answer):
    """(list of number, number) -> dict of {str: list of number}"""
    d = {
        'low': [], 
        'correct': [], 
        'high': []
    }
    
    for num in guesses:
        if num < answer:
            d['low'].append(num)
        elif num > answer:
            d['high'].append(num)
        else:
            d['correct'].append(num)

    if len(d['low']) == 0:
        del d['low']
    
    if len(d['high']) == 0:
        del d['high']
    
    if len(d['correct']) == 0:
        del d['correct']
    
    return d