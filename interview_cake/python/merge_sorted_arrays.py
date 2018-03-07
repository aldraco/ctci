from copy import copy


def _yield_merged_sorted_array_elems(a, b):
    """A generator to assist constructing a merged array.

    :param a: sorted list of ints
    :param b: sorted list of ints
    """
    while (a and b):
        if a[0] <= b[0]:
            yield a.pop(0)
        else:
            yield b.pop(0)

    if not a:
        for item in b:
            yield item
    else:
        for item in a:
            yield item

def merge_arrays(a, b):
    a = copy(a)
    b = copy(b)
    return [i for i in _yield_merged_sorted_array_elems(a, b)]

def main():
    a = [3, 4, 6, 10, 11, 15]
    b = [1, 5, 8, 12, 14, 19]
    merged = merge_arrays(a, b)
    print(merged)

if __name__ == "__main__":
    main()
