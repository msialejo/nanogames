#include <stdbool.h>

#include "include/raylib.h"

typedef struct Entity {
    float x;
    float y;
    int size;
    float speed;
    Color color;
} Entity;

Entity make_player(int winWidth, int winHeight);
Entity make_goal(int winWidth, int winHeight);
void handle_player_input(Entity* p);

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

            // collision logic
            float dx = (goal.x + goal.size / 2.0f) - (player.x + player.size / 2.0f);
            float dy = (goal.y + goal.size / 2.0f) - (player.y + player.size / 2.0f);

            if ((dx * dx + dy * dy) < (player.size * player.size / 4.0f)) {
                gameOver = true;
            }

        } else {
            if (IsKeyPressed(KEY_ENTER)) break;
        }

        BeginDrawing();
        ClearBackground(DARKGRAY);

        DrawRectangle(goal.x, goal.y, goal.size, goal.size, goal.color);
        DrawRectangle(player.x, player.y, player.size, player.size, player.color);

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
        .size = 50,
        .speed = 5.0f,
        .color = RED,
    };
}

Entity make_goal(int winWidth, int winHeight) {
    return (Entity){
        .x = winWidth - 100,
        .y = winHeight / 2 - 25,
        .size = 50,
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
