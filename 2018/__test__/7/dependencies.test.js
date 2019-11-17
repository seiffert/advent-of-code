import { sortDependencies, processDependencies } from "../../7/dependencies";

const INPUT = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`;

test("dependency order", () => {
  expect(sortDependencies(INPUT)).toBe("CABDFE");
});

test("process with workers", () => {
  expect(processDependencies(INPUT, 2, 0)).toBe(15);
});
