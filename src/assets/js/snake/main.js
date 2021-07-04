let snakeCanvas = document.getElementById("snake_canvas");
let paintingCtx = snakeCanvas.getContext("2d");
let canvasWidth = snakeCanvas.width; // 600
let canvasHeight = snakeCanvas.height; // 600

/*Создание ячеек.*/
let cellSize = 20;
let widthInCells = canvasWidth / cellSize; // 30 ячеек в ширину.
let heightInBlocks = canvasHeight / cellSize; // 30 ячеек в высоту.

/*START--------------------------------------------------------------------------------------------*/

let gameScore = 0; // Съеденных яблок.
let snake = new Snake();
let apple = new Apple();
let directions = {37: "left", 65: "left", 38: "up", 87: "up", 39: "right", 68: "right", 40: "down", 83: "down"};
jQuery("body").keydown(function (event) {
    let newDirection = directions[event.keyCode];
    if(newDirection !== undefined) {
        snake.setDirection(newDirection);
    }
});

let intervalId = setInterval(function () {
    paintingCtx.clearRect(0, 0, canvasWidth, canvasHeight);
    drawBorder();
    drawGameText("20px Monospaced", "Black", "top", "left", "Game score: " + gameScore, cellSize, canvasHeight - cellSize);
    snake.snakeMove();
    snake.drawSnake();
    apple.drawApple();
}, 50);

/*FUNCTIONS------------------------------------------------------------------------------------------*/

/*Рисование границ игорового поля.*/
function drawBorder() {
    paintingCtx.fillStyle = "DarkOrange";
    paintingCtx.fillRect(0, 0, canvasWidth, cellSize); //Верхняя граница игрового поля.
    paintingCtx.fillRect(0, canvasHeight - cellSize, canvasWidth, cellSize); // Нижняя граница.
    paintingCtx.fillStyle = "Orange";
    paintingCtx.fillRect(0, 0, cellSize, canvasHeight); // Левая граница.
    paintingCtx.fillRect(canvasWidth - cellSize, 0, cellSize, canvasHeight); // Правая граница.
}

/*Рисование на холсте ТЕКСТА.*/
function drawGameText(font, fillStyle, textBaseLine, textAlign, text, x, y) {
    paintingCtx.font = font; // Monospaced.
    paintingCtx.fillStyle = fillStyle; // Black.
    paintingCtx.textBaseline = textBaseLine;
    paintingCtx.textAlign = textAlign; //
    paintingCtx.fillText(text, x, y);
}

/*В случае окончания игры.*/
function gameOver() {
    clearInterval(intervalId);
    apple.deleteApple();
    drawGameText("80px Monospaced", "Black", "middle", "center", "Game over", canvasWidth / 2, canvasHeight / 2);
}

