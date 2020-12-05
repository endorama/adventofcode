from itertools import permutations
from functools import reduce

if __name__ == "__main__":
    with open('./input.txt') as reader:
        lines = reader.readlines()
        numbers = [int(l) for l in lines]
        print ( numbers )

        values = ()
        for p in permutations(numbers,3):
            #  print(p)
            if sum(p) == 2020:
                values = p
                break

        print("values: {}".format(values))
        print(reduce((lambda x, y: x * y), values))
