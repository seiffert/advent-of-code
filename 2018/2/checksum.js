import "core-js/fn/array/flat-map";

function checksum(input) {
  let twice = 0;
  let thrice = 0;

  input.map(word => {
    var letterCounts = {};
    for (let i = 0; i < word.length; i++) {
      const letter = word[i];
      letterCounts[letter] = letterCounts[letter]
        ? letterCounts[letter] + 1
        : 1;
    }

    if (Object.values(letterCounts).filter(c => c == 2).length) {
      twice++;
    }
    if (Object.values(letterCounts).filter(c => c == 3).length) {
      thrice++;
    }
  });
  return twice * thrice;
}

function closeButDifferent(words) {
  function common(a, b) {
    let c = "";
    for (let i = 0; i < Math.min(a.length, b.length); i++) {
      if (a[i] == b[i]) {
        c += a[i];
      }
    }
    return c;
  }
  const unique = (value, index, self) => self.indexOf(value) === index;

  return words
    .flatMap(word =>
      words
        .map(word2 => common(word, word2))
        .filter(c => word.length - c.length == 1)
    )
    .find(unique);
}

export { checksum, closeButDifferent };
