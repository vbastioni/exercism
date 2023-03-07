#include "isogram.h"

bool is_isogram(const char phrase[]) {
    if (phrase == (void*)0)
        return false;
    long found = 0;
    for (int i = 0; phrase[i] != '\0'; i++) {
        char c = phrase[i];
        if (c == ' ' || c == '-')
            continue;
        c = c > 'A' ? (c - 'A') : (c - 'a');
        if ((found & (1 << c)) != 0)
            return false;
        found |= (1 << c);
    }
    return true;
}
