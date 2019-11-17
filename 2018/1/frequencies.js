function calculateFrequency(s) {
  return eval(s.replace(/,/g, ``));
}

function findRepetition(s) {
  const changes = String(s)
    .replace(/ /g, "")
    .split(",");

  let freq = 0;
  let history = { "0": true };

  for (let i = 0; ; i = (i + 1) % changes.length) {
    freq += (changes[i][0] == "-" ? -1 : 1) * changes[i].substr(1);
    if (history[`${freq}`]) {
      return freq;
    }
    history[`${freq}`] = true;
  }
}

export { calculateFrequency, findRepetition };
