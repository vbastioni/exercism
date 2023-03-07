#include "armstrong_numbers.h"

bool is_armstrong_number(int candidate) {
    if (candidate < 0) return false;
    int count = log10(candidate) + 1, sum = 0;
    for (int loc = candidate; loc != 0; loc /= 10)
        sum += pow(loc % 10, count);
    return sum == candidate;
}
