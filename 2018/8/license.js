let license = input => constructTree(input).sumMetadata();
let rootValue = input => constructTree(input).value();
let constructTree = input => new Tree(input.split` `.map(n => parseInt(n)));

class Tree {
  children = [];
  metadata = [];

  constructor(input) {
    let numChildren = input.shift();
    let numMetadata = input.shift();

    for (let i = 0; i < numChildren; i++) {
      this.children.push(new Tree(input));
    }

    for (let i = 0; i < numMetadata; i++) {
      this.metadata.push(input.shift());
    }
  }

  sumMetadata = () =>
    this.children.reduce(
      (acc, cur) => acc + cur.sumMetadata(),
      this.metadata.reduce((acc, cur) => acc + cur, 0)
    );

  value = () =>
    this.children.length == 0
      ? this.sumMetadata()
      : this.metadata.reduce((acc, cur) => {
          if (this.children[cur - 1]) {
            return acc + this.children[cur - 1].value();
          }
          return acc;
        }, 0);
}

export { license, rootValue };
