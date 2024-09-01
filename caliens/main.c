#include <stdio.h>       
#include "raylib.h"
#include "aliens.h"
// #include "aliens.c"

void InitGame(void) 
{
    

    for (size_t i = 0; i < sizeof(stars) / sizeof(*stars); i++) {
        stars[i] = new_star_field();
    }

    for (int i = 0; i < NUM_BULLETS; i++) {
        bullet[i].active = false;
        bullet[i].rec.height = 0;
        bullet[i].rec.width = 0;
        bullet[i].rec.x = 0;
        bullet[i].rec.y = 0;
        bullet->colour = ORANGE;
    }

    for (int i = 0; i < NUM_ENEMIES; i++) {
        enemies[i].rec.width = enemy_size;
        enemies[i].rec.height = enemy_size;
        enemies[i].rec.x = 10 + (1.3 * enemy_size * (i%10));
        enemies[i].rec.y = (1.3 * enemy_size * i/10);
        enemies[i].colour = GREEN;
    }
}

void handle_inputs(void) 
{
    if (IsKeyDown(KEY_LEFT) && player.rec.x > 0) {
        player.rec.x -= 10;
    }
    if (IsKeyDown(KEY_RIGHT) && player.rec.x < screen_width - player_width) {
        player.rec.x += 10;
    }
    if (IsKeyPressed(KEY_SPACE)) {
        shoot();
    }
}

void shoot(void) 
{
    for (int i = 0; i < NUM_BULLETS; i++) {
        if (!bullet[i].active) {
            bullet[i].active = true;
            bullet[i].rec.height = 20;
            bullet[i].rec.width = 5;
            bullet[i].rec.x = player.rec.x + player_width / 2;
            bullet[i].rec.y = player.rec.y;
            bullet[i].speed = 10;
            printf("New bullet creatored.\n");
            break;
        }
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

int main(void)
{
    InitWindow(screen_width, screen_height, "raylib [core] example - basic window");

    InitGame();

    SetTargetFPS(120);
    
    Texture2D game_texture = LoadTexture("img/SpaceInvaders.png");

    double dt;
    while (!WindowShouldClose())
    {
        dt = GetFrameTime();

        BeginDrawing();

        UpdateGame();
        DrawGame(game_texture);
        
        
        EndDrawing();
    }

    CloseWindow();

    return 0;
}

void UpdateGame(void)
{
    handle_inputs();
    for (int i = 0; i < NUM_BULLETS; i++) {
        if (bullet[i].active) {
            bullet[i].rec.y -= bullet[i].speed;
        }
        if (bullet[i].rec.y + bullet[i].rec.height < 0) {
            bullet[i].active = false;
        }
    }
}

void DrawGame(Texture2D tex) 
{
    player.in_rec.width = tex.width / 7;
    player.in_rec.height = tex.height / 5;
    player.in_rec.x = player.in_rec.width * 4;
    player.in_rec.y = 0;

    Vector2 origin = {0,0};

        ClearBackground(BLACK);
        DrawFPS(12, 36);

        for(size_t i = 0; i < sizeof(stars) / sizeof(*stars); i++) {
            DrawRectangle(stars[i].x, 
            stars[i].y, 
            stars[i].w, 
            stars[i].h,
            stars[i].colour);   
        }

        DrawTexturePro(tex, player.in_rec, player.rec, origin, 0, ORANGE);
        
        for (int i = 0; i < NUM_BULLETS; i++) {
            if (bullet[i].active) {
                DrawRectangleRec(bullet[i].rec, bullet[i].colour);
            }
        }

        for (int i = 0; i < NUM_ENEMIES; i++) {
            DrawRectangleRec(enemies[i].rec, enemies[i].colour);
        }
}
