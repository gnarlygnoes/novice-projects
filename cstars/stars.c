#include <stdio.h>       
#include "raylib.h"

static const int screenWidth = 1920;
static const int screenHeight = 1080;

static const int numStars = 90000;

typedef struct Star{
    int x;
    int y;
    float w, h;
    Color colour;
} Star;

typedef struct Player{
    int x, y, w, h;
    Color colour;
} Player;

Star new_star_field();

int main(void)
{
    InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window");

    struct Star stars[numStars];

    // int i = 0;
    for (size_t i = 0; i < sizeof(stars) / sizeof(*stars); i++) {
        stars[i] = new_star_field();
    }
    while (!WindowShouldClose())
    {
        BeginDrawing();

            ClearBackground(BLACK);
            DrawFPS(12, 36);

            for(size_t i = 0; i < sizeof(stars) / sizeof(*stars); i++) {
                DrawRectangle(stars[i].x, 
                stars[i].y, 
                stars[i].w, 
                stars[i].h,
                stars[i].colour);   
            }
        EndDrawing();
    }

    CloseWindow();

    return 0;
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
