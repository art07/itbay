let Apple = function () {
    this.fillingCell = undefined;
    this.moveApple();
};

Apple.prototype.moveApple = function () {
    let randomColumn = Math.floor(Math.random() * (widthInCells - 2)) + 1;
    let randomRow = Math.floor(Math.random() * (heightInBlocks - 2)) + 1;
    this.fillingCell = new Cell(randomColumn, randomRow);
};

Apple.prototype.drawApple = function () {
    this.fillingCell.drawApple("Red");
};

Apple.prototype.deleteApple = function () {
    this.fillingCell = undefined;
};

