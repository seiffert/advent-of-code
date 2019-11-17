class Sky {
  constructor(input) {
    this.stars = input.split`\n`.map(l => new Star(l));
  }

  moveStars = sec => {
    if (sec > 0) {
      this.stars.forEach(s => s.move());
      return this.moveStars(sec - 1);
    }
    return this;
  };

  toString = (maxHeight = 10) => {
    let [maxX, maxY, minX, minY] = this.stars.reduce(
      ([maxX, maxY, minX, minY], s) => [
        s.x > maxX ? s.x : maxX,
        s.y > maxY ? s.y : maxY,
        s.x < minX ? s.x : minX,
        s.y < minY ? s.y : minY
      ],
      [-1, -1, Number.MAX_SAFE_INTEGER, Number.MAX_SAFE_INTEGER]
    );
    if (maxY - minY > maxHeight) return "";

    let map = new Array(maxY - minY + 1)
      .fill()
      .map(_ => new Array(maxX - minX + 1).fill("."));

    this.stars.forEach(s => {
      map[s.y - minY][s.x - minX] = "#";
    });

    return map.map(row => row.join``).join`\n`;
  };
}

class Star {
  constructor(input) {
    let [_, x, y, vx, vy] = input.match(
      /position=<([- 0-9]+), ([- 0-9]+)> velocity=<([- 0-9]+), ([- 0-9]+)>/
    );
    this.x = parseInt(x);
    this.y = parseInt(y);
    this.vx = parseInt(vx);
    this.vy = parseInt(vy);
  }

  move = () => {
    this.x += this.vx;
    this.y += this.vy;
  };
}

export { Sky };
