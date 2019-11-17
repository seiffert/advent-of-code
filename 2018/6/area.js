import "core-js/fn/array/flatten";

function largestArea(input) {
  return initGrid(input)
    .calculateDistances()
    .calculateAreas()
    .removeInfiniteAreas()
    .countLargestAreaCells();
}

function safeRegion(input, maxDistSum) {
  return initGrid(input)
    .calculateDistances()
    .countSafeCells(maxDistSum);
}

function initGrid(input) {
  return input.split`\n`.reduce((g, line, i) => {
    const [x, y] = line.split`, `.map(n => parseInt(n));
    return g.addLocation(String.fromCharCode(65 + i), x, y);
  }, new Grid());
}

class Grid {
  map = [];
  locations = {};
  maxX = -1;
  maxY = -1;

  addLocation = (loc, x, y) => {
    this.scaleMap(x, y);
    this.map[x][y].loc = loc;
    this.locations[loc] = [x, y];

    return this;
  };

  scaleMap = (x, y) => {
    for (let x2 = this.maxX + 1; x2 <= x; x2++) {
      this.map[x2] = Array(this.maxY == -1 ? 0 : this.maxY + 1)
        .fill()
        .map(() => new Cell());
    }
    for (let y2 = this.maxY + 1; y2 <= y; y2++) {
      this.map.forEach(row => {
        row.push(new Cell());
      });
    }

    this.maxX = Math.max(this.maxX, x);
    this.maxY = Math.max(this.maxY, y);
  };

  calculateAreas = () => {
    for (let x = 0; x <= this.maxX; x++) {
      for (let y = 0; y <= this.maxY; y++) {
        this.calculateArea(x, y);
      }
    }
    return this;
  };

  calculateArea = (x, y) => {
    const cell = this.map[x][y];
    if (cell.loc) {
      cell.markArea(this.map[x][y].loc);
      return;
    }

    cell.distances.sort(([locA, distA], [locB, distB]) => distA - distB);
    if (cell.distances[0][1] != cell.distances[1][1]) {
      cell.markArea(cell.distances[0][0]);
    }
  };

  removeInfiniteAreas = () => {
    let infiAreas = [];
    this.map.map((col, x) => {
      col.map((cell, y) => {
        if (x == 0 || y == 0 || x == this.maxX || y == this.maxY) {
          infiAreas.push(cell.area);
        }
      });
    });
    infiAreas = infiAreas
      .filter((x, i) => infiAreas.lastIndexOf(x) == i)
      .sort();

    this.map.forEach(col => {
      col.forEach(cell => {
        if (infiAreas.indexOf(cell.area) != -1) {
          cell.clear();
        }
      });
    });
    return this;
  };

  countLargestAreaCells = () => {
    return Object.values(
      this.map.flatten().reduce((acc, cell) => {
        if (!cell.area) return acc;

        if (!acc[cell.area]) {
          acc[cell.area] = 0;
        }
        acc[cell.area]++;
        return acc;
      }, {})
    ).sort((a, b) => b - a)[0];
  };

  calculateDistances = () => {
    this.map.forEach((col, x) =>
      col.forEach((cell, y) => {
        cell.distances = this.calculateDistancesForCell(x, y);
      })
    );
    return this;
  };

  calculateDistancesForCell = (x, y) => {
    return Object.entries(this.locations).map(([loc, [locX, locY]]) => {
      return [loc, Math.abs(x - locX) + Math.abs(y - locY)];
    });
  };

  countSafeCells = max => {
    return this.map
      .flatten()
      .map(cell => cell.sumDistances())
      .filter(x => x < max).length;
  };

  print() {
    const rows = [];
    for (let y = 0; y <= this.maxY; y++) {
      rows.push(this.map.map(col => col[y].toString()).join``);
    }
    console.log(rows.join`\n`);
    return this;
  }
}

class Cell {
  loc = "";
  area = "";
  distances = [];
  markArea(loc) {
    this.area = loc;
  }
  clear() {
    this.loc = "";
    this.area = "";
  }
  sumDistances = () => {
    return this.distances.reduce((acc, cur) => acc + cur[1], 0);
  };
  toString() {
    if (this.loc) {
      return this.loc;
    }
    if (this.area) {
      return this.area.toLowerCase();
    }
    return ".";
  }
}

export { largestArea, safeRegion };
