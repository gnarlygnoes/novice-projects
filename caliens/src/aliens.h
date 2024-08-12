#include "raylib.h"
#define NUM_STARS 900

static const int screenWidth = 1920;
static const int screenHeight = 1080;

static const int playerWidth = 60;
static const int playerHeight = 80;

typedef struct Star 
{
    int x, y;
    float w, h;
    Color colour;
} Star;

typedef struct Player 
{
    Rectangle rec;
    Color colour;
} Player;

static Player player = { 0 };

void InitGame(void);
Star new_star_field();
void DrawGame(void);
// void UpdateGame(void);
void HandleInputs(void);

struct Star stars[NUM_STARS];