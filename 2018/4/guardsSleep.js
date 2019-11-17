function guardsSleepS1(input) {
  const guards = parse(input);

  const guard = Object.values(guards).reduce((acc, cur) =>
    cur.total > acc.total ? cur : acc
  );
  const minute = Object.entries(guard.byMinute).reduce((acc, cur) =>
    acc[1] > cur[1] ? acc : cur
  )[0];

  return guard.id * minute;
}

function guardsSleepS2(input) {
  const guards = parse(input);

  const guard = Object.values(guards).reduce((acc, cur) => {
    if (!acc.minuteMostSlept) {
      acc.minuteMostSlept = Object.entries(acc.byMinute).reduce((mAcc, mCur) =>
        mAcc[1] > mCur[1] ? mAcc : mCur
      );
    }
    if (!cur.minuteMostSlept) {
      cur.minuteMostSlept = Object.entries(cur.byMinute).reduce((mAcc, mCur) =>
        mAcc[1] > mCur[1] ? mAcc : mCur
      );
    }

    return acc.minuteMostSlept[1] > cur.minuteMostSlept[1] ? acc : cur;
  });
  return guard.id * guard.minuteMostSlept[0];
}

function parse(input) {
  const actions = input
    .split("\n")
    .map(parseLine)
    .sort((l1, l2) => (l1.timestamp > l2.timestamp ? 1 : -1));

  let onDutyGuard;
  let fellAsleep;
  let guards = {};

  actions.forEach(element => {
    if (element.action.substr(0, 5) == "Guard") {
      onDutyGuard = element.action.match(/#([0-9]+)/)[1];
    } else if (element.action == "falls asleep") {
      fellAsleep = element.timestamp;
    } else if (element.action == "wakes up") {
      if (!guards[onDutyGuard]) {
        guards[onDutyGuard] = {
          id: onDutyGuard,
          byMinute: {},
          total: 0
        };
      }

      for (
        let i = fellAsleep.getUTCMinutes();
        i < element.timestamp.getUTCMinutes();
        i++
      ) {
        if (!guards[onDutyGuard].byMinute[i]) {
          guards[onDutyGuard].byMinute[i] = 0;
        }
        guards[onDutyGuard].byMinute[i]++;
        guards[onDutyGuard].total++;
      }
    }
  });
  return guards;
}

function parseLine(line) {
  const matches = line.match(/\[(.*)\] (.*)/);
  return {
    timestamp: new Date(matches[1] + " GMT+00:00"),
    action: matches[2]
  };
}

export { guardsSleepS1, guardsSleepS2 };
