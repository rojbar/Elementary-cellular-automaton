import {clearCanvas} from "./canvas.js";
import obtainInitialState from "./obtainInitialState.js";

function newAutomatonButton(){
    document.querySelector('#newAutomaton').addEventListener('click', function () {
        clearCanvas();
        let cellAmount = parseInt(document.querySelector('#amountCells').value);
        let rightNeighbors = parseInt(document.querySelector("#rightNeighbors").value);
        let leftNeighbors = parseInt(document.querySelector("#leftNeighbors").value);
        let states = document.querySelector("#states").value.split(",");
        window.document.states = states
        let keys = []
        let values = []
        let transitions = document.querySelector("#transitions").value.split("\n");
        for (let i = 0; i < transitions.length; i++) {
            let keyValue = transitions[i].trim().split("-")
            keys[i] = keyValue[0]
            values[i] = keyValue[1]
        }
    
        createCellsRow(cellAmount, states)
        newAutomaton(cellAmount, leftNeighbors, rightNeighbors, states, keys, values)
        setInitialState(obtainInitialState());
    
    });
}

/**
 * 
 * @param {int} amount 
 * @param {[]string} states 
 */
function createCellsRow(amount, states) {
    let row = document.querySelector('#cellRow');
    row.innerHTML = ""

    for (let i = 0; i < amount; i++) {
        const textNode = document.createElement('div');
        textNode.innerText = states[0];
        textNode.className = "col border border-secondary border-2 h-100 d-flex align-items-center justify-content-center cell";
        textNode.setAttribute("id", "cell" + i);
        textNode.addEventListener('click', cellChangeState);
        row.appendChild(textNode);
    }

}

/**
 * 
 * @param {event} e 
 */
function cellChangeState(e) {
    const cellID = e.target.getAttribute("id");
    const cell   = document.querySelector("#" + cellID);
    const states = window.document.states;

    let posNewState = -1;
    for (let i = 0; i < states.length; i++) {
        if (states[i] == cell.innerText) {
            posNewState = i + 1;
            break
        }
    }

    if (posNewState == states.length) {
        posNewState = 0;
    }

    cell.innerText = states[posNewState];
}

export default newAutomatonButton