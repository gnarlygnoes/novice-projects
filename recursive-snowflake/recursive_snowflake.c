#include <stdio.h>
#include "raylib.h"
#include "raymath.h"

#define SCREEN_WIDTH 1920
#define SCREEN_HEIGHT 1080
#define NUM_BRANCHES 7

int draw_snowflakes(Vector2 centre, int branches, float length, float thickness, int levels);

int main(void)
{
    InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "raylib [core] example - basic window");

    Vector2 centre = {
        .x = SCREEN_WIDTH / 2,
        .y = SCREEN_HEIGHT / 2,
    };

    // printf("%i\n", BRANCH_ANGLE);

    while (!WindowShouldClose())
    {
        BeginDrawing();
            ClearBackground(BLACK);
            draw_snowflakes(centre, NUM_BRANCHES, 200.0, 10.0, 4);
        EndDrawing();
    }

    CloseWindow();

    return 0;
}

int draw_snowflakes(Vector2 centre, int branches, float length, float thickness, int levels)
{
    float branch_angle = 2 * PI / branches;
    Color colour = WHITE;
    if (levels >= 4) {
        colour = RAYWHITE;
    } else if (levels == 3)
    {
        colour = RED;
    } else if (levels == 2) {
        colour = YELLOW;
    } else {
        colour = ORANGE;
    }

    if (levels > 0) {

        for (size_t i = 0; i < branches; i++) {
            Vector2 line = {
                .x = centre.x + cos(branch_angle * i) * length,
                .y = centre.y + sin(branch_angle * i) * length,
            };
            DrawLineEx(centre, line, thickness, colour);
            // levels--;
            draw_snowflakes(line, NUM_BRANCHES, length/1.5, thickness/2.0, levels-1);
        }
    }
    DrawFPS(20, 30);

    return 0;
}
