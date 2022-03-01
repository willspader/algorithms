if __name__ == '__main__':
    n = int(input())

    for i in range(n):
        word = input()

        if len(word) > 10:
            print(word[0] + str(len(word) - 2) + word[len(word) - 1])
        else:
            print(word)
