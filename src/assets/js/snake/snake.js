let Snake = function () {
    this.snakeBodyArray = [new Cell(3, 1), new Cell(2, 1), new Cell(1, 1)];
    this.direction = "right";
    this.newDirection = "right";
};

Snake.prototype.snakeMove = function () {
    let snakeHead = this.snakeBodyArray[0];
    let newSnakeHead;

    this.direction = this.newDirection;

    if (this.direction === "right") {
        newSnakeHead = new Cell(snakeHead.column + 1, snakeHead.row);
    } else if (this.direction === "down") {
        newSnakeHead = new Cell(snakeHead.column, snakeHead.row + 1);
    } else if (this.direction === "left") {
        newSnakeHead = new Cell(snakeHead.column - 1, snakeHead.row);
    } else if (this.direction === "up") {
        newSnakeHead = new Cell(snakeHead.column, snakeHead.row - 1);
    }

    if (checkCollision(newSnakeHead)) {
        gameOver();
        return;
    }

    /*Если добавленная голова не столкнулась со стеной, то добавляю ячейку к массиву.*/
    this.snakeBodyArray.unshift(newSnakeHead);

    /*Если позиция головы и яблока в одной клетке, то новая позиция для яблока, хвост без изменений,
    * в противном случае в массиве змейки удаляю последний элемент.*/
    if (newSnakeHead.isEqualCells(apple.fillingCell)) {
        gameScore++;

        let b = true;
        label1:
            while (b) {
                apple.moveApple();
                for (let i = 0; i < this.snakeBodyArray.length; i++) {
                    if (apple.fillingCell.isEqualCells(this.snakeBodyArray[i])) {
                        continue label1;
                    }
                }
                b = false;
            }
    } else {
        this.snakeBodyArray.pop();
    }
}
;

Snake.prototype.drawSnake = function () {
    for (let i = 0; i < this.snakeBodyArray.length; i++) {
        if (i !== 0) {
            i % 2 === 0 ? this.snakeBodyArray[i].drawBodySquare("LightGreen") : this.snakeBodyArray[i].drawBodySquare("Yellow");
        } else {
            this.snakeBodyArray[i].drawHead("Green", this.direction);
        }
    }
};

Snake.prototype.setDirection = function (newDirection) {
    if (this.direction === "up" && newDirection === "down") return;
    else if (this.direction === "right" && newDirection === "left") return;
    else if (this.direction === "down" && newDirection === "up") return;
    else if (this.direction === "left" && newDirection === "right") return;

    this.newDirection = newDirection;
};

/*FUNCTIONS------------------------------------------------------------------*/

function checkCollision(newSnakeHead) {
    let leftWallCollision = (newSnakeHead.column === 0);
    let topWallCollision = (newSnakeHead.row === 0);
    let rightWallCollision = (newSnakeHead.column === widthInCells - 1);
    let bottomWallCollision = (newSnakeHead.row === heightInBlocks - 1);
    let wallCollision = leftWallCollision || topWallCollision || rightWallCollision || bottomWallCollision;

    let selfCollision = false;
    snake.snakeBodyArray.forEach(function (fillingCell) {
        if (newSnakeHead.isEqualCells(fillingCell)) {
            selfCollision = true;
        }
    });

    return wallCollision || selfCollision;
}