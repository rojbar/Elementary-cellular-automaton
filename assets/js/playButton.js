import drawGeneration from "./drawGeneration.js";

function playButton(){
    document.querySelector("#playButton").addEventListener('click', function () {
        for (let i = 0; i < 100; i++) {
            if (getCurrentGeneration() == 0) {
                setInitialState(obtainInitialState());
            }

            calculateNextGeneration();
            let currentState = getCurrentState();
            drawGeneration(currentState.join(" "), 15, getCurrentGeneration());
        }
    })
}

export default playButton
