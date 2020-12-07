const forms = require("fs").readFileSync("input.txt", "utf8").split("\n");
var textByLine = forms.map(x => x.split("\n"));


part1 = forms.reduce(
    (total, input) =>total + new Set(input.replace(/\n/g, '')).size,
    0
);
console.log(part1)


function part2(textByLine) {
    var yesCount = 0;
    textByLine.forEach(group => {
        var originalArray = group[0].split("");
        for (let i = 1; i < group.length; i++) {
            var compareArray = group[i].split("");

            originalArray.forEach(element => {
                if (compareArray.indexOf(element) === -1) {
                    originalArray = originalArray.filter(x => x !== element);
                }
            })
        }
        yesCount += originalArray.length;
    });
    return yesCount;
}

console.log(part2(textByLine));


