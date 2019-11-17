function processPolymer(p) {
  return p.split``.reduce((acc, cur) => {
    const prev = acc.pop();
    if (typeof prev != "undefined" && react(prev, cur)) {
      return acc;
    }
    return acc.concat([prev, cur]);
  }, []).join``;
}

function improvedProcessPolymer(p) {
  return "abcdefghijklmnopqrstuvwxyz".split``.reduce((acc, cur) => {
    const processed = processPolymer(p.replace(new RegExp(`${cur}`, "ig"), ""));
    return processed.length < acc.length ? processed : acc;
  }, p);
}

function react(a, b) {
  return a != b && a.toLowerCase() == b.toLowerCase();
}

export { processPolymer, improvedProcessPolymer };
