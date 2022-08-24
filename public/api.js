function requestBarcodeTypes() {
    // TODO: Make API request
    return [{
        value: "code-128",
        name: "Code-128"
    }, {
        value: "gs1-128",
        name: "GS1-128"
    }]
}

function requestBarcodeOfType(type, text) {
    // TODO: Make API request
    const img = new Image(312, 100)
    img.src = "barcode.png"
    return {bi: img, bt: text}
}
