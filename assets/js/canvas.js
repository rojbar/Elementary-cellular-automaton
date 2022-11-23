export function resizeCanvas() {
    const canvas = document.getElementById('canvas');
    canvas.style.width = "100%";
    canvas.style.height = "1000%";
    canvas.width = canvas.offsetWidth;
    canvas.height = canvas.offsetHeight;
}

export function clearCanvas() {
    const canvas = document.getElementById('canvas');
    if (canvas.getContext) {
        const ctx = canvas.getContext('2d');
        ctx.clearRect(0, 0, canvas.width, canvas.height);
    }
}
