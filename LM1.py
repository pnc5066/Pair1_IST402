# Punit Chudasama
# LM1: Assignment Brute Force Ceasar Cipher

def brute_force(message):
    for shift in range(26):
        decrypted_message = ""
        for char in message:
            if char.isalpha():
                if char.islower():
                    decrypted_char = chr((ord(char) - ord('a') - shift) % 26 + ord('a'))
                else:
                    decrypted_char = chr((ord(char) - ord('A') - shift) % 26 + ord('A'))
                decrypted_message += decrypted_char
            else:
                decrypted_message += char
        print(f"Shift {shift}: {decrypted_message}")


message = "Ugew gnwj zwjw Oslkgf"
brute_force(message)
