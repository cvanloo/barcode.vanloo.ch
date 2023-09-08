async function supported_types() {
    const res = await fetch("/api/supported_types", {
        method: 'GET'
    })
    return res.ok ? (await res.json()) : []
}

async function request(type, text) {
    const url = new URL(`${window.location}api/create_barcode`)
    url.searchParams.set('type', type)
    url.searchParams.set('text', text)

    const res = await fetch(url.href, { method: 'GET' })
    if (!res.ok) throw new Error(`${res.status} ${res.statusText}: ${res.body}`)
    return URL.createObjectURL(await res.blob())
}

export { supported_types, request }
