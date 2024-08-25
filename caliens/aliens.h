#include "raylib.h"
#define NUM_STARS 900

static const int screen_width = 1920;
static const int screen_height = 1080;

static const int player_width = 80;
static const int player_height = 80;

typedef struct Star 
{
    int x, y;
    float w, h;
    Color colour;
} Star;

typedef struct Player 
{
    Texture player_texture;
    Rectangle in_rec;
    Rectangle rec;
    Vector2 pos;
    Color colour;
} Player;

// Initialise player
static Player player = {
    // .in_rec.height = 
    .rec.width = player_width,
    .rec.height = player_height,
    .rec.x = (screen_width / 2) - (player_width / 2),
    .rec.y = screen_height - player_height,
    .colour = RED,
    .pos.x = 0,
    .pos.y = 0,
};

void InitGame(void);
Star new_star_field();
void DrawGame(void);
// void UpdateGame(void);
void handle_inputs(void);

struct Star stars[NUM_STARS];