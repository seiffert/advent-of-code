function sortDependencies(input) {
  return initGraph(input).flat();
}

function processDependencies(input, workers, delay) {
  return initGraph(input).process(workers, delay);
}

function initGraph(input) {
  return input.split`\n`
    .map(line =>
      line
        .match(
          /Step ([a-zA-Z]+) must be finished before step ([a-zA-Z]+) can begin./
        )
        .slice(1, 3)
    )
    .reduce(
      (acc, [requirement, step]) => acc.addDependency(step, requirement),
      new DependencyGraph()
    );
}

class DependencyGraph {
  dependencies = new Map();
  length = 0;

  addDependency = (step, requirement) => {
    this.addIfNew(step);
    this.addIfNew(requirement);

    this.dependencies.get(step).push(requirement);
    return this;
  };

  addIfNew = step => {
    if (!this.dependencies.has(step)) {
      this.dependencies.set(step, []);
      this.length++;
    }
  };

  flat = () => {
    const result = [];
    while (this.length != 0) {
      const step = this.shift();
      result.push(step);
      this.ack(step);
    }
    return result.join``;
  };

  shift = () => {
    const sources = [];
    for (let [key, value] of this.dependencies) {
      if (value.length == 0) {
        sources.push(key);
      }
    }
    if (sources.length == 0) {
      return null;
    }
    sources.sort();
    this.dependencies.delete(sources[0]);
    this.length--;
    return sources[0];
  };

  ack = step => {
    for (let [key, value] of this.dependencies) {
      const depIndex = value.indexOf(step);
      if (depIndex >= 0) {
        value.splice(depIndex, 1);
      }
    }
  };

  process = (num, delay) => {
    const workers = new Workers(num, delay);
    let seconds = 0;
    while (this.length > 0) {
      let todo;
      while (!workers.busy() && (todo = this.shift()) !== null) {
        workers.assign(todo);
      }
      let result = [];
      while (result.length === 0) {
        result = workers.tick();
        seconds++;
      }
      result.forEach(task => this.ack(task));
    }
    while (!workers.idle()) {
      workers.tick().forEach(task => this.ack(task));
      seconds++;
    }
    return seconds;
  };
}

class Workers {
  workers = [];
  delay = 0;
  constructor(num, delay) {
    this.workers = new Array(num).fill().map((w, i) => ({
      worker: String.fromCharCode(65 + i),
      task: "",
      secondsToDo: 0
    }));
    this.delay = delay;
  }
  assign = step =>
    Object.assign(this.workers.find(w => w.task === ""), {
      secondsToDo: this.delay + step.charCodeAt() - 64,
      task: step
    });
  busy = () => this.workers.every(w => w.task !== "");
  idle = () => this.workers.every(w => w.task === "");
  tick = () =>
    this.workers.reduce((done, w) => {
      if (w.task !== "") {
        w.secondsToDo--;
        if (w.secondsToDo == 0) {
          done.push(w.task);
          w.task = "";
        }
      }
      return done;
    }, []);
}

export { sortDependencies, processDependencies };
