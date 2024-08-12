#include <stdio.h>       
#include "raylib.h"
#include "aliens.h"

void HandleInputs(void)
{
    if (IsKeyDown(KEY_RIGHT)) {
        player.rec.x += 10;
    }
}
