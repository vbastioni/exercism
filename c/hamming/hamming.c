#include "hamming.h"

int compute(const char *lhs, const char *rhs) {
    int errors = 0, i = 0;
    for (; lhs[i] != 0 && rhs[i] != 0; i++)
        errors += (lhs[i] != rhs[i]);
    return (lhs[i] == rhs[i]) ? errors : -1;
}
