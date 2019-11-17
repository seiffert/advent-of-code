class Grid {
  squaresBySize = { 1: [] };
  constructor(sn) {
    this.cells = new Array(300).fill().map((_, y) =>
      new Array(300).fill().map((_, x) => {
        let rackId = 10 + x;
        let powerLevel = rackId * (rackId * y + sn);
        powerLevel = ((powerLevel % 1000) - (powerLevel % 100)) / 100 - 5;

        const s = new Square(x, y, 1, powerLevel);
        this.squaresBySize[1].push(s);
        return s;
      })
    );

    this.calculateSquares();
  }

  largestSquare = () => {
    return Object.values(this.squaresBySize).reduce(
      (largest, squares) =>
        squares.reduce(
          (largest, curSquare) =>
            curSquare.powerLevel > largest.powerLevel ? curSquare : largest,
          largest
        ),
      new Square(0, 0, 0, 0)
    );
  };

  largest3x3Square = () =>
    this.squaresBySize[3].reduce((largest, cur) => {
      return cur.powerLevel > largest.powerLevel ? cur : largest;
    });

  calculateSquares = () => {
    for (let n = 2; n <= 300; n++) {
      this.calculateSquaresOfSize(n);
    }
  };

  calculateSquaresOfSize = size => {
    this.squaresBySize[size] = [];

    this.squaresBySize[size - 1].forEach(square => {
      if (square.x + size > 300 || square.y + size > 300) {
        return;
      }
      this.squaresBySize[size].push(
        new Square(
          square.x,
          square.y,
          size,
          square.powerLevel +
            this.cells[square.y + size - 1]
              .slice(square.x, square.x + size - 1)
              .reduce((acc, cur) => cur.powerLevel + acc, 0) +
            this.cells
              .slice(square.y, square.y + size)
              .reduce(
                (acc, cur) => cur[square.x + size - 1].powerLevel + acc,
                0
              )
        )
      );
    });
  };
}

class Square {
  constructor(x, y, size, powerLevel) {
    this.x = x;
    this.y = y;
    this.size = size;
    this.powerLevel = powerLevel;
  }
}

export { Grid };
