/**
 * 
 * @returns []string
 */
function obtainInitialState() {
    let cells = document.querySelector('#cellRow').children;

    let states = []
    for (let i = 0; i < cells.length; i++) {
        states[i] = cells[i].innerText
    }
    return states
}

export default obtainInitialState