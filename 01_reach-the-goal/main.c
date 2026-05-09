#include <stdbool.h>

#include "include/raylib.h"

int main() {
    const int winWidth = 960;
    const int winHeight = 600;

    InitWindow(winWidth, winHeight, "Reach The Goal");
    SetTargetFPS(60);

    int x = 50;
    int y = 250;

    int speed = 5;

    int goal_x = 900;
    int goal_y = 250;

    bool gameOver = false;

    while (!WindowShouldClose()) {
        if (!gameOver) {
            if (IsKeyDown(KEY_A) || IsKeyDown(KEY_LEFT)) x -= speed;
            if (IsKeyDown(KEY_D) || IsKeyDown(KEY_RIGHT)) x += speed;
            if (IsKeyDown(KEY_W) || IsKeyDown(KEY_UP)) y -= speed;
            if (IsKeyDown(KEY_S) || IsKeyDown(KEY_DOWN)) y += speed;

            /* collision logic */
            int dx = (goal_x + 25) - (x + 25);
            int dy = (goal_y + 25) - (y + 25);

            if ((dx * dx + dy * dy) < 625) {
                gameOver = true;
            }

        } else {
            if (IsKeyPressed(KEY_ENTER)) break;
        }

        BeginDrawing();
        ClearBackground(DARKGRAY);

        DrawRectangle(goal_x, goal_y, 50, 50, YELLOW);
        DrawRectangle(x, y, 50, 50, RED);

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
