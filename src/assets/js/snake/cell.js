let Cell = function (column, row) {
    this.column = column;
    this.row = row;
};

Cell.prototype.drawHead = function(color, direction) {
    let x = this.column * cellSize + (cellSize / 2);
    let y = this.row * cellSize + (cellSize / 2);
    paintingCtx.beginPath();
    paintingCtx.arc(x, y, cellSize / 1.5, 0, Math.PI * 2, false);
    paintingCtx.fillStyle = color;
    paintingCtx.fill();
    drawSnakeFace(x, y, direction);
};

Cell.prototype.drawBodySquare = function (color) {
    let x = this.column * cellSize;
    let y = this.row * cellSize;
    paintingCtx.fillStyle = color;
    paintingCtx.fillRect(x, y, cellSize, cellSize);
    paintingCtx.strokeRect(x, y, cellSize, cellSize);

    paintingCtx.beginPath();
    let centerX = this.column * cellSize + cellSize / 2;
    let centerY = this.row * cellSize + cellSize / 2;
    paintingCtx.arc(centerX, centerY, cellSize / 4, 0, Math.PI * 2, false);
    paintingCtx.fillStyle = "Black";
    paintingCtx.fill();
};

Cell.prototype.drawApple = function (color) {
    let centerX = this.column * cellSize + cellSize / 2;
    let centerY = this.row * cellSize + cellSize / 2;

    /*Яблоко.*/
    paintingCtx.fillStyle = color;
    paintingCtx.beginPath();
    paintingCtx.arc(centerX, centerY, cellSize / 2, 0, Math.PI * 2, false);
    paintingCtx.fill();
    paintingCtx.stroke();

    /*Хвостик.*/
    paintingCtx.beginPath();
    paintingCtx.moveTo(this.column * cellSize + cellSize / 2, this.row * cellSize);
    paintingCtx.lineTo(this.column * cellSize + cellSize / 2, this.row * cellSize - cellSize / 2);
    paintingCtx.stroke();

    /*Л.листик.*/
    paintingCtx.fillStyle = "LightGreen";
    let ellipseX = this.column * cellSize + cellSize / 4;
    let ellipseY = this.row * cellSize - cellSize / 4;
    paintingCtx.beginPath();
    paintingCtx.ellipse(/*x*/ellipseX, /*y*/ellipseY, /*толщина*/3, /*длинна*/7, 45 * Math.PI/180, 0, 2 * Math.PI);
    paintingCtx.fill();

    /*П.листик.*/
    ellipseX = this.column * cellSize + ((cellSize / 4) * 3);
    ellipseY = this.row * cellSize - cellSize / 4;
    paintingCtx.beginPath();
    paintingCtx.ellipse(/*x*/ellipseX, /*y*/ellipseY, /*толщина*/7, /*длинна*/3, 45 * Math.PI/180, 0, 2 * Math.PI);
    paintingCtx.fill();
};

Cell.prototype.isEqualCells = function (otherBlock) {
    return this.column === otherBlock.column && this.row === otherBlock.row;
};

/*FUNCTIONS-----------------------------------------------------------------------------------------*/

function drawSnakeFace(x, y, direction) {
    paintingCtx.fillStyle = "Black";
    if (direction === "down") {
        /*Рот.*/
        paintingCtx.beginPath();
        paintingCtx.arc(x, y, cellSize / 2, 0, Math.PI, false);
        paintingCtx.stroke();
        /*Л.глаз.*/
        paintingCtx.beginPath();
        paintingCtx.arc(x - 4, y - 3, cellSize / 10, 0, Math.PI * 2, false);
        paintingCtx.fill();
        /*П.глаз.*/
        paintingCtx.beginPath();
        paintingCtx.arc(x + 4, y - 3, cellSize / 10, 0, Math.PI * 2, false);
        paintingCtx.fill();
    } else if (direction === "up"){
        /*Рот.*/
        paintingCtx.beginPath();
        paintingCtx.arc(x, y, cellSize / 2, 0, Math.PI, true);
        paintingCtx.stroke();
        /*Л.глаз.*/
        paintingCtx.beginPath();
        paintingCtx.arc(x - 4, y + 3, cellSize / 10, 0, Math.PI * 2, false);
        paintingCtx.fill();
        /*П.глаз.*/
        paintingCtx.beginPath();
        paintingCtx.arc(x + 4, y + 3, cellSize / 10, 0, Math.PI * 2, false);
        paintingCtx.fill();
    } else if(direction === "right") {
        /*Рот.*/
        paintingCtx.beginPath();
        paintingCtx.arc(x, y, cellSize / 2, Math.PI * 1.5, Math.PI * 0.5, false);
        paintingCtx.stroke();
        /*Л.глаз.*/
        paintingCtx.beginPath();
        paintingCtx.arc(x - 3, y - 4, cellSize / 10, 0, Math.PI * 2, false);
        paintingCtx.fill();
        /*П.глаз.*/
        paintingCtx.beginPath();
        paintingCtx.arc(x - 3, y + 4, cellSize / 10, 0, Math.PI * 2, false);
        paintingCtx.fill();
    } else {
        /*Рот.*/
        paintingCtx.beginPath();
        paintingCtx.arc(x, y, cellSize / 2, Math.PI * 0.5, Math.PI * 1.5, false);
        paintingCtx.stroke();
        /*Л.глаз.*/
        paintingCtx.beginPath();
        paintingCtx.arc(x + 3, y - 4, cellSize / 10, 0, Math.PI * 2, false);
        paintingCtx.fill();
        /*П.глаз.*/
        paintingCtx.beginPath();
        paintingCtx.arc(x + 3, y + 4, cellSize / 10, 0, Math.PI * 2, false);
        paintingCtx.fill();
    }
}