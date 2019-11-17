class Pots {
  rules = {};
  firstPot = 0;

  constructor(input) {
    const lines = input.split`\n`;
    this.pots = this.padString(
      lines.shift().match(/initial state: ([#.]+)$/)[1]
    );
    lines.shift();

    lines.forEach(line => {
      const [_, pattern, result] = line.match(/([#.]{5}) => ([#.])$/);
      this.rules[pattern] = result;
    });
  }

  proceed = gens => {
    for (let gen = 0; gen < gens; gen++) {
      const firstPotBefore = this.firstPot;
      const newPots = this.spread(this.pots);

      if (newPots == this.pots) {
        // repetition found, we just need to fix the firstPot pointer now
        this.firstPot += (this.firstPot - firstPotBefore) * (gens - gen - 1);
        break;
      }

      this.pots = newPots;
    }
    return this;
  };

  spread = pots => {
    let newGen = "";
    for (let i = 0; i < pots.length - 4; i++) {
      newGen += this.rules[pots.substring(i, i + 5)];
    }
    this.firstPot -= 2;
    newGen = this.padString(newGen);

    return newGen;
  };

  sum = () => {
    let curValue = -this.firstPot;

    let res = 0;
    for (let i = 0; i < this.pots.length; i++) {
      if (this.pots[i] == "#") {
        res += curValue;
      }
      curValue++;
    }
    return res;
  };

  padString = s => {
    return this.padStringLeft(this.padStringRight(s));
  };

  padStringLeft = s => {
    if (s.startsWith("....")) {
      return s;
    }
    this.firstPot++;
    return this.padStringLeft("." + s);
  };

  padStringRight = s => {
    if (s.endsWith("....")) {
      return s;
    }
    return this.padStringRight(s + ".");
  };
}

export default Pots;
