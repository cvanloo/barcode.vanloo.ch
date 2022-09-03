Storage.prototype.getObject = function(key) {
    const val = this.getItem(key)
    return val && JSON.parse(val)
}

Storage.prototype.setObject = function(key, obj) {
    this.setItem(key, JSON.stringify(obj))
}

function uuidv4() {
    return ([1e7] + -1e3 + -4e3 + -8e3 + -1e11).replace(/[018]/g, c =>
        (c ^ crypto.getRandomValues(new Uint8Array(1))[0] & 15 >> c / 4).toString(16)
    );
}

Storage.prototype.getSession = function() {
    let u = this.getItem('uuid')
    if (null === u) {
        u = uuidv4()
        this.setItem('uuid', u)
    }
    return u
}

Storage.prototype.newSession = function() {
    let u = uuidv4()
    this.setItem('uuid', u)
    return u
}
