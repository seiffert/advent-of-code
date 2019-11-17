class Board {
  constructor(initialRecipes) {
    this.recipes = initialRecipes;
    this.recipesArr = initialRecipes.split``.map(n => parseInt(n));
    this.elves = [0, 1];
  }

  scoreAfter = n => {
    while (this.recipes.length < n + 10) {
      this.generate();
    }
    return this.recipes.substring(n, n + 10);
  };

  recipesBefore = seq => {
    let i;
    do {
      for (let j = 0; j < 100000; j++) this.generate();
      i = this.recipes.indexOf(seq);
    } while (i === -1);
    return i;
  };

  generate = () => {
    const recipe1 = this.recipesArr[this.elves[0]];
    const recipe2 = this.recipesArr[this.elves[1]];
    const newRecipes = `` + (recipe1 + recipe2);

    newRecipes.split``.forEach(r => this.recipesArr.push(parseInt(r)));
    this.recipes += newRecipes;

    this.elves = [
      (recipe1 + 1 + this.elves[0]) % this.recipesArr.length,
      (recipe2 + 1 + this.elves[1]) % this.recipesArr.length
    ];
  };
}

export default Board;
