#include <stdio.h>       
#include "raylib.h"
#include "aliens.h"
// #include "aliens.c"

int main(void)
{
    InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window");

    InitGame();

    while (!WindowShouldClose())
    {
        BeginDrawing();
        DrawGame();
        HandleInputs();
        printf("Player x value: %f\n", player.rec.width);
        EndDrawing();
    }

    CloseWindow();

    return 0;
}

void DrawGame(void) {

        ClearBackground(BLACK);
        DrawFPS(12, 36);

        for(size_t i = 0; i < sizeof(stars) / sizeof(*stars); i++) {
            DrawRectangle(stars[i].x, 
            stars[i].y, 
            stars[i].w, 
            stars[i].h,
            stars[i].colour);   
        }

        DrawRectangleRec(player.rec, player.colour);
}

Star new_star_field() {
    int r = GetRandomValue(100, 255);
    int g = GetRandomValue(100, 255);
    int b = GetRandomValue(100, 255);

    Color c = {
        .r=r, .g=g, .b=b, .a=255
    };
    
    Star star = {
        .x =  GetRandomValue(0, GetScreenWidth()),
        .y =  GetRandomValue(0, GetScreenHeight()),
        .w = GetRandomValue(1, 5) / 1.3,
        .h = star.w,
        .colour = c,
    };

    return star;
}
