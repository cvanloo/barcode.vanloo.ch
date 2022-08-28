async function requestBarcodeTypes() {
    const res = await fetch("http://barcode.vanloo.ch/api/supported_types", {
        method: 'GET'
    })
    if (res.ok) {
        const p = await res.json()
        return p
    }
    return []
}

async function requestBarcodeOfType(type, text) {
    const url = new URL("http://barcode.vanloo.ch/api/create_barcode")
    url.searchParams.set('type', type)
    url.searchParams.set('text', text)

    const res = await fetch(url.href, {
        method: 'GET'
    })

    if (res.ok) {
        const blob = await res.blob()
        const imgURL = URL.createObjectURL(blob)
        return imgURL
    }

    return null // throw an exception?
}
