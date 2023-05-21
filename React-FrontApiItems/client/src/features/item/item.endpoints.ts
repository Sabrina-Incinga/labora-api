export async function getItems(page: number, itemsPerPage:number) {
    const res = await fetch(`http://localhost:8000/items?page=${page}&itemsPerPage=${itemsPerPage}`)
    const data = await res.json()
  
    return await data
  }