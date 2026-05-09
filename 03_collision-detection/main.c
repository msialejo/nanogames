#include <stdbool.h>

#include "include/raylib.h"

typedef struct Entity {
    float x;
    float y;
    int width;
    int height;
    float speed;
    Color color;
} Entity;

Entity make_player(int winWidth, int winHeight);
Entity make_goal(int winWidth, int winHeight);
void handle_player_input(Entity* p);
bool check_collision_AABB(Entity* a, Entity* b);

int main() {
    const int winWidth = 960;
    const int winHeight = 600;

    InitWindow(winWidth, winHeight, "Reach The Goal");
    SetTargetFPS(60);

    Entity player = make_player(winWidth, winHeight);
    Entity goal = make_goal(winWidth, winHeight);
    bool gameOver = false;

    while (!WindowShouldClose()) {
        if (!gameOver) {
            handle_player_input(&player);

            if (check_collision_AABB(&player, &goal)) {
                gameOver = true;
            }

        } else {
            if (IsKeyPressed(KEY_ENTER)) break;
        }

        BeginDrawing();
        ClearBackground(DARKGRAY);

        DrawRectangle(goal.x, goal.y, goal.width, goal.height, goal.color);
        DrawRectangle(player.x, player.y, player.width, player.height, player.color);

        if (gameOver) {
            DrawRectangle(0, 0, winWidth, winHeight, Fade(BLACK, 0.4f));
            DrawText("YOU WIN!", 400, 250, 40, GREEN);
            DrawText("Press ENTER to Exit", 365, 310, 20, RAYWHITE);
        }

        EndDrawing();
    }

    CloseWindow();
    return 0;
}

Entity make_player(int winWidth, int winHeight) {
    return (Entity){
        .x = 50,
        .y = winHeight / 2 - 25,
        .width = 50,
        .height = 50,
        .speed = 5.0f,
        .color = RED,
    };
}

Entity make_goal(int winWidth, int winHeight) {
    return (Entity){
        .x = winWidth - 100,
        .y = winHeight / 2 - 25,
        .width = 50,
        .height = 50,
        .speed = 0,
        .color = YELLOW,
    };
}

void handle_player_input(Entity* p) {
    if (IsKeyDown(KEY_A) || IsKeyDown(KEY_LEFT)) p->x -= p->speed;
    if (IsKeyDown(KEY_D) || IsKeyDown(KEY_RIGHT)) p->x += p->speed;
    if (IsKeyDown(KEY_W) || IsKeyDown(KEY_UP)) p->y -= p->speed;
    if (IsKeyDown(KEY_S) || IsKeyDown(KEY_DOWN)) p->y += p->speed;
}

bool check_collision_AABB(Entity* a, Entity* b) {
    return (a->x < b->x + b->width &&   // player left < goal right
            a->x + a->width > b->x &&   // player right > goal left
            a->y < b->y + b->height &&  // player top < goal bottom
            a->y + a->height > b->y);   // player bottom > goal top
}
