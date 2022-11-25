import obtainInitialState from "./obtainInitialState.js"
import drawGeneration from "./drawGeneration.js"

function nexGenerationButton(){
    document.querySelector("#nextGeneration").addEventListener('click', function () {
        if (getCurrentGeneration() == 0) {
            setInitialState(obtainInitialState());
        }
    
        calculateNextGeneration();
        let currentState = getCurrentState();
        drawGeneration(currentState.join(" "), 10, getCurrentGeneration());
    });
}
export default nexGenerationButton


