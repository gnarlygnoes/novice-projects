#include <stdio.h>       
#include "raylib.h"
#include "aliens.h"
// #include "aliens.c"

void InitGame(void) 
{
    Texture2D game_texture = LoadTexture("img/SpaceInvaders.png");

    for (size_t i = 0; i < sizeof(stars) / sizeof(*stars); i++) {
        stars[i] = new_star_field();
    }

    // Initialise player
    player.in_rec.width = game_texture.width / 7;
    player.in_rec.height = game_texture.height / 5;
    player.in_rec.x = player.in_rec.width * 4;
    player.in_rec.y = 0;
}



int main(void)
{
    InitWindow(screen_width, screen_height, "raylib [core] example - basic window");

    InitGame();

    SetTargetFPS(60);
    
    double dt;
    while (!WindowShouldClose())
    {
        dt = GetFrameTime();

        BeginDrawing();

        DrawGame();
        handle_inputs();
        
        EndDrawing();
    }

    CloseWindow();

    return 0;
}

void DrawGame(void) 
{

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
        // DrawTexturePro();
}

void handle_inputs(void) 
{
    if (IsKeyDown(KEY_LEFT) && player.rec.x > 0) {
        player.rec.x -= 10;
    }
    if (IsKeyDown(KEY_RIGHT) && player.rec.x < screen_width - player_width) {
        player.rec.x += 10;
    }
}

Star new_star_field() 
{
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
