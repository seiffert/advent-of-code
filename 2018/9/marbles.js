let highscore = input => {
  let [_, numPlayers, numRounds] = input
    .match(new RegExp(/([0-9]+) players; last marble is worth ([0-9]+) points/))
    .map(n => parseInt(n));

  let game = new Game(numPlayers);
  game.play(numRounds);

  return game.scores[game.getLeader()];
};

class Game {
  currentMarble = 0;
  firstMarble = this.currentMarble;
  currentPlayer = 1;
  round = 1;

  next = { 0: 0 };
  prev = { 0: 0 };

  constructor(numPlayers) {
    this.scores = new Array(numPlayers + 1).fill().map(_ => 0);
    this.numPlayers = numPlayers;
  }

  play = rounds => {
    for (let i = 0; i < rounds; this.round++ && i++) {
      if (this.round % 23 == 0) {
        const seventhPrev = this.counterClockwise(this.currentMarble, 7);

        this.scores[this.currentPlayer] += this.round + seventhPrev;

        this.currentMarble = this.next[seventhPrev];
        this.next[this.prev[seventhPrev]] = this.next[seventhPrev];
        this.prev[this.next[seventhPrev]] = this.prev[seventhPrev];
      } else {
        this.next[this.round] = this.clockwise(this.currentMarble, 2);
        this.prev[this.clockwise(this.currentMarble, 2)] = this.round;
        this.prev[this.round] = this.clockwise(this.currentMarble, 1);
        this.next[this.clockwise(this.currentMarble, 1)] = this.round;
        this.currentMarble = this.round;
      }
      this.currentPlayer = (this.currentPlayer + 1) % this.numPlayers;
    }
  };

  getLeader = _ =>
    this.scores.reduce(
      (leader, score, player) =>
        this.scores[leader] < score ? player : leader,
      0
    );

  clockwise = (from, n) =>
    n > 0 ? this.clockwise(this.next[from], n - 1) : from;
  counterClockwise = (from, n) =>
    n > 0 ? this.counterClockwise(this.prev[from], n - 1) : from;
}

export { highscore };
