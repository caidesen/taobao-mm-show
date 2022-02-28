
export function getNextPicture() {
    return fetch('/api/list?len=5').then(res => res.json()).then(res => res.data as MmPicture[])
}
