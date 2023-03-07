#include "resistor_color.h"

int color_code(resistor_band_t color)
{
    return (int)color;
}

const resistor_band_t *colors()
{
    static const resistor_band_t colors[10] = {
        BLACK, BROWN, RED, ORANGE, YELLOW, GREEN, BLUE, VIOLET, GREY, WHITE
    };
    return colors;
}
