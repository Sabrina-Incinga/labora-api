import { itemDTO } from "../types/item.types"

const endpoint = 'http://localhost:8000/items'

export async function getItems(page: number, itemsPerPage:number) {
    const res = await fetch(`${endpoint}?page=${page}&itemsPerPage=${itemsPerPage}`)
    const data = await res.json()
  
    return await data
}

export async function updateItem(id: number, payload: itemDTO) {
  const config = {
    method: "PUT",
    body: JSON.stringify(payload),
    withCredentials: true,
    crossdomain: true,
    headers: { "Content-Type": "application/json" },
 }
  const res = await fetch(`${endpoint}/${id}`, config)
  const data = res.headers.get('X-Message')

  return data
}

export async function deleteItem(id: number) {
  const config = {
    method: "DELETE",
    withCredentials: true,
    crossdomain: true,
    headers: { "Content-Type": "application/json" },
  }

  const res = await fetch(`${endpoint}/${id}`,config)
  const data = res.headers.get('X-Message')

  return data
}