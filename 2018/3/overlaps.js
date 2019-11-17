import "core-js/fn/array/flat-map";

function overlaps(input) {
  return new Fabric(input.map(Square.parse)).countOverlaps();
}

function notOverlapping(input) {
  const squares = input.map(Square.parse);
  const f = new Fabric(squares);

  return squares.filter(s => !s.overlapping)[0].id;
}

class Fabric {
  area = {};

  constructor(squares) {
    squares.forEach(s => this.add(s));
  }
  add = square => {
    for (let x = square.left; x < square.left + square.width; x++) {
      for (let y = square.top; y < square.top + square.height; y++) {
        this.area[x] = this.area[x] || {};
        this.area[x][y] = this.area[x][y] || [];

        if (this.area[x][y].push(square) > 1) {
          this.area[x][y].forEach(s => (s.overlapping = true));
        }
      }
    }
    return this;
  };
  countOverlaps = () =>
    Object.values(this.area)
      .flatMap(Object.values)
      .filter(s => s.length > 1).length;
}

class Square {
  static parse(desc) {
    const [id, left, top, width, height] = desc
      .replace(/ /g, "")
      .split(/[#:@,x]/)
      .filter(Boolean)
      .map(n => parseInt(n));

    return new Square(id, left, top, width, height);
  }
  constructor(id, left, top, width, height) {
    this.id = id;
    this.left = left;
    this.top = top;
    this.width = width;
    this.height = height;
    this.overlapping = false;
  }
}

export { overlaps, notOverlapping };
