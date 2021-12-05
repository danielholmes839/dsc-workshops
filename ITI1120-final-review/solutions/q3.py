def decode(secret_message, encoding):
    """(str, dict of {str: str}) -> str """
    final = ""
    for e in secret_message:
        new_letter = encoding[e]
        final += new_letter
    return final