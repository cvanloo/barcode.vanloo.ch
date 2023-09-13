async function supported_types() {
    const res = await fetch("http://localhost:8080/api/supported_types", {
        method: 'GET'
    })
    return res.ok ? (await res.json()) : []
}

async function request(type, text) {
    const url = new URL(`http://localhost:8080/api/create_barcode`)
    url.searchParams.set('type', type)
    url.searchParams.set('text', text)

    const res = await fetch(url.href, { method: 'GET' })
    if (!res.ok) throw new Error(`${res.status} ${res.statusText}: ${res.body}`)
    return URL.createObjectURL(await res.blob())
}

export { supported_types, request }
