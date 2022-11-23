function generateRulesButton(){
    document.querySelector("#generateRules").addEventListener('click', function () {
        let states = document.querySelector("#states").value.split(",");
        let rightNeighbors = parseInt(document.querySelector("#rightNeighbors").value);
        let leftNeighbors = parseInt(document.querySelector("#leftNeighbors").value);
    
        if (states.length >= 8 || rightNeighbors+leftNeighbors >= 12) {
            window.alert("too many states, too many rulesets!")
            return
        }

        let perm = configurations(states, rightNeighbors + leftNeighbors + 1);
        let transitions = document.querySelector("#transitions")
    
        let value = "";
        for (let i = 0; i < perm.length; i++) {
            if (i === perm.length - 1) {
                value += perm[i].join(";") + `;-${states[0]}`;
                break
            }
            value += perm[i].join(";") + `;-${states[0]}\n`;
        }
        transitions.value = value;
    });
}

/**
 * 
 * @param {[]string} states 
 * @param {int} cellsQuantity 
 * @returns []string
 */
function configurations(states, cellsQuantity) {
    let answer = states.map(x => [x])
    for (let i = 1; i < cellsQuantity; i++) {
        answer = multiplyArrays(answer, states.map(x => [x]))
    }

    return answer;
};

/**
 * 
 * @param {[]any} a 
 * @param {[]any} b 
 * @returns []any
 */
function multiplyArrays(a, b) {
    let answer = []
    for (let i = 0; i < a.length; i++) {
        for (let j = 0; j < b.length; j++) {
            answer.push(a[i].concat(b[j]))
        }
    }

    return answer
}

export default generateRulesButton
