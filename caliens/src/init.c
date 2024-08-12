#include <stdio.h>
#include "raylib.h"
#include "aliens.h"

void InitGame(void) {
    // Initialise Player
    player.rec.width = playerWidth;
    player.rec.height = playerHeight;
    player.rec.x = (screenWidth / 2) + (playerWidth / 2);
    player.rec.y = screenHeight - playerHeight;
    player.colour = RED;

    for (size_t i = 0; i < sizeof(stars) / sizeof(*stars); i++) {
        stars[i] = new_star_field();
    }
}