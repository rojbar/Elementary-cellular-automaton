import drawGeneration from "./drawGeneration.js";
import obtainInitialState from "./obtainInitialState.js";

function playButton(){
    document.querySelector("#playButton").addEventListener('click', function () {
        for (let i = 0; i < 100; i++) {
            if (getCurrentGeneration() == 0) {
                setInitialState(obtainInitialState());
            }

            calculateNextGeneration();
            let currentState = getCurrentState();
            drawGeneration(currentState.join(" "), 10, getCurrentGeneration());
        }
    })
}

export default playButton
