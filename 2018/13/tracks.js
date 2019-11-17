import "core-js/fn/array/flatten";

const cartSymbolToTrack = {
  ">": "-",
  "<": "-",
  v: "|",
  "^": "|"
};
const leftTurns = {
  ">": "^",
  "<": "v",
  v: ">",
  "^": "<"
};
const rightTurns = {
  ">": "v",
  "<": "^",
  v: "<",
  "^": ">"
};

class Tracks {
  constructor(input) {
    this.tracks = input.split`\n`.map((line, y) =>
      line.split``.map((symbol, x) => new Track(x, y, symbol))
    );
    this.carts = this.tracks
      .flatten()
      .map(t => t.cart)
      .filter(c => c);
  }

  firstCollision = () => {
    while (!this.findCollision()) {
      this.tick();
    }

    const c = this.findCollision();
    return `${c.x},${c.y}`;
  };

  lastCart = () => {
    while (this.carts.length > 1) {
      this.tick();
    }

    const c = this.carts[0];
    return `${c.x},${c.y}`;
  };

  findCollision = () =>
    this.tracks.reduce(
      (acc, row) => (acc ? acc : row.find(t => t.collision)),
      null
    );

  tick = () => {
    this.carts
      .sort((a, b) => {
        const diff = a.y - b.y;
        if (diff !== 0) {
          return diff;
        }
        return a.x - b.x;
      })
      .forEach(this.moveCart);
    this.carts = this.carts.filter(c => !c.broken);
  };

  moveCart = cart => {
    const track = this.tracks[cart.y][cart.x];
    if (track.symbol == "+") {
      cart.turn();
    }
    switch ([cart.symbol, [track.symbol]].join``) {
      case ">/":
      case "</":
      case "^\\":
      case "v\\":
        cart.turnLeft();
        break;
      case ">\\":
      case "<\\":
      case "^/":
      case "v/":
        cart.turnRight();
        break;
    }
    cart.move();

    track.cart = null;
    const newTrack = this.tracks[cart.y][cart.x];
    if (newTrack.cart && !newTrack.cart.broken) {
      newTrack.collision = true;

      newTrack.cart.broken = true;
      cart.broken = true;
      newTrack.cart = null;
      return;
    }
    newTrack.cart = cart;
  };
}

class Track {
  cart;
  collision = false;
  constructor(x, y, symbol) {
    this.x = x;
    this.y = y;

    switch (symbol) {
      case ">":
      case "<":
      case "v":
      case "^":
        this.cart = new Cart(x, y, symbol);
        symbol = cartSymbolToTrack[symbol];
        break;
    }
    this.symbol = symbol;
  }
}

class Cart {
  broken = false;
  turnCount = 0;

  constructor(x, y, symbol) {
    this.x = x;
    this.y = y;
    this.symbol = symbol;
  }
  turn = () => {
    this.turnCount++;
    if (this.turnCount % 3 == 1) {
      this.turnLeft();
    } else if (this.turnCount % 3 == 0) {
      this.turnRight();
    }
  };

  turnLeft = () => {
    this.symbol = leftTurns[this.symbol];
  };

  turnRight = () => {
    this.symbol = rightTurns[this.symbol];
  };

  move = () => {
    switch (this.symbol) {
      case ">":
        this.x++;
        break;
      case "<":
        this.x--;
        break;
      case "v":
        this.y++;
        break;
      case "^":
        this.y--;
        break;
    }
  };
}

export default Tracks;
